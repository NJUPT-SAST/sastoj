package data

import (
	"context"
	"encoding/json"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent/contest"
	"sastoj/ent/problem"
	"sastoj/ent/submission"
	"sastoj/pkg/util"
	"sort"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
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

func (r *rankRepo) Find(ctx context.Context, contest *biz.Contest) (*biz.Rank, error) {
	if contest.EndTime.After(time.Now()) {
		const prefix = "admin:contest:rank:"
		key := prefix + strconv.FormatInt(contest.Id, 10)
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
	} else {
		return nil, nil
	}
}

func (r *rankRepo) Save(ctx context.Context, contest *biz.Contest, rank *biz.Rank) error {
	if contest.EndTime.After(time.Now()) {
		const prefix = "admin:contest:rank:"
		key := prefix + strconv.FormatInt(contest.Id, 10)
		rankData, err := json.Marshal(rank)
		if err != nil {
			return err
		}
		err = r.data.redis.Set(ctx, key, rankData, contest.EndTime.Sub(time.Now())).Err()
		if err != nil {
			return err
		}
	} else {
		return nil
	}
	return nil
}

func (r *rankRepo) GetSubmissions(ctx context.Context, contestId int64) (map[int64]*biz.UserRank, error) {
	submissions, err := r.data.db.Submission.Query().
		Where(submission.HasProblemsWith(problem.HasContestWith(contest.ID(contestId)))).
		WithUsers().
		WithProblems().
		All(ctx)
	if err != nil {
		return nil, err
	}
	sort.Slice(submissions, func(i, j int) bool {
		return submissions[i].CreateTime.Before(submissions[j].CreateTime)
	})
	// flag if the problem is accepted
	var problemAccepted = make(map[int64]struct{})
	var res = make(map[int64]*biz.UserRank)
	for _, v := range submissions {
		if _, ok := res[v.UserID]; !ok {
			res[v.UserID] = &biz.UserRank{
				UserId:       v.UserID,
				UserName:     v.Edges.Users.Username,
				AchievedTime: time.Unix(0, 0),
				Problems:     make(map[int64]*biz.UserProblemResult),
			}
		}
		var status int32 = biz.ProblemUnaccepted
		if v.State == util.Accepted {
			if _, ok := problemAccepted[v.ProblemID]; !ok {
				problemAccepted[v.ProblemID] = struct{}{}
				status = biz.ProblemFirstAccepted
			} else {
				status = biz.ProblemAccepted
			}
		}
		if _, ok := res[v.UserID].Problems[v.ProblemID]; !ok {
			res[v.UserID].Problems[v.ProblemID] = &biz.UserProblemResult{
				SubmissionId: v.ID,
				ProblemId:    v.ProblemID,
				State:        status,
				Point:        int32(v.Point),
				TriedCount:   0,
				SubmitTime:   v.CreateTime,
			}
		} else {
			if res[v.UserID].Problems[v.ProblemID].State != biz.ProblemUnaccepted {
				continue
			}
			res[v.UserID].Problems[v.ProblemID].TriedCount += 1
			// only update submission when score is higher
			if int64(v.Point) > int64(res[v.UserID].Problems[v.ProblemID].Point) {
				res[v.UserID].Problems[v.ProblemID].Point = int32(v.Point)
				res[v.UserID].Problems[v.ProblemID].SubmissionId = v.ID
				res[v.UserID].Problems[v.ProblemID].SubmitTime = v.CreateTime
				res[v.UserID].Problems[v.ProblemID].State = status
			}
		}
	}
	for i, v := range res {
		for _, p := range v.Problems {
			if p.State != biz.ProblemUnaccepted {
				res[i].AcceptCount += 1
			}
			res[i].TotalPoint += p.Point
			if v.AchievedTime.Before(p.SubmitTime) {
				v.AchievedTime = p.SubmitTime
			}
		}
	}
	return res, nil
}
