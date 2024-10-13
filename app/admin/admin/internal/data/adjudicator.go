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

func (r *adjudicatorRepo) Save(ctx context.Context, problemID int64, groupIds []int64) error {
	return r.data.db.Problem.Update().Where(problem.IDEQ(problemID)).AddAdjudicatorIDs(groupIds...).Exec(ctx)
}

func (r *adjudicatorRepo) Update(ctx context.Context, problemID int64, groupIds []int64) error {
	return r.data.db.Problem.Update().Where(problem.IDEQ(problemID)).ClearAdjudicators().AddAdjudicatorIDs(groupIds...).Exec(ctx)
}

func (r *adjudicatorRepo) FindByID(ctx context.Context, problemId int64) (*biz.Adjudicator, error) {
	groups, err := r.data.db.Problem.Query().Where(problem.IDEQ(problemId)).QueryAdjudicators().All(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Adjudicator{
		Groups: groups,
	}, nil
}

func (r *adjudicatorRepo) FindProblemById(ctx context.Context, problemId int64) (*ent.Problem, error) {
	v, err := r.data.db.Problem.Query().Where(problem.ID(problemId)).First(ctx)
	if err != nil {
		return nil, err
	}
	return v, nil
}
