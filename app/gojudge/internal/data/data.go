package data

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
	Ch           *amqp.Channel
	Ent          *ent.Client
	FileLocation string
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {

	// conn ch by amqp
	MqConn, err := amqp.Dial(c.Data.Mq)
	if err != nil {
		log.Errorf("failed connecting to mq: %v", err)
	}
	ch, err := MqConn.Channel()
	if err != nil {
		log.Errorf("failed opening a channel")
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
		log.Errorf("failed opening connection to database: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	// func tp cleanup resources
	cleanup := func() {
		//TODO delay until judge was closed
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		Ch:           ch,
		Ent:          client,
		FileLocation: c.Data.Load.ProblemCasesLocation,
	}, cleanup, nil
}
