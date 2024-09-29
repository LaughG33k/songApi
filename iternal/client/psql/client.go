package psql

import (
	"context"
	"time"

	"github.com/LaughG33k/songApi/iternal/config"
	"github.com/LaughG33k/songApi/pkg"
	"github.com/jackc/pgx"
)

type PsqlClient interface {
	ExecEx(ctx context.Context, sql string, options *pgx.QueryExOptions, arguments ...interface{}) (commandTag pgx.CommandTag, err error)
	QueryEx(ctx context.Context, sql string, options *pgx.QueryExOptions, args ...interface{}) (*pgx.Rows, error)
	QueryRowEx(ctx context.Context, sql string, options *pgx.QueryExOptions, args ...interface{}) *pgx.Row
	BeginEx(ctx context.Context, txOptions *pgx.TxOptions) (*pgx.Tx, error)
}

func NewClient(ctx context.Context, cfg config.DBCfg) (*pgx.ConnPool, error) {

	var pool *pgx.ConnPool

	err := pkg.Retry(func() error {

		conn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
			ConnConfig: pgx.ConnConfig{
				Host:     cfg.Host,
				Port:     cfg.Port,
				User:     cfg.User,
				Password: cfg.Password,
				Database: cfg.DB,
			},
			MaxConnections: 100,
		})

		if err != nil {
			return err
		}

		pool = conn

		return nil
	}, 5, 1*time.Second)

	if err != nil {
		return nil, err
	}
	return pool, nil
}
