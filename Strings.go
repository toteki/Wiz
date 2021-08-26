package wiz

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"strings"
)

// Converts string to lowercase
func Lowercase(s string) string {
	return strings.ToLower(s)
}

// Converts string to upperxase
func Uppercase(s string) string {
	return strings.ToUpper(s)
}

// Attmepts to convert an object to string. Tries type assertion to string first, then looks for a .Stringer interface, then handles byte slices and arrays as base64, and finally attempts json marshaling. If none of those workes, returns a string representation of the input's type using reflect package.
func String(input interface{}) string {
	//First check for interface support
	s1, ok := input.(fmt.Stringer)
	if ok {
		return s1.String()
	}
	//Then try type assertion
	s2, ok := input.(string)
	if ok {
		return s2
	}
	//Then check for a pre-json special case (byte slices and arrays)
	byat := reflect.TypeOf(byte(0))
	if reflect.TypeOf(input) == reflect.SliceOf(byat) {
		return base64.StdEncoding.EncodeToString(input.([]byte))
	}
	//Byte arrays are harder to detect than slices
	v := reflect.ValueOf(input)
	k := v.Kind()
	if k == reflect.Array && reflect.TypeOf(input).Elem() == byat {
		vbs := reflect.ValueOf(make([]byte, v.Len()))
		reflect.Copy(vbs, v)
		return base64.StdEncoding.EncodeToString(vbs.Bytes())
	}
	//Then attempt JSON marshaling
	//(note: Oddly, byte slices become base64 but byte arrays become arrays of ints)
	j4, err := Marshal(input)
	if err == nil {
		return string(j4)
	}
	//If not, just string the data type
	return reflect.TypeOf(input).String()
}
