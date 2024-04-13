package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/user/contest/internal/biz"
	"sastoj/ent/problem"
)

type problemRepo struct {
	data *Data
	log  *log.Helper
}

func (p problemRepo) ListProblem(ctx context.Context, contestID int) ([]*biz.Problem, error) {
	po, err := p.data.db.Problem.Query().
		Select(problem.FieldID, problem.FieldTitle, problem.FieldPoint, problem.FieldIndex).
		Where(problem.ContestIDEQ(contestID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var problems []*biz.Problem
	for _, v := range po {
		problems = append(problems, &biz.Problem{
			ID:    int64(v.ID),
			Title: v.Title,
			Point: v.Point,
			Index: v.Index,
		})
	}
	return problems, nil
}

func (p problemRepo) GetProblem(ctx context.Context, problemID, contestID int) (*biz.Problem, error) {
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
		Index:   po.Index,
	}, nil
}

func NewProblemRepo(data *Data, logger log.Logger) biz.ProblemRepo {
	return &problemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
