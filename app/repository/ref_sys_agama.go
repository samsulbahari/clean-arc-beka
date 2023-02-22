package repository

import "grpc-beka/domain"

type RefSysAgama interface {
	GetData() ([]domain.RefSysAgama, error)
}
