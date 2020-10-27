package logrus

import (
	"errors"
	"github.com/sirupsen/logrus"
	"logger/internal"
	"logger/pkg/structs"
	"os"
)

type Log struct {
	logrus *logrus.Logger
	level  structs.Level
}

func (l *Log) New() internal.Logger {
	l.logrus = logrus.New()
	l.logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	l.logrus.Out = os.Stdout

	return l
}

func (l *Log) LogLevel() structs.Level {
	return l.level
}
func (l *Log) SetLogLevel(level structs.Level) error {
	l.level = level

	logrusLevel, err := l.ConvertLevel(l.level)
	if err != nil {
		return err
	}
	l.logrus.SetLevel(logrusLevel.(logrus.Level))

	return nil
}

func (l *Log) ConvertLevel(level structs.Level) (interface{}, error) {
	switch level {
	case structs.LevelPanic:
		return logrus.PanicLevel, nil
	case structs.LevelFatal:
		return logrus.FatalLevel, nil
	case structs.LevelError:
		return logrus.ErrorLevel, nil
	case structs.LevelWarn:
		return logrus.WarnLevel, nil
	case structs.LevelInfo:
		return logrus.InfoLevel, nil
	case structs.LevelDebug:
		return logrus.DebugLevel, nil
	case structs.LevelTrace:
		return logrus.TraceLevel, nil
	default:
		return nil, errors.New("unknown log level")
	}
}
