package service

import (
	"github.com/google/wire"
	pb "sastoj/api/sastoj/admin/admin/service/v1"
	"sastoj/app/admin/admin/internal/biz"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	ctsc  *biz.ContestUsecase
	casec *biz.CaseUsecase
	rc    *biz.RankUsecase
	jc    *biz.JudgeUsecase
	jrc   *biz.JudgerUsecase
	uc    *biz.UserUsecase
	gc    *biz.GroupUsecase
	pc    *biz.ProblemUsecase
}

func NewAdminService(problemCase *biz.CaseUsecase, contest *biz.ContestUsecase, rank *biz.RankUsecase, judge *biz.JudgeUsecase, judger *biz.JudgerUsecase, user *biz.UserUsecase, group *biz.GroupUsecase, problem *biz.ProblemUsecase) *AdminService {
	return &AdminService{
		casec: problemCase,
		ctsc:  contest,
		rc:    rank,
		jc:    judge,
		jrc:   judger,
		uc:    user,
		gc:    group,
		pc:    problem,
	}
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdminService)
