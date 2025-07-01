package biz

import (
	"context"
	"sastoj/pkg/mq"

	"github.com/go-kratos/kratos/v2/log"
)

// MQRepo defines the repository interface for MQ operations
type MQRepo interface {
	GetQueueStats(ctx context.Context) ([]mq.QueueStats, error)
	GetMQOverview(ctx context.Context) (*mq.RabbitMQOverview, error)
	GetQueueStatsForVHost(ctx context.Context, vhost string) ([]mq.QueueStats, error)
}

// MQUsecase is the MQ usecase
type MQUsecase struct {
	repo   MQRepo
	logger *log.Helper
}

// NewMQUsecase creates a new MQ usecase
func NewMQUsecase(repo MQRepo, logger log.Logger) *MQUsecase {
	return &MQUsecase{
		repo:   repo,
		logger: log.NewHelper(logger),
	}
}

// GetQueueStats retrieves statistics for all queues from RabbitMQ
func (uc *MQUsecase) GetQueueStats(ctx context.Context) ([]mq.QueueStats, error) {
	return uc.repo.GetQueueStats(ctx)
}

// GetMQOverview retrieves overview statistics from RabbitMQ
func (uc *MQUsecase) GetMQOverview(ctx context.Context) (*mq.RabbitMQOverview, error) {
	return uc.repo.GetMQOverview(ctx)
}

// GetQueueStatsForVHost retrieves statistics for all queues in a specific virtual host
func (uc *MQUsecase) GetQueueStatsForVHost(ctx context.Context, vhost string) ([]mq.QueueStats, error) {
	return uc.repo.GetQueueStatsForVHost(ctx, vhost)
}
