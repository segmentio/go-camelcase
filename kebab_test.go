package strcase

import (
	"fmt"
	"testing"
)

func TestKebabCase(t *testing.T) {
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
		if out := KebabCase(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func ExampleKebabCase() {
	fmt.Println(KebabCase("sample text"))
	fmt.Println(KebabCase("sample-text"))
	fmt.Println(KebabCase("sample_text"))
	fmt.Println(KebabCase("sample___text"))
	fmt.Println(KebabCase("sampleText"))
	fmt.Println(KebabCase(">> samPLE text <<"))
	fmt.Println(KebabCase("sample 2 Text"))
	fmt.Println(KebabCase("SAMPLE 2 TEXT"))
	fmt.Println(KebabCase("inviteYourCustomersAddInvites"))
	fmt.Println(KebabCase("2FA Enabled"))
	fmt.Println(KebabCase("Enabled 2FA"))
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

func ExampleKebabCase_withSpace() {
	fmt.Println(KebabCase("   sample   2    Text   "))
	// Output:
	// sample-2-text
}

func ExampleKebabCase_long() {
	fmt.Println(KebabCase("super long sentence that is not really necessary but we need to try it"))
	// Output:
	// super-long-sentence-that-is-not-really-necessary-but-we-need-to-try-it
}

func ExampleKebabCase_multiline() {
	fmt.Println(KebabCase("here\nis\na\nmultiline\nstring"))
	// Output:
	// here-is-a-multiline-string
}

func ExampleKebabCase_quoted() {
	fmt.Println(KebabCase("\"hello world\""))
	// Output:
	// hello-world
}

func ExampleKebabCase_swedish() {
	// Swedish and non ASCII char are not supported :(
	// open an issue if this is something you need
	//
	// want:
	// när-såg-du-en-kråka-väl-bita-en-man
	fmt.Println(KebabCase("När såg du en kråka väl bita en man?"))
	// Output:
	// n-r-s-g-du-en-kr-ka-v-l-bita-en-man
}

func ExampleKebabCase_nonASCII() {
	fmt.Println(KebabCase("   $#$sample   2    Text   "))
	fmt.Println(KebabCase("___$$Base64Encode"))
	fmt.Println(KebabCase("FOO#BAR#BAZ"))
	fmt.Println(KebabCase("FOO:BAR$BAZ"))
	fmt.Println(KebabCase("FOO#BAR#BAZ"))
	fmt.Println(KebabCase("something.com"))
	fmt.Println(KebabCase("$something%"))
	fmt.Println(KebabCase("something.com"))
	fmt.Println(KebabCase("•¶§ƒ˚foo˙∆˚¬"))
	fmt.Println(KebabCase("•¶§ƒ˚foo˙∆˚¬bar"))
	fmt.Println(KebabCase("•¶§ƒ˚foo˙∆˚¬ bar"))
	fmt.Println(KebabCase("CStringRef"))

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

func ExampleKebabCase_cryptic() {
	fmt.Println(KebabCase("5test"))
	fmt.Println(KebabCase("test5"))
	fmt.Println(KebabCase("THE5r"))
	fmt.Println(KebabCase("5TEst"))
	fmt.Println(KebabCase("_5TEst"))
	fmt.Println(KebabCase("@%#&5TEst"))
	fmt.Println(KebabCase("edf_6N"))
	fmt.Println(KebabCase("f_pX9"))
	fmt.Println(KebabCase("p_z9Rg"))
	fmt.Println(KebabCase("@49L0S145_¬fwHƒ0TSLNVp"))
	fmt.Println(KebabCase("lk0B@bFmjrLQ_Z6YL"))
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
