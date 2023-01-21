package strcase

import "testing"

func TestPascal(t *testing.T) {
	samples := []TT{
		{"sample text", "SampleText"},
		{"sample-text", "SampleText"},
		{"sample_text", "SampleText"},
		{"sample___text", "SampleText"},
		{"sampleText", "SampleText"},
		{"inviteYourCustomersAddInvites", "InviteYourCustomersAddInvites"},
		{"sample 2 Text", "Sample2Text"},
		{"   sample   2    Text   ", "Sample2Text"},
		{"   $#$sample   2    Text   ", "Sample2Text"},
		{"SAMPLE 2 TEXT", "Sample2Text"},
		{"___$$Base64Encode", "Base64Encode"},
		{"FOO:BAR$BAZ", "FooBarBaz"},
		{"FOO#BAR#BAZ", "FooBarBaz"},
		{"something.com", "SomethingCom"},
		{"$something%", "Something"},
		{"something.com", "SomethingCom"},
		{"•¶§ƒ˚foo˙∆˚¬", "Foo"},
		{"CStringRef", "CstringRef"},
		{"5test", "5Test"},
		{"test5", "Test5"},
		{"THE5r", "The5R"},
		{"5TEst", "5Test"},
		{"_5TEst", "5Test"},
		{"@%#&5TEst", "5Test"},
		{"edf_6N", "Edf6N"},
		{"f_pX9", "FPX9"},
		{"p_z9Rg", "PZ9Rg"},
		{"2FA Enabled", "2FaEnabled"},
		{"Enabled 2FA", "Enabled2Fa"},
	}

	for _, sample := range samples {
		if out := Pascal(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}
