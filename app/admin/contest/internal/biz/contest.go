package biz

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/go-kratos/kratos/v2/log"
)

type Contest struct {
	Id          int64
	Title       string
	Description string
	Status      int32
	Type        int32
	StartTime   timestamp.Timestamp
	EndTime     timestamp.Timestamp
	Language    string
	ExtraTime   int32
	CreateTime  timestamp.Timestamp
}

type ContestRepo interface {
	Save(context.Context, *Contest) (*Contest, error)
	Update(context.Context, *Contest) (*Contest, error)
	FindByID(context.Context, int64) (*Contest, error)
	ListByHello(context.Context, string) ([]*Contest, error)
	ListAll(context.Context) ([]*Contest, error)
}

type ContestUsecase struct {
	repo ContestRepo
	log  *log.Helper
}

func NewContestUsecase(repo ContestRepo, logger log.Logger) *ContestUsecase {
	return &ContestUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ContestUsecase) CreateContest(ctx context.Context, g *Contest) (*Contest, error) {
	return uc.repo.Save(ctx, g)
}
