package vError

import (
	"bytes"
	"strings"
	"text/template"
)

type ValidationError struct {
	Data   map[string]interface{}
	Format string
}

func (e ValidationError) Error() string {
	buf := &bytes.Buffer{}
	template.Must(template.New("ValidationError").Parse(e.Format)).Execute(buf, e.Data)

	return buf.String()
}

type Errors []error

func (e Errors) Error() string {
	strs := []string{}
	for _, v := range e {
		strs = append(strs, v.Error())
	}
	return strings.Join(strs, "\n")
}

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
	if err != nil {
		e.hasError = true
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
