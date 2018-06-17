package vString

import (
	"regexp"
	"testing"
)

func TestRegExPatterns(t *testing.T) {
	type Data struct {
		Value    string
		Expected bool
	}

	comparedPattern := func(t *testing.T, pattern *regexp.Regexp, dataSet []Data) {
		for _, data := range dataSet {
			if pattern.MatchString(data.Value) != data.Expected {
				t.Errorf("Not equal to expected '%t'", data.Expected)
			}
		}
	}

	t.Run("AlphaPattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(AlphaPattern), []Data{
			{
				Value:    "abc",
				Expected: true,
			},
			{
				Value:    "abc-_",
				Expected: false,
			},
			{
				Value:    "123",
				Expected: false,
			},
			{
				Value:    "123-_",
				Expected: false,
			},
		})
	})

	t.Run("AlphaDashPattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(AlphaDashPattern), []Data{
			{
				Value:    "abc",
				Expected: true,
			},
			{
				Value:    "abc-_",
				Expected: true,
			},
			{
				Value:    "123",
				Expected: true,
			},
			{
				Value:    "123-_",
				Expected: true,
			},
		})
	})

	t.Run("AlphaNumericPattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(AlphaNumericPattern), []Data{
			{
				Value:    "abc",
				Expected: true,
			},
			{
				Value:    "abc-_",
				Expected: false,
			},
			{
				Value:    "123",
				Expected: true,
			},
			{
				Value:    "123-_",
				Expected: false,
			},
		})
	})

	t.Run("CreditCardPattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(CreditCardPattern), []Data{
			{
				// American Express 1
				Value:    "378282246310005",
				Expected: true,
			},
			{
				// American Express 2
				Value:    "371449635398431",
				Expected: true,
			},
			{
				// American Express Corporate
				Value:    "378734493671000",
				Expected: true,
			},
			{
				// Visa 1
				Value:    "4111111111111111",
				Expected: true,
			},
			{
				// Visa 2
				Value:    "4012888888881881",
				Expected: true,
			},
			{
				// Visa 3
				Value:    "4222222222222",
				Expected: true,
			},
			{
				// Diners Club 1
				Value:    "30569309025904",
				Expected: true,
			},
			{
				// Diners Club 2
				Value:    "38520000023237",
				Expected: true,
			},
			{
				// Discover 1
				Value:    "6011111111111117",
				Expected: true,
			},
			{
				// Discover 2
				Value:    "6011000990139424",
				Expected: true,
			},
			{
				// JBC 1
				Value:    "3530111333300000",
				Expected: true,
			},
			{
				// JBC 2
				Value:    "3566002020360505",
				Expected: true,
			},
			{
				// MasterCard 1
				Value:    "5555555555554444",
				Expected: true,
			},
			{
				// MasterCard 2
				Value:    "5105105105105100",
				Expected: true,
			},
			{
				Value:    "12341234123412341234",
				Expected: false,
			},
		})
	})

	t.Run("CSSColorPattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(CSSColorPattern), []Data{
			{
				Value:    "#111",
				Expected: true,
			},
			{
				Value:    "#222222",
				Expected: true,
			},
			{
				Value:    "rgb(3,3,3)",
				Expected: true,
			},
			{
				Value:    "rgba(4%,4,4%,0.4)",
				Expected: true,
			},
			{
				Value:    "hsl(5,5,5)",
				Expected: true,
			},
			{
				Value:    "hsla(6,6,6,0.6)",
				Expected: true,
			},
			{
				Value:    "#11",
				Expected: false,
			},
			{
				Value:    "rgb(2,2,2,2)",
				Expected: false,
			},
			{
				Value:    "rgba(3,3,3,33)",
				Expected: false,
			},
			{
				Value:    "hsla(4,4,4,4)",
				Expected: false,
			},
		})
	})

	t.Run("EmailPattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(EmailPattern), []Data{
			{
				Value:    "test@example.com",
				Expected: true,
			},
			{
				Value:    "test@example",
				Expected: false,
			},
		})
	})

	t.Run("IPPattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(IPPattern), []Data{
			{
				Value:    "127.0.0.1",
				Expected: true,
			},
			{
				Value:    "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
				Expected: false,
			},
		})
	})

	t.Run("IPV4Pattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(IPV4Pattern), []Data{
			{
				Value:    "127.0.0.1",
				Expected: true,
			},
			{
				Value:    "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
				Expected: false,
			},
		})
	})

	t.Run("IPV6Pattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(IPV6Pattern), []Data{
			{
				Value:    "127.0.0.1",
				Expected: false,
			},
			{
				Value:    "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
				Expected: true,
			},
		})
	})

	t.Run("URLPattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(URLPattern), []Data{
			{
				Value:    "https://www.example.com",
				Expected: true,
			},
			{
				Value:    "http://www.example.com",
				Expected: true,
			},
			{
				Value:    "//www.example.com",
				Expected: false,
			},
		})
	})

	t.Run("UUIDPattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(UUIDPattern), []Data{
			{
				Value:    "3f3e3573-15a2-49dd-87ec-2015c2a1de45",
				Expected: true,
			},
			{
				Value:    "2c1d43b8-e6d7-376e-af7f-d4bde997cc3f",
				Expected: false,
			},
			{
				Value:    "39888f87-fb62-5988-a425-b2ea63f5b81e",
				Expected: false,
			},
		})
	})

	t.Run("UUID3Pattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(UUID3Pattern), []Data{
			{
				Value:    "3f3e3573-15a2-49dd-87ec-2015c2a1de45",
				Expected: false,
			},
			{
				Value:    "2c1d43b8-e6d7-376e-af7f-d4bde997cc3f",
				Expected: true,
			},
			{
				Value:    "39888f87-fb62-5988-a425-b2ea63f5b81e",
				Expected: false,
			},
		})
	})

	t.Run("UUID4Pattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(UUID4Pattern), []Data{
			{
				Value:    "3f3e3573-15a2-49dd-87ec-2015c2a1de45",
				Expected: true,
			},
			{
				Value:    "2c1d43b8-e6d7-376e-af7f-d4bde997cc3f",
				Expected: false,
			},
			{
				Value:    "39888f87-fb62-5988-a425-b2ea63f5b81e",
				Expected: false,
			},
		})
	})

	t.Run("UUID5Pattern", func(t *testing.T) {
		comparedPattern(t, regexp.MustCompile(UUID5Pattern), []Data{
			{
				Value:    "3f3e3573-15a2-49dd-87ec-2015c2a1de45",
				Expected: false,
			},
			{
				Value:    "2c1d43b8-e6d7-376e-af7f-d4bde997cc3f",
				Expected: false,
			},
			{
				Value:    "39888f87-fb62-5988-a425-b2ea63f5b81e",
				Expected: true,
			},
		})
	})
}
