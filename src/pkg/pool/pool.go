package pool

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

var postgresClient *pgxpool.Pool

func InitPool(ctx context.Context, config *Config) error {
	var err error
	postgresClient, err = pgxpool.Connect(ctx, config.CreateString())
	if err != nil {
		return err
	}

	if err = postgresClient.Ping(ctx); err != nil {
		return err
	}
	return nil
}

func GetConnection() (*pgxpool.Pool, error) {
	if postgresClient == nil {
		err := InitPool(context.Background(), NewConfig())
		if err != nil {
			return nil, err
		}
	}
	return postgresClient, nil
}

func ClosePool(ctx context.Context) error {
	if postgresClient != nil {
		postgresClient.Close()
		return nil
	}

	return fmt.Errorf("pool is nil")
}
