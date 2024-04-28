package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/user/gateway/internal/biz"
	"strconv"
)

type problemRepo struct {
	data *Data
	log  *log.Helper
}

func (p *problemRepo) GetProblems(ctx context.Context, contestID int64) ([]*biz.Problem, error) {
	po := p.data.redis.Get(ctx, "contest:"+strconv.FormatInt(contestID, 10))
	if po.Err() != nil {
		return nil, po.Err()
	}
	problemsJson := po.Val()
	var problems []*biz.Problem
	err := json.Unmarshal([]byte(problemsJson), &problems)
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func (p *problemRepo) GetProblem(ctx context.Context, problemID, contestID int64) (*biz.Problem, error) {
	po := p.data.redis.Get(ctx, "problem:"+strconv.FormatInt(problemID, 10))
	if po.Err() != nil {
		return nil, po.Err()
	}
	var problem biz.Problem
	err := json.Unmarshal([]byte(po.Val()), &problem)
	if err != nil {
		return nil, err
	}
	return &problem, nil
}

func NewProblemRepo(data *Data, logger log.Logger) biz.ProblemRepo {
	return &problemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
