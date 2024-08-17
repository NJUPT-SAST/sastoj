package job

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewContestJob)

type JobFunc func(ctx context.Context)
