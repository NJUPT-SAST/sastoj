package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/admin/user/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepo) FindByID(ctx context.Context, i int64) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepo) ListByHello(ctx context.Context, s string) ([]*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepo) ListAll(ctx context.Context) ([]*biz.User, error) {
	//TODO implement me
	panic("implement me")
}
