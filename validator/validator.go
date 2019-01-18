package validator

import (
	"github.com/aereal/go-envschema/envloader"
	"github.com/aereal/go-envschema/result"
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

func (v *Validator) Validate() (result.Result, error) {
	loader := envloader.New()
	return v.ValidateValue(loader)
}

func (v *Validator) ValidateValue(input ValueLoader) (result.Result, error) {
	jsResult, err := v.schema.Validate(input)
	if err != nil {
		return nil, err
	}
	return result.From(jsResult), nil
}
