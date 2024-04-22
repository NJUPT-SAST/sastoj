package data

import (
	"context"
	"sastoj/ent/contest"

	"sastoj/app/admin/contest/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type ContestRepo struct {
	data *Data
	log  *log.Helper
}

func NewContestRepo(data *Data, logger log.Logger) biz.ContestRepo {
	return &ContestRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ContestRepo) Save(ctx context.Context, g *biz.Contest) (*biz.Contest, error) {
	res, err := r.data.db.Contest.Create().
		SetTitle(g.Title).
		SetDescription(g.Description).
		SetStatus(int16(g.Status)).
		SetType(int16(g.Type)).
		SetStartTime(g.StartTime).
		SetEndTime(g.EndTime).
		SetLanguage(g.Language).
		SetExtraTime(int16(g.ExtraTime)).
		SetCreateTime(g.CreateTime).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	g.Id = res.ID
	return g, nil
}

func (r *ContestRepo) Update(ctx context.Context, g *biz.Contest) error {
	_, err := r.data.db.Contest.Update().
		SetTitle(g.Title).
		SetDescription(g.Description).
		SetStatus(int16(g.Status)).
		SetType(int16(g.Type)).
		SetStartTime(g.StartTime).
		SetEndTime(g.EndTime).
		SetLanguage(g.Language).
		SetExtraTime(int16(g.ExtraTime)).
		SetCreateTime(g.CreateTime).
		Where(contest.ID(g.Id)).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContestRepo) FindByID(ctx context.Context, id int64) (*biz.Contest, error) {
	po, err := r.data.db.Contest.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Contest{
		Title:       po.Title,
		Description: po.Description,
		Status:      int32(po.Status),
		Type:        int32(po.Type),
		StartTime:   po.StartTime,
		EndTime:     po.EndTime,
		Language:    po.Language,
		ExtraTime:   int32(po.ExtraTime),
		CreateTime:  po.CreateTime,
	}, nil
}
func (r *ContestRepo) Delete(ctx context.Context, id int64) error {
	err := r.data.db.Contest.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContestRepo) ListPages(ctx context.Context, current int64, size int64) ([]*biz.Contest, error) {
	res, err := r.data.db.Contest.Query().Offset(int((current - 1) * size)).Limit(int(size)).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Contest, 0)
	for _, po := range res {
		rv = append(rv, &biz.Contest{
			Title:       po.Title,
			Description: po.Description,
			Status:      int32(po.Status),
			Type:        int32(po.Type),
			StartTime:   po.StartTime,
			EndTime:     po.EndTime,
			Language:    po.Language,
			ExtraTime:   int32(po.ExtraTime),
			CreateTime:  po.CreateTime,
		})
	}
	return rv, nil
}