package cleanservice

import "grpc-beka/domain"

type RefSysAgamaPort struct {
	rs RefSysAgamas
}

func (x *RefSysAgamaPort) Convert(data []domain.RefSysAgama) {

	result := make([]*RefSysAgama, 0)
	var ref_sys_agama RefSysAgama
	for _, v := range data {

		ref_sys_agama.Id = v.Id
		ref_sys_agama.Keterangan = v.Keterangan
		ref_sys_agama.Kode = v.Kode
		result = append(result, &ref_sys_agama)
	}

	x.rs.Data = result
	x.rs.Message = "success"
}

func NewRefSysAgamaPortOutPort() RefSysAgamaPort {
	return RefSysAgamaPort{}
}
