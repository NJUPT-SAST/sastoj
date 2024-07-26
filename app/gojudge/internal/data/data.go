package data

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	amqp "github.com/rabbitmq/amqp091-go"
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
	FileManage      *FileManage
	Client          *pbc.ExecutorClient
	Logger          *log.Helper
	Commands        *Commands
	JudgeMiddleware *conf.JudgeMiddleware
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {

	logHelper := log.NewHelper(log.With(logger, "module", "data"))

	//commands to judge
	command := NewCommands(
		c.JudgeMiddleware.Language.Enable,
		c.JudgeMiddleware.Language.Compile,
		c.JudgeMiddleware.Language.Run,
		c.JudgeMiddleware.Language.Source,
		c.JudgeMiddleware.Language.Target,
		c.JudgeMiddleware.Language.ExecConfig,
	)
	fmt.Println("Read Command Conf Success")
	for ck, cv := range command {
		fmt.Printf("Language %v:  %v\n", ck, cv)
	}

	// conn ch by amqp
	MqConn, err := amqp.Dial(c.Data.Mq)
	if err != nil {
		logHelper.Errorf("failed connecting to mq: %v", err)
	}
	ch, err := MqConn.Channel()
	if err != nil {
		logHelper.Errorf("failed opening a channel")
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

	middle := c.JudgeMiddleware
	logHelper.Infof("scanned judge config: judge.endpoint=%s judge.compile=%v judge.run=%v", middle.Endpoint, middle.Language.Compile, middle.Language.Run)

	//conn
	ClientConn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(middle.Endpoint),
		grpc.WithHealthCheck(false))

	if err != nil {
		logHelper.Errorf("failed creating clients: %v", err)
	}
	exec := pbc.NewExecutorClient(ClientConn)

	// func tp cleanup resources
	cleanup := func() {
		//TODO delay until judge was closed
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		Ch:              ch,
		Ent:             client,
		FileManage:      &FileManage{FileLocation: c.Data.Load.ProblemCasesLocation},
		Client:          &exec,
		Commands:        &command,
		Logger:          log.NewHelper(logger),
		JudgeMiddleware: c.JudgeMiddleware,
	}, cleanup, nil
}
