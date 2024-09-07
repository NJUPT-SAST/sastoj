package data

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sastoj/app/judge/gojudge/internal/biz"
	"sastoj/app/judge/gojudge/pkg/gojudge"
	"sastoj/ent"
	"sastoj/ent/problem"
	"sastoj/pkg/mq"
	pJudge "sastoj/pkg/problem"
	"sastoj/pkg/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

const (
	// ProblemKey .
	ProblemKey    = "problem"
	CaseKey       = "case"
	SubmissionKey = "submission"
	SelfTestKey   = "self-test"
)

type submissionRepo struct {
	data *Data
	log  *log.Helper
}

func (r *submissionRepo) JudgeSelfTest(ctx context.Context, test *mq.SelfTest) error {
	test.IsCompiled = false
	test.Stdout = ""
	test.Stderr = ""
	test.Time = 0
	test.Memory = 0

	defer func() {
		marshal, err := json.Marshal(test)
		if err != nil {
			r.log.Errorf("marshal error: %v", err)
			return
		}
		err = r.data.redis.Set(ctx, fmt.Sprintf("%s:%d:%s", SelfTestKey, test.UserID, test.ID), marshal, 2*time.Hour).Err()
		if err != nil {
			r.log.Errorf("cache redis error: %v", err)
		}
	}()

	command := *r.data.gojudge.Commands
	testConfig, ok := command[test.Language]
	if !ok {
		return errors.New("language not supported")
	}

	fileID, result, err := r.data.gojudge.Compile([]byte(test.Code), test.Language, uuid.NewString())
	if err != nil {
		if result != nil {
			test.Stderr = "compile error"
			stderr, ok := result.Files["stderr"]
			if ok {
				test.Stderr = string(stderr)
			}
		}
		return err
	}

	test.IsCompiled = true

	defer func() {
		if err := r.data.gojudge.DeleteFile(fileID); err != nil {
			r.log.Errorf("delete file error: %v", err)
		}
	}()

	result, err = r.data.gojudge.ClassicJudge([]byte(test.Input), test.Language, fileID, uuid.NewString(), testConfig.CompileConfig.CpuTimeLimit, testConfig.CompileConfig.ClockTimeLimit, testConfig.CompileConfig.MemoryLimit, testConfig.CompileConfig.StdoutMaxSize)
	if err != nil {
		return err
	}

	test.Time = result.Time
	test.Memory = result.Memory
	test.Stderr = string(result.Files["stderr"])
	test.Stdout = string(result.Files["stdout"])
	return nil
}

