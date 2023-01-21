package strcase

// Camel transforms the given string as `camelCase`.
func Camel(s string) string {
	b := make([]byte, 0, 64)
	l := len(s)
	i := 0

	for i < l {

		// skip leading bytes that aren't letters or digits
		for i < l && !isWord(s[i]) {
			i++
		}

		// set the first byte to uppercase if it needs to
		if i < l {
			c := s[i]

			// simply append contiguous digits
			if isDigit(c) {
				for i < l {
					if c = s[i]; !isDigit(c) {
						break
					}
					b = append(b, c)
					i++
				}
				continue
			}

			// the sequence starts with and uppercase letter, we append
			// all following uppercase letters as equivalent lowercase's
			if isUpper(c) {
				b = append(b, c)
				i++

				for i < l {
					if c = s[i]; !isUpper(c) {
						break
					}
					b = append(b, toLower(c))
					i++
				}

			} else {
				b = append(b, toUpper(c))
				i++
			}

			// append all trailing lowercase letters
			for i < l {
				if c = s[i]; !isLower(c) {
					break
				}
				b = append(b, c)
				i++
			}
		}
	}

	// the first byte must always be lowercase
	if len(b) != 0 {
		b[0] = toLower(b[0])
	}

	return string(b)
}
