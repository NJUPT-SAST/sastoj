package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

const ContestTypeIOI = 1
const ContestTypeACM = 2

type Contest struct {
	ID          int64
	Title       string
	Description string
	Status      int32
	Type        int32
	StartTime   time.Time
	EndTime     time.Time
	Language    string
	ExtraTime   int32
	CreateTime  time.Time
}

type ContestRepo interface {
	Save(context.Context, *Contest) (*Contest, error)
	Update(context.Context, *Contest) error
	Delete(context.Context, int64) error
	FindByID(context.Context, int64) (*Contest, error)
	ListPages(ctx context.Context, current int64, size int64) ([]*Contest, error)
	AddContestants(ctx context.Context, contestID int64, groupID int64, role int32) error
	GetRacingContests(ctx context.Context) ([]*Contest, error)
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
func (uc *ContestUsecase) UpdateContest(ctx context.Context, g *Contest) error {
	return uc.repo.Update(ctx, g)
}
func (uc *ContestUsecase) DeleteContest(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
func (uc *ContestUsecase) FindContest(ctx context.Context, id int64) (*Contest, error) {
	return uc.repo.FindByID(ctx, id)
}
func (uc *ContestUsecase) ListContest(ctx context.Context, current int64, size int64) ([]*Contest, error) {
	return uc.repo.ListPages(ctx, current, size)
}
func (uc *ContestUsecase) AddContestants(ctx context.Context, contestID int64, groupID int64, role int32) error {
	_, err := uc.repo.FindByID(ctx, contestID)
	if err != nil {
		return err
	}

	return uc.repo.AddContestants(ctx, contestID, groupID, role)
}
func (uc *ContestUsecase) GetRacingContests(ctx context.Context) ([]*Contest, error) {
	return uc.repo.GetRacingContests(ctx)
}
