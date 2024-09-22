package data

import (
	"context"
	"sastoj/app/user/contest/internal/biz"
	"sastoj/ent"
	"sastoj/ent/contest"
	"sastoj/ent/problem"
	"sastoj/ent/problemtype"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type problemRepo struct {
	data *Data
	log  *log.Helper
}

func (p problemRepo) GetProblemCaseVer(ctx context.Context, problemId int64) (int8, error) {
	result, err := p.data.redis.Get(ctx, "problem:"+strconv.Itoa(int(problemId))+":"+"case_version").Result()
	if err == nil {
		ver, _ := strconv.Atoi(result)
		return int8(ver), nil
	}
	po, err := p.data.db.Problem.Query().Where(problem.IDEQ(problemId)).Only(ctx)
	if err != nil {
		return 0, err
	}
	p.data.redis.Set(ctx, "problem:"+strconv.Itoa(int(problemId))+":"+"case_version", strconv.Itoa(int(po.CaseVersion)), 2*time.Hour)
	return int8(po.CaseVersion), nil
}

func (p problemRepo) ListProblem(ctx context.Context, contestID int64) ([]*biz.Problem, error) {
	po, err := p.data.db.Problem.Query().
		Select(problem.FieldID, problem.FieldTitle, problem.FieldScore, problem.FieldIndex).
		Where(problem.HasContestWith(contest.IDEQ(contestID))).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var problems []*biz.Problem
	for _, v := range po {
		problems = append(problems, &biz.Problem{
			ID:    v.ID,
			Title: v.Title,
			Point: v.Score,
			Index: v.Index,
		})
	}
	return problems, nil
}

func (p problemRepo) GetProblem(ctx context.Context, problemID, contestID int64) (*biz.Problem, error) {
	po, err := p.data.db.Problem.Query().
		Select(problem.FieldID, problem.FieldTitle, problem.FieldContent, problem.FieldScore).
		WithProblemType(func(q *ent.ProblemTypeQuery) {
			q.Select(problemtype.FieldDisplayName)
		}).
		Where(problem.IDEQ(problemID), problem.ContestIDEQ(contestID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Problem{
		ID:      po.ID,
		Title:   po.Title,
		Type:    po.Edges.ProblemType.DisplayName,
		Content: po.Content,
		Index:   po.Index,
	}, nil
}

func NewProblemRepo(data *Data, logger log.Logger) biz.ProblemRepo {
	return &problemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
