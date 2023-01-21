package strcase

import (
	"testing"
)

func BenchmarkPascalCase(b *testing.B) {
	s := "some sample text here_noething:too$amazing"
	b.SetBytes(int64(len(s)))
	for i := 0; i < b.N; i++ {
		Pascal(s)
	}
}

func BenchmarkPascalUnchangedLong(b *testing.B) {
	s := "invite_your_customers_add_invites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Pascal(s)
	}
}

func BenchmarkPascalUnchangedSimple(b *testing.B) {
	s := "sample_text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Pascal(s)
	}
}

func BenchmarkPascalModifiedUnicode(b *testing.B) {
	s := "ß_ƒ_foo"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Pascal(s)
	}
}

func BenchmarkPascalCaseASCIILong(b *testing.B) {
	s := "some sample text here_noething:too$amazing"
	b.SetBytes(int64(len(s)))
	for i := 0; i < b.N; i++ {
		Pascal(s)
	}
}

func BenchmarkPascalModifiedLong(b *testing.B) {
	s := "inviteYourCustomersAddInvites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Pascal(s)
	}
}

func BenchmarkPascalModifiedLongSpecialChars(b *testing.B) {
	s := "FOO:BAR$BAZ__Sample    Text___"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Pascal(s)
	}
}

func BenchmarkPascalModifiedSimple(b *testing.B) {
	s := "sample text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Pascal("sample text")
	}
}

func BenchmarkPascalModifiedUnicode2(b *testing.B) {
	s := "ẞ•¶§ƒ˚foo˙∆˚¬"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Pascal(s)
	}
}

func BenchmarkPascalLeadingUnderscoresDigitUpper(b *testing.B) {
	s := "_5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Pascal(s)
	}
}

func BenchmarkPascalDigitUpper(b *testing.B) {
	s := "5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Pascal(s)
	}
}

func BenchmarkPascalDigitUpper2(b *testing.B) {
	s := "lk0B@bFmjrLQ_Z6YL"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Pascal(s)
	}
}
