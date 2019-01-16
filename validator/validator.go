package validator

import (
	"errors"
	"os"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

type SchemaLoader gojsonschema.JSONLoader

type ValueLoader gojsonschema.JSONLoader

type Validator struct {
	schema *gojsonschema.Schema
}

func New(sl SchemaLoader) (*Validator, error) {
	sc, err := gojsonschema.NewSchema(sl)
	if err != nil {
		return nil, err
	}
	return &Validator{
		schema: sc,
	}, nil
}

func (v *Validator) Validate() (*Result, error) {
	value := getEnvs()
	input := gojsonschema.NewGoLoader(value)
	return v.ValidateValue(input)
}

func (v *Validator) ValidateValue(input ValueLoader) (*Result, error) {
	jsResult, err := v.schema.Validate(input)
	if err != nil {
		return nil, err
	}
	if jsResult.Valid() {
		return &Result{isValid: true}, nil
	}

	result := &Result{isValid: false}
	for _, err := range jsResult.Errors() {
		result.errors = append(result.errors, errors.New(err.String()))
	}
	return result, nil
}

type Result struct {
	errors  []error
	isValid bool
}

func (r *Result) IsValid() bool {
	return r.isValid
}

func (r *Result) Errors() []error {
	return r.errors
}

func getEnvs() map[string]string {
	return environAsMap(os.Environ())
}

func environAsMap(pairs []string) map[string]string {
	envs := map[string]string{}
	for _, env := range pairs {
		pair := strings.SplitN(env, "=", 2)
		envs[pair[0]] = pair[1]
	}
	return envs
}
