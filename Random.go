package wiz

import (
	"crypto/rand"
	"github.com/pkg/errors"
)

// RandomBytes returns byte slice of length n filled with random data (uses crypto/rand)
func RandomBytes(len int) ([]byte, error) {
	if len < 0 {
		return []byte{}, errors.New("wiz.RandomBytes: Cannot pass negative length")
	}
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return []byte{}, errors.Wrap(err, "wiz.RandomBytes")
	}
	return b, nil
}
