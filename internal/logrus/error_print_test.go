package logrus

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/totoval/framework/helpers/toto"
	"github.com/totoval/logger/pkg/structs"
	"testing"
)

func TestLog_errPrint(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	type args struct {
		err       error
		startFrom int
		fields    toto.V
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"standard",
			fields{
				logrus: logrus.New(),
				level:  structs.LevelTrace,
			},
			args{
				err:       nil,
				startFrom: 0,
				fields: toto.V{
					"mode": "standard",
				},
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				logrus: tt.fields.logrus,
				level:  tt.fields.level,
			}
			if got := l.errPrint(tt.args.err, tt.args.startFrom, tt.args.fields); got != tt.want {
				t.Errorf("errPrint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_errPrintln(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	type args struct {
		err    error
		fields toto.V
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"standard",
			fields{
				logrus: logrus.New(),
				level:  structs.LevelTrace,
			},
			args{
				err:    errors.New("standard test"),
				fields: toto.V{"mode": "standard"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &Log{
				logrus: tt.fields.logrus,
				level:  tt.fields.level,
			}
		})
	}
}
