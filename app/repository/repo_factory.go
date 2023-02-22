package repository

type RepoFactory interface {
	GetRepoTx() (RepoTx, error)
	GetRepoNonTx() RepoNonTx
}
