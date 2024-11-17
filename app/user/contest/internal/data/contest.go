package data

import (
	"context"
	"encoding/json"
	"sastoj/app/user/contest/internal/biz"
	"sastoj/ent"
	"sastoj/ent/contest"
	"sastoj/ent/group"
	"sastoj/ent/user"
	"sastoj/pkg/util"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type contestRepo struct {
	data *Data
	log  *log.Helper
}

const (
	userStatePrefix = "user:contest:state:"
	redisPrefix     = "user:contest:contest:"
)

func (c *contestRepo) ListContest(ctx context.Context, userID int64) ([]*biz.Contest, error) {
	var (
		po  []*ent.Contest
		err error
	)
	if userID != 0 {
		po, err = c.data.db.Contest.
			Query().
			Where(
				contest.HasContestantsWith(
					group.HasUsersWith(
						user.IDEQ(userID),
					),
				),
				contest.StateEQ(contest.StateNORMAL),
			).
			Order(ent.Desc(contest.FieldStartTime)).
			All(ctx)
	} else {
		po, err = c.data.db.Contest.
			Query().
			Where(
				contest.StartTimeLT(time.Now().Add(24*time.Hour)),
				contest.EndTimeGT(time.Now()),
				contest.StateEQ(contest.StateNORMAL),
			).
			All(ctx)
	}
	if err != nil {
		return nil, err
	}
	// cache contest
	for _, v := range po {
		marshal, err := json.Marshal(v)
		if err != nil {
			c.log.Errorf("marshal contest failed: %v", err)
		}
		_, err = c.data.redis.SetNX(ctx, redisPrefix+strconv.FormatInt(v.ID, 10), marshal, 2*time.Hour).Result()
		if err != nil {
			c.log.Errorf("cache contest failed: %v", err)
		}
	}
	var ret []*biz.Contest
	for _, v := range po {
		var groups []int64
		for _, g := range v.Edges.Contestants {
			groups = append(groups, g.ID)
		}
		if v.State != contest.StateNORMAL {
			continue
		}
		ret = append(ret, &biz.Contest{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Status:      timeToState(v.StartTime, v.EndTime),
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

func (c *contestRepo) UserStateCache(ctx context.Context, userId int64) (string, error) {
	state, err := c.data.redis.Get(ctx, userStatePrefix+strconv.FormatInt(userId, 10)).Int()
	if err != nil {
		return "", err
	}
	stateStr, err := util.UserStateToEnt(int16(state))
	if err != nil {
		return "", err
	}
	return stateStr.String(), nil
}

func (c *contestRepo) UserState(ctx context.Context, userId int64) (string, error) {
	return c.data.db.User.Query().Where(user.IDEQ(userId)).Select(user.FieldState).String(ctx)
}

func (c *contestRepo) UpdateUserStateCache(ctx context.Context, userId int64, state string) error {
	return c.data.redis.Set(ctx, userStatePrefix+strconv.FormatInt(userId, 10), util.UserStateToInt(user.State(state)), 1*time.Hour).Err()
}

func (c *contestRepo) IsUserBanned(state string) bool {
	return state != user.StateNORMAL.String()
}

// NewContestRepo .
func NewContestRepo(data *Data, logger log.Logger) biz.ContestRepo {
	return &contestRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func timeToState(s, e time.Time) int16 {
	now := time.Now()
	if now.Before(s) {
		return 0
	}
	if now.After(e) {
		return 2
	}
	return 1
}
