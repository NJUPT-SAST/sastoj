package data

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"sastoj/app/judge/freshcup/internal/conf"
	"sastoj/app/judge/freshcup/internal/server"
	"sastoj/ent"
	"sastoj/ent/problemtype"
	"sastoj/pkg/file"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewSubmissionRepo)

// Data .
type Data struct {
	db     *ent.Client
	redis  *redis.Client
	fm     *file.FcConfigManager
	logger *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	ctx := context.Background()
	// New logHelper
	logHelper := log.NewHelper(logger)

	// Connect to database
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	drv, err := sql.Open(
		c.Database.Driver,
		c.Database.Source,
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
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	// Check if problemTypes are created
	exists, err := client.ProblemType.Query().Where(
		problemtype.SlugName("freshcup-single-choice"),
		problemtype.SlugName("freshcup-multiple-choice"),
		problemtype.SlugName("freshcup-short-answer"),
	).Exist(ctx)
	if err != nil {
		log.Errorf("failed checking problemTypes: %v", err)
		return nil, nil, err
	}
	if !exists {
		_, err := client.ProblemType.Create().
			SetSlugName("freshcup-single-choice").
			SetDisplayName("Single-Choice").
			SetDescription("Single Choice Problem powered by Freshcup").
			SetSubmissionChannelName(server.SubmissionQueueName).
			SetSelfTestChannelName(server.SelfTestQueueName).
			SetJudge("freshcup").
			Save(ctx)
		if err != nil {
			log.Errorf("failed creating problemTypes: %v", err)
			return nil, nil, err
		}
		_, err = client.ProblemType.Create().
			SetSlugName("freshcup-multiple-choice").
			SetDisplayName("Multiple-Choice").
			SetDescription("Multiple Choice Problem powered by Freshcup").
			SetSubmissionChannelName(server.SubmissionQueueName).
			SetSelfTestChannelName(server.SelfTestQueueName).
			SetJudge("freshcup").
			Save(ctx)
		if err != nil {
			log.Errorf("failed creating problemTypes: %v", err)
			return nil, nil, err
		}
		_, err = client.ProblemType.Create().
			SetSlugName("freshcup-short-answer").
			SetDisplayName("Short-Answer").
			SetDescription("Short Answer Problem powered by Freshcup").
			SetSubmissionChannelName(server.SubmissionQueueName).
			SetSelfTestChannelName(server.SelfTestQueueName).
			SetJudge("freshcup").
			Save(ctx)
		if err != nil {
			log.Errorf("failed creating problemTypes: %v", err)
			return nil, nil, err
		}
	}
	log.Errorf("failed checking problemTypes: %v", err)

	// connect to redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
		DB:   int(c.Redis.Db),
	})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Errorf("failed connecting to redis: %v", err)
		return nil, nil, err
	}

	// Create a file manager
	fm := file.NewFcConfigManager(c.Load.ProblemCasesLocation)

	return &Data{
		db:     client,
		redis:  redisClient,
		fm:     fm,
		logger: logHelper,
	}, cleanup, nil
}
