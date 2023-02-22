package outport

import (
	sys_agama "grpc-beka/app/services/ref_sys_agama"
	"grpc-beka/domain"

	drivenoutport "grpc-beka/interface_adapter/presenter/grpc_presenter/driven_outport"
)

type ref_sys_agama_outport struct {
	p drivenoutport.RefSysAgamaOutPort
}

// Convert implements changepassword.OutX
func (x *ref_sys_agama_outport) Convert(ref []domain.RefSysAgama) {

	x.p.Convert(ref)

}

func NewRefSysOutPort(p drivenoutport.RefSysAgamaOutPort) sys_agama.OutX {
	return &ref_sys_agama_outport{
		p: p,
	}
}
