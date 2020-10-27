package logrus

import (
	"github.com/sirupsen/logrus"
	"github.com/totoval/framework/helpers/toto"
	"logger/pkg/structs"
)

func (l *Log) println(level structs.Level, msg interface{}, fields toto.V) {
	logrusLevel, err := l.ConvertLevel(level)
	if err != nil {
		panic(err)
	}

	if fields == nil {
		l.logrus.Log(logrusLevel.(logrus.Level), msg)
	} else {
		l.logrus.WithFields(logrus.Fields(fields)).Log(logrusLevel.(logrus.Level), msg)
	}

	// if level <= l.LogLevel() {
	// 	_fields := make(map[string]interface{})
	// 	if fields != nil {
	// 		_fields = fields
	// 	}
	//
	// 	switch level {
	// 	case pkg.LevelPanic:
	// 		sentry.CaptureError(errors.New(fmt.Sprintf("%v - %v", msg, fields)))
	// 	case pkg.LevelFatal:
	// 		sentry.CaptureError(errors.New(fmt.Sprintf("%v - %v", msg, fields)))
	// 	case pkg.LevelError:
	// 		sentry.CaptureError(errors.New(fmt.Sprintf("%v - %v", msg, fields)))
	// 	case pkg.LevelWarn:
	// 		_fields["level"] = "WARN"
	// 		sentry.CaptureMsg(fmt.Sprintf("%v", msg), _fields)
	// 		//case INFO:
	// 		//	_fields["level"] = "INFO"
	// 		//	sentry.CaptureMsg(msg, _fields)
	// 		//case DEBUG:
	// 		//	_fields["level"] = "DEBU"
	// 		//	sentry.CaptureMsg(msg, _fields)
	// 		//case TRACE:
	// 		//	_fields["level"] = "TRAC"
	// 		//	sentry.CaptureMsg(msg, _fields)
	// 	}
	// }
}

func (l *Log) Debug(msg interface{}, v ...toto.V) {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	l.println(structs.LevelDebug, msg, fields)
}

func (l *Log) Error(err error, v ...toto.V) error {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	l.errPrintln(err, fields)
	return err
}

func (l *Log) Info(msg interface{}, v ...toto.V) {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	l.println(structs.LevelInfo, msg, fields)
}

func (l *Log) Warn(msg interface{}, v ...toto.V) {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	l.println(structs.LevelWarn, msg, fields)
}

func (l *Log) Fatal(msg interface{}, v ...toto.V) {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	l.println(structs.LevelFatal, msg, fields)
}

func (l *Log) Panic(msg interface{}, v ...toto.V) {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	l.println(structs.LevelPanic, msg, fields)
}

func (l *Log) Trace(msg interface{}, v ...toto.V) {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	l.println(structs.LevelTrace, msg, fields)
}

func (l *Log) ErrorStr(err error, v ...toto.V) string {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	return l.errPrint(err, 2, fields)
}
