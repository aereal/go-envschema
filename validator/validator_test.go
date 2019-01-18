package validator

import (
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
			if got.Valid() != tt.wantValid {
				t.Errorf("Validator.ValidateValue() = %v, want %v", got.Valid(), tt.wantValid)
			}
		})
	}
}
