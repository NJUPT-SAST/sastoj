package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Problem struct {
	ID       int64
	Title    string
	Type     string
	Content  string
	Score    int32
	Index    int16
	Metadata map[string]string
}

type ProblemRepo interface {
	ListProblem(ctx context.Context, contestID int64) ([]*Problem, error)
	GetProblem(ctx context.Context, problemID, contestID int64) (*Problem, error)
	GetProblemCaseVer(ctx context.Context, problemId int64) (int8, error)
}

type ProblemUsecase struct {
	repo ProblemRepo
	log  *log.Helper
}

func NewProblemUsecase(repo ProblemRepo, logger log.Logger) *ProblemUsecase {
	return &ProblemUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ProblemUsecase) ListProblem(ctx context.Context, contestID int64) ([]*Problem, error) {
	problems, err := uc.repo.ListProblem(ctx, contestID)
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func (uc *ProblemUsecase) GetProblem(ctx context.Context, contestID int64, problemID int64) (*Problem, error) {
	problem, err := uc.repo.GetProblem(ctx, contestID, problemID)
	if err != nil {
		return nil, err
	}
	return problem, nil
}

func (uc *ProblemUsecase) GetProblemCaseVer(ctx context.Context, problemId int64) (int8, error) {
	return uc.repo.GetProblemCaseVer(ctx, problemId)
}
