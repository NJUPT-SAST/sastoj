package service

import (
	"context"
	v1 "sastoj/api/sastoj/admin/admin/service/v1"
)

// GetMQQueueStats retrieves the status of RabbitMQ queues
func (s *AdminService) GetMQQueueStats(ctx context.Context, req *v1.GetMQQueueStatsRequest) (*v1.GetMQQueueStatsReply, error) {
	queueStats, err := s.mqc.GetQueueStats(ctx)
	if err != nil {
		return nil, err
	}

	// Convert queue stats to proto messages
	pbStats := make([]*v1.GetMQQueueStatsReply_QueueStats, 0, len(queueStats))
	for _, q := range queueStats {
		pbStat := &v1.GetMQQueueStatsReply_QueueStats{
			QueueName:       q.Name,
			MessagesReady:   q.MessagesReady,
			MessagesUnacked: q.MessagesUnacked,
			Consumers:       int64(q.Consumers),
			Memory:          q.Memory,
			MessageRateIn:   int64(q.MessageRateIn),
			MessageRateOut:  int64(q.MessageRateOut),
			MessageRateAck:  int64(q.MessageRateAck),
		}
		pbStats = append(pbStats, pbStat)
	}

	return &v1.GetMQQueueStatsReply{
		Stats: pbStats,
	}, nil
}
