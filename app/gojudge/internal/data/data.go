package data

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	pbc "sastoj/api/sastoj/gojudge/judger/gojudge/v1"
	"sastoj/ent"

	"sastoj/app/gojudge/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data data
type Data struct {
	Ch              *amqp.Channel
	Ent             *ent.Client
	Redis           *redis.Client
	FileManage      *FileManage
	Client          *pbc.ExecutorClient
	Logger          *log.Helper
	Commands        *Commands
	JudgeMiddleware *conf.JudgeMiddleware
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	logHelper := log.NewHelper(log.With(logger, "module", "data"))

	logHelper.Infof("Read Go-Judge Conf Success: Endpoint=%s", c.JudgeMiddleware.Endpoint)
	//commands to judge
	command := NewCommands(
		c.JudgeMiddleware.Language.Enable,
		c.JudgeMiddleware.Language.Compile,
		c.JudgeMiddleware.Language.Run,
		c.JudgeMiddleware.Language.Source,
		c.JudgeMiddleware.Language.Target,
		c.JudgeMiddleware.Language.ExecConfig,
	)
	logHelper.Infof("Read Command Conf Success")
	for ck, cv := range command {
		logHelper.Infof("Language=%v Config=%v", ck, cv)
	}
	//conn go-judge
	ClientConn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(c.JudgeMiddleware.Endpoint),
		grpc.WithHealthCheck(false))
	if err != nil {
		logHelper.Errorf("failed creating go-judge clients: %v", err)
	}
	exec := pbc.NewExecutorClient(ClientConn)

	// conn ch by amqp
	MqConn, err := amqp.Dial(c.Data.Mq)
	if err != nil {
		logHelper.Errorf("failed connecting to mq: %v", err)
	}
	ch, err := MqConn.Channel()
	if err != nil {
		logHelper.Errorf("failed opening a channel")
	}
	logHelper.Infof("MQ Channel run Success: %v", ch)

	// connect to redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: c.Data.Redis.Addr,
		DB:   int(c.Data.Redis.Db),
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Errorf("failed connecting to redis: %v", err)
		return nil, nil, err
	}

	// ent client
	drv, err := sql.Open(
		c.Data.Database.Driver,
		c.Data.Database.Source,
	)
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		log.WithContext(ctx, logger)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		logHelper.Errorf("failed opening connection to database: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		logHelper.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}
	logHelper.Infof("ent run Success: %v", ch)

	// func tp cleanup resources
	cleanup := func() {
		//TODO delay until judge was closed
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		Ch:              ch,
		Ent:             client,
		Redis:           redisClient,
		FileManage:      &FileManage{FileLocation: c.Data.Load.ProblemCasesLocation},
		Client:          &exec,
		Commands:        &command,
		Logger:          log.NewHelper(logger),
		JudgeMiddleware: c.JudgeMiddleware,
	}, cleanup, nil
}
