package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/gfdmit/subscription-service/config"
	"github.com/gfdmit/subscription-service/internal/model"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct {
	pool *pgxpool.Pool
}

func New(conf config.Config, ctx context.Context) (*postgresRepository, error) {
	connString := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v?sslmode=%v",
		conf.Postgres.User,
		url.QueryEscape(conf.Postgres.Pass),
		conf.Postgres.Host,
		conf.Postgres.Port,
		conf.Postgres.DB,
		conf.Postgres.SSL,
	)
	poolConf, err := initPoolConf(conf, connString)
	if err != nil {
		return nil, fmt.Errorf("initPoolConf: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConf)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pool.Ping: %w", err)
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("postgers.WithInstance: %v", err)
	}

	migrations := fmt.Sprintf("file://%v", conf.Migrations)
	m, err := migrate.NewWithDatabaseInstance(migrations, conf.DB, driver)
	if err != nil {
		return nil, fmt.Errorf("migrate.NewWithDatabaseInstance: %v", err)
	}
	defer m.Close()

	log.Println("applying migrations...")
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("nothing to migrate")
		} else {
			return nil, fmt.Errorf("error when migrating: %v", err)
		}
	} else {
		log.Println("migrated successfully!")
	}

	return &postgresRepository{
		pool: pool,
	}, nil
}

func (pr postgresRepository) CreateSubscription(ctx context.Context, subscription model.Subscription) (*model.Subscription, error) {
	var (
		subs model.Subscription
	)
	row := pr.pool.QueryRow(
		ctx,
		createSubscriptionQuery,
		subscription.ServiceName,
		subscription.Price,
		subscription.UserID,
		subscription.StartDate,
		subscription.EndDate,
	)
	err := row.Scan(
		&subs.ID,
		&subs.ServiceName,
		&subs.Price,
		&subs.UserID,
		&subs.StartDate,
		&subs.EndDate,
	)
	if err != nil {
		return nil, fmt.Errorf("row.Scan: %v", err)
	}
	return &subs, nil
}

func (pr postgresRepository) GetSubscription(ctx context.Context, id int) (*model.Subscription, error) {
	subscription := model.Subscription{}
	row := pr.pool.QueryRow(
		ctx,
		getSubscriptionQuery,
		id,
	)
	err := row.Scan(
		&subscription.ID,
		&subscription.ServiceName,
		&subscription.Price,
		&subscription.UserID,
		&subscription.StartDate,
		&subscription.EndDate,
	)
	if err != nil {
		return nil, fmt.Errorf("row.Scan: %v", err)
	}
	return &subscription, nil
}

func (pr postgresRepository) GetSubscriptions(ctx context.Context) ([]model.Subscription, error) {
	subscriptions := []model.Subscription{}
	rows, err := pr.pool.Query(
		ctx,
		getSubscriptionsQuery,
	)
	if err != nil {
		return nil, fmt.Errorf("pool.Query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var subscription model.Subscription
		err := rows.Scan(
			&subscription.ID,
			&subscription.ServiceName,
			&subscription.Price,
			&subscription.UserID,
			&subscription.StartDate,
			&subscription.EndDate,
		)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %v", err)
		}
		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions, nil
}

func (pr postgresRepository) UpdateSubscription(ctx context.Context, id int, subscription model.Subscription) (*model.Subscription, error) {
	pr.pool.QueryRow(
		ctx,
		updateSubscriptionQuery,
		subscription.ServiceName,
		subscription.Price,
		subscription.UserID,
		subscription.StartDate,
		subscription.EndDate,
		id,
	)
	subscription.ID = strconv.Itoa(id)
	return &subscription, nil
}

func (pr postgresRepository) DeleteSubscription(ctx context.Context, id int) (bool, error) {
	pr.pool.QueryRow(
		ctx,
		deleteSubscriptionQuery,
		id,
	)
	return true, nil
}

func initPoolConf(conf config.Config, connString string) (*pgxpool.Config, error) {
	poolConf, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig: %v", err)
	}
	poolConf.MaxConns = conf.MaxConns
	poolConf.MinConns = conf.MinConns
	poolConf.MaxConnLifetime = conf.MaxLifetime
	poolConf.HealthCheckPeriod = conf.HealthCheck
	return poolConf, err
}
