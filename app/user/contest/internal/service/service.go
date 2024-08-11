package service

import (
	pb "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/contest/internal/biz"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewContestService)

type ContestService struct {
	pb.UnimplementedContestServer
	contestUc  *biz.ContestUsecase
	problemUc  *biz.ProblemUsecase
	submitUc   *biz.SubmissionUsecase
	registerUc *biz.RegisterUsecase
}

func NewContestService(contestUsecase *biz.ContestUsecase, problemUsecase *biz.ProblemUsecase, submitUsecase *biz.SubmissionUsecase, registerUsecase *biz.RegisterUsecase) *ContestService {
	return &ContestService{
		contestUc:  contestUsecase,
		problemUc:  problemUsecase,
		submitUc:   submitUsecase,
		registerUc: registerUsecase,
	}
}
