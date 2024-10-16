package data

import (
	"context"
	"errors"
	"fmt"
	"sastoj/app/admin/admin/internal/conf"
	"sastoj/ent"
	"sastoj/pkg/file"
	"sastoj/pkg/util"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewProblemCaseRepo, NewContestRepo, NewJudgeRepo, NewAdjudicatorRepo, NewProblemRepo, NewRankRepo, NewUserRepo, NewGroupRepo)

// Data .
type Data struct {
	db    *ent.Client
	redis *redis.Client
	jcm   *file.JudgeConfigManager
	fcm   *file.FcConfigManager
}

func (d *Data) GetConfig(p *ent.Problem) (string, error) {
	var config string
	var err error
	switch p.Edges.ProblemType.Judge {
	case "freshcup":
		config, err = d.fcm.GetConfigString(p.ID)
	case "gojudge":
		config, err = d.jcm.GetConfigString(p.ID)
	}
	if err != nil {
		return "", err
	}
	return config, nil
}

func (d *Data) SetConfig(p *ent.Problem, config string) error {
	var err error
	switch p.Edges.ProblemType.Judge {
	case "freshcup":
		err = d.fcm.SetConfigString(p.ID, config)
	case "gojudge":
		err = d.jcm.SetConfigString(p.ID, config)
	}
	return err
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
	return &Data{
		db:    client,
		redis: redisClient,
		jcm:   file.NewJudgeConfigManager(c.Load.GetProblemCasesLocation()),
		fcm:   file.NewFcConfigManager(c.Load.GetProblemCasesLocation()),
	}, cleanup, nil
}
