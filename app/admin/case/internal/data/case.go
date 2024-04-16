package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/admin/case/internal/biz"
	"sastoj/ent"
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
	var points []int64
	var indexes []int64
	var isAutos []bool
	for _, c := range cs {
		points = append(points, c.Point)
		indexes = append(indexes, c.Index)
		isAutos = append(isAutos, c.IsAuto)
	}
	rcs, err := r.data.db.ProblemCase.MapCreateBulk(points, func(c *ent.ProblemCaseCreate, i int) {
		c.SetPoint(int(points[i])).SetIndex(int(indexes[i])).SetIsAuto(isAutos[i]).SetProblemID(int(pi))
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
	_, err := r.data.db.ProblemCase.UpdateOneID(int(c.Id)).SetPoint(int(c.Point)).SetIndex(int(c.Index)).SetIsAuto(c.IsAuto).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *caseRepo) DeleteByCaseIds(ctx context.Context, cis []int64) error {
	intSlice := make([]int, len(cis))
	for i, ci := range cis {
		intSlice[i] = int(ci)
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
		problemcase.ProblemIDEQ(int(pi))).SetIsDeleted(true).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *caseRepo) FindByProblemId(ctx context.Context, pi int64) ([]*biz.Case, error) {
	problemcase, err := r.data.db.ProblemCase.Query().Where(problemcase.ProblemIDEQ(int(pi))).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Case, 0)
	for _, p := range problemcase {
		rv = append(rv, &biz.Case{
			Id:     int64(p.ID),
			Point:  int64(p.Point),
			Index:  int64(p.Index),
			IsAuto: p.IsAuto,
		})
	}
	return rv, nil
}
