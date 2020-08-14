package wiz

import (
	"encoding/hex"
	"github.com/pkg/errors"
	"strings"
)

//		Bytes to hexadecimal and back. We output uppercase.

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed functions:
//			BytesToHex(data []byte) string
//			HexToBytes(data string) ([]byte, error)

//		HexToBytes error means string wasn't an even number of 0-9,a-f,A-F chars

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

// BytesToHex Converts byte slice to string representation (hexadecimal).
func BytesToHex(data []byte) string {
	str := hex.EncodeToString(data)
	return strings.ToUpper(str)
}

// HexToBytes Converts string to byte slice if it is valid hexadecimal. Slice length is zero on fail.
func HexToBytes(data string) ([]byte, error) {
	b, err := hex.DecodeString(data)
	if err != nil {
		return []byte{}, errors.Wrap(err, "HexToBytes")
	}
	return b, nil
}
