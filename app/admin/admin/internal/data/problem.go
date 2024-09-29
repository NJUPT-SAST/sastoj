package data

import (
	"context"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent/contest"
	"sastoj/ent/problem"
	"sastoj/pkg/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type problemRepo struct {
	data *Data
	log  *log.Helper
}

// NewProblemRepo problem
func NewProblemRepo(data *Data, logger log.Logger) biz.ProblemRepo {
	return &problemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *problemRepo) Save(ctx context.Context, g *biz.Problem) (*int64, error) {
	res, err := r.data.db.Problem.Create().
		SetProblemTypeID(g.TypeId).
		SetTitle(g.Title).
		SetContent(g.Content).
		SetScore(int16(g.Point)).
		SetContestID(g.ContestId).
		SetCaseVersion(1).
		SetIndex(int16(g.Index)).
		SetOwnerID(getUserID(ctx)).
		SetVisibility(util.VisToEnt(g.Visibility)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &res.ID, nil
}

func (r *problemRepo) Update(ctx context.Context, g *biz.Problem) (*int64, error) {
	res, err := r.data.db.Problem.Update().
		SetProblemTypeID(g.TypeId).
		SetTitle(g.Title).
		SetContent(g.Content).
		SetScore(int16(g.Point)).
		SetContestID(g.ContestId).
		AddCaseVersion(1).
		SetIndex(int16(g.Index)).
		SetOwnerID(getUserID(ctx)).
		SetVisibility(util.VisToEnt(g.Visibility)).
		Where(problem.ID(g.Id)).
		Where(problem.IsDeleted(false)).
		Where(problem.HasContestWith(
			contest.IDEQ(g.ContestId), contest.StartTimeGT(time.Now()))).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	res64 := int64(res)
	return &res64, nil
}

func (r *problemRepo) FindByID(ctx context.Context, id int64) (*biz.Problem, error) {
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
	vis := util.VisToPb(p.Visibility)
	return &biz.Problem{
		Id:          p.ID,
		Title:       p.Title,
		Content:     p.Content,
		Point:       int32(p.Score),
		ContestId:   p.ContestID,
		CaseVersion: int32(p.CaseVersion),
		Index:       int32(p.Index),
		OwnerId:     owner.ID,
		Visibility:  vis,
	}, nil
}

func (r *problemRepo) Delete(ctx context.Context, id int64) (*int64, error) {
	res, err := r.data.db.Problem.Update().
		SetIsDeleted(true).
		Where(problem.ID(id)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	res64 := int64(res)
	return &res64, nil
}

func (r *problemRepo) ListPages(ctx context.Context, currency int32, size int32) ([]*biz.Problem, error) {
	res, err := r.data.db.Problem.Query().
		Limit(int(size)).Offset(int((currency - 1) * size)).
		WithOwner().
		All(ctx)
	if err != nil {
		return nil, err
	}
	list := make([]*biz.Problem, 0)
	for _, v := range res {
		vis := util.VisToPb(v.Visibility)
		owner, err := v.QueryOwner().First(ctx)
		if err != nil {
			return nil, err
		}
		list = append(list, &biz.Problem{
			Id:          v.ID,
			TypeId:      v.ProblemTypeID,
			Title:       v.Title,
			Content:     v.Content,
			Point:       int32(v.Score),
			ContestId:   v.ContestID,
			CaseVersion: int32(v.CaseVersion),
			Index:       int32(v.Index),
			OwnerId:     owner.ID,
			Visibility:  vis,
		})
	}
	return list, nil
}

func getUserID(ctx context.Context) int64 {
	return util.GetUserInfoFromCtx(ctx).UserId
}
