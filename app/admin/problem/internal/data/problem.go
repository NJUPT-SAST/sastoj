package data

import (
	"context"
	problem2 "sastoj/api/sastoj/admin/problem/service/v1"
	"sastoj/ent/problem"

	"github.com/go-kratos/kratos/v2/log"
)

type ProblemRepo struct {
	data *Data
	log  *log.Helper
}

// NewProblemRepo problem
func NewProblemRepo(data *Data, logger log.Logger) *ProblemRepo {
	return &ProblemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ProblemRepo) Save(ctx context.Context, g *problem2.CreateProblemRequest) (*problem2.CreateProblemReply, error) {
	res, err := r.data.db.Problem.Create().
		SetTitle(g.Title).
		SetContent(g.Content).
		SetPoint(int16(g.Point)).
		SetContestsID(g.ContestId).
		SetCaseVersion(int16(g.CaseVersion)).
		SetIndex(int16(g.Index)).
		SetConfig(g.Config).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &problem2.CreateProblemReply{Id: res.ID}, nil
}

func (r *ProblemRepo) Update(ctx context.Context, g *problem2.UpdateProblemRequest) (*int, error) {
	res, err := r.data.db.Problem.Update().
		SetTitle(g.Title).
		SetContent(g.Content).
		SetPoint(int16(g.Point)).
		SetContestsID(g.ContestId).
		SetCaseVersion(int16(g.CaseVersion)).
		SetIndex(int16(g.Index)).
		SetConfig(g.Config).
		Where(problem.ID(g.Id)).
		Where(problem.IsDeleted(false)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *ProblemRepo) FindByID(ctx context.Context, id int64) (*problem2.GetProblemReply, error) {
	v, err := r.data.db.Problem.Query().
		Where(problem.ID(id)).
		Where(problem.IsDeleted(false)).
		First(ctx) //return nil while delete
	if err != nil {
		return nil, err
	}
	return &problem2.GetProblemReply{
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
	res, err := r.data.db.Problem.Update().
		SetIsDeleted(true).
		Where(problem.ID(id)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *ProblemRepo) ListPages(ctx context.Context, currency int32, size int32) ([]*problem2.ListProblemReply_Problem, error) {
	res, err := r.data.db.Problem.Query().Limit(int(size)).Offset(int((currency - 1) * size)).All(ctx)
	if err != nil {
		return nil, err
	}
	list := make([]*problem2.ListProblemReply_Problem, 0)
	for _, v := range res {
		list = append(list, &problem2.ListProblemReply_Problem{
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
