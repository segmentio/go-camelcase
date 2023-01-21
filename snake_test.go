package strcase

import (
	"fmt"
	"testing"
)

func TestSnake(t *testing.T) {
	samples := []TT{
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
		{"@49L0S145_¬fwHƒ0TSLNVp", "49l0s145_fw_h_0tslnvp"},
		{"lk0B@bFmjrLQ_Z6YL", "lk0_b_b_fmjr_lq_z6yl"},
		{"samPLE text", "sam_ple_text"},
		{"here\nis\na\nmultiline\nstring", "here_is_a_multiline_string"},
		{"\"hello world\"", "hello_world"},
		// {
		// Swedish and non ASCII char are not supported :(
		// open an issue if this is something you need
		// 	input: "När såg du en kråka väl bita en man?"
		// 	got: n_r_s_g_du_en_kr_ka_v_l_bita_en_man
		// 	want: "när_såg_du_en_kråka_väl_bita_en",
		// },
		{
			"super long sentence that is not really necessary but we need to try it",
			"super_long_sentence_that_is_not_really_necessary_but_we_need_to_try_it",
		},
	}

	for _, sample := range samples {
		if out := Snake(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func ExampleSnake() {
	fmt.Println(Snake("sample text"))
	fmt.Println(Snake("sample-text"))
	fmt.Println(Snake("sample_text"))
	fmt.Println(Snake("sample___text"))
	fmt.Println(Snake("sampleText"))
	fmt.Println(Snake(">> samPLE text <<"))
	fmt.Println(Snake("sample 2 Text"))
	fmt.Println(Snake("SAMPLE 2 TEXT"))
	fmt.Println(Snake("inviteYourCustomersAddInvites"))
	fmt.Println(Snake("2FA Enabled"))
	fmt.Println(Snake("Enabled 2FA"))
	// Output:
	// sample_text
	// sample_text
	// sample_text
	// sample_text
	// sample_text
	// sam_ple_text
	// sample_2_text
	// sample_2_text
	// invite_your_customers_add_invites
	// 2fa_enabled
	// enabled_2fa
}

func ExampleSnake_withSpace() {
	fmt.Println(Snake("   sample   2    Text   "))
	// Output:
	// sample_2_text
}

func ExampleSnake_multiline() {
	fmt.Println(Snake("here\nis\na\nmultiline\nstring"))
	// Output:
	// here_is_a_multiline_string
}

func ExampleSnake_quoted() {
	fmt.Println(Snake("\"hello world\""))
	// Output:
	// hello_world
}

func ExampleSnake_swedish() {
	// Swedish and non ASCII char are not supported :(
	// open an issue if this is something you need
	//
	// want:
	// när-såg-du-en-kråka-väl-bita-en-man
	fmt.Println(Snake("När såg du en kråka väl bita en man?"))
	// Output:
	// n_r_s_g_du_en_kr_ka_v_l_bita_en_man
}

func ExampleSnake_nonASCII() {
	fmt.Println(Snake("   $#$sample   2    Text   "))
	fmt.Println(Snake("___$$Base64Encode"))
	fmt.Println(Snake("FOO#BAR#BAZ"))
	fmt.Println(Snake("FOO:BAR$BAZ"))
	fmt.Println(Snake("FOO#BAR#BAZ"))
	fmt.Println(Snake("something.com"))
	fmt.Println(Snake("$something%"))
	fmt.Println(Snake("something.com"))
	fmt.Println(Snake("•¶§ƒ˚foo˙∆˚¬"))
	fmt.Println(Snake("•¶§ƒ˚foo˙∆˚¬bar"))
	fmt.Println(Snake("•¶§ƒ˚foo˙∆˚¬ bar"))
	fmt.Println(Snake("CStringRef"))

	// Output:
	// sample_2_text
	// base64_encode
	// foo_bar_baz
	// foo_bar_baz
	// foo_bar_baz
	// something_com
	// something
	// something_com
	// foo
	// foo_bar
	// foo_bar
	// cstring_ref
}
