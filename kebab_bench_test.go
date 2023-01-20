package strcase

import (
	"testing"
)

func BenchmarkKebabUnchangedLong(b *testing.B) {
	var s = "invite_your_customers_add_invites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		KebabCase(s)
	}
}

func BenchmarkKebabUnchangedSimple(b *testing.B) {
	var s = "sample_text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		KebabCase(s)
	}
}

func BenchmarkKebabModifiedUnicode(b *testing.B) {
	var s = "ß_ƒ_foo"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		KebabCase(s)
	}
}

func BenchmarkKebabCaseASCIILong(b *testing.B) {
	var s = "some sample text here_noething:too$amazing"
	b.SetBytes(int64(len(s)))
	for i := 0; i < b.N; i++ {
		KebabCase(s)
	}
}

func BenchmarkKebabModifiedLong(b *testing.B) {
	var s = "inviteYourCustomersAddInvites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		KebabCase(s)
	}
}

func BenchmarkKebabModifiedLongSpecialChars(b *testing.B) {
	var s = "FOO:BAR$BAZ__Sample    Text___"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		KebabCase(s)
	}
}

func BenchmarkKebabModifiedSimple(b *testing.B) {
	var s = "sample text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		KebabCase("sample text")
	}
}

func BenchmarkKebabModifiedUnicode2(b *testing.B) {
	var s = "ẞ•¶§ƒ˚foo˙∆˚¬"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		KebabCase(s)
	}
}

func BenchmarkKebabLeadingUnderscoresDigitUpper(b *testing.B) {
	var s = "_5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		KebabCase(s)
	}
}

func BenchmarkKebabDigitUpper(b *testing.B) {
	var s = "5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		KebabCase(s)
	}
}

func BenchmarkKebabDigitUpper2(b *testing.B) {
	var s = "lk0B@bFmjrLQ_Z6YL"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		KebabCase(s)
	}
}
