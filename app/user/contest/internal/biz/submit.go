package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Submit struct {
	ID          int       `json:"id,omitempty"`
	UserID      int       `json:"user_id,omitempty"`
	ProblemID   int       `json:"problem_id,omitempty"`
	Code        string    `json:"code,omitempty"`
	State       int       `json:"state,omitempty"`
	Point       int       `json:"point,omitempty"`
	CreateTime  time.Time `json:"create_time"`
	TotalTime   int       `json:"total_time,omitempty"`
	MaxMemory   int       `json:"max_memory,omitempty"`
	Language    string    `json:"language,omitempty"`
	CaseVersion int       `json:"case_version,omitempty"`
}

type Pretest struct {
	ID       string `json:"id,omitempty"`
	UserID   int    `json:"user_id,omitempty"`
	Code     string `json:"code,omitempty"`
	Language string `json:"language,omitempty"`
}

type SubmitRepo interface {
	CreateSubmit(ctx context.Context, submit *Submit) (int, error)
	UpdateState(ctx context.Context, submitID int, state int) error
	UpdateSubmit(ctx context.Context, submit *Submit) error
	GetSubmission(ctx context.Context, submitID int, userID int) (*Submit, error)
	CreatePretest(ctx context.Context, pretest *Pretest) error
}

type SubmitUsecase struct {
	repo SubmitRepo
	log  *log.Helper
}

func NewSubmitUsecase(repo SubmitRepo, logger log.Logger) *SubmitUsecase {
	return &SubmitUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SubmitUsecase) CreateSubmit(ctx context.Context, submit *Submit) (int, error) {
	submitID, err := uc.repo.CreateSubmit(ctx, submit)
	if err != nil {
		return 0, err
	}
	return submitID, nil
}

func (uc *SubmitUsecase) UpdateState(ctx context.Context, submitID int, state int) error {
	return uc.repo.UpdateState(ctx, submitID, state)
}

func (uc *SubmitUsecase) UpdateSubmit(ctx context.Context, submit *Submit) error {
	return uc.repo.UpdateSubmit(ctx, submit)
}

func (uc *SubmitUsecase) GetSubmission(ctx context.Context, submitID int, userID int) (*Submit, error) {
	return uc.repo.GetSubmission(ctx, submitID, userID)
}

func (uc *SubmitUsecase) PretestProblem(ctx context.Context, pretest *Pretest) error {
	return uc.repo.CreatePretest(ctx, pretest)
}
