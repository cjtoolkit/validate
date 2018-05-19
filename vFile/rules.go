package vFile

import (
	"io"
	"mime/multipart"
	"path"

	"github.com/cjtoolkit/validate/vError"
)

func Mandatory() ValidationRule {
	return func(value *multipart.FileHeader, hasError bool) error {
		if nil == value {
			return vError.ValidationError{
				Type:   Type,
				Data:   nil,
				Format: MandatoryFormat,
			}
		}

		return nil
	}
}

func Optional(rules ...ValidationRule) ValidationRule {
	return func(value *multipart.FileHeader, hasError bool) error {
		if nil == value {
			return nil
		}

		collector := vError.NewErrorCollector()

		for _, rule := range rules {
			collector.Collect(rule(value, collector.HasError()))
		}

		return collector.GetErrors()
	}
}

func OverrideErrorMsg(validationError vError.ValidationError, rules ...ValidationRule) ValidationRule {
	return func(value *multipart.FileHeader, hasError bool) error {
		collector := vError.NewErrorCollector()
		for _, rule := range rules {
			collector.Collect(rule(value, hasError))
		}

		if collector.HasError() {
			return validationError
		}

		return nil
	}
}

func AcceptMime(mimes ...string) ValidationRule {
	return Optional(func(value *multipart.FileHeader, hasError bool) error {
		fileType := value.Header.Get("Content-Type")

		for _, mime := range mimes {
			if matched, _ := path.Match(mime, fileType); matched {
				return nil
			}
		}

		return vError.ValidationError{
			Type: Type,
			Data: map[string]interface{}{
				"mimes": mimes,
			},
			Format: AcceptMimeFormat,
		}
	})
}

func MaxSize(maxBytes int64) ValidationRule {
	return Optional(func(value *multipart.FileHeader, hasError bool) error {
		file, err := value.Open()
		if nil != err {
			return nil
		}
		defer file.Close()

		size, err := file.Seek(0, io.SeekEnd)
		if nil != err {
			return nil
		}

		file.Seek(0, io.SeekStart)

		if size > maxBytes {
			return vError.ValidationError{
				Type: Type,
				Data: map[string]interface{}{
					"maxBytes": maxBytes,
				},
				Format: MaxSizeFormat,
			}
		}

		return nil
	})
}
