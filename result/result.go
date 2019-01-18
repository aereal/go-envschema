package result

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

type ResultError interface {
	Field() string
	Type() string
	Description() string
	DescriptionFormat() string
	Value() interface{}
	String() string
}

type Result interface {
	Valid() bool
	Errors() []ResultError
	CombinedError() error
}

func Success() Result {
	return &result{valid: true, errors: []ResultError{}}
}

func Failure(errs []ResultError) Result {
	return &result{valid: false, errors: errs}
}

func From(res *gojsonschema.Result) Result {
	errs := make([]ResultError, len(res.Errors()))
	for i, e := range res.Errors() {
		errs[i] = e
	}
	return &result{valid: res.Valid(), errors: errs}
}

type result struct {
	valid  bool
	errors []ResultError
}

func (r *result) Valid() bool {
	return r.valid
}

func (r *result) Errors() []ResultError {
	return r.errors
}

func (r *result) CombinedError() error {
	if len(r.errors) == 0 {
		return nil
	}

	msgs := make([]string, len(r.errors)+1)
	msgs[0] = "failed to following validations:"
	for i, e := range r.errors {
		msgs[i+1] = fmt.Sprintf("%q; actual=%v", e, e.Value())
	}
	return errors.New(strings.Join(msgs, "\n"))
}
