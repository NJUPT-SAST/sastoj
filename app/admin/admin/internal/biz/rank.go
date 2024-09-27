package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sort"
	"time"
)

type Rank struct {
	ContestID   int64
	RefreshTime time.Time
	Users       []*UserContestResult
}

type UserContestResult struct {
	UserID      int64
	UserName    string
	Rank        int32
	Penalty     int32
	Score       int32
	ScoreTime   int32
	AcceptCount int32
	Problems    []*UserProblemResult
}

type UserProblemResult struct {
	SubmissionID int64
	ProblemID    int64
	Index        int16
	SubmitTime   time.Time
	Status       int16
	Point        int16
	TriedCount   int32
	AcceptTime   int32
}

type RankRepo interface {
	Save(context.Context, *Contest, *Rank) error
	Update(context.Context, *Contest, *Rank) error
	Delete(ctx context.Context, contest *Contest) error
	Find(ctx context.Context, contest *Contest) (*Rank, error)
	ListPage(ctx context.Context, contest *Contest, current int64, size int64) (*Rank, error)
	FindNewSubmissions(ctx context.Context, contestId int64, startTime time.Time) (map[int64]*UserContestResult, error)
}

type RankUsecase struct {
	repo RankRepo
	log  *log.Helper
}

func NewRankUsecase(repo RankRepo, logger log.Logger) *RankUsecase {
	return &RankUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (r *RankUsecase) FindRank(ctx context.Context, contest *Contest) (*Rank, error) {
	return r.repo.Find(ctx, contest)
}

func (r *RankUsecase) ListPage(ctx context.Context, contest *Contest, current int64, size int64) (*Rank, error) {
	return r.repo.ListPage(ctx, contest, current, size)
}

func (r *RankUsecase) SaveRank(ctx context.Context, contest *Contest, rank *Rank) error {
	return r.repo.Save(ctx, contest, rank)
}

func (r *RankUsecase) UpdateRank(ctx context.Context, contest *Contest, rank *Rank) error {
	return r.repo.Update(ctx, contest, rank)
}

func (r *RankUsecase) DeleteRank(ctx context.Context, contest *Contest) error {
	return r.repo.Delete(ctx, contest)
}

func (r *RankUsecase) RefreshRank(ctx context.Context, contest *Contest, oldRank *Rank) (*Rank, error) {
	if _, ok := ContestRankRuleMap[contest.Type]; !ok {
		// TODO: define error
		return nil, nil
	}
	submission, err := r.repo.FindNewSubmissions(ctx, contest.Id, oldRank.RefreshTime)
	if err != nil {
		return nil, err
	}
	rank := ContestRankRuleMap[contest.Type].Update(oldRank, contest, submission)
	return rank, nil
}

type ContestRankRule interface {
	Update(*Rank, *Contest, map[int64]*UserContestResult) *Rank
}

var ContestRankRuleMap = map[int32]ContestRankRule{
	ContestTypeIOI: &IOIRank{},
	ContestTypeACM: &ACMRank{},
}

type IOIRank struct{}

func (i *IOIRank) Update(rank *Rank, contest *Contest, submission map[int64]*UserContestResult) *Rank {
	const SubmissionAccepted int16 = 1
	var userProblem = make(map[int64]map[int64]*UserProblemResult)
	var rankResult = Rank{
		ContestID:   rank.ContestID,
		RefreshTime: time.Now(),
	}
	var userResult = make(map[int64]*UserContestResult)
	for _, v := range rank.Users {
		userProblem[v.UserID] = make(map[int64]*UserProblemResult)
		for _, p := range v.Problems {
			userProblem[v.UserID][p.ProblemID] = p
		}
		userResult[v.UserID] = v
	}
	for _, v := range submission {
		if _, ok := userProblem[v.UserID]; !ok {
			userProblem[v.UserID] = make(map[int64]*UserProblemResult)
		}
		if _, ok := userResult[v.UserID]; !ok {
			userResult[v.UserID] = &UserContestResult{
				UserID:   v.UserID,
				UserName: v.UserName,
				Problems: make([]*UserProblemResult, 0),
			}
		}
		for _, p := range v.Problems {
			if _, ok := userProblem[v.UserID][p.ProblemID]; !ok {
				userProblem[v.UserID][p.ProblemID] = p
			} else {
				if p.Point > userProblem[v.UserID][p.ProblemID].Point ||
					(p.Point == userProblem[v.UserID][p.ProblemID].Point && p.SubmitTime.Before(userProblem[v.UserID][p.ProblemID].SubmitTime)) {
					userProblem[v.UserID][p.ProblemID] = p
				}
			}
		}
	}
	for id, v := range userProblem {
		for _, p := range v {
			userResult[id].Problems = append(userResult[id].Problems, p)
			userResult[id].Score += int32(p.Point)
			userResult[id].ScoreTime = max(userResult[id].ScoreTime, int32(p.SubmitTime.Sub(contest.StartTime).Minutes()))
			if p.Status == SubmissionAccepted {
				userResult[id].AcceptCount++
			}
		}
		sort.Slice(userResult[id].Problems, func(i, j int) bool {
			return userResult[id].Problems[i].Index < userResult[id].Problems[j].Index
		})
		rankResult.Users = append(rankResult.Users, userResult[id])
	}
	sort.Slice(rankResult.Users, func(i, j int) bool {
		if rankResult.Users[i].Score != rankResult.Users[j].Score {
			return rankResult.Users[i].Score > rankResult.Users[j].Score
		} else if rankResult.Users[i].AcceptCount != rankResult.Users[j].AcceptCount {
			return rankResult.Users[i].AcceptCount > rankResult.Users[j].AcceptCount
		}
		return rankResult.Users[i].ScoreTime < rankResult.Users[j].ScoreTime
	})
	if len(rankResult.Users) == 0 {
		return rank
	}
	var rankID int32 = 1
	rankResult.Users[0].Rank = rankID
	for id := 1; id < len(rankResult.Users); id++ {
		if rankResult.Users[id].Score == rankResult.Users[id-1].Score {
			rankResult.Users[id].Rank = rankID
		} else {
			rankID = int32(id) + 1
			rankResult.Users[id].Rank = rankID
		}
	}
	return &rankResult
}

type ACMRank struct{}

func (a *ACMRank) Update(rank *Rank, contest *Contest, submission map[int64]*UserContestResult) *Rank {
	const SubmissionAccepted int16 = 1
	var userProblem = make(map[int64]map[int64]*UserProblemResult)
	var rankResult = Rank{
		ContestID:   rank.ContestID,
		RefreshTime: time.Now(),
	}
	var userResult = make(map[int64]*UserContestResult)
	for _, v := range rank.Users {
		userProblem[v.UserID] = make(map[int64]*UserProblemResult)
		for _, p := range v.Problems {
			userProblem[v.UserID][p.ProblemID] = p
		}
		userResult[v.UserID] = v
	}
	for _, v := range submission {
		if _, ok := userProblem[v.UserID]; !ok {
			userProblem[v.UserID] = make(map[int64]*UserProblemResult)
		}
		if _, ok := userResult[v.UserID]; !ok {
			userResult[v.UserID] = &UserContestResult{
				UserID:   v.UserID,
				UserName: v.UserName,
				Problems: make([]*UserProblemResult, 0),
			}
		}
		for _, p := range v.Problems {
			if _, ok := userProblem[v.UserID][p.ProblemID]; !ok {
				userProblem[v.UserID][p.ProblemID] = p
				if p.Status == SubmissionAccepted {
					userProblem[v.UserID][p.ProblemID].AcceptTime = int32(p.SubmitTime.Sub(contest.StartTime).Minutes())
				} else {
					userProblem[v.UserID][p.ProblemID].Point = 0
					userProblem[v.UserID][p.ProblemID].TriedCount = 1
				}
			} else {
				if userProblem[v.UserID][p.ProblemID].Status != SubmissionAccepted {
					if p.Status == SubmissionAccepted {
						userProblem[v.UserID][p.ProblemID].AcceptTime = int32(p.SubmitTime.Sub(rank.RefreshTime).Minutes())
						userProblem[v.UserID][p.ProblemID].Point = p.Point
						userProblem[v.UserID][p.ProblemID].Status = p.Status
					} else {
						userProblem[v.UserID][p.ProblemID].TriedCount++
						userProblem[v.UserID][p.ProblemID].SubmitTime = p.SubmitTime
					}
				}
			}
		}
	}
	for id, v := range userProblem {
		for _, p := range v {
			userResult[id].Problems = append(userResult[id].Problems, p)
			userResult[id].Score += int32(p.Point)
			userResult[id].ScoreTime = max(userResult[id].ScoreTime, p.AcceptTime)
			if p.Status == SubmissionAccepted {
				userResult[id].AcceptCount += 1
				userResult[id].Penalty += p.TriedCount*20 + p.AcceptTime
			}
		}
		sort.Slice(userResult[id].Problems, func(i, j int) bool {
			return userResult[id].Problems[i].Index < userResult[id].Problems[j].Index
		})
		rankResult.Users = append(rankResult.Users, userResult[id])
	}
	sort.Slice(rankResult.Users, func(i, j int) bool {
		if rankResult.Users[i].AcceptCount != rankResult.Users[j].AcceptCount {
			return rankResult.Users[i].AcceptCount > rankResult.Users[j].AcceptCount
		}
		if rankResult.Users[i].Penalty != rankResult.Users[j].Penalty {
			return rankResult.Users[i].Penalty < rankResult.Users[j].Penalty
		}
		return rankResult.Users[i].ScoreTime < rankResult.Users[j].ScoreTime
	})
	if len(rankResult.Users) == 0 {
		return rank
	}
	var rankID int32 = 1
	rankResult.Users[0].Rank = rankID
	for id := 1; id < len(rankResult.Users); id++ {
		if rankResult.Users[id].AcceptCount == rankResult.Users[id-1].AcceptCount &&
			rankResult.Users[id].Penalty == rankResult.Users[id-1].Penalty &&
			rankResult.Users[id].ScoreTime == rankResult.Users[id-1].ScoreTime {
			rankResult.Users[id].Rank = rankID
		} else {
			rankID = int32(id) + 1
			rankResult.Users[id].Rank = rankID
		}
	}
	return rank
}
