package loglumberjack

import (
	"log"

	drivenlog "grpc-beka/interface_adapter/gateway/utilx/driven_log"

	"gopkg.in/natefinch/lumberjack.v2"
)

type logx struct {
	l *log.Logger
}

// WriteLn implements drivenlog.DrivenLog
func (x *logx) WriteLn(err error) {
	x.l.Println(err)
}

func New() drivenlog.DrivenLog {
	l := log.New(&lumberjack.Logger{
		Filename:   "log/server.log",
		MaxSize:    3, // megabytes
		MaxBackups: 3,
		MaxAge:     60,   //days
		Compress:   true, // disabled by default
	}, "", log.Ldate|log.Ltime|log.Llongfile)

	return &logx{
		l: l,
	}
}
