package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/user/gateway/internal/biz"
	"strconv"
)

type contestRepo struct {
	data *Data
	log  *log.Helper
}

func (c *contestRepo) GetContests(ctx context.Context, groupID int64) ([]*biz.Contest, error) {
	po := c.data.redis.Get(ctx, "group:"+strconv.FormatInt(groupID, 10))
	if po.Err() != nil {
		return nil, po.Err()
	}
	contestsJson := po.Val()
	var contests []*biz.Contest
	err := json.Unmarshal([]byte(contestsJson), &contests)
	if err != nil {
		return nil, err
	}
	return contests, nil
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
