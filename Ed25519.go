package wiz

import (
	"crypto/ed25519"
	"github.com/pkg/errors"
)

//		TODO: Add ephemeral key generation and diffie hellmann, if needed.
//			I'm sure the parent package has it, so just go there first.

//		Ed25519 cryptography. (Elliptic curve)

//		Seed size: 32 bytes (key generation)
//		Public key size: 32 bytes
//		Private key size: 64 bytes
//		Signature size: 64 bytes
//		Shared secret size: 32 bytes (diffie hellman output)

//		Note that in ECC there is no "encrypt/decrypt". This is not RSA.
//		The only things you can	do with an elliptic keypair are
//			-	For one thing, create the public/private key pair (requires randomness)
//			- Create a signature of a piece of data, using your private key
//			- Verify a signature of a piece of data, using the signer's public key
//			- Knowing your own private key and someone else's public key, generate
//				a shared secret based on the two (diffie hellmann). It's the same
//				output every time you do it though, so in practice "ephemeral" keys
//				are used, where each side generates a random single-use keypair,
//				signs the public key with their real (long term) public key, and then
//				using each side's ephemeral key pair, do diffie hellmann as above.
//				End result is both sides have the same shared 32 byte secret which
//				can be used for anything, including symmetric encryption like AES.
//				That secret is then called the session key.

//		If you're doing ECC, Ed25519 is beter than ECDSA (on any curve) because
//			ECDSA requires a random seed for every signature, and any weakness in
//			the random generator leaks the private key (happened in the playstation
//			breach). Ed25519 manages to be a secure signature scheme based on the
//			same math (it's still ECC) but using a deterministic input. You still
//			need good randoms when generating the keys but after that no entropy
//			is required and shitty hardware randoms can't kill you.

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed functions:
//			NewEdKeyPair(seed []byte) ([]byte, []byte, error)
//				The returns are public key, then private key. Don't mess up the order.
//			EdSign(data, publicKey, privateKey []byte) ([]byte, error)
//				Does auto-verify the signature based on public key provided.
//			EdVerify(data, signature, publicKey []byte) error
//				Signature is only valid if error is nil

//		It's crypto, so error checking is critical. EdVerify results in nil error if the
//			signature was valid for the provided public key and exact data.

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

func NewEdKeyPair(seed []byte) ([]byte, []byte, error) {
	example := []byte("Noctu Orfei Aude Fraetor")
	if len(seed) != 32 {
		return []byte{}, []byte{}, errors.New("NewEdKeyPair: seed size did not match requirement")
	}
	pri := ed25519.NewKeyFromSeed(seed)
	pub := make([]byte, 32)
	copy(pub, pri[32:])
	_, err := EdSign(example, pub, pri)
	if err != nil {
		return []byte{}, []byte{}, errors.Wrap(err, "NewEdKeyPair")
	}
	return pub, pri, nil
}

func EdSign(data, publicKey, privateKey []byte) ([]byte, error) {
	if len(privateKey) != 64 {
		return []byte{}, errors.New("EdSign: private key size did not match requirement")
	}
	if len(publicKey) != 32 {
		return []byte{}, errors.New("EdSign: public key size did not match requirement")
	}
	sig := ed25519.Sign(privateKey, data)
	err := EdVerify(data, sig, publicKey)
	if err != nil {
		//It is a security risk to leak an invalid signature if nonrandom (return empty []byte instead)
		return []byte{}, errors.Wrap(err, "EdSign (self-check)")
	}
	return sig, nil
}

func EdVerify(data, signature, publicKey []byte) error {
	if len(publicKey) != 32 {
		return errors.New("EdVerify: public key size did not match requirement")
	}
	if ed25519.Verify(publicKey, data, signature) {
		return nil //Signature valid
	}
	return errors.New("EdVerify: Signature invalid")
}
