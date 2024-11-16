package biz

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Contest is a Contest model.
type Contest struct {
	ID          int64
	Title       string
	Description string
	Status      int16
	Type        int16
	StartTime   time.Time
	EndTime     time.Time
	Language    string
	ExtraTime   int16
	CreateTime  time.Time
	Groups      []int64
}

// ContestRepo is a Contest repo.
type ContestRepo interface {
	ListContest(ctx context.Context, userID int64) ([]*Contest, error)
	JoinContest(ctx context.Context, userID, contestID int64, isJoin bool) error
	UserStateCache(ctx context.Context, userId int64) (string, error)
	UpdateUserStateCache(ctx context.Context, userId int64, state string) error
	UserState(ctx context.Context, userId int64) (string, error)
	IsUserBanned(state string) bool
}

// ContestUsecase is a Contest usecase.
type ContestUsecase struct {
	repo ContestRepo
	log  *log.Helper
}

// NewContestUsecase new a Contest usecase.
func NewContestUsecase(repo ContestRepo, logger log.Logger) *ContestUsecase {
	return &ContestUsecase{repo: repo, log: log.NewHelper(logger)}
}

// ListContest list all contests.
func (uc *ContestUsecase) ListContest(ctx context.Context, userID int64) ([]*Contest, error) {
	contests, err := uc.repo.ListContest(ctx, userID)
	if err != nil {
		return nil, err
	}
	return contests, nil
}

// JoinContest joins a contest.
func (uc *ContestUsecase) JoinContest(ctx context.Context, userID, contestID int64, isJoin bool) error {
	// log
	uc.log.Infof("userID: %d, contestID: %d, isJoin: %v", userID, contestID, isJoin)
	// redis set isJoin state
	return uc.repo.JoinContest(ctx, userID, contestID, isJoin)
}

func (uc *ContestUsecase) getUserStateWithCache(ctx context.Context, userId int64) (string, error) {
	state, err := uc.repo.UserStateCache(ctx, userId)
	if err == redis.Nil {
		state, err = uc.repo.UserState(ctx, userId)
		if err != nil {
			return "", err
		}
		err = uc.repo.UpdateUserStateCache(ctx, userId, state)
	}
	return state, err
}

func (uc *ContestUsecase) CheckBanned(ctx context.Context, userId int64) (bool, error) {
	state, err := uc.getUserStateWithCache(ctx, userId)
	if err != nil {
		return false, err
	}
	return uc.repo.IsUserBanned(state), nil
}
