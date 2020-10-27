package pkg

import (
	"logger/internal"
	"logger/internal/logrus"
	"logger/pkg/structs"
	"reflect"
	"testing"
)

func TestLog_Component(t *testing.T) {
	_logrus := &logrus.Log{}

	type fields struct {
		log internal.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			"standard",
			fields{log: _logrus},
			_logrus,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				log: tt.fields.log,
			}
			if got := l.Component(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Component() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_Config(t *testing.T) {
	_logrus := &logrus.Log{}

	type fields struct {
		log internal.Logger
	}
	type args struct {
		configuration map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"standard",
			fields{log: _logrus},
			args{configuration: map[string]interface{}{"level": structs.LevelTrace}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				log: tt.fields.log,
			}
			if err := l.Config(tt.args.configuration); (err != nil) != tt.wantErr {
				t.Errorf("Config() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLog_Use(t *testing.T) {
	_logrus := &logrus.Log{}
	type fields struct {
		log internal.Logger
	}
	type args struct {
		driver string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Componentor
	}{
		{
			"standard",
			fields{_logrus},
			args{driver: DriverLogrus},
			&Log{_logrus},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				log: tt.fields.log,
			}
			if got := l.Use(tt.args.driver); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Use() = %v, want %v", got, tt.want)
			}
		})
	}
}
