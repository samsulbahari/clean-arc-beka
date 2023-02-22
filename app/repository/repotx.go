package repository

type RepoTx interface {
	Save() error
	Discard() error
	Repo
}
