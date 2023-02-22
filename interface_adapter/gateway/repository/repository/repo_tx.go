package repository

import (
	"grpc-beka/app/repository"

	drivenrepository "grpc-beka/interface_adapter/gateway/repository/driven_repository"
)

type repo_tx struct {
	tx   drivenrepository.Tx
	repo repository.Repo
}

func (x *repo_tx) GetData() repository.RefSysAgama {
	return x.repo.GetData()
}

// Commit implements repository.RepoTx
func (x *repo_tx) Save() error {
	return x.tx.Commit()
}

// Rollback implements repository.RepoTx
func (x *repo_tx) Discard() error {
	return x.tx.Rollback()
}

func new_repo_tx(tx drivenrepository.Tx) repository.RepoTx {
	return &repo_tx{
		tx:   tx,
		repo: new_repo(true, tx, nil),
	}
}
