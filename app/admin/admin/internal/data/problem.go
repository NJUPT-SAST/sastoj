package data

import (
	"context"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent"
	"sastoj/ent/contest"
	"sastoj/ent/problem"
	"sastoj/ent/user"
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
		SetProblemTypeID(g.TypeID).
		SetTitle(g.Title).
		SetContent(g.Content).
		SetScore(g.Score).
		SetContestID(g.ContestID).
		SetCaseVersion(1).
		SetIndex(g.Index).
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
		SetProblemTypeID(g.TypeID).
		SetTitle(g.Title).
		SetContent(g.Content).
		SetScore(g.Score).
		SetContestID(g.ContestID).
		AddCaseVersion(1).
		SetIndex(g.Index).
		SetVisibility(util.VisToEnt(g.Visibility)).
		Where(problem.ID(g.ID)).
		Where(problem.IsDeleted(false)).
		Where(problem.HasContestWith(
			contest.IDEQ(g.ContestID), contest.StartTimeGT(time.Now()))).
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
		WithOwner(func(q *ent.UserQuery) {
			q.Select(user.FieldID)
		}).
		First(ctx)
	if err != nil {
		return nil, err
	}
	vis := util.VisToPb(p.Visibility)
	return &biz.Problem{
		ID:          p.ID,
		Title:       p.Title,
		Content:     p.Content,
		Score:       p.Score,
		ContestID:   p.ContestID,
		CaseVersion: p.CaseVersion,
		Index:       p.Index,
		OwnerID:     p.Edges.Owner.ID,
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
		WithOwner(func(q *ent.UserQuery) {
			q.Select(user.FieldID)
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}
	list := make([]*biz.Problem, 0)
	for _, v := range res {
		vis := util.VisToPb(v.Visibility)

		list = append(list, &biz.Problem{
			ID:          v.ID,
			TypeID:      v.ProblemTypeID,
			Title:       v.Title,
			Content:     v.Content,
			Score:       v.Score,
			ContestID:   v.ContestID,
			CaseVersion: v.CaseVersion,
			Index:       v.Index,
			OwnerID:     v.Edges.Owner.ID,
			Visibility:  vis,
		})
	}
	return list, nil
}

func getUserID(ctx context.Context) int64 {
	return util.GetUserInfoFromCtx(ctx).UserID
}
