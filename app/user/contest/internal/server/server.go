package server

import (
	v1 "sastoj/api/sastoj/user/contest/service/v1"
	"time"

	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

var apiLimiter = map[string]time.Duration{
	v1.Contest_Submit_FullMethodName:   10 * time.Second,
	v1.Contest_SelfTest_FullMethodName: 5 * time.Second,
}
