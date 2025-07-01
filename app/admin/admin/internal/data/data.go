package data

import (
	"context"
	"errors"
	"fmt"
	"sastoj/app/admin/admin/internal/conf"
	"sastoj/ent"
	"sastoj/pkg/file"
	"sastoj/pkg/mq"
	"sastoj/pkg/util"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewProblemCaseRepo, NewContestRepo, NewJudgeRepo, NewAdjudicatorRepo, NewProblemRepo, NewRankRepo, NewUserRepo, NewGroupRepo, NewMqRepo)

// Data .
type Data struct {
	db    *ent.Client
	redis *redis.Client
	ex    *mq.OjExchange
	jcm   *file.JudgeConfigManager
	mqc   *mq.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
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
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}
	// Set default root group and root admin if it does not exist in the database.
	rootGroup, err := util.InsertDefaultGroup(context.TODO(), client)
	if err != nil {
		return nil, nil, err
	}
	if !util.ExistRootUser(context.TODO(), client) {
		if c.Auth == nil || c.Auth.RootName == "" || c.Auth.RootPassword == "" {
			return nil, nil, errors.New("unable to get root admin in database and unable to read root admin in the config. please have at least one root admin")
		}
		err = util.InsertDefaultUser(context.TODO(), client, rootGroup, &c.Auth.RootName, &c.Auth.RootPassword)
		if err != nil {
			return nil, nil, err
		}
	}
	// connect to redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		DB:       int(c.Redis.Db),
		Password: c.Redis.Password,
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Errorf("failed connecting to redis: %v", err)
		return nil, nil, err
	}
	// connect to mq
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s", c.Mq.User, c.Mq.Passwd, c.Mq.Url, c.Mq.AmqpPort)
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Errorf("failed connecting to mq: %v", err)
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("failed opening a channel")
		return nil, nil, err
	}
	ex, err := mq.NewExchange(ch, "update_problem")
	if err != nil {
		log.Errorf("failed creating exchange")
		return nil, nil, err
	}
	return &Data{
		db:    client,
		redis: redisClient,
		ex:    ex,
		jcm:   file.NewJudgeConfigManager(c.Load.GetProblemCasesLocation()),
		mqc:   mq.NewClient(c.Mq.Url+":"+c.Mq.HttpPort, c.Mq.User, c.Mq.Passwd),
	}, cleanup, nil
}
