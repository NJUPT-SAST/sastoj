package biz

import (
	"context"
	//"github.com/go-kratos/kratos/v2/errors"

	//v1 "sastoj/api/sastoj/admin/problem/service/v1"
	//
	//"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1., "user not found")
)

type Problem struct {
	Id          int64
	Title       string
	Content     string
	Point       int32
	ContestId   int64
	CaseVersion int32
	Index       int32
	Config      string
}

type ProblemRepo interface {
	Save(context.Context, *Problem) (*Problem, error)
	Update(context.Context, *Problem) (*int, error)
	Delete(context.Context, int64) (*int, error)
	FindByID(context.Context, int64) (*Problem, error)
	ListPages(ctx context.Context, currency int32, size int32) ([]*Problem, error)
}

type ProblemUsecase struct {
	repo ProblemRepo
	log  *log.Helper
}

func NewProblemUsecase(repo ProblemRepo, logger log.Logger) *ProblemUsecase {
	return &ProblemUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateProblem creates a Problem, and returns the new Problem.
func (uc *ProblemUsecase) CreateProblem(ctx context.Context, g *Problem) (*Problem, error) {
	//TODO check validation of problem: (code?)
	uc.log.WithContext(ctx).Infof("CreateProblem: %v", g)
	rv, err := uc.repo.Save(ctx, g)
	if err != nil {
		return nil, err
	}
	return rv, nil

}

func (uc *ProblemUsecase) UpdateProblem(ctx context.Context, g *Problem) (bool, error) {
	//TODO check validation of problem: (code?)
	uc.log.WithContext(ctx).Infof("UpdateProblem: %v", g)
	rv, err := uc.repo.Update(ctx, g)
	if err != nil || *rv == 0 {
		return false, err
	}
	return true, nil
}

func (uc *ProblemUsecase) DeleteProblem(ctx context.Context, id int64) (bool, error) {
	uc.log.WithContext(ctx).Infof("DeleteProblem: %v", id)
	rv, err := uc.repo.Delete(ctx, id)
	if err != nil || *rv == 0 {
		return false, err
	}
	return true, nil
}

func (uc *ProblemUsecase) FindProblem(ctx context.Context, id int64) (*Problem, error) {
	uc.log.WithContext(ctx).Infof("FindProblem: %v", id)
	rv, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *ProblemUsecase) ListProblem(ctx context.Context, currency int32, size int32) ([]*Problem, error) {
	uc.log.WithContext(ctx).Infof("ListProblem: %v %v", currency, size)
	rv, err := uc.repo.ListPages(ctx, currency, size)
	if err != nil {
		return nil, err
	}
	return rv, nil
}
