package strcase

import "testing"

type TT struct {
	str, out string
}

func TestCamel(t *testing.T) {
	samples := []TT{
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
		{"CStringRef", "cstringRef"},
		{"5test", "5Test"},
		{"test5", "test5"},
		{"THE5r", "the5R"},
		{"5TEst", "5Test"},
		{"_5TEst", "5Test"},
		{"@%#&5TEst", "5Test"},
		{"edf_6N", "edf6N"},
		{"f_pX9", "fPX9"},
		{"p_z9Rg", "pZ9Rg"},
		{"2FA Enabled", "2FaEnabled"},
		{"Enabled 2FA", "enabled2Fa"},
	}

	for _, sample := range samples {
		if out := Camel(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}
