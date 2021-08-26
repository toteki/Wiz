package wiz

import (
	"regexp"
	"unicode"
)

// Converts binary to ASCII - returns false if non-ASCII bytes found
func ASCII(b []byte) (string, bool) {
	for _, c := range b {
		if c > unicode.MaxASCII {
			return "", false
		}
	}
	return string(b), true
}

// Converts binary to printable string - returns false if non-printable characters found
func Printable(b []byte) (string, bool) {
	for _, c := range b {
		if c > unicode.MaxASCII || !unicode.IsGraphic(rune(c)) {
			return "", false
		}
	}
	return string(b), true
}

// Strips non-ascii characters from a string
func StipNonASCII(in string) string {
	re := regexp.MustCompile("[[:^ascii:]]")
	return re.ReplaceAllLiteralString(in, "")
}

// Strips non-printable and non-ascii characters from a string
func StripNonPrintableASCII(in string) string {
	b := []byte(in)
	out := []byte{}
	for _, c := range b {
		if c <= unicode.MaxASCII && unicode.IsGraphic(rune(c)) {
			out = append(out, c)
		}
	}
	return string(out)
}
