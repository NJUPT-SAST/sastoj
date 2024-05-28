package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Submit struct {
	ID          int64     `json:"id,omitempty"`
	UserID      int64     `json:"user_id,omitempty"`
	ProblemID   int64     `json:"problem_id,omitempty"`
	Code        string    `json:"code,omitempty"`
	Status      int16     `json:"state,omitempty"`
	Point       int16     `json:"point,omitempty"`
	CreateTime  time.Time `json:"create_time"`
	TotalTime   int32     `json:"total_time,omitempty"`
	MaxMemory   int32     `json:"max_memory,omitempty"`
	Language    string    `json:"language,omitempty"`
	CaseVersion int8      `json:"case_version,omitempty"`
}

type Pretest struct {
	ID       string `json:"id,omitempty"`
	UserID   int64  `json:"user_id,omitempty"`
	Code     string `json:"code,omitempty"`
	Language string `json:"language,omitempty"`
	Input    string `json:"input,omitempty"`
	Output   string `json:"output,omitempty"`
}

type Case struct {
	ID    int64 `json:"id,omitempty"`
	Index int16 `json:"index,omitempty"`
	State int16 `json:"state,omitempty"`
	Point int16 `json:"point,omitempty"`
}
type SubmitRepo interface {
	CreateSubmit(ctx context.Context, submit *Submit) (int64, error)
	UpdateStatus(ctx context.Context, submitID int64, status int16) error
	UpdateSubmit(ctx context.Context, submit *Submit) error
	GetSubmission(ctx context.Context, submitID int64, userID int64) (*Submit, error)
	CreatePretest(ctx context.Context, pretest *Pretest) error
	GetCases(submissionID int64, userID int64) ([]*Case, error)
}

type SubmitUsecase struct {
	repo SubmitRepo
	log  *log.Helper
}

func NewSubmitUsecase(repo SubmitRepo, logger log.Logger) *SubmitUsecase {
	return &SubmitUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SubmitUsecase) CreateSubmit(ctx context.Context, submit *Submit) (int64, error) {
	submitID, err := uc.repo.CreateSubmit(ctx, submit)
	if err != nil {
		return 0, err
	}
	return submitID, nil
}

func (uc *SubmitUsecase) UpdateStatus(ctx context.Context, submitID int64, status int16) error {
	return uc.repo.UpdateStatus(ctx, submitID, status)
}

func (uc *SubmitUsecase) UpdateSubmit(ctx context.Context, submit *Submit) error {
	return uc.repo.UpdateSubmit(ctx, submit)
}

func (uc *SubmitUsecase) GetSubmission(ctx context.Context, submitID int64, userID int64) (*Submit, error) {
	return uc.repo.GetSubmission(ctx, submitID, userID)
}

func (uc *SubmitUsecase) PretestProblem(ctx context.Context, pretest *Pretest) error {
	return uc.repo.CreatePretest(ctx, pretest)
}

func (uc *SubmitUsecase) GetCases(submissionID int64, userID int64) ([]*Case, error) {
	return uc.repo.GetCases(submissionID, userID)
}
