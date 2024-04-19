package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/user/contest/internal/biz"
	"sastoj/ent/contest"
	"sastoj/ent/problem"
	"strconv"
	"time"
)

type problemRepo struct {
	data *Data
	log  *log.Helper
}

func (p problemRepo) GetProblemCaseVer(ctx context.Context, problemId int64) (int32, error) {
	result, err := p.data.redis.Get(ctx, "problem:"+strconv.Itoa(int(problemId))+":"+"case_version").Result()
	if err == nil {
		ver, _ := strconv.Atoi(result)
		return int32(ver), nil
	}
	po, err := p.data.db.Problem.Query().Where(problem.IDEQ(problemId)).Only(ctx)
	if err != nil {
		return 0, err
	}
	p.data.redis.Set(ctx, "problem:"+strconv.Itoa(int(problemId))+":"+"case_version", strconv.Itoa(int(po.CaseVersion)), 2*time.Hour)
	return int32(po.CaseVersion), nil
}

func (p problemRepo) ListProblem(ctx context.Context, contestID int64) ([]*biz.Problem, error) {
	po, err := p.data.db.Problem.Query().
		Select(problem.FieldID, problem.FieldTitle, problem.FieldPoint, problem.FieldIndex).
		Where(problem.HasContestsWith(contest.IDEQ(contestID))).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var problems []*biz.Problem
	for _, v := range po {
		problems = append(problems, &biz.Problem{
			ID:    v.ID,
			Title: v.Title,
			Point: int32(v.Point),
			Index: int32(v.Index),
		})
	}
	return problems, nil
}

func (p problemRepo) GetProblem(ctx context.Context, problemID, contestID int64) (*biz.Problem, error) {
	po, err := p.data.db.Problem.Query().
		Select(problem.FieldID, problem.FieldTitle, problem.FieldContent, problem.FieldPoint).
		Where(problem.IDEQ(problemID), problem.ContestIDEQ(contestID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Problem{
		ID:      int64(po.ID),
		Title:   po.Title,
		Content: po.Content,
		Index:   int32(po.Index),
	}, nil
}

func NewProblemRepo(data *Data, logger log.Logger) biz.ProblemRepo {
	return &problemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
