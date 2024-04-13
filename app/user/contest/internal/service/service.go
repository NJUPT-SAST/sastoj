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
	contestUc *biz.ContestUsecase
	problemUc *biz.ProblemUsecase
}

func NewUserContestService(contestUsecase *biz.ContestUsecase, problemUsecase *biz.ProblemUsecase) *UserContestService {
	return &UserContestService{
		contestUc: contestUsecase,
		problemUc: problemUsecase,
	}
}
