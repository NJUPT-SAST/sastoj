package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/ent"
)

// Judger is a Judger model.
type Judger struct {
	Groups []*ent.Group
}

// JudgerRepo is a Greater repo.
type JudgerRepo interface {
	Save(ctx context.Context, groupIds []int64) error
	Update(ctx context.Context, groupIds []int64) error
	FindByID(ctx context.Context, problemId int64) (*Judger, error)
	FindProblemById(ctx context.Context, problemId int64) (*ent.Problem, error)
}

// JudgerUsecase is a Judger usecase.
type JudgerUsecase struct {
	repo JudgerRepo
	log  *log.Helper
}

// NewJudgerUsecase new a Judger usecase.
func NewJudgerUsecase(repo JudgerRepo, logger log.Logger) *JudgerUsecase {
	return &JudgerUsecase{repo: repo, log: log.NewHelper(logger)}
}

// UpdateJudger creates a Judger, and returns the new Judger.
func (uc *JudgerUsecase) UpdateJudger(ctx context.Context, problemId int64, groupIds []int64) error {
	uc.log.WithContext(ctx).Infof("CreateJudger: %v")
	_, err := uc.repo.FindProblemById(ctx, problemId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			//如何类型相同则说明没有该id的problem
			return errors.New("not find the specified Id of problem")
		}
		return err
	}
	judgerGroups, err := uc.repo.FindByID(ctx, problemId)
	if len(judgerGroups.Groups) == 0 {
		return uc.repo.Save(ctx, groupIds)
	} else {
		return uc.repo.Update(ctx, groupIds)

	}
}

func (uc *JudgerUsecase) GetJudger(ctx context.Context, problemId int64) (*Judger, error) {
	return uc.repo.FindByID(ctx, problemId)
}
