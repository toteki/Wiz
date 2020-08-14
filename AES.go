package wiz

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/pkg/errors"
)

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed functions:
//		  AESEncrypt(data []byte, key []byte) ([]byte, error)
//				Uses AES256, GCM mode. crypto/rand generates nonce each time.
//			AESDecrypt(stream []byte, key []byte) ([]byte, error)
//				Uses AES256, GCM mode. Expects nonce to be first 32 bytes of stream.

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

func AESEncrypt(data []byte, key []byte) ([]byte, error) {

	if len(key) != 32 {
		return []byte{}, errors.New("AESEncrypt: Key should be 32 bytes long")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, errors.Wrap(err, "AESEncrypt: Failed to create cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, errors.Wrap(err, "AESEncrypt: Failed to create gcm")
	}

	nonce := make([]byte, 32)
	_, err = rand.Read(nonce)
	if err != nil {
		return []byte{}, errors.Wrap(err, "AESEncrypt: Failed to create nonce")
	}

	encrypted := gcm.Seal(nil, nonce, data, nil)

	return encrypted, nil
}

func AESDecrypt(stream []byte, key []byte) ([]byte, error) {

	if len(key) != 32 {
		return []byte{}, errors.New("AESDecrypt: Key should be 32 bytes long")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, errors.Wrap(err, "AESDecrypt: Failed to create cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, errors.Wrap(err, "AESDecrypt: Failed to create gcm")
	}

	if len(stream) < gcm.NonceSize() {
		return []byte{}, errors.Wrap(err, "AESDecrypt: stream is shorter than gcm.NonceSize()")
	}

	nonce := stream[:gcm.NonceSize()]
	stream = stream[gcm.NonceSize():]

	decrypted, err := gcm.Open(nil, nonce, stream, nil)
	if err != nil {
		return []byte{}, errors.Wrap(err, "AESDecrypt: Failed to decrypt")
	}
	return decrypted, nil
}
