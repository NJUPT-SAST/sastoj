package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/admin/group/internal/biz"
	"sastoj/ent/group"
)

type groupRepo struct {
	data *Data
	log  *log.Helper
}

// NewGroupRepo .
func NewGroupRepo(data *Data, logger log.Logger) biz.GroupRepo {
	return &groupRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *groupRepo) Save(ctx context.Context, g *biz.Group) (*biz.Group, error) {
	po, err := r.data.db.Group.
		Create().
		SetGroupName(g.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Group{Id: int64(po.ID), Name: po.GroupName}, nil
}

func (r *groupRepo) Update(ctx context.Context, g *biz.Group) (*biz.Group, error) {
	return g, nil
}

func (r *groupRepo) FindByID(ctx context.Context, id int64) (*biz.Group, error) {
	group, err := r.data.db.Group.Query().Where(group.IDEQ(int(id))).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Group{
		Id:   int64(group.ID),
		Name: group.GroupName,
	}, nil
}

func (r *groupRepo) ListByPage(context.Context, string) ([]*biz.Group, error) {
	return nil, nil
}

func (r *groupRepo) DeleteById(context.Context, int64) error {
	return nil
}
