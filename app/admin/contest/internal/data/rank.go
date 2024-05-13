package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/admin/contest/internal/biz"
	"sastoj/ent/contest"
	"sastoj/ent/contestresult"
	"sastoj/ent/problem"
	"sastoj/ent/submission"
	"strconv"
	"time"
)

type RankRepo struct {
	data *Data
	log  *log.Helper
}

func NewRankRepo(data *Data, logger log.Logger) biz.RankRepo {
	return &RankRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *RankRepo) Save(ctx context.Context, contest *biz.Contest, rank *biz.Rank) error {
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
		for _, v := range rank.Users {
			_, err := r.data.db.ContestResult.Create().
				SetContestID(contest.Id).
				SetUserID(v.UserID).
				SetRank(v.Rank).
				SetScore(v.Score).
				SetPenalty(v.Penalty).
				SetScoreTime(v.ScoreTime).
				Save(ctx)
			if err != nil {
				return err
			}
			submissionIds := make([]int64, len(v.Problems))
			for id, p := range v.Problems {
				submissionIds[id] = p.SubmissionID
			}
			_, err = r.data.db.ContestResult.Update().
				Where(contestresult.UserID(v.UserID), contestresult.ContestID(contest.Id)).
				AddSubmissionIDs(submissionIds...).Save(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *RankRepo) Update(ctx context.Context, contest *biz.Contest, rank *biz.Rank) error {
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
		for _, v := range rank.Users {
			_, err := r.data.db.ContestResult.Update().
				Where(contestresult.UserID(v.UserID), contestresult.ContestID(contest.Id)).
				SetRank(v.Rank).
				SetScore(v.Score).
				SetPenalty(v.Penalty).
				SetScoreTime(v.ScoreTime).
				Save(ctx)
			if err == nil {
				return err
			}
			err = r.data.db.ContestResult.Update().Where(contestresult.ContestID(contest.Id)).ClearSubmissions().Exec(ctx)
			if err != nil {
				return err
			}
			submissionIds := make([]int64, len(v.Problems))
			for id, p := range v.Problems {
				submissionIds[id] = p.SubmissionID
			}
			_, err = r.data.db.ContestResult.Update().
				Where(contestresult.UserID(v.UserID), contestresult.ContestID(contest.Id)).
				AddSubmissionIDs(submissionIds...).Save(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *RankRepo) Delete(ctx context.Context, contest *biz.Contest) error {
	if contest.EndTime.Before(time.Now()) {
		const prefix = "admin:contest:rank:"
		key := prefix + strconv.FormatInt(contest.Id, 10)
		err := r.data.redis.Del(ctx, key).Err()
		if err != nil {
			return err
		}
	} else {
		_, err := r.data.db.ContestResult.Delete().Where(contestresult.ContestID(contest.Id)).Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RankRepo) Find(ctx context.Context, contest *biz.Contest) (*biz.Rank, error) {
	if contest.EndTime.After(time.Now()) {
		const prefix = "admin:contest:rank:"
		key := prefix + strconv.FormatInt(contest.Id, 10)
		var rankData string
		err := r.data.redis.Get(ctx, key).Scan(&rankData)
		if err != nil {
			return nil, err
		}
		var rank biz.Rank
		err = json.Unmarshal([]byte(rankData), &rank)
		if err != nil {
			return nil, err
		}
		return &rank, nil
	} else {
		results, err := r.data.db.ContestResult.Query().
			Where(contestresult.ContestID(contest.Id)).
			WithUser().WithSubmissions().
			All(ctx)
		if err != nil {
			return nil, err
		}
		problemList, err := r.data.db.Problem.Query().Where(problem.ContestID(contest.Id)).All(ctx)
		if err != nil {
			return nil, err
		}
		var problemId = make(map[int64]int16)
		for _, v := range problemList {
			problemId[v.ID] = v.Index
		}
		var rank = &biz.Rank{
			ContestID: contest.Id,
			Users:     make([]*biz.UserContestResult, 0),
		}
		for _, v := range results {
			const SubmissionAccepted int16 = 1
			var problems []*biz.UserProblemResult
			var acceptCount int32 = 0
			for _, p := range v.Edges.Submissions {
				problems = append(problems, &biz.UserProblemResult{
					SubmissionID: p.ID,
					ProblemID:    p.ProblemID,
					SubmitTime:   p.CreateTime,
					Status:       p.Status,
					Point:        p.Point,
					Index:        problemId[p.ProblemID],
				})
				if p.Status == SubmissionAccepted {
					acceptCount += 1
				}
			}
			rank.Users = append(rank.Users, &biz.UserContestResult{
				UserID:      v.UserID,
				UserName:    v.Edges.User.Username,
				Rank:        v.Rank,
				Score:       v.Score,
				ScoreTime:   v.ScoreTime,
				Penalty:     v.Penalty,
				AcceptCount: acceptCount,
				Problems:    problems,
			})
		}
		return rank, nil
	}
}

func (r *RankRepo) FindNewSubmissions(ctx context.Context, contestId int64, startTime time.Time) (map[int64]*biz.UserContestResult, error) {
	res, err := r.data.db.Submission.Query().
		Where(submission.HasProblemsWith(problem.HasContestsWith(contest.ID(contestId))), submission.CreateTimeGT(startTime)).
		WithUsers().WithProblems().
		All(ctx)
	if err != nil {
		return nil, err
	}
	var rv = make(map[int64]*biz.UserContestResult)
	for _, s := range res {
		if _, ok := rv[s.UserID]; !ok {
			rv[s.UserID] = &biz.UserContestResult{
				UserID:   s.UserID,
				UserName: s.Edges.Users.Username,
			}
		}
		rv[s.UserID].Problems = append(rv[s.UserID].Problems, &biz.UserProblemResult{
			SubmissionID: s.ID,
			ProblemID:    s.ProblemID,
			SubmitTime:   s.CreateTime,
			Status:       s.Status,
			Point:        s.Point,
			Index:        s.Edges.Problems.Index,
		})
	}
	return rv, nil
}
