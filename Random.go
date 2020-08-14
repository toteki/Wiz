package wiz

import (
	"crypto/rand"
	"github.com/pkg/errors"
)

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed functions:
//			RandomBytes(len int) ([]byte, error)
//				Generates random binary of a given length in bytes

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

// RandomBytes returns byte slice of length n filled with random data
func RandomBytes(len int) ([]byte, error) {
	if len < 0 {
		return []byte{}, errors.New("RandomBytes: Cannot pass negative length")
	}
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return []byte{}, errors.Wrap(err, "RandomBytes")
	}
	return b, nil
}
