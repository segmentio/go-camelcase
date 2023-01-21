package strcase

import (
	"bytes"
	"testing"
)

func BenchmarkBufferSnakeUnchangedLong(b *testing.B) {
	s := "invite_your_customers_add_invites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		snakeByteBuffer(s)
	}
}

func BenchmarkBufferSnakeUnchangedSimple(b *testing.B) {
	s := "sample_text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		snakeByteBuffer(s)
	}
}

func BenchmarkBufferSnakeModifiedUnicode(b *testing.B) {
	s := "ß_ƒ_foo"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		snakeByteBuffer(s)
	}
}

func BenchmarkBufferSnakeCaseASCIILong(b *testing.B) {
	s := "some sample text here_noething:too$amazing"
	b.SetBytes(int64(len(s)))
	for i := 0; i < b.N; i++ {
		snakeByteBuffer(s)
	}
}

func BenchmarkBufferSnakeModifiedLong(b *testing.B) {
	s := "inviteYourCustomersAddInvites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		snakeByteBuffer(s)
	}
}

func BenchmarkBufferSnakeModifiedLongSpecialChars(b *testing.B) {
	s := "FOO:BAR$BAZ__Sample    Text___"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		snakeByteBuffer(s)
	}
}

func BenchmarkBufferSnakeModifiedSimple(b *testing.B) {
	s := "sample text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		snakeByteBuffer("sample text")
	}
}

func BenchmarkBufferSnakeModifiedUnicode2(b *testing.B) {
	s := "ẞ•¶§ƒ˚foo˙∆˚¬"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		snakeByteBuffer(s)
	}
}

func BenchmarkBufferSnakeLeadingUnderscoresDigitUpper(b *testing.B) {
	s := "_5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		snakeByteBuffer(s)
	}
}

func BenchmarkBufferSnakeDigitUpper(b *testing.B) {
	s := "5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		snakeByteBuffer(s)
	}
}

func BenchmarkBufferSnakeDigitUpper2(b *testing.B) {
	s := "lk0B@bFmjrLQ_Z6YL"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		snakeByteBuffer(s)
	}
}

// snakeByteBuffer returns the snake-case of the given string.
// it is the same function as Snake but uses a bytes.Buffer instead of []byte
// in order to test the performance of the bytes.Buffer vs []byte.
func snakeByteBuffer(s string) string { // nolint: unparam
	idx := 0
	hasLower := false
	hasUnderscore := false
	lowercaseSinceUnderscore := false

	// loop through all good characters:
	// - lowercase
	// - digit
	// - underscore (as long as the next character is lowercase or digit)
	for ; idx < len(s); idx++ {
		if isLower(s[idx]) {
			hasLower = true
			if hasUnderscore {
				lowercaseSinceUnderscore = true
			}
			continue
		} else if isDigit(s[idx]) {
			continue
		} else if s[idx] == underscoreByte && idx > 0 && idx < len(s)-1 && (isLower(s[idx+1]) || isDigit(s[idx+1])) {
			hasUnderscore = true
			lowercaseSinceUnderscore = false
			continue
		}
		break
	}

	if idx == len(s) {
		return s // no changes needed, can just borrow the string
	}

	// if we get here then we must need to manipulate the string
	var b bytes.Buffer
	b.WriteString(s[:idx])

	if isUpper(s[idx]) && (!hasLower || hasUnderscore && !lowercaseSinceUnderscore) {
		for idx < len(s) && (isUpper(s[idx]) || isDigit(s[idx])) {
			b.WriteByte(asciiLowercaseArray[s[idx]])
			idx++
		}

		for idx < len(s) && (isLower(s[idx]) || isDigit(s[idx])) {
			b.WriteByte(s[idx])
			idx++
		}
	}

	for idx < len(s) {
		if !isAlphanumeric(s[idx]) {
			idx++
			continue
		}

		if b.Len() > 0 {
			b.WriteByte(underscoreByte)
		}

		for idx < len(s) && (isUpper(s[idx]) || isDigit(s[idx])) {
			b.WriteByte(asciiLowercaseArray[s[idx]])
			idx++
		}

		for idx < len(s) && (isLower(s[idx]) || isDigit(s[idx])) {
			b.WriteByte(s[idx])
			idx++
		}
	}
	return b.String()
}
