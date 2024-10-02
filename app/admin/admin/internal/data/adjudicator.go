package data

import (
	"context"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent"
	"sastoj/ent/problem"

	"github.com/go-kratos/kratos/v2/log"
)

type adjudicatorRepo struct {
	data *Data
	log  *log.Helper
}

// NewAdjudicatorRepo .
func NewAdjudicatorRepo(data *Data, logger log.Logger) biz.AdjudicatorRepo {
	return &adjudicatorRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *adjudicatorRepo) Save(ctx context.Context, groupIDs []int64) error {
	return r.data.db.Problem.Update().AddAdjudicatorIDs(groupIDs...).Exec(ctx)
}

func (r *adjudicatorRepo) Update(ctx context.Context, groupIDs []int64) error {
	return r.data.db.Problem.Update().ClearAdjudicators().AddAdjudicatorIDs(groupIDs...).Exec(ctx)
}

func (r *adjudicatorRepo) FindByID(ctx context.Context, problemID int64) (*biz.Adjudicator, error) {
	groups, err := r.data.db.Problem.Query().Where(problem.IDEQ(problemID)).QueryAdjudicators().All(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Adjudicator{
		Groups: groups,
	}, nil
}

func (r *adjudicatorRepo) FindProblemByID(ctx context.Context, problemID int64) (*ent.Problem, error) {
	v, err := r.data.db.Problem.Query().Where(problem.ID(problemID)).First(ctx)
	if err != nil {
		return nil, err
	}
	return v, nil
}
