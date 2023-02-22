package ref_sys_agama

import (
	"grpc-beka/app/repository"

	"grpc-beka/app/utilx"
)

type interactor struct {
	logx         utilx.LogX
	repo_factory repository.RepoFactory
}

// Process implements Caller
func (x *interactor) Process(out OutX) error {

	repo := x.repo_factory.GetRepoNonTx()

	user, err := repo.GetData().GetData()
	if err != nil {
		x.logx.WriteLn(err)
		return err
	}

	out.Convert(user)

	return nil
}

func New(logx utilx.LogX, repo repository.RepoFactory) Caller {
	return &interactor{
		logx:         logx,
		repo_factory: repo,
	}
}
