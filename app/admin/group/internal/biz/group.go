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
	Update(context.Context, *Group) (*int64, error)
	FindByID(context.Context, int64) (*Group, error)
	ListByPage(ctx context.Context, current int64, size int64) ([]*Group, error)
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

func (uc *GroupUsecase) UpdateGroup(ctx context.Context, g *Group) (bool, error) {
	uc.log.WithContext(ctx).Infof("UpdateGroup: %v", g.Id)
	rv, err := uc.repo.Update(ctx, g)
	if err != nil || *rv == 0 {
		return false, err
	}
	return true, nil
}

func (uc *GroupUsecase) ListGroup(ctx context.Context, current int64, size int64) ([]*Group, error) {
	uc.log.WithContext(ctx).Infof("ListGroup: %v", current)
	out, err := uc.repo.ListByPage(ctx, current, size)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (uc *GroupUsecase) DeleteGroup(ctx context.Context, id int64) (bool, error) {
	uc.log.WithContext(ctx).Infof("GetGroup: %v", id)
	err := uc.repo.DeleteById(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
