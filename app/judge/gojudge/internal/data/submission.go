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
	"sync"
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
	// set submission status to judging
	s.Status = util.Judging
	marshal, _ := json.Marshal(s)
	r.data.redis.Set(ctx, fmt.Sprintf("%s:%d:%s", SubmissionKey, s.UserID, s.ID), marshal, 2*time.Hour)

	// begin getting basic info of the problem
	// set submission status to system error
	s.Status = util.SystemError

	// cache submission after judging
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

	// get problem from redis or db
	// TODO: add case version for cache
	var p *ent.Problem
	pBytes, err := r.data.redis.Get(ctx, fmt.Sprintf("%s:%d", ProblemKey, s.ProblemID)).Result()
	if err != nil {
		// get problem from ent
		p, err = r.data.db.Problem.Query().
			Where(problem.ID(s.ProblemID)).
			First(ctx)
		if err != nil {
			return err
		}
		pStr, _ := json.Marshal(p)
		r.data.redis.Set(ctx, fmt.Sprintf("%s:%d", ProblemKey, s.ProblemID), pStr, 2*time.Hour)
	} else {
		_ = json.Unmarshal([]byte(pBytes), &p)
	}

	// get judge config
	config, err := r.data.fm.GetConfig(p.ID)
	if err != nil {
		return err
	}

	// set case version and submission info
	s.CaseVer = p.CaseVersion
	s.TotalTime = 0
	s.MaxMemory = 0
	s.Point = 0

	// init submission cases
	subtaskCreates := make([]*ent.SubmissionSubtaskCreate, len(config.Task.Subtasks))
	cases := make([][]*ent.SubmissionCaseCreate, len(config.Task.Subtasks))

	// save submission to database after judging
	defer func() {
		// save submission to database
		submission, err := r.data.db.Submission.Create().
			SetUserID(s.UserID).
			SetProblemID(s.ProblemID).
			SetCode(s.Code).
			SetState(s.Status).
			SetPoint(s.Point).
			SetTotalTime(s.TotalTime).
			SetMaxMemory(s.MaxMemory).
			SetLanguage(s.Language).
			SetCompileStderr(s.Stderr).
			SetCaseVersion(int8(s.CaseVer)).
			SetCreateTime(s.CreateTime).
			Save(ctx)
		if err != nil {
			r.log.Errorf("save submission error: %v", err)
			return
		}

		// skip save cases and subtasks to database when system error
		if s.Status == util.SystemError || s.Status == util.CompileError {
			return
		}

		// save subtasks and cases to database
		for _, builder := range subtaskCreates {
			builder.SetSubmissionID(submission.ID)
		}
		subtasks, err := r.data.db.SubmissionSubtask.CreateBulk(subtaskCreates...).Save(ctx)
		if err != nil {
			r.log.Errorf("save subtasks error: %v", err)
			return
		}

		submissionCase := make([]*ent.SubmissionCaseCreate, 0)
		for i, submissionCaseCreates := range cases {
			for _, create := range submissionCaseCreates {
				create.SetSubmissionSubtaskID(subtasks[i].ID)
				submissionCase = append(submissionCase, create)
			}
		}
		_, err = r.data.db.SubmissionCase.CreateBulk(submissionCase...).Save(ctx)
		if err != nil {
			r.log.Errorf("save cases error: %v", err)
		}
	}()

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

	// delete compiled file after judging
	defer func() {
		if err := r.data.gojudge.DeleteFile(fileID); err != nil {
			r.log.Errorf("delete file error: %v", err)
		}
	}()

	// begin judging
	// init submission info
	s.Status = util.Accepted

	// test each case
	for i, subtask := range config.Task.Subtasks {
		var subtaskState int16 = util.Accepted
		var totalTime uint64 = 0
		var maxMemory uint64 = 0
		submissionCaseCreates := make([]*ent.SubmissionCaseCreate, len(subtask.Cases))
		casesResult := make([]bool, len(subtask.Cases))
		subtaskBuilder := r.data.db.SubmissionSubtask.Create().
			SetState(util.Waiting).
			SetTotalTime(0).
			SetMaxMemory(0).
			SetPoint(0)
		// wait for all cases finished
		wg := sync.WaitGroup{}
		wg.Add(len(subtask.Cases))
		for j, c := range subtask.Cases {
			err := func() error {
				casesResult[j] = false

				builder := r.data.db.SubmissionCase.Create().
					SetState(util.Waiting).
					SetTime(0).
					SetMemory(0).
					SetPoint(0)

				// TODO: cache case

				// delete test file and set submission cases
				defer func() {
					submissionCaseCreates[j] = builder
					wg.Done()
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
					builder.SetStderr(err.Error())
					s.Status = util.RuntimeError
					s.Stderr = err.Error()
					return nil
				}
				r.log.Infof("submission: %s result: %+v", s.ID, result)

				// get result state
				state := gojudge.Convert(result.Status)

				if out := result.Files["stdout"]; state == util.Accepted && out != nil {
					switch p.CompareType {
					case problem.CompareTypeIGNORE_LINE_END_SPACE_AND_TEXT_END_ENTER:
						if !util.StringCompareIgnoreLineEndSpaceAndTextEndEnter(string(out), string(ans)) {
							state = util.WrongAnswer
						}
					case problem.CompareTypeSTRICT:
						if !bytes.Equal(out, ans) {
							if util.BytesMatchIgnoringSpacesAndNewlines(out, ans) {
								state = util.PresentationError
							} else {
								state = util.WrongAnswer
							}
						}
					default:
						state = util.SystemError
						r.log.Errorf("unsupported compare type: %v", p.CompareType)
					}
				}

				// set submission case
				builder.SetState(state)
				builder.SetTime(result.Time)
				builder.SetMemory(result.Memory)
				// TODO: support more contest type

				// set subtask info
				totalTime += result.Time
				maxMemory = max(maxMemory, result.Memory)

				// set submission info
				s.TotalTime += result.Time
				s.MaxMemory = max(s.MaxMemory, result.Memory)

				// set status and point
				if state != util.Accepted {
					builder.SetPoint(0)
					if s.Status == util.Accepted {
						s.Status = state
					}
					if subtaskState == util.Accepted {
						subtaskState = state
					}
					msg, ok := result.Files["stderr"]
					if ok {
						builder.SetStderr(string(msg))
					} else {
						builder.SetStdout(result.Error)
					}
				} else {
					builder.SetPoint(c.Score)
					casesResult[j] = true
				}

				return nil
			}()
			if err != nil {
				s.Status = util.SystemError
				return err
			}
		}
		wg.Wait()
		taskPoint, casesPoint, err := pJudge.Judging(casesResult, config.Task.TaskType, subtask)
		if err != nil {
			r.log.Errorf("judging score error: %v", err)
			s.Status = util.SystemError
			return err
		}
		for i, casesPoint := range casesPoint {
			submissionCaseCreates[i].SetPoint(casesPoint)
		}
		cases[i] = submissionCaseCreates
		subtaskBuilder.SetPoint(taskPoint)
		subtaskBuilder.SetState(subtaskState)
		subtaskBuilder.SetTotalTime(totalTime)
		subtaskBuilder.SetMaxMemory(maxMemory)
		subtaskCreates[i] = subtaskBuilder

		s.Point += taskPoint
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
