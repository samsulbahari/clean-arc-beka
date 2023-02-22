package postgresql

import (
	drivenrepository "grpc-beka/interface_adapter/gateway/repository/driven_repository"

	"github.com/jackc/pgx/v5"
)

type pg_row struct {
	row pgx.Row
}

// Scan implements persistence.Row
func (x *pg_row) Scan(dest ...any) error {
	return x.row.Scan(dest...)
}

func new_pg_row(data pgx.Row) drivenrepository.Row {
	return &pg_row{
		row: data,
	}
}