func (r *submissionRepo) JudgeSubmission(ctx context.Context, s *mq.Submission) error {
	s.Status = util.Judging

	marshal, _ := json.Marshal(s)
	r.data.redis.Set(ctx, fmt.Sprintf("%s:%d:%s", SubmissionKey, s.UserID, s.ID), marshal, 2*time.Hour)

	s.Status = util.SystemError

	// cache submission
	defer func() {
		marshal, err := json.Marshal(s)
		if err != nil {
			r.log.Errorf("marshal error: %v", err)
		}
		err = r.data.redis.Set(ctx, fmt.Sprintf("%s:%d:%s", SubmissionKey, s.UserID, s.ID), marshal, 2*time.Hour).Err()
		if err != nil {
			r.log.Errorf("cache redis error: %v", err)
		}
	}()

	// get problem from redis
	// TODO: add case version for cache
	var p *ent.Problem
	pBytes, err := r.data.redis.Get(ctx, fmt.Sprintf("%s:%d", ProblemKey, s.ProblemID)).Result()
	if err != nil {
		// get problem from ent
		p, err = r.data.db.Problem.Query().
			Where(problem.ID(s.ProblemID)).
			WithProblemCases().
			First(ctx)
		if err != nil {
			return err
		}
		pStr, _ := json.Marshal(p)
		r.data.redis.Set(ctx, fmt.Sprintf("%s:%d", ProblemKey, s.ProblemID), pStr, 2*time.Hour)
	} else {
		_ = json.Unmarshal([]byte(pBytes), &p)
	}

	// set case version
	s.CaseVer = p.CaseVersion

	// compile submission
	fileID, result, err := r.data.gojudge.Compile([]byte(s.Code), s.Language, uuid.NewString())
	if err != nil {
		if result != nil {
			s.Status = util.CompileError
			s.Stderr = "Compile error without stderr"
			stderr, ok := result.Files["stderr"]
			if ok {
				s.Stderr = string(stderr)
			}
		}
		return err
	}

	// get judge config
	config, err := r.data.fm.GetJudgeConfig(p.ID)
	if err != nil {
		return err
	}

	// init submission info
	s.TotalTime = 0
	s.MaxMemory = 0
	s.Point = 0
	s.Status = util.Accepted

	// init submission cases
	casesResult := make([]bool, len(config.Task.Cases))
	builders := make([]*ent.SubmissionCaseCreate, len(config.Task.Cases))

	// save submission
	defer func() {
		// delete compiled file
		if err := r.data.gojudge.DeleteFile(fileID); err != nil {
			r.log.Errorf("delete file error: %v", err)
		}

		// skip save to database when system error
		if s.Status == util.SystemError {
			return
		}

		scores, err := pJudge.Judging(casesResult, &config.Task)
		if err != nil {
			r.log.Errorf("judging score error: %v", err)
			return
		}

		s.Point = 0
		for i, score := range scores {
			s.Point += score
			builders[i].SetPoint(score)
		}

		submission, err := r.data.db.Submission.Create().
			SetUserID(s.UserID).
			SetProblemID(s.ProblemID).
			SetCode(s.Code).
			SetStatus(s.Status).
			SetPoint(s.Point).
			SetTotalTime(int32(s.TotalTime)).
			SetMaxMemory(int32(s.MaxMemory)).
			SetLanguage(s.Language).
			SetCompileMessage(s.Stderr).
			SetCaseVersion(int8(s.CaseVer)).
			Save(ctx)
		if err != nil {
			r.log.Errorf("save error: %v", err)
			return
		}

		for _, builder := range builders {
			builder.SetSubmissionID(submission.ID)
		}
		_, err = r.data.db.SubmissionCase.CreateBulk(builders...).Save(ctx)
		if err != nil {
			r.log.Errorf("save error: %v", err)
		}
	}()

	// test each case
	for _, c := range config.Task.Cases {
		err := func() error {
			// get file index
			fileIndex := util.GetCaseIndex(c.Input)

			casesResult[fileIndex-1] = false

			// get problem cases ID
			problemCasesID := int64(-1)
			for _, problemCase := range p.Edges.ProblemCases {
				if problemCase.Index == int16(fileIndex) {
					problemCasesID = problemCase.ID
				}
			}

			if problemCasesID == -1 {
				return errors.New("problem case not found")
			}

			builder := r.data.db.SubmissionCase.Create().
				SetProblemCasesID(0).
				SetState(util.Waiting).
				SetTime(0).
				SetMemory(0).
				SetPoint(0).
				SetMessage("")

			// TODO: cache case

			// delete test file and set submission cases
			defer func() {
				builders[fileIndex-1] = builder
			}()

			// get case input and answer
			in, ans, err := r.data.fm.FetchCase(s.ProblemID, c.Input, c.Answer)
			if err != nil {
				return err
			}
			in = util.RemoveCr(in)
			ans = util.RemoveCr(ans)

			// judge case
			result, err := r.data.gojudge.ClassicJudge(in, s.Language, fileID, uuid.NewString(), uint64(config.ResourceLimits.Time), uint64(config.ResourceLimits.Time*2), uint64(config.ResourceLimits.Memory), int64(len(ans)))
			if err != nil {
				builder.SetState(util.RuntimeError)
				builder.SetMessage(err.Error())
				s.Status = util.RuntimeError
				s.Stderr = err.Error()
				return nil
			}
			r.log.Infof("submission: %s result: %+v", s.ID, result)

			// get result state
			state := gojudge.Convert(result.Status)

			// TODO: support more compare method
			if out := result.Files["stdout"]; state == util.Accepted && out != nil {
				if p.RestrictPresentation && !bytes.Equal(out, ans) {
					if util.BytesMatchIgnoringSpacesAndNewlines(out, ans) {
						state = util.PresentationError
					} else {
						state = util.WrongAnswer
					}
				} else if !util.BytesMatchIgnoringSpacesAndNewlines(out, ans) {
					state = util.WrongAnswer
				}
			}

			// set submission case
			builder.SetProblemCasesID(problemCasesID)
			builder.SetState(state)
			builder.SetTime(int32(result.Time))
			builder.SetMemory(int32(result.Memory))
			// TODO: support more contest type
			builder.SetPoint(c.Score)

			// set submission info
			s.TotalTime += result.Time
			s.MaxMemory = max(s.MaxMemory, result.Memory)

			// set status
			if state != util.Accepted {
				if s.Status == util.Accepted {
					s.Status = state
				}

				msg, ok := result.Files["stderr"]
				if ok {
					builder.SetMessage(string(msg))
				} else {
					builder.SetMessage(result.Error)
				}
			}

			// set result when AC
			casesResult[fileIndex-1] = true

			return nil
		}()
		if err != nil {
			s.Status = util.SystemError
			return err
		}
	}
	return nil
}

// NewSubmissionRepo .
func NewSubmissionRepo(data *Data, logger log.Logger) biz.SubmissionRepo {
	return &submissionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
