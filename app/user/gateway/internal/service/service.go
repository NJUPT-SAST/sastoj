package service

import (
	pb "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/gateway/internal/biz"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGatewayService)

type GatewayService struct {
	pb.UnimplementedContestServer
	contestUc    *biz.ContestUsecase
	problemUc    *biz.ProblemUsecase
	submissionUc *biz.SubmissionUsecase
}

func NewGatewayService(contestUsecase *biz.ContestUsecase, problemUsecase *biz.ProblemUsecase, submissionUsecase *biz.SubmissionUsecase) *GatewayService {
	return &GatewayService{
		contestUc:    contestUsecase,
		problemUc:    problemUsecase,
		submissionUc: submissionUsecase,
	}
}
