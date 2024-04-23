package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/user/gateway/internal/biz"
	"strconv"
	"time"
)

type contestRepo struct {
	data *Data
	log  *log.Helper
}

func (c *contestRepo) GetContests(ctx context.Context, groupID int64) ([]*biz.Contest, error) {
	po := c.data.redis.Get(ctx, "po:"+strconv.FormatInt(groupID, 10))
	if po.Err() != nil {
		return nil, po.Err()
	}
	var ret []*biz.Contest
	for _, _ = range po.Val() {
		ret = append(ret, &biz.Contest{
			// TODO: get contest from redis
			ID:          0,
			Title:       "",
			Description: "",
			Status:      0,
			Type:        0,
			StartTime:   time.Time{},
			EndTime:     time.Time{},
			Language:    "",
			ExtraTime:   0,
			CreateTime:  time.Time{},
		})
	}
	return ret, nil
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
