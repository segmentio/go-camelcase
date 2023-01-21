package strcase

import (
	"fmt"
	"testing"
)

func TestKebab(t *testing.T) {
	samples := []TT{
		{"sample text", "sample-text"},
		{"sample-text", "sample-text"},
		{"sample_text", "sample-text"},
		{"sample___text", "sample-text"},
		{"sampleText", "sample-text"},
		{"inviteYourCustomersAddInvites", "invite-your-customers-add-invites"},
		{"sample 2 Text", "sample-2-text"},
		{"   sample   2    Text   ", "sample-2-text"},
		{"   $#$sample   2    Text   ", "sample-2-text"},
		{"SAMPLE 2 TEXT", "sample-2-text"},
		{"___$$Base64Encode", "base64-encode"},
		{"FOO:BAR$BAZ", "foo-bar-baz"},
		{"FOO#BAR#BAZ", "foo-bar-baz"},
		{"something.com", "something-com"},
		{"$something%", "something"},
		{"something.com", "something-com"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
		{"•¶§ƒ˚foo˙∆˚¬bar", "foo-bar"},
		{"•¶§ƒ˚foo˙∆˚¬ bar", "foo-bar"},
		{"CStringRef", "cstring-ref"},
		{"5test", "5test"},
		{"test5", "test5"},
		{"THE5r", "the5r"},
		{"5TEst", "5test"},
		{"_5TEst", "5test"},
		{"@%#&5TEst", "5test"},
		{"edf_6N", "edf-6n"},
		{"f_pX9", "f-p-x9"},
		{"p_z9Rg", "p-z9-rg"},
		{"2FA Enabled", "2fa-enabled"},
		{"Enabled 2FA", "enabled-2fa"},
		{"@49L0S145_¬fwHƒ0TSLNVp", "49l0s145-fw-h-0tslnvp"},
		{"lk0B@bFmjrLQ_Z6YL", "lk0-b-b-fmjr-lq-z6yl"},
		{"samPLE text", "sam-ple-text"},
		{"here\nis\na\nmultiline\nstring", "here-is-a-multiline-string"},
		{"\"hello world\"", "hello-world"},
		// {
		// Swedish and non ASCII char are not supported :(
		// open an issue if this is something you need
		// 	input: "När såg du en kråka väl bita en man?"
		// 	got: n-r-s-g-du-en-kr-ka-v-l-bita-en-man
		// 	want: "när-såg-du-en-kråka-väl-bita-en",
		// },
		{
			"super long sentence that is not really necessary but we need to try it",
			"super-long-sentence-that-is-not-really-necessary-but-we-need-to-try-it",
		},
	}

	for _, sample := range samples {
		if out := Kebab(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func ExampleKebab() {
	fmt.Println(Kebab("sample text"))
	fmt.Println(Kebab("sample-text"))
	fmt.Println(Kebab("sample_text"))
	fmt.Println(Kebab("sample___text"))
	fmt.Println(Kebab("sampleText"))
	fmt.Println(Kebab(">> samPLE text <<"))
	fmt.Println(Kebab("sample 2 Text"))
	fmt.Println(Kebab("SAMPLE 2 TEXT"))
	fmt.Println(Kebab("inviteYourCustomersAddInvites"))
	fmt.Println(Kebab("2FA Enabled"))
	fmt.Println(Kebab("Enabled 2FA"))
	// Output:
	// sample-text
	// sample-text
	// sample-text
	// sample-text
	// sample-text
	// sam-ple-text
	// sample-2-text
	// sample-2-text
	// invite-your-customers-add-invites
	// 2fa-enabled
	// enabled-2fa
}

func ExampleKebab_withSpace() {
	fmt.Println(Kebab("   sample   2    Text   "))
	// Output:
	// sample-2-text
}

func ExampleKebab_long() {
	fmt.Println(Kebab("super long sentence that is not really necessary but we need to try it"))
	// Output:
	// super-long-sentence-that-is-not-really-necessary-but-we-need-to-try-it
}

func ExampleKebab_multiline() {
	fmt.Println(Kebab("here\nis\na\nmultiline\nstring"))
	// Output:
	// here-is-a-multiline-string
}

func ExampleKebab_quoted() {
	fmt.Println(Kebab("\"hello world\""))
	// Output:
	// hello-world
}

func ExampleKebab_swedish() {
	// Swedish and non ASCII char are not supported :(
	// open an issue if this is something you need
	//
	// want:
	// när-såg-du-en-kråka-väl-bita-en-man
	fmt.Println(Kebab("När såg du en kråka väl bita en man?"))
	// Output:
	// n-r-s-g-du-en-kr-ka-v-l-bita-en-man
}

func ExampleKebab_nonASCII() {
	fmt.Println(Kebab("   $#$sample   2    Text   "))
	fmt.Println(Kebab("___$$Base64Encode"))
	fmt.Println(Kebab("FOO#BAR#BAZ"))
	fmt.Println(Kebab("FOO:BAR$BAZ"))
	fmt.Println(Kebab("FOO#BAR#BAZ"))
	fmt.Println(Kebab("something.com"))
	fmt.Println(Kebab("$something%"))
	fmt.Println(Kebab("something.com"))
	fmt.Println(Kebab("•¶§ƒ˚foo˙∆˚¬"))
	fmt.Println(Kebab("•¶§ƒ˚foo˙∆˚¬bar"))
	fmt.Println(Kebab("•¶§ƒ˚foo˙∆˚¬ bar"))
	fmt.Println(Kebab("CStringRef"))

	// Output:
	// sample-2-text
	// base64-encode
	// foo-bar-baz
	// foo-bar-baz
	// foo-bar-baz
	// something-com
	// something
	// something-com
	// foo
	// foo-bar
	// foo-bar
	// cstring-ref
}

func ExampleKebab_cryptic() {
	fmt.Println(Kebab("5test"))
	fmt.Println(Kebab("test5"))
	fmt.Println(Kebab("THE5r"))
	fmt.Println(Kebab("5TEst"))
	fmt.Println(Kebab("_5TEst"))
	fmt.Println(Kebab("@%#&5TEst"))
	fmt.Println(Kebab("edf_6N"))
	fmt.Println(Kebab("f_pX9"))
	fmt.Println(Kebab("p_z9Rg"))
	fmt.Println(Kebab("@49L0S145_¬fwHƒ0TSLNVp"))
	fmt.Println(Kebab("lk0B@bFmjrLQ_Z6YL"))
	// Output:
	// 5test
	// test5
	// the5r
	// 5test
	// 5test
	// 5test
	// edf-6n
	// f-p-x9
	// p-z9-rg
	// 49l0s145-fw-h-0tslnvp
	// lk0-b-b-fmjr-lq-z6yl
}
