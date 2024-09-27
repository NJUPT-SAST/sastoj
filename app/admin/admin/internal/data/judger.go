package data

import (
	"context"
	"sastoj/ent"
	"sastoj/ent/problem"

	"sastoj/app/admin/admin/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type judgerRepo struct {
	data *Data
	log  *log.Helper
}

// NewjudgerRepo .
func NewjudgerRepo(data *Data, logger log.Logger) biz.JudgerRepo {
	return &judgerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *judgerRepo) Save(ctx context.Context, groupIds []int64) error {
	return r.data.db.Problem.Update().AddJudgerIDs(groupIds...).Exec(ctx)
}

func (r *judgerRepo) Update(ctx context.Context, groupIds []int64) error {
	return r.data.db.Problem.Update().ClearJudgers().AddJudgerIDs(groupIds...).Exec(ctx)
}

func (r *judgerRepo) FindByID(ctx context.Context, problemId int64) (*biz.Judger, error) {
	groups, err := r.data.db.Problem.Query().Where(problem.IDEQ(problemId)).QueryJudgers().All(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Judger{
		Groups: groups,
	}, nil
}

func (r *judgerRepo) FindProblemById(ctx context.Context, problemId int64) (*ent.Problem, error) {
	v, err := r.data.db.Problem.Query().Where(problem.ID(problemId)).First(ctx)
	if err != nil {
		return nil, err
	}
	return v, nil
}
