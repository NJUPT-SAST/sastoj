package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Submission struct {
	ID         string    `json:"id,omitempty"`
	UserID     int64     `json:"user_id,omitempty"`
	ProblemID  int64     `json:"problem_id,omitempty"`
	Code       string    `json:"code,omitempty"`
	Status     int16     `json:"state,omitempty"`
	Point      int16     `json:"point,omitempty"`
	CreateTime time.Time `json:"create_time"`
	TotalTime  int32     `json:"total_time,omitempty"`
	MaxMemory  int32     `json:"max_memory,omitempty"`
	Language   string    `json:"language,omitempty"`
}

type SelfTest struct {
	ID       string `json:"id,omitempty"`
	UserID   int64  `json:"user_id,omitempty"`
	Code     string `json:"code,omitempty"`
	Language string `json:"language,omitempty"`
	Input    string `json:"input,omitempty"`
}

type SubmissionRepo interface {
	CreateSubmission(ctx context.Context, submission *Submission) (string, error)
	GetSubmission(ctx context.Context, submissionID string, userID int64) (*Submission, error)
	CreateSelfTest(ctx context.Context, selfTest *SelfTest) error
}

type SubmissionUsecase struct {
	repo SubmissionRepo
	log  *log.Helper
}

func (uc *SubmissionUsecase) CreateSubmission(ctx context.Context, submission *Submission) (string, error) {
	submissionID, err := uc.repo.CreateSubmission(ctx, submission)
	if err != nil {
		return "", err
	}
	return submissionID, nil
}

func (uc *SubmissionUsecase) GetSubmission(ctx context.Context, submissionID string, userID int64) (*Submission, error) {
	submission, err := uc.repo.GetSubmission(ctx, submissionID, userID)
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (uc *SubmissionUsecase) CreateSelfTest(ctx context.Context, selfTest *SelfTest) error {
	return uc.repo.CreateSelfTest(ctx, selfTest)
}

func NewSubmissionUsecase(repo SubmissionRepo, logger log.Logger) *SubmissionUsecase {
	return &SubmissionUsecase{repo: repo, log: log.NewHelper(logger)}
}
