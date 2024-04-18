package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/user/contest/internal/biz"
	"sastoj/ent/contest"
	"sastoj/ent/contestgroup"
	"strconv"
	"time"
)

type contestRepo struct {
	data *Data
	log  *log.Helper
}

func (c *contestRepo) ListContest(ctx context.Context, userID int64) ([]*biz.Contest, error) {
	uPo, err := c.data.db.User.
		Get(ctx, int(userID))
	if err != nil {
		return nil, err
	}
	gID := uPo.GroupID

	po, err := c.data.db.Contest.
		Query().
		Where(
			contest.HasContestGroupWith(
				contestgroup.ID(gID),
			),
		).
		All(ctx)

	if err != nil {
		return nil, err
	}
	var ret []*biz.Contest
	for _, v := range po {
		ret = append(ret, &biz.Contest{
			ID:          int64(v.ID),
			Title:       v.Title,
			Description: v.Description,
			State:       v.State,
			Type:        v.Type,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
			Language:    v.Language,
			ExtraTime:   v.ExtraTime,
			CreateTime:  v.CreateTime,
		})
	}
	return ret, nil
}

func (c *contestRepo) JoinContest(ctx context.Context, userID, contestID int64, isJoin bool) error {
	const prefix = "user:contest:"
	// TODO set dynamic expire time
	if !isJoin {
		return c.data.redis.Del(ctx, prefix+strconv.FormatInt(userID, 10)).Err()
	}
	//joinContest, err := c.data.db.Contest.Get(ctx, int(contestID))
	//if err != nil {
	//	return err
	//}
	//expireAt := joinContest.EndTime
	err := c.data.redis.SetNX(ctx, prefix+strconv.FormatInt(userID, 10), isJoin, 2*time.Hour).Err()
	if err != nil {
		return err
	}
	//err = c.data.redis.ExpireAt(ctx, prefix+strconv.FormatInt(userID, 10), expireAt).Err()
	//if err != nil {
	//	return err
	//}
	return nil
}

// NewContestRepo .
func NewContestRepo(data *Data, logger log.Logger) biz.ContestRepo {
	return &contestRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
