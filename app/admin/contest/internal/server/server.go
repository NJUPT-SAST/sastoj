package server

import (
	v1 "sastoj/api/sastoj/admin/contest/service/v1"
	"sastoj/pkg/middleware/auth"

	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewCronWorker)

// apiMap is a map of method to group for auth middleware.
var apiMap = map[string]string{
	v1.Contest_GetContest_FullMethodName:  auth.UserGroup,
	v1.Contest_ListContest_FullMethodName: auth.UserGroup,
}
