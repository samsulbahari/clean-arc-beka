package postgresql

import (
	"context"

	drivenrepository "grpc-beka/interface_adapter/gateway/repository/driven_repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

const url_pg = "user=postgres password=sintara23 host=localhost port=5432 dbname=cbs  pool_max_conns=10"

type pool_handler struct {
	pool *pgxpool.Pool
}

// BeginTx implements persistence.PoolHandler
func (x *pool_handler) BeginTx() (drivenrepository.Tx, error) {
	tx, err := x.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	return NewTx(tx), nil
}

// Exec implements persistence.PoolHandler
func (x *pool_handler) Exec(query string, args ...interface{}) error {
	_, err := x.pool.Exec(context.Background(), query, args...)
	return err
}

// Get implements persistence.PoolHandler
func (x *pool_handler) Get(query string, args ...interface{}) drivenrepository.Row {
	return new_pg_row(x.pool.QueryRow(context.Background(), query, args...))
}

// Select implements persistence.PoolHandler
func (x *pool_handler) Select(query string, args ...interface{}) (drivenrepository.Rows, error) {
	rows, err := x.pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	return new_pg_rows(rows), nil
}

func NewPool() (drivenrepository.PoolHandler, error) {
	poolx, err := pgxpool.New(context.Background(), url_pg)
	if err != nil {
		return nil, err
	}

	return &pool_handler{
		pool: poolx,
	}, nil
}
