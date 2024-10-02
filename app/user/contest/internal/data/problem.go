package data

import (
	"context"
	"encoding/json"
	"fmt"
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

const ProblemPrefix = "user:contest:problem:"

func (p *problemRepo) GetProblemCaseVer(ctx context.Context, problemID int64) (int8, error) {
	result, err := p.data.redis.Get(ctx, "problem:"+strconv.Itoa(int(problemID))+":"+"case_version").Result()
	if err == nil {
		ver, _ := strconv.Atoi(result)
		return int8(ver), nil
	}
	po, err := p.data.db.Problem.Query().Where(problem.IDEQ(problemID)).Only(ctx)
	if err != nil {
		return 0, err
	}
	p.data.redis.Set(ctx, "problem:"+strconv.Itoa(int(problemID))+":"+"case_version", strconv.Itoa(int(po.CaseVersion)), 2*time.Hour)
	return int8(po.CaseVersion), nil
}

func (p *problemRepo) ListProblem(ctx context.Context, contestID int64) ([]*biz.Problem, error) {
	po, err := p.data.db.Problem.Query().
		Select(problem.FieldID, problem.FieldTitle, problem.FieldScore, problem.FieldIndex).
		Where(problem.HasContestWith(contest.IDEQ(contestID))).
		WithProblemType(func(q *ent.ProblemTypeQuery) {
			q.Select(problemtype.FieldDisplayName)
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var problems []*biz.Problem
	for _, v := range po {
		problems = append(problems, &biz.Problem{
			ID:    v.ID,
			Type:  v.Edges.ProblemType.DisplayName,
			Title: v.Title,
			Score: int32(v.Score),
			Index: v.Index,
		})
	}
	return problems, nil
}

func (p *problemRepo) GetProblem(ctx context.Context, problemID, contestID int64) (*biz.Problem, error) {
	key := ProblemPrefix + strconv.FormatInt(problemID, 10)
	result, err := p.data.redis.Get(ctx, key).Result()
	if err == nil {
		var po *ent.Problem
		err = json.Unmarshal([]byte(result), &po)
		if err != nil {
			p.log.Errorf("unmarshal problem failed: %v", err)
		} else {
			return &biz.Problem{
				ID:      po.ID,
				Title:   po.Title,
				Type:    po.Edges.ProblemType.DisplayName,
				Content: po.Content,
				Index:   po.Index,
				Score:   int32(po.Score),
			}, nil
		}
	}
	now := time.Now()
	po, err := p.data.db.Problem.Query().
		Select(problem.FieldID, problem.FieldTitle, problem.FieldContent, problem.FieldScore).
		Where(
			problem.IDEQ(problemID),
			problem.ContestIDEQ(contestID),
			problem.HasContestWith(
				contest.StateEQ(contest.StateNORMAL),
				contest.StartTimeLT(now),
				contest.EndTimeGT(now))).
		WithProblemType(func(q *ent.ProblemTypeQuery) {
			q.Select(
				problemtype.FieldDisplayName,
				problemtype.FieldSubmissionChannelName,
				problemtype.FieldSelfTestChannelName)
		}).
		WithContest(func(q *ent.ContestQuery) {
			q.Select(contest.FieldEndTime).WithContestants()
		}).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("problem %d is not exist", problemID)
	}

	marshal, err := json.Marshal(po)
	if err != nil {
		p.log.Errorf("marshal problem failed: %v", err)
	} else {
		duration := po.Edges.Contest.EndTime.Sub(now)
		_, err = p.data.redis.SetNX(ctx, key, marshal, duration).Result()
		if err != nil {
			p.log.Errorf("cache problem failed: %v", err)
		}
	}

	return &biz.Problem{
		ID:      po.ID,
		Title:   po.Title,
		Type:    po.Edges.ProblemType.DisplayName,
		Content: po.Content,
		Index:   po.Index,
		Score:   int32(po.Score),
	}, nil
}

func NewProblemRepo(data *Data, logger log.Logger) biz.ProblemRepo {
	return &problemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
