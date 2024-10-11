package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Submission is a Submit model.
type Submission struct {
	Id         int64
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
	GetReferenceAnswer(context.Context, int64) (*string, error)
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

func (uc *JudgeUsecase) SubmitJudge(ctx context.Context, submissionId int64, point int32) error {
	uc.log.WithContext(ctx).Infof("SubmitJudge : %d", submissionId)
	return uc.repo.SubmitJudge(ctx, submissionId, point)
}

func (uc *JudgeUsecase) GetJudgableProblems(ctx context.Context, userId int64) ([]*Problem, error) {
	uc.log.WithContext(ctx).Infof("GetJudgableProblems from user %d", userId)
	return uc.repo.GetJudgableProblems(ctx, userId)
}
func (uc *JudgeUsecase) GetSubmissions(ctx context.Context, problemId int64, status int32) ([]*Submission, error) {
	uc.log.WithContext(ctx).Infof("GetSubmissions from problem %d", problemId)
	if status == 0 {
		return uc.repo.GetSubmissions(ctx, problemId)
	}
	return uc.repo.GetSubmissionsWithStatus(ctx, problemId, status)
}

func (uc *JudgeUsecase) GetReferenceAnswer(ctx context.Context, problemId int64) (*string, error) {
	uc.log.WithContext(ctx).Infof("GetReferenceAnswer from problem %d", problemId)
	return uc.repo.GetReferenceAnswer(ctx, problemId)
}
