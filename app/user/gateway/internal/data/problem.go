package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/user/gateway/internal/biz"
)

type problemRepo struct {
	data *Data
	log  *log.Helper
}

func (p *problemRepo) GetProblems(ctx context.Context, contestID int64) ([]*biz.Problem, error) {
	return p.data.cache.contest2problems[contestID], nil
}

func (p *problemRepo) GetProblem(ctx context.Context, problemID, contestID int64) (*biz.Problem, error) {
	return p.data.cache.problems[problemID], nil
}

func NewProblemRepo(data *Data, logger log.Logger) biz.ProblemRepo {
	return &problemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
