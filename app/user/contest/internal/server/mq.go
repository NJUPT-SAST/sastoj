package server

import (
	"context"
	"sastoj/app/user/contest/internal/conf"
	"sastoj/app/user/contest/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/rabbitmq"
)

const ProblemExchangeName = "update_problem"

func NewMQServer(c *conf.Server, s *service.ContestService, logger log.Logger) *rabbitmq.Server {
	var opts = []rabbitmq.ServerOption{
		rabbitmq.WithAddress([]string{c.Mq}),
		rabbitmq.WithCodec("json"),
		rabbitmq.WithExchange(ProblemExchangeName, true),
	}

	srv := rabbitmq.NewServer(opts...)

	_ = rabbitmq.RegisterSubscriber(srv, context.Background(), ProblemExchangeName, s.UpdateProblemHandle)

	return srv
}
