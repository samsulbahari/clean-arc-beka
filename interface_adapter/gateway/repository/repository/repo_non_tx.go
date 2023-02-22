package repository

import (
	"grpc-beka/app/repository"

	drivenrepository "grpc-beka/interface_adapter/gateway/repository/driven_repository"
)

type repo_non_tx struct {
	repo repository.Repo
}

// GetUser implements repository.RepoNonTx
func (x *repo_non_tx) GetData() repository.RefSysAgama {
	return x.repo.GetData()
}

func new_repo_non_tx(pool drivenrepository.PoolHandler) repository.RepoNonTx {
	return &repo_non_tx{
		repo: new_repo(false, nil, pool),
	}
}
