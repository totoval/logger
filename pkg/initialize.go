package pkg

import (
	"errors"
	"logger/internal"
	"logger/pkg/structs"

	"logger/internal/logrus"
)

type Log struct {
	log internal.Logger
}

func (l *Log) Component() interface{} {
	return l.log
}

func (l *Log) Use(driver string) Componentor {
	switch driver {
	case DriverLogrus:
		l.log = &logrus.Log{}
	default:
		l.log = &logrus.Log{}
	}
	return l
}

func (l *Log) Config(configuration map[string]interface{}) error {
	level, ok := configuration["level"].(structs.Level)
	if !ok {
		return errors.New("unknown configuration log level")
	}

	return l.log.New().SetLogLevel(level)
}
