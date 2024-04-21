package data

import (
	"context"

	"sastoj/app/admin/contest/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type ContestRepo struct {
	data *Data
	log  *log.Helper
}

func NewContestRepo(data *Data, logger log.Logger) biz.ContestRepo {
	return &ContestRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ContestRepo) Save(ctx context.Context, g *biz.Contest) (*biz.Contest, error) {
	return g, nil
}

func (r *ContestRepo) Update(ctx context.Context, g *biz.Contest) (*biz.Contest, error) {
	return g, nil
}

func (r *ContestRepo) FindByID(context.Context, int64) (*biz.Contest, error) {
	return nil, nil
}

func (r *ContestRepo) ListByHello(context.Context, string) ([]*biz.Contest, error) {
	return nil, nil
}

func (r *ContestRepo) ListAll(context.Context) ([]*biz.Contest, error) {
	return nil, nil
}
