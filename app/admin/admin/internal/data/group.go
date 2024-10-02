package data

import (
	"context"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent/group"

	"github.com/go-kratos/kratos/v2/log"
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
	r.log.WithContext(ctx).Infof("Save Group: %v", g.Name)
	return &biz.Group{ID: po.ID, Name: po.GroupName}, nil
}

func (r *groupRepo) Update(ctx context.Context, g *biz.Group) (*int64, error) {
	po, err := r.data.db.Group.Update().
		SetGroupName(g.Name).
		Where(group.IDEQ(g.ID)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	res64 := int64(po)
	return &res64, nil
}

func (r *groupRepo) FindByID(ctx context.Context, id int64) (*biz.Group, error) {
	po, err := r.data.db.Group.Query().Where(group.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Group{
		ID:   po.ID,
		Name: po.GroupName,
	}, nil
}

func (r *groupRepo) ListByPage(ctx context.Context, current int64, size int64) ([]*biz.Group, error) {
	po, err := r.data.db.Group.Query().Limit(int(size)).Offset(int((current - 1) * size)).All(ctx)
	if err != nil {
		return nil, err
	}
	groups := make([]*biz.Group, len(po))
	for _, g := range po {
		groups = append(groups, &biz.Group{
			ID:   g.ID,
			Name: g.GroupName,
		})
	}
	return groups, nil
}

func (r *groupRepo) DeleteByID(ctx context.Context, id int64) error {
	err := r.data.db.Group.DeleteOneID(id).Exec(ctx)
	return err
}
