package loader

import (
	"reflect"
	"testing"

	"github.com/aereal/go-envschema/validator"
	"github.com/xeipuuv/gojsonschema"
)

func TestLoader_LoadValue(t *testing.T) {
	type fields struct {
		validator *validator.Validator
	}
	type args struct {
		loader gojsonschema.JSONLoader
		dest   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				validator: testValidator,
			},
			args: args{
				dest:   &testConfig{},
				loader: gojsonschema.NewStringLoader(`{"ADDR":"foo","DSN":"bar"}`),
			},
			want:    &testConfig{Addr: "foo", Dsn: "bar"},
			wantErr: false,
		},
		{
			name: "not valid",
			fields: fields{
				validator: testValidator,
			},
			args: args{
				dest:   &testConfig{},
				loader: gojsonschema.NewStringLoader(`{"ADDR":1,"DSN":"bar"}`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Loader{
				validator: tt.fields.validator,
			}
			if err := l.LoadValue(tt.args.loader, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("Loader.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !reflect.DeepEqual(tt.args.dest, tt.want) {
				t.Errorf("Loader.Load() loaded = %#v, want = %#v", tt.args.dest, tt.want)
			}
		})
	}
}

var testValidator *validator.Validator

type testConfig struct {
	Addr string `json:"ADDR"`
	Dsn  string `json:"DSN"`
}

func init() {
	sc := gojsonschema.NewStringLoader(`
	{
		"definitions": {},
		"$schema": "http://json-schema.org/draft-07/schema#",
		"type": "object",
		"title": "configuration",
		"properties": {
			"ADDR": {
				"type": "string"
			},
			"DSN": {
				"type": "string",
				"minLength": 1
			}
		}
	}
	`)
	v, err := validator.New(sc)
	if err != nil {
		panic(err)
	}
	testValidator = v
}
