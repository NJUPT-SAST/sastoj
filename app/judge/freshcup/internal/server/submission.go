package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	rabbitmqBroker "github.com/tx7do/kratos-transport/broker/rabbitmq"
	"github.com/tx7do/kratos-transport/transport/rabbitmq"
	"sastoj/app/judge/freshcup/internal/conf"
	"sastoj/app/judge/freshcup/internal/service"
)

func NewSubmissionServer(c *conf.Server, s *service.SubmissionService, logger log.Logger) *rabbitmq.Server {
	var opts = []rabbitmq.ServerOption{
		rabbitmq.WithAddress([]string{c.Mq}),
		rabbitmq.WithCodec("json"),
	}

	srv := rabbitmq.NewServer(opts...)

	_ = rabbitmq.RegisterSubscriber(srv, context.Background(), "freshcup-submission", s.SubmissionHandle, broker.WithQueueName("freshcup-submission"), rabbitmqBroker.WithDurableQueue())
	_ = rabbitmq.RegisterSubscriber(srv, context.Background(), "freshcup-self-test", s.SelfTestHandle, broker.WithQueueName("freshcup-self-test"), rabbitmqBroker.WithDurableQueue())

	return srv
}
