package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/judge/freshcup/internal/biz"
	"sastoj/ent"
	"sastoj/ent/problem"
	"sastoj/ent/submission"
	"sastoj/pkg/mq"
	"sastoj/pkg/util"
	"strings"
	"time"
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
	return nil
}

func (r *submissionRepo) JudgeSubmission(ctx context.Context, s *mq.Submission) error {
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
			WithProblemType().
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

	s.Status = util.Accepted

	// get judge config
	config, err := r.data.fm.GetConfig(p.ID)
	if err != nil {
		return err
	}
	if config == nil {
		r.log.Errorf("config load error for problem %d", p.ID)
		return nil
	}

	// save submission
	defer func() {
		// skip save to database when system error
		if s.Status == util.SystemError {
			return
		}
		// save only the latest submission
		result, err := r.data.db.Submission.Update().
			Where(submission.UserIDEQ(s.UserID), submission.ProblemIDEQ(s.ProblemID)).
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
			Save(ctx)
		if err != nil {
			r.log.Errorf("save submission error: %v", err)
			return
		}
		if result == 0 {
			_, err = r.data.db.Submission.Create().
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
				Save(ctx)
		}
	}()

	// judge submission
	switch p.Edges.ProblemType.SlugName {
	case "freshcup-single-choice":
		// Auto judge
		if s.Code == config.ReferenceAnswer {
			s.Point = 100
		} else {
			s.Point = 0
		}
	case "freshcup-multiple-choice":
		// Auto judge
		if s.Code == config.ReferenceAnswer {
			s.Point = 100
		} else if partialContains(s.Code, config.ReferenceAnswer) {
			s.Point = config.PartialScore
		} else {
			s.Point = 0
		}
	case "freshcup-short-answer":
		// Manual judge
		s.Status = util.Waiting
	}

	return nil
}

func partialContains(code, answer string) bool {
	for _, char := range code {
		if !strings.ContainsRune(answer, char) {
			return false
		}
	}
	return true
}

// NewSubmissionRepo .
func NewSubmissionRepo(data *Data, logger log.Logger) biz.SubmissionRepo {
	return &submissionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
