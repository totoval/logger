package internal

import (
	"github.com/totoval/framework/helpers/toto"
	"github.com/totoval/logger/pkg/structs"
)

type Logger interface {
	New() Logger
	SetLogLevel(level structs.Level) error
	ConvertLevel(level structs.Level) (interface{}, error)
	LogLevel() structs.Level
	Debug(msg interface{}, v ...toto.V)
	Error(err error, v ...toto.V) error
	Info(msg interface{}, v ...toto.V)
	Warn(msg interface{}, v ...toto.V)
	Fatal(msg interface{}, v ...toto.V)
	Panic(msg interface{}, v ...toto.V)
	Trace(msg interface{}, v ...toto.V)
	ErrorStr(err error, v ...toto.V) string
}
