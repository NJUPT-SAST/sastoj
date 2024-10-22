package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/user/contest/internal/biz"
	"strconv"
)

type rankRepo struct {
	data *Data
	log  *log.Helper
}

func NewRankRepo(data *Data, logger log.Logger) biz.RankRepo {
	return &rankRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *rankRepo) Find(ctx context.Context, contestId int64) (*biz.Rank, error) {
	const prefix = "admin:contest:rank:"
	key := prefix + strconv.FormatInt(contestId, 10)
	data, err := r.data.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var rank biz.Rank
	err = json.Unmarshal([]byte(data), &rank)
	if err != nil {
		return nil, err
	}
	return &rank, nil
}
