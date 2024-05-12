package data

import (
	"context"
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
	if contest.EndTime.Before(time.Now()) {
		const prefix = "admin:contest:rank:"
		key := prefix + strconv.FormatInt(contest.Id, 10)
		err := r.data.redis.Set(ctx, key, rank, contest.EndTime.Sub(time.Now())).Err()
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
			if err == nil {
				return err
			}
			submissionIds := make([]int64, len(v.Problems))
			for _, p := range v.Problems {
				submissionIds = append(submissionIds, p.SubmissionID)
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
	if contest.EndTime.Before(time.Now()) {
		const prefix = "admin:contest:rank:"
		key := prefix + strconv.FormatInt(contest.Id, 10)
		err := r.data.redis.Set(ctx, key, rank, contest.EndTime.Sub(time.Now())).Err()
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
			for _, p := range v.Problems {
				submissionIds = append(submissionIds, p.SubmissionID)
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
	if contest.EndTime.Before(time.Now()) {
		const prefix = "admin:contest:rank:"
		key := prefix + strconv.FormatInt(contest.Id, 10)
		var rank biz.Rank
		err := r.data.redis.Get(ctx, key).Scan(&rank)
		if err != nil {
			return nil, err
		}
		return &rank, nil
	} else {
		results, err := r.data.db.ContestResult.Query().Where(contestresult.ContestID(contest.Id)).All(ctx)
		if err != nil {
			return nil, err
		}
		var rank = &biz.Rank{
			ContestID: contest.Id,
			Users:     make([]*biz.UserContestResult, 0),
		}
		for _, v := range results {
			const SubmissionAccepted int16 = 1
			var problems []*biz.UserProblemResult
			for _, p := range v.Edges.Submissions {
				problems = append(problems, &biz.UserProblemResult{
					SubmissionID: p.ID,
					ProblemID:    p.ProblemID,
					SubmitTime:   p.CreateTime,
					Status:       p.Status,
					Point:        p.Point,
					Index:        p.Edges.Problems.Index,
				})
				if p.Status == SubmissionAccepted {
					rank.Users[v.UserID].AcceptCount++
				}
			}
			rank.Users = append(rank.Users, &biz.UserContestResult{
				UserID:    v.UserID,
				UserName:  v.Edges.User.Username,
				Rank:      v.Rank,
				Score:     v.Score,
				ScoreTime: v.ScoreTime,
				Penalty:   v.Penalty,
				Problems:  problems,
			})
		}
		return rank, nil
	}
}

func (r *RankRepo) FindNewSubmissions(ctx context.Context, contestId int64, startTime time.Time) (map[int64]*biz.UserContestResult, error) {
	res, err := r.data.db.Submission.Query().Where(submission.HasProblemsWith(problem.HasContestsWith(contest.ID(contestId))), submission.CreateTimeGT(startTime)).All(ctx)
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
