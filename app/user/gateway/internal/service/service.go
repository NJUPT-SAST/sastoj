package service

import (
	"github.com/google/wire"
	pb "sastoj/api/sastoj/user/gateway/service/v1"
	"sastoj/app/user/gateway/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGatewayService)

type GatewayService struct {
	pb.UnimplementedContestServer
	pb.UnimplementedProblemServer
	pb.UnimplementedSubmissionServer
	contestUc *biz.ContestUsecase
}

func NewGatewayService(contestUsecase *biz.ContestUsecase) *GatewayService {
	return &GatewayService{
		contestUc: contestUsecase,
	}
}
