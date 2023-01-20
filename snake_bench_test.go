package strcase

import (
	"testing"
)

func BenchmarkSnakeCase(b *testing.B) {
	s := "some sample text here_noething:too$amazing"
	b.SetBytes(int64(len(s)))
	for i := 0; i < b.N; i++ {
		SnakeCase(s)
	}
}
func BenchmarkSnakeUnchangedLong(b *testing.B) {
	var s = "invite_your_customers_add_invites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		SnakeCase(s)
	}
}

func BenchmarkSnakeUnchangedSimple(b *testing.B) {
	var s = "sample_text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		SnakeCase(s)
	}
}

func BenchmarkSnakeModifiedUnicode(b *testing.B) {
	var s = "ß_ƒ_foo"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		SnakeCase(s)
	}
}

func BenchmarkSnakeModifiedLong(b *testing.B) {
	var s = "inviteYourCustomersAddInvites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		SnakeCase(s)
	}
}

func BenchmarkSnakeModifiedLongSpecialChars(b *testing.B) {
	var s = "FOO:BAR$BAZ__Sample    Text___"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		SnakeCase(s)
	}
}

func BenchmarkSnakeModifiedSimple(b *testing.B) {
	var s = "sample text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		SnakeCase("sample text")
	}
}

func BenchmarkSnakeModifiedUnicode2(b *testing.B) {
	var s = "ẞ•¶§ƒ˚foo˙∆˚¬"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		SnakeCase(s)
	}
}

func BenchmarkSnakeLeadingUnderscoresDigitUpper(b *testing.B) {
	var s = "_5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		SnakeCase(s)
	}
}

func BenchmarkSnakeDigitUpper(b *testing.B) {
	var s = "5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		SnakeCase(s)
	}
}

func BenchmarkSnakeDigitUpper2(b *testing.B) {
	var s = "lk0B@bFmjrLQ_Z6YL"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		SnakeCase(s)
	}
}
