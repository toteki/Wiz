package wiz

import (
	"github.com/pkg/errors"
	"reflect"
	"strconv"
)

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed functions:
//      Uint64(interface{}) (uint64, error)
//				Converts certain number types and strings to uint64.
//				Returns an error on failed conversion explaining why.

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

func Uint64(input interface{}) (uint64, error) {
	n := uint64(0)
	err := error(nil)
	switch v := reflect.ValueOf(input); v.Kind() {
	case reflect.String:
		n, err = strconv.ParseUint(input.(string), 10, 64)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := v.Int()
		if i < 0 {
			err = errors.New("Uint64: negative input")
		} else {
			n = uint64(i)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n = v.Uint()
	default:
		err = errors.New("Uint64: cannot handle type " + reflect.TypeOf(input).String())
	}
	return n, err
}
