package ecdh

import (
	"crypto"
	"io"

	"golang.org/x/crypto/curve25519"
)

type ecdhCurve25519 struct {
}

// GenerateKeyPair generates a pair of private and public key using entropy from rand.
func (e ecdhCurve25519) GenerateKeyPair(rand io.Reader) (crypto.PrivateKey, crypto.PublicKey, error) {
	var private, public [32]byte

	// Compute secret key.
	_, err := rand.Read(private[:])
	if err != nil {
		return nil, nil, err
	}

	private[0] &= 248
	private[31] &= 127
	private[31] |= 64

	// Compute public key.
	curve25519.ScalarBaseMult(&public, &private)

	return &private, &public, nil

}

// ComputeSharedSecret computes shared secret key.
func (e ecdhCurve25519) ComputeSharedSecret(private crypto.PrivateKey, peerPublic crypto.PublicKey) []byte {
	var sharedSecret [32]byte

	curve25519.ScalarMult(&sharedSecret, private.(*[32]byte), peerPublic.(*[32]byte))

	return sharedSecret[:]
}
