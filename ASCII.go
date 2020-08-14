package wiz

import (
	"unicode"
)

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed functions:
//			ASCII(b []byte) (string, bool)
//				Converts binary to ASCII - returns false if non-ASCII bytes found
//			Printable(b []byte) (string, bool)
//				Converts binary to printable string - returns false if non-printable characters found

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
