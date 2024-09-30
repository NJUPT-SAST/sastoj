package job

import (
	"context"
	"sastoj/app/admin/admin/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
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
		rank, err := s.rc.Update(ctx, contest)
		if err != nil {
			log.Warn(err.Error())
			continue
		}
		err = s.rc.Save(ctx, contest, rank)
		if err != nil {
			log.Warn(err.Error())
			continue
		}
		log.Info("cron ranking contest " + contest.Title + " successfully.")
	}
}
