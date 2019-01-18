package loader

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aereal/go-envschema/envloader"
	"github.com/aereal/go-envschema/validator"
	"github.com/xeipuuv/gojsonschema"
)

type Loader struct {
	validator *validator.Validator
}

func New(sl validator.SchemaLoader) (*Loader, error) {
	validator, err := validator.New(sl)
	if err != nil {
		return nil, err
	}
	return &Loader{
		validator: validator,
	}, nil
}

func (l *Loader) LoadValue(loader gojsonschema.JSONLoader, dest interface{}) error {
	result, err := l.validator.ValidateValue(loader)
	if err != nil {
		return fmt.Errorf("validation failed: %s", err)
	}
	if !result.Valid() {
		return result.CombinedError()
	}

	switch src := loader.JsonSource().(type) {
	case string:
		buf := bytes.NewBufferString(src)
		if err := json.NewDecoder(buf).Decode(dest); err != nil {
			return fmt.Errorf("failed to decode JSON string: %s", err)
		}
	default:
		jsonBytes, err := json.Marshal(src)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %s", err)
		}
		if err := json.Unmarshal(jsonBytes, dest); err != nil {
			return fmt.Errorf("failed unmarshal JSON: %s", err)
		}
	}

	return nil
}

func (l *Loader) Load(dest interface{}) error {
	loader := envloader.New()
	return l.LoadValue(loader, dest)
}
