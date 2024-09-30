package biz

import (
	"context"
	"errors"
	"sastoj/ent"

	"github.com/go-kratos/kratos/v2/log"
)

// Adjudicator is an Adjudicator model.
type Adjudicator struct {
	Groups []*ent.Group
}

// AdjudicatorRepo is a Greater repo.
type AdjudicatorRepo interface {
	Save(ctx context.Context, groupIds []int64) error
	Update(ctx context.Context, groupIds []int64) error
	FindByID(ctx context.Context, problemId int64) (*Adjudicator, error)
	FindProblemById(ctx context.Context, problemId int64) (*ent.Problem, error)
}

// AdjudicatorUsecase is a Adjudicator usecase.
type AdjudicatorUsecase struct {
	repo AdjudicatorRepo
	log  *log.Helper
}

// NewAdjudicatorUsecase new a Adjudicator usecase.
func NewAdjudicatorUsecase(repo AdjudicatorRepo, logger log.Logger) *AdjudicatorUsecase {
	return &AdjudicatorUsecase{repo: repo, log: log.NewHelper(logger)}
}

// UpdateAdjudicator creates a Adjudicator, and returns the new Adjudicator.
func (uc *AdjudicatorUsecase) UpdateAdjudicator(ctx context.Context, problemId int64, groupIds []int64) error {
	uc.log.WithContext(ctx).Infof("CreateAdjudicator: %v")
	_, err := uc.repo.FindProblemById(ctx, problemId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			//如何类型相同则说明没有该id的problem
			return errors.New("not find the specified Id of problem")
		}
		return err
	}
	adjudicatorGroups, err := uc.repo.FindByID(ctx, problemId)
	if len(adjudicatorGroups.Groups) == 0 {
		return uc.repo.Save(ctx, groupIds)
	} else {
		return uc.repo.Update(ctx, groupIds)

	}
}

func (uc *AdjudicatorUsecase) GetAdjudicator(ctx context.Context, problemId int64) (*Adjudicator, error) {
	return uc.repo.FindByID(ctx, problemId)
}
