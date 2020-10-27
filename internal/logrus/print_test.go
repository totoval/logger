package logrus

import (
	"github.com/sirupsen/logrus"
	"github.com/totoval/framework/helpers/toto"
	"github.com/totoval/logger/pkg/structs"
	"testing"
)

func TestLog_Debug(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	type args struct {
		msg interface{}
		v   []toto.V
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"standard",
			fields{
				logrus.New(),
				structs.LevelTrace,
			},
			args{
				"test",
				[]toto.V{
					{
						"mode": "standard",
					},
				},
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

func TestLog_Error(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	type args struct {
		err error
		v   []toto.V
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				logrus: tt.fields.logrus,
				level:  tt.fields.level,
			}
			if err := l.Error(tt.args.err, tt.args.v...); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLog_ErrorStr(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	type args struct {
		err error
		v   []toto.V
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				logrus: tt.fields.logrus,
				level:  tt.fields.level,
			}
			if got := l.ErrorStr(tt.args.err, tt.args.v...); got != tt.want {
				t.Errorf("ErrorStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_Fatal(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	type args struct {
		msg interface{}
		v   []toto.V
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"standard",
			fields{
				logrus.New(),
				structs.LevelTrace,
			},
			args{
				"test",
				[]toto.V{
					{
						"mode": "standard",
					},
				},
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

func TestLog_Info(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	type args struct {
		msg interface{}
		v   []toto.V
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"standard",
			fields{
				logrus.New(),
				structs.LevelTrace,
			},
			args{
				"test",
				[]toto.V{
					{
						"mode": "standard",
					},
				},
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

func TestLog_Panic(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	type args struct {
		msg interface{}
		v   []toto.V
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"standard",
			fields{
				logrus.New(),
				structs.LevelTrace,
			},
			args{
				"test",
				[]toto.V{
					{
						"mode": "standard",
					},
				},
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

func TestLog_Trace(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	type args struct {
		msg interface{}
		v   []toto.V
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"standard",
			fields{
				logrus.New(),
				structs.LevelTrace,
			},
			args{
				"test",
				[]toto.V{
					{
						"mode": "standard",
					},
				},
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

func TestLog_Warn(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	type args struct {
		msg interface{}
		v   []toto.V
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"standard",
			fields{
				logrus.New(),
				structs.LevelTrace,
			},
			args{
				"test",
				[]toto.V{
					{
						"mode": "standard",
					},
				},
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

func TestLog_println(t *testing.T) {
	type fields struct {
		logrus *logrus.Logger
		level  structs.Level
	}
	type args struct {
		level  structs.Level
		msg    interface{}
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
				logrus.New(),
				structs.LevelTrace,
			},
			args{
				structs.LevelTrace,
				"test",
				toto.V{
					"mode": "standard",
				},
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
