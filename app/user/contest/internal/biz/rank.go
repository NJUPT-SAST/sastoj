package biz

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Rank struct {
	ContestId int64
	UserRank  []*UserRank
}

type UserRank struct {
	UserId       int64
	UserName     string
	Rank         int32
	Penalty      int32
	TotalPoint   int32
	AcceptCount  int32
	AchievedTime time.Time
	Problems     map[int64]*UserProblemResult
}

type UserProblemResult struct {
	SubmissionId int64
	ProblemId    int64
	State        int32
	Point        int32
	TriedCount   int32
	SubmitTime   time.Time
}

type RankRepo interface {
	Find(ctx context.Context, contest *Contest) (*Rank, error)
}

type RankUsecase struct {
	repo RankRepo
	log  *log.Helper
}

func NewRankUsecase(repo RankRepo, logger log.Logger) *RankUsecase {
	return &RankUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (r *RankUsecase) Find(ctx context.Context, contest *Contest) (*Rank, error) {
	if contest.EndTime.After(time.Now()) {
		return r.repo.Find(ctx, contest)
	} else {
		return nil, errors.New("contest with specified Id not found")
	}
}