package biz

import (
	"context"
	"sastoj/pkg/mq"

	"github.com/go-kratos/kratos/v2/log"
)

// SubmissionRepo is a Greater repo.
type SubmissionRepo interface {
	JudgeSubmission(context.Context, *mq.Submission) error
	JudgeSelfTest(context.Context, *mq.SelfTest) error
}

// SubmissionUsecase is a Submission usecase.
type SubmissionUsecase struct {
	repo SubmissionRepo
	log  *log.Helper
}

// NewSubmissionUsecase new a Submission usecase.
func NewSubmissionUsecase(repo SubmissionRepo, logger log.Logger) *SubmissionUsecase {
	return &SubmissionUsecase{repo: repo, log: log.NewHelper(logger)}
}

// JudgeSubmission creates a Submission, and returns error or nil.
func (uc *SubmissionUsecase) JudgeSubmission(ctx context.Context, s *mq.Submission) error {
	uc.log.WithContext(ctx).Infof("JudgeSubmission: %v", s.ID)
	return uc.repo.JudgeSubmission(ctx, s)
}

// JudgeSelfTest creates a SelfTest, and returns error or nil.
func (uc *SubmissionUsecase) JudgeSelfTest(ctx context.Context, s *mq.SelfTest) error {
	uc.log.WithContext(ctx).Infof("JudgeSelfTest: %v", s.ID)
	return uc.repo.JudgeSelfTest(ctx, s)
}
