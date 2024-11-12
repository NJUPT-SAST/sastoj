package biz

import (
	"context"
	pb "sastoj/api/sastoj/admin/admin/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

// Problem is a Problem model.
type Problem struct {
	ID          int64
	TypeID      int64
	Title       string
	Content     string
	Point       int32
	ContestID   int64
	CaseVersion int32
	Index       int32
	Visibility  pb.Visibility
	OwnerID     int64
	Config      string
	Metadata    map[string]string
}

// ProblemType is a ProblemType model.
type ProblemType struct {
	ID          int64
	SlugName    string
	DisplayName string
	Description string
	Judge       string
}

type ProblemRepo interface {
	Save(context.Context, *Problem) (*int64, error)
	Update(context.Context, *Problem) error
	Delete(context.Context, int64) (*int64, error)
	FindByID(context.Context, int64) (*Problem, error)
	ListPages(context.Context, int32, int32, int64) ([]*Problem, int64, error)
	GetTypes(context.Context) ([]*ProblemType, error)
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

func (uc *ProblemUsecase) ListProblem(ctx context.Context, current int32, size int32, contestID int64) ([]*Problem, int64, error) {
	uc.log.WithContext(ctx).Infof("ListProblem: %v %v", current, size)
	rv, total, err := uc.repo.ListPages(ctx, current, size, contestID)
	if err != nil {
		return nil, 0, err
	}
	return rv, total, nil
}

func (uc *ProblemUsecase) GetProblemTypes(ctx context.Context) ([]*ProblemType, error) {
	uc.log.WithContext(ctx).Infof("GetProblemTypes")
	rv, err := uc.repo.GetTypes(ctx)
	if err != nil {
		return nil, err
	}
	return rv, nil
}
