package strcase

import (
	"testing"
)

func BenchmarkSnakeCase(b *testing.B) {
	s := "some sample text here_noething:too$amazing"
	b.SetBytes(int64(len(s)))
	for i := 0; i < b.N; i++ {
		Snake(s)
	}
}

func BenchmarkSnakeUnchangedLong(b *testing.B) {
	s := "invite_your_customers_add_invites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snake(s)
	}
}

func BenchmarkSnakeUnchangedSimple(b *testing.B) {
	s := "sample_text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snake(s)
	}
}

func BenchmarkSnakeModifiedUnicode(b *testing.B) {
	s := "ß_ƒ_foo"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snake(s)
	}
}

func BenchmarkSnakeModifiedLong(b *testing.B) {
	s := "inviteYourCustomersAddInvites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snake(s)
	}
}

func BenchmarkSnakeModifiedLongSpecialChars(b *testing.B) {
	s := "FOO:BAR$BAZ__Sample    Text___"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snake(s)
	}
}

func BenchmarkSnakeModifiedSimple(b *testing.B) {
	s := "sample text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snake("sample text")
	}
}

func BenchmarkSnakeModifiedUnicode2(b *testing.B) {
	s := "ẞ•¶§ƒ˚foo˙∆˚¬"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snake(s)
	}
}

func BenchmarkSnakeLeadingUnderscoresDigitUpper(b *testing.B) {
	s := "_5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snake(s)
	}
}

func BenchmarkSnakeDigitUpper(b *testing.B) {
	s := "5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snake(s)
	}
}

func BenchmarkSnakeDigitUpper2(b *testing.B) {
	s := "lk0B@bFmjrLQ_Z6YL"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snake(s)
	}
}
