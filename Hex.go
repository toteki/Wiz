package wiz

import (
	"encoding/hex"
	"github.com/pkg/errors"
	"strings"
)

// BytesToHex Converts byte slice to uppercase hexadecimal string representation
func BytesToHex(data []byte) string {
	str := hex.EncodeToString(data)
	return strings.ToUpper(str)
}

// HexToBytes Converts string to byte slice if it is valid hexadecimal. Slice length is zero on fail.
func HexToBytes(data string) ([]byte, error) {
	b, err := hex.DecodeString(data)
	if err != nil {
		return []byte{}, errors.Wrap(err, "wiz.HexToBytes")
	}
	return b, nil
}
