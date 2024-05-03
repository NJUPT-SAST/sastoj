package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Case struct {
	Id        int64
	ProblemId int64
	Point     int32
	Index     int32
	IsAuto    bool
}

// CasesRepo is a Contest repo.
type CasesRepo interface {
	FetchCases(int64) ([]byte, int64, error)
}

// CasesUsecase is a Contest usecase.
type CasesUsecase struct {
	repo CasesRepo
	log  *log.Helper
}

// NewCasesUsecase new a Contest usecase.
func NewCasesUsecase(repo CasesRepo, logger log.Logger) *CasesUsecase {
	return &CasesUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CasesUsecase) FetchCases(problemId int64) ([]byte, int64, error) {
	uc.log.Infof("FetchCases from problem %d to rsjudge", problemId)
	return uc.repo.FetchCases(problemId)
}
