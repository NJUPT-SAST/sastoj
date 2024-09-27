package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewCaseUsecase, NewContestUsecase, NewJudgeUsecase, NewJudgerUsecase, NewProblemUsecase, NewRankUsecase, NewUserUsecase, NewGroupUsecase)