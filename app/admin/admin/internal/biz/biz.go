package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewCaseUsecase, NewContestUsecase, NewJudgeUsecase, NewAdjudicatorUsecase, NewProblemUsecase, NewRankUsecase, NewUserUsecase, NewGroupUsecase)
