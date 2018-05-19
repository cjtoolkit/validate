package vFile

import (
	"bytes"
	"errors"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"testing"

	"github.com/cjtoolkit/validate/vError"
)

func TestRules(t *testing.T) {
	createTestForm := func() *multipart.Form {
		buf := &bytes.Buffer{}
		defer buf.Reset()

		mw := multipart.NewWriter(buf)

		{
			// Create four byte text file
			h := make(textproto.MIMEHeader)

			h.Set("Content-Disposition", `form-data; name="text4"; filename="text4.txt"`)
			h.Set("Content-Type", "text/plain")

			w, _ := mw.CreatePart(h)

			fmt.Fprint(w, "aaaa")
		}

		{
			// Create eight byte text file
			h := make(textproto.MIMEHeader)

			h.Set("Content-Disposition", `form-data; name="text8"; filename="text8.txt"`)
			h.Set("Content-Type", "text/plain")

			w, _ := mw.CreatePart(h)

			fmt.Fprint(w, "aaaaaaaa")
		}

		{
			// Create image file
			h := make(textproto.MIMEHeader)

			h.Set("Content-Disposition", `form-data; name="img"; filename="img.jpg"`)
			h.Set("Content-Type", "image/jpeg")

			w, _ := mw.CreatePart(h)

			fmt.Fprint(w, "iiii")
		}

		boundary := mw.Boundary()
		mw.Close()

		multipartForm, _ := multipart.NewReader(buf, boundary).ReadForm(1 * 1024 * 1024)

		return multipartForm
	}

	t.Run("Mandatory", func(t *testing.T) {
		t.Run("Has File", func(t *testing.T) {
			if Mandatory()(&multipart.FileHeader{}, false) != nil {
				t.Error("Should be nil")
			}
		})

		t.Run("Has no File", func(t *testing.T) {
			if Mandatory()(nil, false) == nil {
				t.Error("Should not be nil")
			}
		})
	})

	t.Run("Optional", func(t *testing.T) {
		rule := ValidationRule(func(value *multipart.FileHeader, hasError bool) error {
			return errors.New("I am error")
		})

		t.Run("Value is empty", func(t *testing.T) {
			if Optional(rule)(nil, false) != nil {
				t.Error("Should be nil")
			}
		})

		t.Run("Value is not empty", func(t *testing.T) {
			if Optional(rule)(&multipart.FileHeader{}, false) == nil {
				t.Error("Should not be nil")
			}
		})
	})

	t.Run("OverrideErrorMsg", func(t *testing.T) {
		ruleWithError := ValidationRule(func(value *multipart.FileHeader, hasError bool) error {
			return errors.New("I am error")
		})
		ruleWithNoError := ValidationRule(func(value *multipart.FileHeader, hasError bool) error {
			return nil
		})

		t.Run("Has Error", func(t *testing.T) {
			if OverrideErrorMsg(vError.ValidationError{}, ruleWithError)(nil, false) == nil {
				t.Error("Should not be nil")
			}
		})

		t.Run("No Error", func(t *testing.T) {
			if OverrideErrorMsg(vError.ValidationError{}, ruleWithNoError)(nil, false) != nil {
				t.Error("Should be nil")
			}
		})
	})

	t.Run("AcceptMime", func(t *testing.T) {
		form := createTestForm()

		t.Run("No file, returns no error", func(t *testing.T) {
			err := AcceptMime("image/jpeg")(nil, false)

			if nil != err {
				t.Error("Should have no error")
			}
		})

		t.Run("File is jpeg, expect jpeg", func(t *testing.T) {
			err := AcceptMime("image/jpeg")(form.File["img"][0], false)

			if nil != err {
				t.Error("Should have no error")
			}
		})

		t.Run("File is jpeg, expect text", func(t *testing.T) {
			err := AcceptMime("text/plain")(form.File["img"][0], false)

			if nil == err {
				t.Error("Should have error")
			}
		})
	})

	t.Run("MaxSize", func(t *testing.T) {
		form := createTestForm()

		t.Run("No file, returns no error", func(t *testing.T) {
			err := MaxSize(6)(nil, false)

			if nil != err {
				t.Error("Should have no error")
			}
		})

		t.Run("File is less than 6 bytes, return no error", func(t *testing.T) {
			err := MaxSize(6)(form.File["text4"][0], false)

			if nil != err {
				t.Error("Should have no error")
			}
		})

		t.Run("File is greater than 6 bytes, return error", func(t *testing.T) {
			err := MaxSize(6)(form.File["text8"][0], false)

			if nil == err {
				t.Error("Should have error")
			}
		})
	})
}
