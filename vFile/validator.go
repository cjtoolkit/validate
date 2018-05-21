package vFile

import (
	"mime/multipart"

	"github.com/cjtoolkit/validate/vError"
)

type ValidationRule func(value *multipart.FileHeader, hasError bool) error

func Validate(file *multipart.FileHeader, rules ...ValidationRule) (*multipart.FileHeader, error) {
	value := file

	collector := vError.NewErrorCollector()

	for _, rule := range rules {
		collector.Collect(rule(value, collector.HasError()))
	}

	return value, collector.GetErrors()
}

func ValidateMultiple(files []*multipart.FileHeader, rules ...ValidationRule) ([]*multipart.FileHeader, error) {
	collector := vError.NewErrorCollector()

	for _, file := range files {
		_, err := Validate(file, rules...)
		collector.Collect(err)
	}

	return files, vError.CleanError(collector.GetErrors())
}

func Must(file *multipart.FileHeader, err error) *multipart.FileHeader {
	vError.Must(err)
	return file
}

func MustMultiple(files []*multipart.FileHeader, err error) []*multipart.FileHeader {
	vError.Must(err)
	return files
}
