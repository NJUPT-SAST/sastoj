package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Register struct {
	UUID     string
	Endpoint string
	Secret   string
}

type RegisterRepo interface {
	RegisterGateway(ctx context.Context, register *Register) error
}

type RegisterUsecase struct {
	registerRepo RegisterRepo
	log          *log.Helper
}

func NewRegisterUsecase(repo RegisterRepo, logger log.Logger) *RegisterUsecase {
	return &RegisterUsecase{registerRepo: repo, log: log.NewHelper(logger)}
}

func (uc *RegisterUsecase) RegisterGateway(ctx context.Context, register *Register) error {
	return uc.registerRepo.RegisterGateway(ctx, register)
}
