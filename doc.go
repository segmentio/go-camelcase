// Package strcase implements fast snake_case, kebab-case, camelCase and Pascal conversions.
// The `strcase` package provides utility functions for converting any strings to various case formats.
// It can convert any string to:
//
// - [camelCase](https://en.wikipedia.org/wiki/Letter_case#Camel_case) or lower camel case:
//   - example: "theQuickBrownFoxJumpsOverTheLazyDog", `type myInternalType struct {}`
//   - usage: Internal (private) variables, functions, methods, and types in Go
//
// - Pascal or [upper camel case](https://en.wikipedia.org/wiki/Letter_case#Camel_case):
//   - example: "TheQuickBrownFoxJumpsOverTheLazyDog", `type MyExportedType struct {}`
//   - usage: Exported (public) variables, functions, methods, and types in Go
//
// - [snake_case](https://en.wikipedia.org/wiki/Letter_case#Snake_case):
//   - example: "the_quick_brown_fox_jumps_over_the_lazy_dog"
//   - usage: naming convention in Python.
//
// - dash-case or [kebab-case](https://en.wikipedia.org/wiki/Letter_case#Kebab_case):
//   - "the-quick-brown-fox-jumps-over-the-lazy-dog"
//   - usage: naming convention in CSS, also used in HTML and kubernetes manifests.
//
// Note:
//
// the "dash" is actually an ASCII hyphen a.k.a "hyphen-minus" a.k.a "minus sign", unicode `U+002D`,
// represented as `&#x002D;` is often confused
// with "hyphen", unicode `U+2010`, represented as `&#x2010;`
// or with "En Dash" unicode `U+2013`,  represented as `&#x2013;` .
//
// see [Wikipedia](https://en.wikipedia.org/wiki/Dash#Unicode) for more details
// and [Unicode ASCII punctuation](https://www.unicode.org/charts/PDF/U0000.pdf) for the full list of dashes.
// I'm no expert in this area, it seems quite complicated, so if you have any suggestions, please open an issue and let us know.
package strcase
