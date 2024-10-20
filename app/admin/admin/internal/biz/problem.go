package biz

import (
	"context"
	pb "sastoj/api/sastoj/admin/admin/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

// Problem is a Problem model.
type Problem struct {
	Id          int64
	TypeId      int64
	Title       string
	Content     string
	Point       int32
	ContestId   int64
	CaseVersion int32
	Index       int32
	Visibility  pb.Visibility
	OwnerId     int64
	Config      string
	Metadata    map[string]string
}

type ProblemRepo interface {
	Save(context.Context, *Problem) (*int64, error)
	Update(context.Context, *Problem) error
	Delete(context.Context, int64) (*int64, error)
	FindByID(context.Context, int64) (*Problem, error)
	ListPages(context.Context, int32, int32) ([]*Problem, error)
}

type ProblemUsecase struct {
	repo ProblemRepo
	log  *log.Helper
}

func NewProblemUsecase(repo ProblemRepo, logger log.Logger) *ProblemUsecase {
	return &ProblemUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateProblem creates a Problem, and returns the new Problem.
func (uc *ProblemUsecase) CreateProblem(ctx context.Context, p *Problem) (*int64, error) {
	uc.log.WithContext(ctx).Infof("CreateProblem: %v", p)
	rv, err := uc.repo.Save(ctx, p)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *ProblemUsecase) UpdateProblem(ctx context.Context, g *Problem) (bool, error) {
	err := uc.repo.Update(ctx, g)
	if err != nil {
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
