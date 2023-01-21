package strcase

import (
	"testing"
	"unicode/utf8"
)

func FuzzKebabCase(f *testing.F) {
	testcases := []string{
		" ",
		"!12345",
		"___$$Base64Encode",
		"FOO:BAR$BAZ",
		"FOO#BAR#BAZ",
		"something.com",
		"$something%",
		"something.com",
		"•¶§ƒ˚foo˙∆˚¬",
		"_5TEst",
		"@%#&5TEst",
		"edf_6N",
		"Enabled 2FA",
		"@49L0S145_¬fwHƒ0TSLNVp",
		"lk0B@bFmjrLQ_Z6YL",
		"samPLE text",
		"\x9c\xdd",
		"\x83\xb3\xe6",
		"泃",
	}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(
		func(t *testing.T, a string) {
			x := Kebab(a)
			if utf8.ValidString(a) && !utf8.ValidString(x) {
				t.Errorf("Kebab produced invalid UTF-8 string %q", x)
			}
		},
	)
}
