package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sastoj/app/user/contest/internal/biz"
	"sastoj/ent"
	"sastoj/ent/submission"
	"sastoj/ent/submissioncase"
	"sastoj/pkg/middleware/auth"
	"sastoj/pkg/mq"
	"sastoj/pkg/util"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type submissionRepo struct {
	data *Data
	log  *log.Helper
}

func (s *submissionRepo) GetCases(ctx context.Context, submissionID string, contestID int64) ([]*biz.Case, error) {
	id, err := strconv.ParseInt(submissionID, 10, 64)
	claim := ctx.Value("userInfo").(*auth.Claims)
	userID := claim.UserId
	var cases []*biz.Case
	if err != nil {
		// get from redis
		po, err := s.data.redis.Get(ctx, fmt.Sprintf("cases:%d:%s", userID, submissionID)).Result()
		if err != nil {
			return nil, fmt.Errorf("no cases found: %s", submissionID)
		}
		err = json.Unmarshal([]byte(po), &cases)
		if err != nil {
			return nil, err
		}
	} else {
		// get from db
		po, err := s.data.db.SubmissionCase.Query().Where(submissioncase.SubmissionIDEQ(id)).All(ctx)
		if err != nil {
			return nil, err
		}
		cases = make([]*biz.Case, 0)
		for i, v := range po {
			cases = append(cases, &biz.Case{
				Index: int32(i),
				State: int32(v.State),
			})
		}
	}
	// TODO: Add permission check, including contest type and userID
	return cases, nil
}

func (s *submissionRepo) CreateSelfTest(ctx context.Context, selfTest *biz.SelfTest) error {
	return s.data.stCh.Publish(ctx, &mq.SelfTest{
		ID:       selfTest.ID,
		UserID:   selfTest.UserID,
		Code:     selfTest.Code,
		Language: selfTest.Language,
		Input:    selfTest.Input,
		Token:    "",
	})
}

func (s *submissionRepo) GetSubmission(ctx context.Context, submissionID string, contestID int64) (*biz.Submission, error) {
	id, err := strconv.ParseInt(submissionID, 10, 64)
	claim := ctx.Value("userInfo").(*auth.Claims)
	userID := claim.UserId
	var res *biz.Submission
	if err != nil {
		// get from redis
		var po *mq.Submission
		result, err := s.data.redis.Get(ctx, fmt.Sprintf("submission:%d:%s", userID, submissionID)).Result()
		if err != nil {
			return nil, fmt.Errorf("no submission found: %s", submissionID)
		}
		err = json.Unmarshal([]byte(result), &po)
		if err != nil {
			return nil, err
		}
		res = &biz.Submission{
			ID:         po.ID,
			UserID:     po.UserID,
			ProblemID:  po.ProblemID,
			Code:       po.Code,
			Status:     po.Status,
			Point:      po.Point,
			CreateTime: po.CreateTime,
			TotalTime:  po.TotalTime,
			MaxMemory:  po.MaxMemory,
			Language:   po.Language,
		}
	} else {
		// get from db
		po, err := s.data.db.Submission.Get(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("no submission found: %s", submissionID)
		}
		if po.UserID != int64(userID) {
			// TODO: ADD Global Error: Permission Denied
			return nil, errors.New("permission denied")
		}
		res = &biz.Submission{
			ID:         strconv.FormatInt(po.ID, 10),
			UserID:     po.UserID,
			ProblemID:  po.ProblemID,
			Code:       po.Code,
			Status:     po.Status,
			Point:      po.Point,
			CreateTime: po.CreateTime,
			TotalTime:  po.TotalTime,
			MaxMemory:  po.MaxMemory,
			Language:   po.Language,
		}
	}
	// TODO: Add contest type check
	return res, nil
}

func (s *submissionRepo) GetSubmissions(ctx context.Context, contestID int64, problemId int64) ([]*biz.Submission, error) {
	claim := ctx.Value("userInfo").(*auth.Claims)
	userID := claim.UserId
	po, err := s.data.db.Submission.Query().
		Select(submission.FieldID, submission.FieldStatus, submission.FieldPoint, submission.FieldLanguage, submission.FieldCreateTime, submission.FieldLanguage).
		Where(submission.UserIDEQ(int64(userID)), submission.ProblemIDEQ(problemId)).
		Order(ent.Desc(submission.FieldCreateTime)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var submissions []*biz.Submission
	for _, v := range po {
		submissions = append(submissions, &biz.Submission{
			ID:         strconv.FormatInt(v.ID, 10),
			ProblemID:  problemId,
			Language:   v.Language,
			Status:     v.Status,
			Point:      v.Point,
			CreateTime: v.CreateTime,
		})
	}
	// TODO: Add contest type check
	return submissions, nil
}

func (s *submissionRepo) CreateSubmission(ctx context.Context, submission *biz.Submission) error {
	m := &mq.Submission{
		ID:         submission.ID,
		UserID:     submission.UserID,
		ProblemID:  submission.ProblemID,
		Code:       submission.Code,
		Status:     util.Waiting,
		Point:      0,
		CreateTime: submission.CreateTime,
		TotalTime:  0,
		MaxMemory:  0,
		Language:   submission.Language,
		Token:      "",
	}
	err := s.data.sCh.Publish(ctx, m)
	if err != nil {
		return err
	}
	marshal, err := json.Marshal(m)
	if err != nil {
		return err
	}
	set := s.data.redis.Set(ctx, fmt.Sprintf("submission:%d:%s", submission.UserID, submission.ID), marshal, 2*time.Hour)
	if set.Err() != nil {
		return set.Err()
	}
	return nil
}

//func (s *submissionRepo) UpdateStatus(ctx context.Context, submitID string, status int16) error {
//	_, err := s.data.db.Submission.UpdateOneID(submitID).
//		SetStatus(status).
//		Save(ctx)
//	return err
//}

func (s *submissionRepo) UpdateSubmission(ctx context.Context, submission *biz.Submission) error {
	_, err := s.data.db.Submission.UpdateOneID(util.ParseInt64(submission.ID)).
		SetStatus(submission.Status).
		SetPoint(submission.Point).
		SetTotalTime(submission.TotalTime).
		SetMaxMemory(submission.MaxMemory).
		Save(ctx)
	return err
}

func NewSubmitRepo(data *Data, logger log.Logger) biz.SubmissionRepo {
	return &submissionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
