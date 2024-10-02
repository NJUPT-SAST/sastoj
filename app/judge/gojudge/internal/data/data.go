package data

import (
	"context"
	"fmt"
	pbc "sastoj/api/sastoj/gojudge/judger/gojudge/v1"
	"sastoj/app/judge/gojudge/internal/conf"
	"sastoj/app/judge/gojudge/pkg/gojudge"
	"sastoj/ent"
	"sastoj/ent/problemtype"
	"sastoj/pkg/file"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewSubmissionRepo)

// Data .
type Data struct {
	db      *ent.Client
	redis   *redis.Client
	gojudge *gojudge.GoJudge
	fm      *file.Manager
	logger  *log.Helper
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
	exists, err := client.ProblemType.Query().Where(problemtype.SlugNameEQ("gojudge-classic-algo")).Exist(ctx)
	if err != nil {
		log.Errorf("failed checking problemTypes: %v", err)
	}
	if !exists {
		_, err := client.ProblemType.Create().
			SetSlugName("gojudge-classic-algo").
			SetDisplayName("Classic-Algo").
			SetDescription("Classic Algo Problem powered by Gojudge").
			SetJudge("gojudge").
			Save(ctx)
		if err != nil {
			log.Errorf("failed creating problemTypes: %v", err)
			return nil, nil, err
		}
	}

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
	fm := file.NewManager(c.Load.ProblemCasesLocation)

	// Create a go-judge client
	ClientConn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint(c.Judge.Endpoint),
		grpc.WithHealthCheck(false))
	if err != nil {
		logHelper.Errorf("failed creating go-judge clients: %v", err)
	}
	exec := pbc.NewExecutorClient(ClientConn)

	// Commands
	envs := make(map[string][]string)
	for k, v := range c.Judge.Language.Env {
		envs[k] = v.Env
	}

	commands := gojudge.NewCommands(
		c.Judge.Language.Enable,
		envs,
		c.Judge.Language.Compile,
		c.Judge.Language.Run,
		c.Judge.Language.Source,
		c.Judge.Language.Target,
		c.Judge.Language.ExecConfig,
	)
	return &Data{
		db:    client,
		redis: redisClient,
		gojudge: &gojudge.GoJudge{
			Client:   &exec,
			Commands: &commands,
		},
		fm:     fm,
		logger: logHelper,
	}, cleanup, nil
}
