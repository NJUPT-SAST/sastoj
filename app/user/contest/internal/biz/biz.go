package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewContestUsecase, NewProblemUsecase, NewSubmissionUsecase, NewRegisterUsecase, NewRankUsecase)
