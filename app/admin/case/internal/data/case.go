package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/admin/case/internal/biz"
	"sastoj/ent"
	"sastoj/ent/problem"
	"sastoj/ent/problemcase"
)

type caseRepo struct {
	data *Data
	log  *log.Helper
}

// NewProblemCaseRepo .
func NewProblemCaseRepo(data *Data, logger log.Logger) biz.CaseRepo {
	return &caseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *caseRepo) Save(ctx context.Context, pi int64, cs []*biz.Case) ([]int64, error) {
	var points []int32
	var indexes []int32
	var isAutos []bool
	for _, c := range cs {
		points = append(points, c.Point)
		indexes = append(indexes, c.Index)
		isAutos = append(isAutos, c.IsAuto)
	}
	rcs, err := r.data.db.ProblemCase.MapCreateBulk(points, func(c *ent.ProblemCaseCreate, i int) {
		c.SetPoint(int16(points[i])).SetIndex(int16(indexes[i])).SetIsAuto(isAutos[i]).SetProblemsID(pi)
	}).Save(ctx)
	if err != nil {
		return nil, err
	}
	var ids []int64
	for _, rc := range rcs {
		ids = append(ids, int64(rc.ID))
	}
	return ids, nil
}

func (r *caseRepo) Update(ctx context.Context, c *biz.Case) error {
	_, err := r.data.db.ProblemCase.UpdateOneID(c.Id).SetPoint(int16(c.Point)).SetIndex(int16(c.Index)).SetIsAuto(c.IsAuto).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *caseRepo) DeleteByCaseIds(ctx context.Context, cis []int64) error {
	intSlice := make([]int64, len(cis))
	for i, ci := range cis {
		intSlice[i] = ci
	}
	_, err := r.data.db.ProblemCase.Update().Where(
		problemcase.IDIn(intSlice...)).SetIsDeleted(true).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *caseRepo) DeleteByProblemId(ctx context.Context, pi int64) error {
	_, err := r.data.db.ProblemCase.Update().Where(
		problemcase.HasProblemsWith(problem.ID(pi))).SetIsDeleted(true).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *caseRepo) FindByProblemId(ctx context.Context, pi int64) ([]*biz.Case, error) {
	problemcase, err := r.data.db.ProblemCase.Query().Where(problemcase.HasProblemsWith(problem.ID(pi))).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Case, 0)
	for _, p := range problemcase {
		rv = append(rv, &biz.Case{
			Id:     p.ID,
			Point:  int32(p.Point),
			Index:  int32(p.Index),
			IsAuto: p.IsAuto,
		})
	}
	return rv, nil
}
