package data

import (
	"context"
	"encoding/json"
	"errors"
	"sastoj/app/user/contest/internal/biz"
	"sastoj/ent/submission"
	"sastoj/pkg/mq"
	"sastoj/pkg/util"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

type submissionRepo struct {
	data *Data
	log  *log.Helper
}

func (s *submissionRepo) GetCases(submissionID int64, userID int64) ([]*biz.Case, error) {
	po, err := s.data.redis.Get(context.Background(), "case:"+strconv.FormatInt(userID, 10)+":"+strconv.FormatInt(submissionID, 10)).Result()
	if err != nil {
		return nil, err
	}
	var cases []*biz.Case
	err = json.Unmarshal([]byte(po), &cases)
	if err != nil {
		return nil, err
	}
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

func (s *submissionRepo) GetSubmission(ctx context.Context, submitID int64, userID int64) (*biz.Submission, error) {
	po, err := s.data.db.Submission.Get(ctx, submitID)
	if err != nil {
		return nil, err
	}
	if po.UserID != userID {
		// TODO: ADD Global Error: Permission Denied
		return nil, errors.New("permission denied")
	}
	return &biz.Submission{
		ID:          strconv.FormatInt(po.ID, 10),
		UserID:      po.UserID,
		ProblemID:   po.ProblemID,
		Code:        po.Code,
		Status:      po.Status,
		Point:       po.Point,
		CreateTime:  po.CreateTime,
		TotalTime:   po.TotalTime,
		MaxMemory:   po.MaxMemory,
		Language:    po.Language,
		CaseVersion: po.CaseVersion,
	}, nil
}

func (s *submissionRepo) GetSubmissions(ctx context.Context, userId int64, problemId int64) ([]*biz.Submission, error) {
	submissions, err := s.data.db.Submission.Query().
		Select(submission.FieldID, submission.FieldStatus, submission.FieldPoint, submission.FieldLanguage, submission.FieldCreateTime).
		Where(submission.UserIDEQ(userId), submission.ProblemIDEQ(problemId)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var submits []*biz.Submission
	for _, v := range submissions {
		submits = append(submits, &biz.Submission{
			ID:         strconv.FormatInt(v.ID, 10),
			UserID:     userId,
			ProblemID:  problemId,
			Status:     v.Status,
			Point:      v.Point,
			CreateTime: v.CreateTime,
		})
	}
	return submits, nil
}

func (s *submissionRepo) CreateSubmission(ctx context.Context, submission *biz.Submission) error {
	return s.data.sCh.Publish(ctx, &mq.Submission{
		ID:        submission.ID,
		UserID:    submission.UserID,
		ProblemID: submission.ProblemID,
		Code:      submission.Code,
	})
}

func (s *submissionRepo) UpdateStatus(ctx context.Context, submitID int64, status int16) error {
	_, err := s.data.db.Submission.UpdateOneID(submitID).
		SetStatus(status).
		Save(ctx)
	return err
}

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
