package data

import (
	"context"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent/problemtype"
	"sastoj/pkg/mq"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
)

// mqRepo is the implementation of biz.MQRepo
type mqRepo struct {
	data   *Data
	logger *log.Helper
}

// NewMqRepo creates a new MQ repository
func NewMqRepo(data *Data, logger log.Logger) biz.MQRepo {
	return &mqRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

// GetQueueStats retrieves statistics for all queues from RabbitMQ
func (r *mqRepo) GetQueueStats(ctx context.Context) ([]mq.QueueStats, error) {
	judges, err := r.data.db.ProblemType.Query().Select(problemtype.FieldJudge).Unique(true).All(ctx)
	if err != nil {
		return nil, err
	}
	queueStats, err := r.data.mqc.GetQueueStats(ctx)
	if err != nil {
		return nil, err
	}
	var filteredStats []mq.QueueStats
	for _, stat := range queueStats {
		for _, judge := range judges {
			if judge.Judge != "" && strings.HasPrefix(stat.Name, judge.Judge) {
				filteredStats = append(filteredStats, stat)
				break
			}
		}
	}
	return filteredStats, nil

}

// GetMQOverview retrieves overview statistics from RabbitMQ
func (r *mqRepo) GetMQOverview(ctx context.Context) (*mq.RabbitMQOverview, error) {
	return r.data.mqc.GetMQOverview(ctx)
}

// GetQueueStatsForVHost retrieves statistics for all queues in a specific virtual host
func (r *mqRepo) GetQueueStatsForVHost(ctx context.Context, vhost string) ([]mq.QueueStats, error) {
	return r.data.mqc.GetQueueStatsForVHost(ctx, vhost)
}
