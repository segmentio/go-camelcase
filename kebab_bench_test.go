package strcase

import (
	"testing"
)

func BenchmarkKebabUnchangedLong(b *testing.B) {
	s := "invite_your_customers_add_invites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Kebab(s)
	}
}

func BenchmarkKebabUnchangedSimple(b *testing.B) {
	s := "sample_text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Kebab(s)
	}
}

func BenchmarkKebabModifiedUnicode(b *testing.B) {
	s := "ß_ƒ_foo"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Kebab(s)
	}
}

func BenchmarkKebabCaseASCIILong(b *testing.B) {
	s := "some sample text here_noething:too$amazing"
	b.SetBytes(int64(len(s)))
	for i := 0; i < b.N; i++ {
		Kebab(s)
	}
}

func BenchmarkKebabModifiedLong(b *testing.B) {
	s := "inviteYourCustomersAddInvites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Kebab(s)
	}
}

func BenchmarkKebabModifiedLongSpecialChars(b *testing.B) {
	s := "FOO:BAR$BAZ__Sample    Text___"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Kebab(s)
	}
}

func BenchmarkKebabModifiedSimple(b *testing.B) {
	s := "sample text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Kebab("sample text")
	}
}

func BenchmarkKebabModifiedUnicode2(b *testing.B) {
	s := "ẞ•¶§ƒ˚foo˙∆˚¬"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Kebab(s)
	}
}

func BenchmarkKebabLeadingUnderscoresDigitUpper(b *testing.B) {
	s := "_5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Kebab(s)
	}
}

func BenchmarkKebabDigitUpper(b *testing.B) {
	s := "5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Kebab(s)
	}
}

func BenchmarkKebabDigitUpper2(b *testing.B) {
	s := "lk0B@bFmjrLQ_Z6YL"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Kebab(s)
	}
}
