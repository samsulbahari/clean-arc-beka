package repository

import (
	"grpc-beka/app/repository"

	drivenrepository "grpc-beka/interface_adapter/gateway/repository/driven_repository"
)

type repo_factory struct {
	pool drivenrepository.PoolHandler
}

// GetRepo implements repository.RepoFactory
func (x *repo_factory) GetRepoNonTx() repository.RepoNonTx {
	return new_repo_non_tx(x.pool)
}

// GetRepoTx implements repository.RepoFactory
func (x *repo_factory) GetRepoTx() (repository.RepoTx, error) {
	tx, err := x.pool.BeginTx()
	if err != nil {
		return nil, err
	}

	return new_repo_tx(tx), nil
}

func NewRepoFactory(pool drivenrepository.PoolHandler) repository.RepoFactory {
	return &repo_factory{
		pool: pool,
	}
}
