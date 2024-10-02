package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/judge/freshcup/internal/biz"
	"sastoj/pkg/mq"
)

type submissionRepo struct {
	data *Data
	log  *log.Helper
}

func (r *submissionRepo) JudgeSelfTest(ctx context.Context, test *mq.SelfTest) error {
	return nil
}

func (r *submissionRepo) JudgeSubmission(ctx context.Context, s *mq.Submission) error {
	return nil
}

// NewSubmissionRepo .
func NewSubmissionRepo(data *Data, logger log.Logger) biz.SubmissionRepo {
	return &submissionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
