package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"sort"
	"time"
)

type Rank struct {
	ContestId int64
	UserRank  []*UserRank
}

type UserRank struct {
	UserId       int64
	UserName     string
	Rank         int32
	Penalty      int32
	TotalPoint   int32
	AcceptCount  int32
	AchievedTime time.Time
	Problems     map[int64]*UserProblemResult
}

type UserProblemResult struct {
	SubmissionId int64
	ProblemId    int64
	State        int32
	Point        int32
	TriedCount   int32
	SubmitTime   time.Time
}

const (
	ProblemUnaccepted    = 1
	ProblemAccepted      = 2
	ProblemFirstAccepted = 3
)

type RankRepo interface {
	GetSubmissions(ctx context.Context, contestId int64) (map[int64]*UserRank, error)
	Save(ctx context.Context, contest *Contest, rank *Rank) error
	Find(ctx context.Context, contest *Contest) (*Rank, error)
}

type RankUsecase struct {
	repo RankRepo
	log  *log.Helper
}

func NewRankUsecase(repo RankRepo, logger log.Logger) *RankUsecase {
	return &RankUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (r *RankUsecase) Find(ctx context.Context, contest *Contest) (*Rank, error) {
	if contest.EndTime.After(time.Now()) {
		return r.repo.Find(ctx, contest)
	} else {
		return r.Update(ctx, contest)
	}
}

func (r *RankUsecase) Save(ctx context.Context, contest *Contest, rank *Rank) error {
	return r.repo.Save(ctx, contest, rank)
}

func (r *RankUsecase) Update(ctx context.Context, contest *Contest) (*Rank, error) {
	userSubmissions, err := r.repo.GetSubmissions(ctx, contest.Id)
	if err != nil {
		return nil, err
	}
	switch contest.Type {
	case ContestTypeIOI:
		return r.rankByIOI(userSubmissions, contest)
	case ContestTypeACM:
		return r.rankByACM(userSubmissions, contest)
	default:
		return nil, fmt.Errorf("unknown contest type, find contest type: %d", contest.Type)
	}
}

func (r *RankUsecase) rankByIOI(userSubmissions map[int64]*UserRank, contest *Contest) (*Rank, error) {
	var res = &Rank{
		ContestId: contest.Id,
		UserRank:  make([]*UserRank, 0),
	}
	if len(userSubmissions) == 0 {
		return res, nil
	}
	for _, v := range userSubmissions {
		res.UserRank = append(res.UserRank, v)
	}
	sort.Slice(res.UserRank, func(i, j int) bool {
		if res.UserRank[i].TotalPoint != res.UserRank[j].TotalPoint {
			return res.UserRank[i].TotalPoint > res.UserRank[j].TotalPoint
		}
		if !res.UserRank[i].AchievedTime.Equal(res.UserRank[j].AchievedTime) {
			return res.UserRank[i].AchievedTime.Before(res.UserRank[j].AchievedTime)
		}
		return res.UserRank[i].AcceptCount < res.UserRank[j].AcceptCount
	})
	var rank int32 = 1
	res.UserRank[0].Rank = rank
	for i := 1; i < len(res.UserRank); i++ {
		if res.UserRank[i].TotalPoint == res.UserRank[i-1].TotalPoint &&
			res.UserRank[i].AcceptCount == res.UserRank[i-1].AcceptCount &&
			res.UserRank[i].AchievedTime.Equal(res.UserRank[i-1].AchievedTime) {
			res.UserRank[i].Rank = rank
		} else {
			rank = int32(i + 1)
			res.UserRank[i].Rank = rank
		}
	}
	return res, nil
}

func (r *RankUsecase) rankByACM(userSubmissions map[int64]*UserRank, contest *Contest) (*Rank, error) {
	var res = &Rank{
		ContestId: contest.Id,
		UserRank:  make([]*UserRank, 0),
	}
	if len(userSubmissions) == 0 {
		return res, nil
	}
	for _, v := range userSubmissions {
		problems := make(map[int64]*UserProblemResult)
		pent := int32(0)
		score := int32(0)
		achieveTime := time.Unix(0, 0)
		for _, p := range v.Problems {
			prob := &UserProblemResult{
				SubmissionId: p.SubmissionId,
				ProblemId:    p.ProblemId,
				State:        p.State,
				TriedCount:   p.TriedCount,
				SubmitTime:   p.SubmitTime,
			}
			if p.State != ProblemAccepted {
				prob.Point = 0
			} else {
				prob.Point = p.Point
				score += p.Point
				pent += 20*p.TriedCount + int32(p.SubmitTime.Sub(contest.StartTime).Minutes())
				if p.SubmitTime.After(achieveTime) {
					achieveTime = p.SubmitTime
				}
			}
			problems[p.ProblemId] = prob
		}
		res.UserRank = append(res.UserRank, &UserRank{
			UserId:       v.UserId,
			UserName:     v.UserName,
			Penalty:      pent,
			TotalPoint:   score,
			AcceptCount:  v.AcceptCount,
			AchievedTime: achieveTime,
			Problems:     problems,
		})
	}
	sort.Slice(res.UserRank, func(i, j int) bool {
		if res.UserRank[i].TotalPoint != res.UserRank[j].TotalPoint {
			return res.UserRank[i].TotalPoint > res.UserRank[j].TotalPoint
		}
		if res.UserRank[i].Penalty != res.UserRank[j].Penalty {
			return res.UserRank[i].Penalty < res.UserRank[j].Penalty
		}
		return res.UserRank[i].AchievedTime.Before(res.UserRank[j].AchievedTime)
	})
	var rank int32 = 1
	res.UserRank[0].Rank = rank
	for i := 1; i < len(res.UserRank); i++ {
		if res.UserRank[i].TotalPoint == res.UserRank[i-1].TotalPoint &&
			res.UserRank[i].Penalty == res.UserRank[i-1].Penalty &&
			res.UserRank[i].AchievedTime.Equal(res.UserRank[i-1].AchievedTime) {
			res.UserRank[i].Rank = rank
		} else {
			rank = int32(i + 1)
			res.UserRank[i].Rank = rank
		}
	}
	return res, nil
}
