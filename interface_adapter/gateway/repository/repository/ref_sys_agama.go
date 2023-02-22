package repository

import (
	"grpc-beka/app/repository"

	"grpc-beka/domain"

	drivenrepository "grpc-beka/interface_adapter/gateway/repository/driven_repository"
)

type ref_sys_agama struct {
	is_tx bool
	tx    drivenrepository.Tx
	pool  drivenrepository.PoolHandler
}

const query = `SELECT id, nomor,kode,keterangan
FROM ref_sys_agama;`

// GetByName implements repository.User
func (x *ref_sys_agama) GetData() ([]domain.RefSysAgama, error) {
	var row drivenrepository.Rows

	if x.is_tx {
		row, _ = x.tx.Select(query)

	} else {
		row, _ = x.pool.Select(query)
	}

	result := make([]domain.RefSysAgama, 0)

	for row.Next() {
		t := new(domain.RefSysAgama)
		err := row.Scan(&t.Id,
			&t.Nomor,
			&t.Kode,
			&t.Keterangan,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, *t)
	}

	return result, nil
}

func get_data(is_tx bool, tx drivenrepository.Tx, pool_handler drivenrepository.PoolHandler) repository.RefSysAgama {
	return &ref_sys_agama{
		is_tx: is_tx,
		tx:    tx,
		pool:  pool_handler,
	}
}
