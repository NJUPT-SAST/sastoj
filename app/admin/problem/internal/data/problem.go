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
		SetScore(int16(g.Point)).
		SetContestsID(g.ContestId).
		SetCaseVersion(int16(g.CaseVersion)).
		SetIndex(int16(g.Index)).
		SetOwnerID(g.OwnerId).
		SetVisibility(problem.Visibility(g.Visibility)).
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
		SetScore(int16(g.Point)).
		SetContestsID(g.ContestId).
		SetCaseVersion(int16(g.CaseVersion)).
		SetIndex(int16(g.Index)).
		SetOwnerID(g.OwnerId).
		SetVisibility(problem.Visibility(g.Visibility)).
		Where(problem.ID(g.Id)).
		Where(problem.IsDeleted(false)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *ProblemRepo) FindByID(ctx context.Context, id int64) (*problem2.GetProblemReply, error) {
	p, err := r.data.db.Problem.Query().
		Where(problem.ID(id)).
		Where(problem.IsDeleted(false)).
		WithOwner().
		First(ctx)
	if err != nil {
		return nil, err
	}
	owner, err := p.QueryOwner().First(ctx)
	if err != nil {
		return nil, err
	}
	var vis problem2.Visibility
	switch n := p.Visibility; n {
	case problem.VisibilityPRIVATE:
		vis = problem2.Visibility_Private
	case problem.VisibilityPUBLIC:
		vis = problem2.Visibility_Public
	case problem.VisibilityCONTEST:
		vis = problem2.Visibility_Contest
	}
	return &problem2.GetProblemReply{
		Id:          p.ID,
		Title:       p.Title,
		Content:     p.Content,
		Point:       int32(p.Score),
		ContestId:   p.ContestID,
		CaseVersion: int32(p.CaseVersion),
		Index:       int32(p.Index),
		OwnerId:     owner.ID,
		Visibility:  vis,
		Config:      "",
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
	res, err := r.data.db.Problem.Query().
		Limit(int(size)).Offset(int((currency - 1) * size)).
		WithOwner().
		All(ctx)
	if err != nil {
		return nil, err
	}
	list := make([]*problem2.ListProblemReply_Problem, 0)
	for _, v := range res {
		var vis problem2.Visibility
		switch n := v.Visibility; n {
		case problem.VisibilityPRIVATE:
			vis = problem2.Visibility_Private
		case problem.VisibilityPUBLIC:
			vis = problem2.Visibility_Public
		case problem.VisibilityCONTEST:
			vis = problem2.Visibility_Contest
		}
		owner, err := v.QueryOwner().First(ctx)
		if err != nil {
			return nil, err
		}
		list = append(list, &problem2.ListProblemReply_Problem{
			Id:          v.ID,
			Title:       v.Title,
			Content:     v.Content,
			Point:       int32(v.Score),
			ContestId:   v.ContestID,
			CaseVersion: int32(v.CaseVersion),
			Index:       int32(v.Index),
			OwnerId:     owner.ID,
			Visibility:  vis,
			Config:      "",
		})
	}
	return list, nil
}
