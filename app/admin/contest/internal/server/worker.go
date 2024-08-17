package server

import (
	"context"
	"github.com/go-co-op/gocron/v2"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/admin/contest/internal/conf"
	"sastoj/app/admin/contest/internal/job"
	"time"
)

type CronWorker struct {
	scheduler gocron.Scheduler
}

func NewCronWorker(c *conf.Job, contestJob *job.ContestJob, logger log.Logger) (*CronWorker, error) {
	if !c.GetRankingCron() {
		log.Info("disable creating ranking cron job.")
		return nil, nil
	}
	ctx := context.Background()
	jobs := contestJob.Init()
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}
	for name, task := range *jobs {
		_, err = scheduler.NewJob(
			gocron.DurationJob(
				time.Duration(c.GetDuration())*time.Second,
			),
			gocron.NewTask(
				func() {
					task(ctx)
				},
			),
		)
		if err != nil {
			log.Warnf("can not create ranking cron job for task %s.", name)
			return nil, err
		}
	}
	return &CronWorker{scheduler: scheduler}, nil
}

func (s *CronWorker) Start(c context.Context) error {
	if s == nil {
		return nil
	}
	s.scheduler.Start()
	log.Info("create ranking cron job successfully.")
	return nil
}

func (s *CronWorker) Stop(c context.Context) error {
	return s.scheduler.Shutdown()
}
