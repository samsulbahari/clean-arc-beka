package controllers

import (
	/*initapp "github.com/clean_arch/app/init_app"*/
	"grpc-beka/app/repository"
	ref_sys "grpc-beka/app/services/ref_sys_agama"
	"grpc-beka/app/utilx"
	drivenrepository "grpc-beka/interface_adapter/gateway/repository/driven_repository"
	rep "grpc-beka/interface_adapter/gateway/repository/repository"
	drivenlog "grpc-beka/interface_adapter/gateway/utilx/driven_log"
	"grpc-beka/interface_adapter/gateway/utilx/log"
)

type Controllers struct {
	l    utilx.LogX
	repo repository.RepoFactory
}

// ChangePassword implements initapp.Services
func (x *Controllers) GetData() ref_sys.Caller {
	return ref_sys.New(x.l, x.repo)
}

func NewControllers(l drivenlog.DrivenLog, db_handler drivenrepository.PoolHandler) *Controllers /*initapp.Services*/ {

	return &Controllers{
		l:    log.NewLogx(l),
		repo: rep.NewRepoFactory(db_handler),
	}
}
