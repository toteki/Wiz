package wiz

import (
	"regexp"
	"unicode"
)

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed functions:
//			ASCII(b []byte) (string, bool)
//				Converts binary to ASCII - returns false if non-ASCII bytes found
//			Printable(b []byte) (string, bool)
//				Converts binary to printable string - returns false if non-printable characters found
//			StripNonASCII(in string) string
//				Strips non-ascii characters from a string
//			StripNonPrintableASCII(in string) string
//				Strips non-printable and non-ascii characters from a string

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

func ASCII(b []byte) (string, bool) {
	for _, c := range b {
		if c > unicode.MaxASCII {
			return "", false
		}
	}
	return string(b), true
}

func Printable(b []byte) (string, bool) {
	for _, c := range b {
		if c > unicode.MaxASCII || !unicode.IsGraphic(rune(c)) {
			return "", false
		}
	}
	return string(b), true
}

func StipNonASCII(in string) string {
	re := regexp.MustCompile("[[:^ascii:]]")
	return re.ReplaceAllLiteralString(in, "")
}

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
