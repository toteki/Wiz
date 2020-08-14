package wiz

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
)

//		This package is json.Marshal but with some extra formatting.
//		Basically proper indenting (increments of three ASCII spaces) and
//			newlines are added. Output is normally bytes but can be string using
//			the shortcut function.

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed functions:
//			CompactJSON(data []byte) ([]byte, error)
//				Takes some json and removes unnecessary whitespace
//			NeatJSON(data []byte) ([]byte, error)
//				Takes some json and adds newlines and indentation for readability
//			Marshal(payload interface{}) ([]byte, error)
//				Takes a payload, and turns it into JSON
//			MarshalNeat(payload interface{})  ([]byte, error)
//				Takes a payload, and turns it into neatly formatted JSON

//		Parameters of type interface{} basically mean "just pass any type"

//From docs on encoding/json.Marshal:

// So that the JSON will be safe to embed inside HTML <script> tags, the string
// is encoded using HTMLEscape, which replaces "<", ">", "&", U+2028, and U+2029
// are escaped to "\u003c","\u003e", "\u0026", "\u2028", and "\u2029". This
// replacement can be disabled when using an Encoder,
// by calling SetEscapeHTML(false).

// Which we shall do.

var jsonencoder *json.Encoder = nil
var jsoninitialized = false
var jsonencbuf = new(bytes.Buffer)
var jsonlocked = false

func jsonInit() {
	if !jsoninitialized {
		jsonencoder = json.NewEncoder(jsonencbuf)
		jsonencoder.SetEscapeHTML(false)
	}
	jsoninitialized = true
}

func jsonLock() {
	for jsonlocked {

	}
	jsonInit()
	jsonlocked = true //Lock encbuf
	jsonencbuf.Reset()
}

func jsonUnlock() {
	jsonlocked = false
}

////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////

func NeatJSON(input []byte) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := json.Indent(buff, input, "", "   ")
	if err != nil {
		return []byte{}, errors.Wrap(err, "NeatJSON")
	}
	return buff.Bytes(), nil
}

func CompactJSON(input []byte) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := json.Compact(buff, input)
	if err != nil {
		return []byte{}, errors.Wrap(err, "CompactJSON")
	}
	return buff.Bytes(), nil
}

func Marshal(payload interface{}) ([]byte, error) {
	jsonLock()
	defer jsonUnlock()
	err := jsonencoder.Encode(payload)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Marshal")
	}
	b := jsonencbuf.Bytes()
	return b, nil
}

func MarshalNeat(payload interface{}) ([]byte, error) {
	b, err := Marshal(payload)
	if err != nil {
		return []byte{}, errors.Wrap(err, "MarshalNeat")
	}
	n, err := NeatJSON(b)
	if err != nil {
		return []byte{}, errors.Wrap(err, "MarshalNeat")
	}
	return n, nil
}
