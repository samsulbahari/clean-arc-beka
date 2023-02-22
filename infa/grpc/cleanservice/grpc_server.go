package cleanservice

import (
	context "context"
	drivenrepository "grpc-beka/interface_adapter/gateway/repository/driven_repository"

	"grpc-beka/interface_adapter/controllers"
	drivenlog "grpc-beka/interface_adapter/gateway/utilx/driven_log"

	servicesoutport "grpc-beka/interface_adapter/presenter/grpc_presenter/outport"
)

type grpc_server struct {
	controllers *controllers.Controllers
	UnimplementedRefSysAgamaServerServer
}

// ChangePassword implements CleanArchServer
func (x *grpc_server) GetData(context.Context, *Empty) (*RefSysAgamas, error) {
	o := NewRefSysAgamaPortOutPort()

	err := x.controllers.GetData().Process(servicesoutport.NewRefSysOutPort(&o))
	if err != nil {
		return nil, err
	}

	return &o.rs, nil
}

func NewGRPCService(l drivenlog.DrivenLog, db_hander drivenrepository.PoolHandler) RefSysAgamaServerServer {

	return &grpc_server{
		controllers: controllers.NewControllers(l, db_hander),
	}
}
