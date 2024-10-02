package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Submission is a Submit model.
type Submission struct {
	ID         int64
	Code       string
	Status     int32
	Point      int32
	CreateTime time.Time
}

// JudgeRepo is a Greater repo.
type JudgeRepo interface {
	SubmitJudge(context.Context, int64, int32) error
	GetJudgableProblems(context.Context, int64) ([]*Problem, error)
	GetSubmissions(context.Context, int64) ([]*Submission, error)
	GetSubmissionsWithStatus(context.Context, int64, int32) ([]*Submission, error)
}

// JudgeUsecase is a Judge usecase.
type JudgeUsecase struct {
	repo JudgeRepo
	log  *log.Helper
}

// NewJudgeUsecase new a Judge usecase.
func NewJudgeUsecase(repo JudgeRepo, logger log.Logger) *JudgeUsecase {
	return &JudgeUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *JudgeUsecase) SubmitJudge(ctx context.Context, submissionID int64, point int32) error {
	uc.log.WithContext(ctx).Infof("SubmitJudge : %d", submissionID)
	return uc.repo.SubmitJudge(ctx, submissionID, point)
}

func (uc *JudgeUsecase) GetJudgableProblems(ctx context.Context, userID int64) ([]*Problem, error) {
	uc.log.WithContext(ctx).Infof("GetJudgableProblems from user %d", userID)
	return uc.repo.GetJudgableProblems(ctx, userID)
}
func (uc *JudgeUsecase) GetSubmissions(ctx context.Context, problemID int64, status int32) ([]*Submission, error) {
	uc.log.WithContext(ctx).Infof("GetSubmissions from problem %d", problemID)
	if status == 0 {
		return uc.repo.GetSubmissions(ctx, problemID)
	}
	return uc.repo.GetSubmissionsWithStatus(ctx, problemID, status)
}
