package service

import (
	"github.com/google/wire"
	pb "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/contest/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserContestService)

type UserContestService struct {
	pb.UnimplementedContestServiceServer
	contestUc  *biz.ContestUsecase
	problemUc  *biz.ProblemUsecase
	submitUc   *biz.SubmitUsecase
	registerUc *biz.RegisterUsecase
}

func NewUserContestService(contestUsecase *biz.ContestUsecase, problemUsecase *biz.ProblemUsecase, submitUsecase *biz.SubmitUsecase, registerUsecase *biz.RegisterUsecase) *UserContestService {
	return &UserContestService{
		contestUc:  contestUsecase,
		problemUc:  problemUsecase,
		submitUc:   submitUsecase,
		registerUc: registerUsecase,
	}
}
