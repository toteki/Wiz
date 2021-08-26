package wiz

import sha3 "golang.org/x/crypto/sha3"

//		SHA3-512

// Hash returns the SHA3-512 of a given byte slice
func Hash(data []byte) []byte {
	//Takes a byte slice and outputs the SHA3-512 hash of it
	//Output is a 64 byte slice
	array := sha3.Sum512(data)
	return array[:]
}

// HashMatch checks if the SHA3-512 of a given byte slice matches a given 64 byte array
func HashMatch(data []byte, hash []byte) bool {
	//Does the hash of byteslice 'data' equal the byte slice 'hash'?
	if len(hash) != 64 {
		return false
	}
	a := [64]byte{}
	b := [64]byte{}
	copy(a[:], Hash(data))
	copy(b[:], hash)
	//Arrays can be compared directly using ==.
	return (a == b)
}
