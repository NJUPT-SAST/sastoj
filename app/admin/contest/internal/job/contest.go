package job

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/admin/contest/internal/biz"
)

type ContestJob struct {
	uc *biz.ContestUsecase
	rc *biz.RankUsecase
}

func NewContestJob(uc *biz.ContestUsecase, rc *biz.RankUsecase) *ContestJob {
	job := &ContestJob{
		uc: uc,
		rc: rc,
	}
	return job
}

func (s *ContestJob) Init() *map[string]JobFunc {
	return &map[string]JobFunc{
		"one": s.Ranking,
	}
}

func (s *ContestJob) Ranking(ctx context.Context) {
	contests, err := s.uc.GetRacingContests(ctx)
	if err != nil {
		log.Warn(err.Error())
	}
	for _, contest := range contests {
		rank, err := s.rc.FindRank(ctx, contest)
		if err != nil && rank != nil {
			log.Warn(err.Error())
			return
		}
		rank, err = s.rc.RefreshRank(ctx, contest, rank)
		if err != nil {
			log.Warn(err.Error())
			return
		}
		err = s.rc.SaveRank(ctx, contest, rank)
		if err != nil {
			log.Warn(err.Error())
			return
		}
		log.Info("cron ranking contest " + contest.Title + " successfully.")
	}
}
