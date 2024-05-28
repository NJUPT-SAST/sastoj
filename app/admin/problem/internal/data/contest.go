package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/ent"
	"sastoj/ent/contest"
)

type ContestRepo struct {
	data *Data
	log  *log.Helper
}

func NewContestRepo(data *Data, logger log.Logger) *ContestRepo {
	return &ContestRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ContestRepo) FindByID(ctx context.Context, id int64) (*ent.Contest, error) {
	v, err := r.data.db.Contest.Query().
		Where(contest.ID(id)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return v, nil
}
