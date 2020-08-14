package wiz

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed functions:
//			Lowercase(string) string
//			Uppercase(string) string
//      String(interface{}) string

//		Never returns errors. However, String(input) will, on
//      encountering an input which cannot be handled, merely print
//      its type instead of its value. This should be extremely rare.

//		Note: Contains unexported supporting functions from JSON.go

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

func Lowercase(s string) string {
	return strings.ToLower(s)
}

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func String(input interface{}) string {
	//First try type assertion
	s1, ok := input.(string)
	if ok {
		return s1
	}
	//Then check for interface support
	s2, ok := input.(fmt.Stringer)
	if ok {
		return s2.String()
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
	j4, err := sMarshal(input)
	if err == nil {
		return string(j4)
	}
	//If not, just string the data type
	return reflect.TypeOf(input).String()
}

//
//	Below: an entire clone of the JSON.go file
//	Used in the json marshaling case above
//	This close exists so that the Strings.go file
//	Remains independent (can be copypasted at will)
//

var sjsonencoder *json.Encoder = nil
var sjsoninitialized = false
var sjsonencbuf = new(bytes.Buffer)
var sjsonlocked = false

func sjsonInit() {
	if !sjsoninitialized {
		sjsonencoder = json.NewEncoder(sjsonencbuf)
		sjsonencoder.SetEscapeHTML(false)
	}
	sjsoninitialized = true
}

func sjsonLock() {
	for sjsonlocked {

	}
	sjsonInit()
	sjsonlocked = true //Lock encbuf
	sjsonencbuf.Reset()
}

func sjsonUnlock() {
	sjsonlocked = false
}

func sMarshal(payload interface{}) ([]byte, error) {
	sjsonLock()
	defer sjsonUnlock()
	err := sjsonencoder.Encode(payload)
	if err != nil {
		return []byte{}, errors.Wrap(err, "sMarshal")
	}
	b := sjsonencbuf.Bytes()
	return b, nil
}
