package logrus

import (
	"github.com/sirupsen/logrus"
	"logger/internal"
	"logger/pkg/structs"
	"reflect"
	"testing"
)

func TestLog_ConvertLevel(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	type args struct {
		level structs.Level
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"standard",
			fields{
				logrus: logrus.New(),
				level:  structs.LevelTrace,
			},
			args{level: structs.LevelTrace},
			logrus.TraceLevel,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				logrus: tt.fields.logrus,
				level:  tt.fields.level,
			}
			got, err := l.ConvertLevel(tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertLevel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_LogLevel(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Level
	}{
		{
			"standard",
			fields{
				logrus: logrus.New(),
				level:  structs.LevelTrace,
			},
			structs.LevelTrace,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				logrus: tt.fields.logrus,
				level:  tt.fields.level,
			}
			if got := l.LogLevel(); got != tt.want {
				t.Errorf("LogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_New(t *testing.T) {
	_logrus := logrus.New()
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	tests := []struct {
		name   string
		fields fields
		want   internal.Logger
	}{
		{
			"standard",
			fields{
				logrus: _logrus,
				level:  structs.LevelTrace,
			},
			&Log{
				logrus: _logrus,
				level:  structs.LevelTrace,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				logrus: tt.fields.logrus,
				level:  tt.fields.level,
			}
			if got := l.New(); !reflect.DeepEqual(got.LogLevel(), tt.want.LogLevel()) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_SetLogLevel(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	type args struct {
		level structs.Level
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"standard",
			fields{
				logrus: logrus.New(),
				level:  structs.LevelTrace,
			},
			args{level: structs.LevelTrace},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				logrus: tt.fields.logrus,
				level:  tt.fields.level,
			}
			if err := l.SetLogLevel(tt.args.level); (err != nil) != tt.wantErr {
				t.Errorf("SetLogLevel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
