package strcase

const dashByte = '-'

// Kebab transform the given string as `kebab-case`.
func Kebab(s string) string {
	idx := 0
	hasLower := false
	hasDash := false
	lowercaseSinceDash := false

	// loop through all good characters:
	// - lowercase
	// - digit
	// - underscore (as long as the next character is lowercase or digit)
	for ; idx < len(s); idx++ {
		if isLower(s[idx]) {
			hasLower = true
			if hasDash {
				lowercaseSinceDash = true
			}
			continue
		} else if isDigit(s[idx]) {
			continue
		} else if s[idx] == dashByte && idx > 0 && idx < len(s)-1 && (isLower(s[idx+1]) || isDigit(s[idx+1])) {
			hasDash = true
			lowercaseSinceDash = false
			continue
		}
		break
	}

	if idx == len(s) {
		return s // no changes needed, can just borrow the string
	}

	// if we get here then we must need to manipulate the string
	b := make([]byte, 0, 64)
	b = append(b, s[:idx]...)

	if isUpper(s[idx]) && (!hasLower || hasDash && !lowercaseSinceDash) {
		for idx < len(s) && (isUpper(s[idx]) || isDigit(s[idx])) {
			b = append(b, asciiLowercaseArray[s[idx]])
			idx++
		}

		for idx < len(s) && (isLower(s[idx]) || isDigit(s[idx])) {
			b = append(b, s[idx])
			idx++
		}
	}

	for idx < len(s) {
		if !isAlphanumeric(s[idx]) {
			idx++
			continue
		}

		if len(b) > 0 {
			b = append(b, dashByte)
		}

		for idx < len(s) && (isUpper(s[idx]) || isDigit(s[idx])) {
			b = append(b, asciiLowercaseArray[s[idx]])
			idx++
		}

		for idx < len(s) && (isLower(s[idx]) || isDigit(s[idx])) {
			b = append(b, s[idx])
			idx++
		}
	}
	return string(b)
}
