//
// Fast camel-case implementation.
//
package camelcase

import "strings"

// camelCase the given `str`.
func Camelcase(str string) string {
	var b [1024]byte
	max := 1024
	l := len(str)
	ret := ""
	bi := 0
	i := 0
	first := true

	for i < l {

		for i < l && !isWord(str[i]) {
			i++
		}

		for i < l && isUpper(str[i]) {
			if bi < max {
				b[bi] = str[i]
				bi++
			}
			i++
		}

		for i < l && isPart(str[i]) {
			if bi < max {
				b[bi] = str[i]
				bi++
			}
			i++
		}

		for i < l && !isWord(str[i]) {
			i++
		}

		if first {
			ret += strings.ToLower(string(b[:bi]))
			first = false
		} else {
			// .ToTitle is weird in Go
			ret += strings.ToUpper(string(b[:1]))
			ret += strings.ToLower(string(b[1:bi]))
		}

		bi = 0
	}

	if len(ret) > 0 {
		ret = ret[:len(ret)]
	}

	return ret
}

func isPart(c byte) bool {
	return isLower(c) || isDigit(c)
}

func isWord(c byte) bool {
	return isLetter(c) || isDigit(c)
}

func isLetter(c byte) bool {
	return isLower(c) || isUpper(c)
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
