package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// Problem is a Problem model.
type Problem struct {
	Id          int64
	Title       string
	Content     string
	Point       int32
	ContestId   int64
	CaseVersion int32
	Index       int32
	Config      string
}

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

func (uc *JudgeUsecase) SubmitJudge(ctx context.Context, submitId int64, point int32) error {
	uc.log.WithContext(ctx).Infof("SubmitJudge : %d", submitId)
	return uc.repo.SubmitJudge(ctx, submitId, point)
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
