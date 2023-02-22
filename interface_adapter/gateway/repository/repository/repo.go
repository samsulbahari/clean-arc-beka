package repository

import (
	"grpc-beka/app/repository"

	drivenrepository "grpc-beka/interface_adapter/gateway/repository/driven_repository"
)

type repo struct {
	is_tx bool
	tx    drivenrepository.Tx
	pool  drivenrepository.PoolHandler
}

func (x *repo) GetData() repository.RefSysAgama {
	return get_data(x.is_tx, x.tx, x.pool)
}

func new_repo(is_tx bool, tx drivenrepository.Tx, pool_handler drivenrepository.PoolHandler) repository.Repo {
	return &repo{
		is_tx: is_tx,
		tx:    tx,
		pool:  pool_handler,
	}
}
