package data

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"sastoj/app/user/gateway/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewContestRepo)

// Data .
type Data struct {
	redis *redis.Client
	ch    *amqp.Channel
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// connect to redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
		DB:   int(c.Redis.Db),
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Errorf("failed connecting to redis: %v", err)
		return nil, nil, err
	}
	// connect to mq
	conn, err := amqp.Dial(c.Mq)
	if err != nil {
		log.Errorf("failed connecting to mq: %v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("failed opening a channel")
	}
	return &Data{
		redis: redisClient,
		ch:    ch,
	}, cleanup, nil
}
