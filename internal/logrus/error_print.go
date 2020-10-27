package logrus

import (
	"fmt"
	"github.com/totoval/framework/helpers/toto"
	"github.com/ztrue/tracerr"
	"logger/pkg/structs"
)

func (l *Log) errPrintln(err error, fields toto.V) {
	startFrom := 2
	if err == nil {
		return
	}
	traceErr := tracerr.Wrap(err)
	frameList := tracerr.StackTrace(traceErr)
	if startFrom > len(frameList) || len(frameList)-2 <= 0 {
		l.println(structs.LevelError, err.Error(), fields)
	}

	traceErr = tracerr.CustomError(err, frameList[startFrom:len(frameList)-2])
	traceErr = tracerr.CustomError(err, frameList)

	if fields == nil {
		fields = toto.V{}
	}
	fields["totoval_trace"] = tracerr.SprintSource(traceErr, 0)
	l.println(structs.LevelError, err.Error(), fields)
}

func (l *Log) errPrint(err error, startFrom int, fields toto.V) string {
	if err == nil {
		return ""
	}
	traceErr := tracerr.Wrap(err)
	frameList := tracerr.StackTrace(traceErr)
	//if startFrom > len(frameList) {
	//	return fmt.Sprint(err.Error(), fields)
	//}
	//
	//traceErr = tracerr.CustomError(err, frameList[startFrom:len(frameList)-2])
	traceErr = tracerr.CustomError(err, frameList)

	if fields == nil {
		fields = toto.V{}
	}
	fields["totoval_trace"] = tracerr.SprintSource(traceErr)
	return fmt.Sprint(err.Error(), fields)
}
