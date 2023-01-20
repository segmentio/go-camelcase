package strcase

import "testing"

var ops int = 1e6

type sample struct {
	str, out string
}

func TestSnakecase(t *testing.T) {
	samples := []sample{
		{"@49L0S145_¬fwHƒ0TSLNVp", "49l0s145_fw_h_0tslnvp"},
		{"lk0B@bFmjrLQ_Z6YL", "lk0_b_b_fmjr_lq_z6yl"},
		{"samPLE text", "sam_ple_text"},
		{"sample text", "sample_text"},
		{"sample-text", "sample_text"},
		{"sample_text", "sample_text"},
		{"sample___text", "sample_text"},
		{"sampleText", "sample_text"},
		{"inviteYourCustomersAddInvites", "invite_your_customers_add_invites"},
		{"sample 2 Text", "sample_2_text"},
		{"   sample   2    Text   ", "sample_2_text"},
		{"   $#$sample   2    Text   ", "sample_2_text"},
		{"SAMPLE 2 TEXT", "sample_2_text"},
		{"___$$Base64Encode", "base64_encode"},
		{"FOO:BAR$BAZ", "foo_bar_baz"},
		{"FOO#BAR#BAZ", "foo_bar_baz"},
		{"something.com", "something_com"},
		{"$something%", "something"},
		{"something.com", "something_com"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
		{"CStringRef", "cstring_ref"},
		{"5test", "5test"},
		{"test5", "test5"},
		{"THE5r", "the5r"},
		{"5TEst", "5test"},
		{"_5TEst", "5test"},
		{"@%#&5TEst", "5test"},
		{"edf_6N", "edf_6n"},
		{"f_pX9", "f_p_x9"},
		{"p_z9Rg", "p_z9_rg"},
		{"2FA Enabled", "2fa_enabled"},
		{"Enabled 2FA", "enabled_2fa"},
	}

	for _, sample := range samples {
		if out := SnakeCase(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}
