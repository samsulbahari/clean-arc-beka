package postgresql

import (
	drivenrepository "grpc-beka/interface_adapter/gateway/repository/driven_repository"

	"github.com/jackc/pgx/v5"
)

type pg_rows struct {
	rows pgx.Rows
}

// Close implements persistence.Rows
func (x *pg_rows) Close() {
	x.rows.Close()
}

// Next implements persistence.Rows
func (x *pg_rows) Next() bool {
	return x.rows.Next()
}

// Scan implements persistence.Rows
func (x *pg_rows) Scan(dest ...any) error {
	return x.rows.Scan(dest...)
}

func new_pg_rows(rows pgx.Rows) drivenrepository.Rows {
	return &pg_rows{
		rows: rows,
	}
}
