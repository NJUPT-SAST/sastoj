package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Group struct {
	Id   int64
	Name string
}

type GroupRepo interface {
	Save(context.Context, *Group) (*Group, error)
	Update(context.Context, *Group) (*Group, error)
	FindByID(context.Context, int64) (*Group, error)
	ListByPage(context.Context, string) ([]*Group, error)
	DeleteById(context.Context, int64) error
}

type GroupUsecase struct {
	repo GroupRepo
	log  *log.Helper
}

func NewGroupUsecase(repo GroupRepo, logger log.Logger) *GroupUsecase {
	return &GroupUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GroupUsecase) CreateGroup(ctx context.Context, g *Group) (*Group, error) {
	uc.log.WithContext(ctx).Infof("CreateGroup: %v", g.Name)
	out, err := uc.repo.Save(ctx, g)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (uc *GroupUsecase) GetGroup(ctx context.Context, g *Group) (*Group, error) {
	uc.log.WithContext(ctx).Infof("GetGroup: %v", g.Id)
	out, err := uc.repo.FindByID(ctx, g.Id)
	if err != nil {
		return nil, err
	}
	return out, nil
}
