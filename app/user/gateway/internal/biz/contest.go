package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
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
}

// ContestRepo is a Contest repo.
type ContestRepo interface {
	GetContests(ctx context.Context, userID int64) ([]*Contest, error)
	JoinContest(ctx context.Context, userID, contestID int64, isJoin bool) error
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

// GetContests list all contests.
func (uc *ContestUsecase) GetContests(ctx context.Context, userID int64) ([]*Contest, error) {
	uc.log.Infof("userID: %d", userID)
	contests, err := uc.repo.GetContests(ctx, userID)
	if err != nil {
		return nil, err
	}
	uc.log.Infof("contests: %v", contests)
	return contests, nil
}

// JoinContest joins a contest.
func (uc *ContestUsecase) JoinContest(ctx context.Context, userID, contestID int64, isJoin bool) error {
	// log
	uc.log.Infof("userID: %d, contestID: %d, isJoin: %v", userID, contestID, isJoin)
	// redis set isJoin state
	return uc.repo.JoinContest(ctx, userID, contestID, isJoin)
}
