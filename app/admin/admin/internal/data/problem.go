package data

import (
	"context"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent"
	"sastoj/ent/contest"
	"sastoj/ent/problem"
	"sastoj/ent/problemtype"
	"sastoj/pkg/mq"
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
		SetScore(int16(g.Point)).
		SetContestID(g.ContestID).
		SetCaseVersion(1).
		SetIndex(int16(g.Index)).
		SetOwnerID(getUserID(ctx)).
		SetVisibility(util.VisToEnt(g.Visibility)).
		SetMetadata(g.Metadata).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	err = r.data.jcm.WriteString(res.ID, g.Config)
	if err != nil {
		return nil, err
	}
	return &res.ID, nil
}

func (r *problemRepo) Update(ctx context.Context, g *biz.Problem) error {
	// not allow to update contestID, ownerID, caseVersion and problemTypeID
	_, err := r.data.db.Problem.Update().
		Where(problem.IDEQ(g.ID)).
		// SetProblemTypeID(g.TypeID).
		SetTitle(g.Title).
		SetContent(g.Content).
		// SetContestID(g.ContestID).
		SetScore(int16(g.Point)).
		AddCaseVersion(1).
		SetIndex(int16(g.Index)).
		// SetOwnerID(getUserID(ctx)).
		SetVisibility(util.VisToEnt(g.Visibility)).
		SetMetadata(g.Metadata).
		Where(problem.IsDeleted(false)).
		Save(ctx)
	if err != nil {
		return err
	}

	if g.Config != "" {
		err = r.data.jcm.WriteString(g.ID, g.Config)
		if err != nil {
			return err
		}
	}

	po, err := r.data.db.Problem.Query().
		Where(problem.IDEQ(g.ID)).
		WithContest(func(q *ent.ContestQuery) {
			q.Where(contest.StateEQ(contest.StateNORMAL),
				contest.StartTimeLT(time.Now()),
				contest.EndTimeGT(time.Now()),
			)
		}).
		WithProblemType(func(q *ent.ProblemTypeQuery) {
			q.Select(problemtype.FieldDisplayName)
		}).
		Only(ctx)
	if err == nil {
		err = r.data.ex.Publish(ctx, mq.Problem{
			ID:       po.ID,
			Title:    po.Title,
			Type:     po.Edges.ProblemType.DisplayName,
			Content:  po.Content,
			Score:    po.Score,
			Index:    po.Index,
			Metadata: po.Metadata,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *problemRepo) FindByID(ctx context.Context, id int64) (*biz.Problem, error) {
	p, err := r.data.db.Problem.Query().
		Where(problem.ID(id)).
		Where(problem.IsDeleted(false)).
		WithOwner().
		WithProblemType().
		First(ctx)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	vis := util.VisToPb(p.Visibility)
	config, err := r.data.jcm.ReadString(p.ID)
	if err != nil {
		return nil, err
	}

	return &biz.Problem{
		ID:          p.ID,
		TypeID:      p.ProblemTypeID,
		Title:       p.Title,
		Content:     p.Content,
		Point:       int32(p.Score),
		ContestID:   p.ContestID,
		CaseVersion: int32(p.CaseVersion),
		Index:       int32(p.Index),
		OwnerID:     p.Edges.Owner.ID,
		Visibility:  vis,
		Config:      config,
		Metadata:    p.Metadata,
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

func (r *problemRepo) ListPages(ctx context.Context, current int32, size int32, contestID int64) ([]*biz.Problem, int64, error) {
	query := r.data.db.Problem.Query().Where(problem.ContestIDEQ(contestID))
	res, err := query.
		Limit(int(size)).Offset(int((current - 1) * size)).
		WithOwner().
		WithProblemType().
		All(ctx)
	if err != nil {
		return nil, 0, err
	}
	total, err := query.Count(ctx)
	list := make([]*biz.Problem, 0)
	for _, v := range res {
		owner := v.Edges.Owner
		vis := util.VisToPb(v.Visibility)
		config, err := r.data.jcm.ReadString(v.ID)
		if err != nil {
			return nil, 0, err
		}

		list = append(list, &biz.Problem{
			ID:          v.ID,
			TypeID:      v.ProblemTypeID,
			Title:       v.Title,
			Content:     v.Content,
			Point:       int32(v.Score),
			ContestID:   v.ContestID,
			CaseVersion: int32(v.CaseVersion),
			Index:       int32(v.Index),
			OwnerID:     owner.ID,
			Visibility:  vis,
			Config:      config,
			Metadata:    v.Metadata,
		})
	}
	return list, int64(total), nil
}

func (r *problemRepo) GetTypes(ctx context.Context) ([]*biz.ProblemType, error) {
	res, err := r.data.db.ProblemType.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	types := make([]*biz.ProblemType, 0)
	for _, v := range res {
		types = append(types, &biz.ProblemType{
			ID:          v.ID,
			SlugName:    v.SlugName,
			DisplayName: v.DisplayName,
			Description: v.Description,
			Judge:       v.Judge,
		})
	}
	return types, nil
}

func getUserID(ctx context.Context) int64 {
	return util.GetUserInfoFromCtx(ctx).UserId
}
