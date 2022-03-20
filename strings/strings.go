package strings

import "unicode"

func ToUpper(x string) (string, error) {
	r := []rune(x)
	for i := range r {
		r[i] = unicode.ToUpper(r[i])
	}
	return string(r), nil
}
