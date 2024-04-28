package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/user/contest/internal/biz"
	"sastoj/ent"
	"sastoj/ent/contest"
	"sastoj/ent/group"
	"sastoj/ent/user"
	"strconv"
	"time"
)

type contestRepo struct {
	data *Data
	log  *log.Helper
}

func (c *contestRepo) ListContest(ctx context.Context, userID int64) ([]*biz.Contest, error) {
	var (
		po  []*ent.Contest
		err error
	)
	if userID != 0 {
		po, err = c.data.db.Contest.
			Query().
			Where(
				contest.HasContestWith(
					group.HasUsersWith(
						user.IDEQ(userID),
					),
				),
			).
			All(ctx)
	} else {
		po, err = c.data.db.Contest.
			Query().
			Where(
				contest.StartTimeGT(time.Now()),
				contest.StartTimeLT(time.Now().Add(24*time.Hour)),
			).
			All(ctx)

	}
	if err != nil {
		return nil, err
	}
	var ret []*biz.Contest
	for _, v := range po {
		var groups []int64
		for _, g := range v.Edges.Contest {
			groups = append(groups, g.ID)
		}
		ret = append(ret, &biz.Contest{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Status:      v.Status,
			Type:        v.Type,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
			Language:    v.Language,
			ExtraTime:   v.ExtraTime,
			CreateTime:  v.CreateTime,
			Groups:      groups,
		})
	}
	return ret, nil
}

func (c *contestRepo) JoinContest(ctx context.Context, userID, contestID int64, isJoin bool) error {
	const prefix = "user:contest:"
	if !isJoin {
		return c.data.redis.Del(ctx, prefix+strconv.FormatInt(userID, 10)).Err()
	}
	joinContest, err := c.data.db.Contest.Get(ctx, contestID)
	if err != nil {
		return err
	}
	expireAt := joinContest.EndTime
	err = c.data.redis.SetNX(ctx, prefix+strconv.FormatInt(userID, 10), isJoin, 2*time.Hour).Err()
	if err != nil {
		return err
	}
	err = c.data.redis.ExpireAt(ctx, prefix+strconv.FormatInt(userID, 10), expireAt).Err()
	if err != nil {
		return err
	}
	return nil
}

// NewContestRepo .
func NewContestRepo(data *Data, logger log.Logger) biz.ContestRepo {
	return &contestRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
