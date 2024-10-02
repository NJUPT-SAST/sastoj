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

// AdjudicatorRepo is an Adjudicator repository.
type AdjudicatorRepo interface {
	Save(ctx context.Context, groupIDs []int64) error
	Update(ctx context.Context, groupIDs []int64) error
	FindByID(ctx context.Context, problemID int64) (*Adjudicator, error)
	FindProblemByID(ctx context.Context, problemID int64) (*ent.Problem, error)
}

// AdjudicatorUsecase is an Adjudicator usecase.
type AdjudicatorUsecase struct {
	repo AdjudicatorRepo
	log  *log.Helper
}

// NewAdjudicatorUsecase new an Adjudicator usecase.
func NewAdjudicatorUsecase(repo AdjudicatorRepo, logger log.Logger) *AdjudicatorUsecase {
	return &AdjudicatorUsecase{repo: repo, log: log.NewHelper(logger)}
}

// UpdateAdjudicator creates an Adjudicator, and returns the new Adjudicator.
func (uc *AdjudicatorUsecase) UpdateAdjudicator(ctx context.Context, problemID int64, groupIDs []int64) error {
	uc.log.WithContext(ctx).Infof("CreateAdjudicator: %v", groupIDs)
	_, err := uc.repo.FindProblemByID(ctx, problemID)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			// 如何类型相同则说明没有该id的problem
			return errors.New("not find the specified ID of problem")
		}
		return err
	}
	adjudicatorGroups, err := uc.repo.FindByID(ctx, problemID)
	if len(adjudicatorGroups.Groups) == 0 {
		return uc.repo.Save(ctx, groupIDs)
	} else {
		return uc.repo.Update(ctx, groupIDs)

	}
}

func (uc *AdjudicatorUsecase) GetAdjudicator(ctx context.Context, problemID int64) (*Adjudicator, error) {
	return uc.repo.FindByID(ctx, problemID)
}
