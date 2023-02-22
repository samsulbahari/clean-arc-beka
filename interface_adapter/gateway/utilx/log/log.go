package log

import (
	drivenlog "grpc-beka/interface_adapter/gateway/utilx/driven_log"

	"grpc-beka/app/utilx"
)

type logx struct {
	l drivenlog.DrivenLog
}

// WriteLn implements utilx.LogX
func (x *logx) WriteLn(err error) {
	x.l.WriteLn(err)
}

func NewLogx(l drivenlog.DrivenLog) utilx.LogX {
	return &logx{
		l: l,
	}
}
