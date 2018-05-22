package vError

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"
)

type ValidationError struct {
	Type   string
	Data   map[string]interface{}
	Format string
}

func (e ValidationError) bytes() []byte {
	buf := &bytes.Buffer{}
	maps := template.FuncMap{"join": join}
	template.Must(template.New("ValidationError").Funcs(maps).Parse(e.Format)).Execute(buf, e.Data)

	return buf.Bytes()
}

func (e ValidationError) Error() string { return string(e.bytes()) }

func (e ValidationError) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string{string(e.bytes())})
}

type Errors []error

func (e Errors) strings() []string {
	strs := []string{}
	for _, v := range e {
		strs = append(strs, v.Error())
	}
	return strs
}

func (e Errors) Error() string                { return strings.Join(e.strings(), "\n") }
func (e Errors) MarshalJSON() ([]byte, error) { return json.Marshal(e.strings()) }

type ErrorCollector struct {
	hasError bool
	errs     []error
}

func NewErrorCollector() *ErrorCollector {
	return &ErrorCollector{
		hasError: false,
		errs:     []error{},
	}
}

func (e *ErrorCollector) Collect(err error) {
	if err == nil {
		return
	}

	e.hasError = true

	switch err := err.(type) {
	case Errors:
		e.errs = append(e.errs, err...)
	default:
		e.errs = append(e.errs, err)
	}
}

func (e *ErrorCollector) HasError() bool { return e.hasError }

func (e *ErrorCollector) GetErrors() error {
	if !e.hasError {
		return nil
	}

	return Errors(e.errs)
}
