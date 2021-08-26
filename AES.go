package wiz

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/pkg/errors"
)

// Encrypts given data using a key. AES256, GCM mode. Uses crypto/rand for nonce.
func AESEncrypt(data []byte, key []byte) ([]byte, error) {

	if len(key) != 32 {
		return []byte{}, errors.New("wiz.AESEncrypt: Key should be 32 bytes long")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, errors.Wrap(err, "wiz.AESEncrypt: Failed to create cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, errors.Wrap(err, "wiz.AESEncrypt: Failed to create gcm")
	}

	nonce := make([]byte, 32)
	_, err = rand.Read(nonce)
	if err != nil {
		return []byte{}, errors.Wrap(err, "wiz.AESEncrypt: Failed to create nonce")
	}

	encrypted := gcm.Seal(nil, nonce, data, nil)

	return encrypted, nil
}

// Decrypts given data using a key. AES256, GCM mode.
func AESDecrypt(stream []byte, key []byte) ([]byte, error) {

	if len(key) != 32 {
		return []byte{}, errors.New("wiz.AESDecrypt: Key should be 32 bytes long")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, errors.Wrap(err, "wiz.AESDecrypt: Failed to create cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, errors.Wrap(err, "wiz.AESDecrypt: Failed to create gcm")
	}

	if len(stream) < gcm.NonceSize() {
		return []byte{}, errors.Wrap(err, "wiz.AESDecrypt: stream is shorter than gcm.NonceSize()")
	}

	nonce := stream[:gcm.NonceSize()]
	stream = stream[gcm.NonceSize():]

	decrypted, err := gcm.Open(nil, nonce, stream, nil)
	if err != nil {
		return []byte{}, errors.Wrap(err, "wiz.AESDecrypt: Failed to decrypt")
	}
	return decrypted, nil
}
