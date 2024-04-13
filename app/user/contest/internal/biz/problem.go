package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Problem struct {
	ID      int64
	Title   string
	Content string
	Point   int
	Index   int
}

type ProblemRepo interface {
	ListProblem(ctx context.Context, contestID int) ([]*Problem, error)
	GetProblem(ctx context.Context, problemID, contestID int) (*Problem, error)
}

type ProblemUsecase struct {
	repo ProblemRepo
	log  *log.Helper
}

func NewProblemUsecase(repo ProblemRepo, logger log.Logger) *ProblemUsecase {
	return &ProblemUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ProblemUsecase) ListProblem(ctx context.Context, contestID int) ([]*Problem, error) {
	problems, err := uc.repo.ListProblem(ctx, contestID)
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func (uc *ProblemUsecase) GetProblem(ctx context.Context, contestID int, problemID int) (*Problem, error) {
	problem, err := uc.repo.GetProblem(ctx, contestID, problemID)
	if err != nil {
		return nil, err
	}
	return problem, nil
}
