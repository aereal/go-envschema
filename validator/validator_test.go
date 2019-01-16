package validator

import (
	"errors"
	"reflect"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

var testSchema *gojsonschema.Schema

func init() {
	sc, err := gojsonschema.NewSchema(gojsonschema.NewStringLoader(`
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
	`))
	if err != nil {
		panic(err)
	}
	testSchema = sc
}

func TestValidator_ValidateValue(t *testing.T) {
	type fields struct {
		schema *gojsonschema.Schema
	}
	type args struct {
		input ValueLoader
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValid bool
		wantErr   bool
	}{
		{
			name: "valid",
			fields: fields{
				schema: testSchema,
			},
			args: args{
				input: gojsonschema.NewStringLoader(`{"DSN":"xxx"}`),
			},
			wantValid: true,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{
				schema: tt.fields.schema,
			}
			got, err := v.ValidateValue(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validator.ValidateValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.IsValid() != tt.wantValid {
				t.Errorf("Validator.ValidateValue() = %v, want %v", got.IsValid(), tt.wantValid)
			}
		})
	}
}

func TestResult_IsValid(t *testing.T) {
	type fields struct {
		errors  []error
		isValid bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ok",
			fields: fields{
				isValid: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Result{
				errors:  tt.fields.errors,
				isValid: tt.fields.isValid,
			}
			if got := r.IsValid(); got != tt.want {
				t.Errorf("Result.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_Errors(t *testing.T) {
	type fields struct {
		errors  []error
		isValid bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []error
	}{
		{
			name: "some errors",
			fields: fields{
				errors: []error{errors.New("oops")},
			},
			want: []error{errors.New("oops")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Result{
				errors:  tt.fields.errors,
				isValid: tt.fields.isValid,
			}
			if got := r.Errors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Result.Errors() = %v, want %v", got, tt.want)
			}
		})
	}
}
