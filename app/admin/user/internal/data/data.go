package data

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"sastoj/app/admin/user/internal/conf"
	"sastoj/ent"
	"sastoj/pkg/util"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	db *ent.Client
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
	// Set default root group and root user if it does not exist in the database.
	rootGroup, err := util.InsertDefaultGroup(context.TODO(), client)
	if err != nil {
		return nil, nil, err
	}
	if !util.ExistRootUser(context.TODO(), client) {
		if c.Auth == nil || c.Auth.RootName == "" || c.Auth.RootPassword == "" {
			return nil, nil, errors.New("unable to get root user in database and unable to read root user in the config. please have at least one root user")
		}
		err = util.InsertDefaultUser(context.TODO(), client, rootGroup, &c.Auth.RootName, &c.Auth.RootPassword)
		if err != nil {
			return nil, nil, err
		}
	}
	return &Data{db: client}, cleanup, nil
}
