package camelcase

import "testing"

var ops int = 1e6

type sample struct {
	str, out string
}

func TestCamelcase(t *testing.T) {
	samples := []sample{
		{"sample text", "sampleText"},
		{"sample-text", "sampleText"},
		{"sample_text", "sampleText"},
		{"sample___text", "sampleText"},
		{"sampleText", "sampleText"},
		{"inviteYourCustomersAddInvites", "inviteYourCustomersAddInvites"},
		{"sample 2 Text", "sample2Text"},
		{"   sample   2    Text   ", "sample2Text"},
		{"   $#$sample   2    Text   ", "sample2Text"},
		{"SAMPLE 2 TEXT", "sample2Text"},
		{"___$$Base64Encode", "base64Encode"},
		{"FOO:BAR$BAZ", "fooBarBaz"},
		{"FOO#BAR#BAZ", "fooBarBaz"},
		{"something.com", "somethingCom"},
		{"$something%", "something"},
		{"something.com", "somethingCom"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
	}

	for _, sample := range samples {
		if out := Camelcase(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func BenchmarkCamelcase(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Camelcase("some sample text here_noething:too$amazing")
	}
}
