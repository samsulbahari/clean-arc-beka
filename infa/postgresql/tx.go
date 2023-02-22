package postgresql

import (
	"context"

	drivenrepository "grpc-beka/interface_adapter/gateway/repository/driven_repository"

	"github.com/jackc/pgx/v5"
)

type pg_tx struct {
	tx pgx.Tx
}

// Commit implements persistence.Tx
func (x *pg_tx) Commit() error {
	return x.tx.Commit(context.Background())
}

// Exec implements persistence.Tx
func (x *pg_tx) Exec(query string, args ...interface{}) error {
	_, err := x.tx.Exec(context.Background(), query, args...)
	return err
}

// Get implements persistence.Tx
func (x *pg_tx) Get(query string, args ...interface{}) drivenrepository.Row {

	return new_pg_row(x.tx.QueryRow(context.Background(), query, args...))

}

// Rollback implements persistence.Tx
func (x *pg_tx) Rollback() error {
	return x.tx.Rollback(context.Background())
}

// Select implements persistence.Tx
func (x *pg_tx) Select(query string, args ...interface{}) (drivenrepository.Rows, error) {
	rows, err := x.tx.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	return new_pg_rows(rows), nil
}

func NewTx(tx pgx.Tx) drivenrepository.Tx {
	return &pg_tx{
		tx: tx,
	}
}
