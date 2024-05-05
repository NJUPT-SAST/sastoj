package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/user/contest/internal/biz"
)

type registerRepo struct {
	data *Data
	log  *log.Helper
}

func (r *registerRepo) RegisterGateway(ctx context.Context, register *biz.Register) error {
	// TODO: check secret
	return r.data.redis.Set(ctx, register.UUID, register.Endpoint, 0).Err()
}

func NewRegisterRepo(data *Data, logger log.Logger) biz.RegisterRepo {
	return &registerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
