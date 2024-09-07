package server

import (
	"context"
	"sastoj/app/judge/gojudge/internal/conf"
	"sastoj/app/judge/gojudge/internal/service"
	"sastoj/pkg/mq"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	rabbitmqBroker "github.com/tx7do/kratos-transport/broker/rabbitmq"
	"github.com/tx7do/kratos-transport/transport/rabbitmq"
)

func NewSubmissionServer(c *conf.Server, s *service.SubmissionService, logger log.Logger) *rabbitmq.Server {
	var opts = []rabbitmq.ServerOption{
		rabbitmq.WithAddress([]string{c.Mq}),
		rabbitmq.WithCodec("json"),
	}

	srv := rabbitmq.NewServer(opts...)

	_ = rabbitmq.RegisterSubscriber(srv, context.Background(), mq.SubmissionQueue, s.SubmissionHandle, broker.WithQueueName(mq.SubmissionQueue), rabbitmqBroker.WithDurableQueue())
	_ = rabbitmq.RegisterSubscriber(srv, context.Background(), mq.SelfTestQueue, s.SelfTestHandle, broker.WithQueueName(mq.SelfTestQueue), rabbitmqBroker.WithDurableQueue())

	return srv
}
