package data

import (
	"context"
	"sastoj/app/user/gateway/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type contestRepo struct {
	data *Data
	log  *log.Helper
}

func (c *contestRepo) GetContests(ctx context.Context, groupID int64) ([]*biz.Contest, error) {
	return c.data.cache.contestantsMap[groupID], nil
}

func (c *contestRepo) JoinContest(ctx context.Context, userID, contestID int64, isJoin bool) error {
	// TODO: set contest join state
	return nil
}

// NewContestRepo .
func NewContestRepo(data *Data, logger log.Logger) biz.ContestRepo {
	return &contestRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/contest")),
	}
}
