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
	r.log.WithContext(ctx).Infof("Save Group: %v", g.Name)
	return &biz.Group{Id: po.ID, Name: po.GroupName}, nil
}

func (r *groupRepo) Update(ctx context.Context, g *biz.Group) (*int64, error) {
	res, err := r.data.db.Group.Update().
		SetGroupName(g.Name).
		Where(group.IDEQ(g.Id)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	res64 := int64(res)
	return &res64, nil
}

func (r *groupRepo) FindByID(ctx context.Context, id int64) (*biz.Group, error) {
	group, err := r.data.db.Group.Query().Where(group.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Group{
		Id:   group.ID,
		Name: group.GroupName,
	}, nil
}

func (r *groupRepo) ListByPage(ctx context.Context, current int64, size int64) ([]*biz.Group, error) {
	res, err := r.data.db.Group.Query().Limit(int(size)).Offset(int((current - 1) * size)).All(ctx)
	if err != nil {
		return nil, err
	}
	var groups []*biz.Group
	for _, group := range res {
		groups = append(groups, &biz.Group{
			Id:   group.ID,
			Name: group.GroupName,
		})
	}
	return groups, nil
}

func (r *groupRepo) DeleteById(ctx context.Context, id int64) error {
	err := r.data.db.Group.DeleteOneID(id).Exec(ctx)
	return err
}
