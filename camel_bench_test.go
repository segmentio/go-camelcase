package strcase

import (
	"testing"
)

func BenchmarkCamelCase(b *testing.B) {
	s := "some sample text here_noething:too$amazing"
	b.SetBytes(int64(len(s)))
	for i := 0; i < b.N; i++ {
		Camel(s)
	}
}

func BenchmarkCamelCase2UnchangedLong(b *testing.B) {
	s := "invite_your_customers_add_invites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Camel(s)
	}
}

func BenchmarkCamelCase2UnchangedSimple(b *testing.B) {
	s := "sample_text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Camel(s)
	}
}

func BenchmarkCamelCase2ModifiedUnicode(b *testing.B) {
	s := "ß_ƒ_foo"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Camel(s)
	}
}

func BenchmarkCamelCase2CaseASCIILong(b *testing.B) {
	s := "some sample text here_noething:too$amazing"
	b.SetBytes(int64(len(s)))
	for i := 0; i < b.N; i++ {
		Camel(s)
	}
}

func BenchmarkCamelCase2ModifiedLong(b *testing.B) {
	s := "inviteYourCustomersAddInvites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Camel(s)
	}
}

func BenchmarkCamelCase2ModifiedLongSpecialChars(b *testing.B) {
	s := "FOO:BAR$BAZ__Sample    Text___"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Camel(s)
	}
}

func BenchmarkCamelCase2ModifiedSimple(b *testing.B) {
	s := "sample text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Camel("sample text")
	}
}

func BenchmarkCamelCase2ModifiedUnicode2(b *testing.B) {
	s := "ẞ•¶§ƒ˚foo˙∆˚¬"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Camel(s)
	}
}

func BenchmarkCamelCase2LeadingUnderscoresDigitUpper(b *testing.B) {
	s := "_5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Camel(s)
	}
}

func BenchmarkCamelCase2DigitUpper(b *testing.B) {
	s := "5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Camel(s)
	}
}

func BenchmarkCamelCase2DigitUpper2(b *testing.B) {
	s := "lk0B@bFmjrLQ_Z6YL"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Camel(s)
	}
}
