package drivenout_port

import "grpc-beka/domain"

type RefSysAgamaOutPort interface {
	Convert(data []domain.RefSysAgama)
}
