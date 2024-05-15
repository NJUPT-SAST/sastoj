package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	problem "sastoj/api/sastoj/admin/problem/service/v1"
	"sastoj/app/admin/problem/internal/data"
	"time"
)

type ProblemUsecase struct {
	contestRepo *data.ContestRepo
	repo        *data.ProblemRepo
	log         *log.Helper
}

func NewProblemUsecase(repo *data.ProblemRepo, contestRepo *data.ContestRepo, logger log.Logger) *ProblemUsecase {
	return &ProblemUsecase{repo: repo, contestRepo: contestRepo, log: log.NewHelper(logger)}
}

// CreateProblem creates a Problem, and returns the new Problem.
func (uc *ProblemUsecase) CreateProblem(ctx context.Context, g *problem.CreateProblemRequest) (*problem.CreateProblemReply, error) {
	uc.log.WithContext(ctx).Infof("CreateProblem: %v", g)
	if g.CaseVersion != 1 {
		return nil, fmt.Errorf("caseVersion should be 1, not %d", g.CaseVersion)
	}
	rv, err := uc.repo.Save(ctx, g)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *ProblemUsecase) UpdateProblem(ctx context.Context, g *problem.UpdateProblemRequest) (bool, error) {
	contest, err := uc.contestRepo.FindByID(ctx, g.ContestId)
	if err != nil {
		return false, err
	}
	old, err := uc.repo.FindByID(ctx, g.Id)
	if err != nil {
		return false, err
	}
	if contest.StartTime.Before(time.Now()) {
		g.CaseVersion = old.CaseVersion + 1
	}
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

func (uc *ProblemUsecase) FindProblem(ctx context.Context, id int64) (*problem.GetProblemReply, error) {
	uc.log.WithContext(ctx).Infof("FindProblem: %v", id)
	rv, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *ProblemUsecase) ListProblem(ctx context.Context, currency int32, size int32) ([]*problem.ListProblemReply_Problem, error) {
	uc.log.WithContext(ctx).Infof("ListProblem: %v %v", currency, size)
	rv, err := uc.repo.ListPages(ctx, currency, size)
	if err != nil {
		return nil, err
	}
	return rv, nil
}
