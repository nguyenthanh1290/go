package ecdh

import (
	"crypto"
	"encoding/hex"
	"errors"
	"io"
)

var (
	ErrInvalidKey = errors.New("pointer to [32]byte key is required")
	ErrInvalidData = errors.New("[32]byte data is required")
)

type ECDH interface {
	GenerateKeyPair(rand io.Reader) (crypto.PrivateKey, crypto.PublicKey, error)
	ComputeSharedSecret(private crypto.PrivateKey, peerPublic crypto.PublicKey) []byte
}

var Curve25519 ECDH

func init() {
	Curve25519 = ecdhCurve25519{}
}

func PublicKeyToHexString(public crypto.PublicKey) (string, error) {
	switch buf := public.(type) {
	case *[32]byte:
		return hex.EncodeToString(buf[:]), nil
	default:
		return "", ErrInvalidKey
	}
}

func PublicKeyFromHexString(data string) (crypto.PublicKey, error) {
	var public [32]byte

	buf, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}

	if len(buf) != 32 {
		return nil, ErrInvalidData
	}

	copy(public[:], buf)

	return &public, nil
}
