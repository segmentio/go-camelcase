package strcase

import (
	"testing"
)

func BenchmarkPascalCase(b *testing.B) {
	s := "some sample text here_noething:too$amazing"
	b.SetBytes(int64(len(s)))
	for i := 0; i < b.N; i++ {
		PascalCase(s)
	}
}

func BenchmarkPascalUnchangedLong(b *testing.B) {
	var s = "invite_your_customers_add_invites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		PascalCase(s)
	}
}

func BenchmarkPascalUnchangedSimple(b *testing.B) {
	var s = "sample_text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		PascalCase(s)
	}
}

func BenchmarkPascalModifiedUnicode(b *testing.B) {
	var s = "ß_ƒ_foo"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		PascalCase(s)
	}
}

func BenchmarkPascalCaseASCIILong(b *testing.B) {
	var s = "some sample text here_noething:too$amazing"
	b.SetBytes(int64(len(s)))
	for i := 0; i < b.N; i++ {
		PascalCase(s)
	}
}

func BenchmarkPascalModifiedLong(b *testing.B) {
	var s = "inviteYourCustomersAddInvites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		PascalCase(s)
	}
}

func BenchmarkPascalModifiedLongSpecialChars(b *testing.B) {
	var s = "FOO:BAR$BAZ__Sample    Text___"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		PascalCase(s)
	}
}

func BenchmarkPascalModifiedSimple(b *testing.B) {
	var s = "sample text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		PascalCase("sample text")
	}
}

func BenchmarkPascalModifiedUnicode2(b *testing.B) {
	var s = "ẞ•¶§ƒ˚foo˙∆˚¬"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		PascalCase(s)
	}
}

func BenchmarkPascalLeadingUnderscoresDigitUpper(b *testing.B) {
	var s = "_5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		PascalCase(s)
	}
}

func BenchmarkPascalDigitUpper(b *testing.B) {
	var s = "5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		PascalCase(s)
	}
}

func BenchmarkPascalDigitUpper2(b *testing.B) {
	var s = "lk0B@bFmjrLQ_Z6YL"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		PascalCase(s)
	}
}
