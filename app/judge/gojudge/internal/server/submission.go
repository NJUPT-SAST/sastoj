package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	rabbitmqBroker "github.com/tx7do/kratos-transport/broker/rabbitmq"
	"github.com/tx7do/kratos-transport/transport/rabbitmq"
	"sastoj/app/judge/gojudge/internal/conf"
	"sastoj/app/judge/gojudge/internal/service"
)

const (
	SubmissionQueueName = "gojudge-submission"
	SelfTestQueueName   = "gojudge-self-test"
)

func NewSubmissionServer(c *conf.Server, s *service.SubmissionService, logger log.Logger) *rabbitmq.Server {
	var opts = []rabbitmq.ServerOption{
		rabbitmq.WithAddress([]string{c.Mq}),
		rabbitmq.WithCodec("json"),
	}

	srv := rabbitmq.NewServer(opts...)

	_ = rabbitmq.RegisterSubscriber(srv, context.Background(), SubmissionQueueName, s.SubmissionHandle, broker.WithQueueName(SubmissionQueueName), rabbitmqBroker.WithDurableQueue())
	_ = rabbitmq.RegisterSubscriber(srv, context.Background(), SelfTestQueueName, s.SelfTestHandle, broker.WithQueueName(SelfTestQueueName), rabbitmqBroker.WithDurableQueue())

	return srv
}
