package data

import (
	"context"
	"sastoj/ent/problem"

	"sastoj/app/admin/problem/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type ProblemRepo struct {
	data *Data
	log  *log.Helper
}

// NewProblemRepo .
func NewProblemRepo(data *Data, logger log.Logger) biz.ProblemRepo {
	return &ProblemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ProblemRepo) Save(ctx context.Context, g *biz.Problem) (*biz.Problem, error) {
	res, err := r.data.db.Problem.Create().
		SetTitle(g.Title).
		SetContent(g.Content).
		SetPoint(int16(g.Point)).
		SetContestID(g.ContestId).
		SetCaseVersion(int16(g.CaseVersion)).
		SetIndex(int16(g.Index)).
		SetConfig(g.Config).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	g.Id = res.ID
	return g, nil
}

func (r *ProblemRepo) Update(ctx context.Context, g *biz.Problem) (*int, error) {
	res, err := r.data.db.Problem.Update().
		SetTitle(g.Title).
		SetContent(g.Content).
		SetPoint(int16(g.Point)).
		SetContestID(g.ContestId).
		SetCaseVersion(int16(g.CaseVersion)).
		SetIndex(int16(g.Index)).
		SetConfig(g.Config).
		Where(problem.ID(g.Id)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *ProblemRepo) FindByID(ctx context.Context, id int64) (*biz.Problem, error) {
	v, err := r.data.db.Problem.Query().Where(problem.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Problem{
		Id:          v.ID,
		Title:       v.Title,
		Content:     v.Content,
		Point:       int32(v.Point),
		ContestId:   v.ContestID,
		CaseVersion: int32(v.CaseVersion),
		Index:       int32(v.Index),
		Config:      v.Config,
	}, nil
}

func (r *ProblemRepo) Delete(ctx context.Context, id int64) (*int, error) {
	res, err := r.data.db.Problem.Delete().Where(problem.ID(id)).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *ProblemRepo) ListPages(ctx context.Context, currency int32, size int32) ([]*biz.Problem, error) {
	res, err := r.data.db.Problem.Query().Limit(int(size)).Offset(int((currency - 1) * size)).All(ctx)
	if err != nil {
		return nil, err
	}
	list := make([]*biz.Problem, 0)
	for _, v := range res {
		list = append(list, &biz.Problem{
			Id:          v.ID,
			Title:       v.Title,
			Content:     v.Content,
			Point:       int32(v.Point),
			ContestId:   v.ContestID,
			CaseVersion: int32(v.CaseVersion),
			Index:       int32(v.Index),
			Config:      v.Config,
		})
	}
	return list, nil
}
