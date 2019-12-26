package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// Encryption/decryption errors.
var (
	ErrPlaintextEmpty                   = errors.New("plaintext is empty")
	ErrCiphertextTooShort               = errors.New("ciphertext is too short")
	ErrCiphertextNotMultipleOfBlockSize = errors.New("ciphertext is not a multiple of the block size")
	ErrPlaintextNotMultipleOfBlockSize  = errors.New("plaintext is not a multiple of the block size")

	// ErrInvalidBlockSize indicates hash blocksize <= 0.
	ErrInvalidBlockSize = errors.New("invalid blocksize")

	// ErrInvalidPKCS7Data indicates bad input to PKCS7 pad or unpad.
	ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data (empty or not padded)")

	// ErrInvalidPKCS7Padding indicates PKCS7 unpad fails to bad input.
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

// Encrypt encrypts the given byte slice using AES-CBC algorithm.
func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	if len(plaintext) == 0 {
		return nil, ErrPlaintextEmpty
	}

	plaintext, err := pkcs7Pad(plaintext, aes.BlockSize)
	if err != nil {
		return nil, err
	}

	if len(plaintext)%aes.BlockSize != 0 {
		return nil, ErrPlaintextNotMultipleOfBlockSize
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// Decrypt decrypts the AES-CBC algorithm encrypted cipher text.
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, ErrCiphertextTooShort
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Remove the embedded IV.
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, ErrCiphertextNotMultipleOfBlockSize
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext, err = pkcs7Unpad(plaintext, aes.BlockSize)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// pkcs7Pad right-pads the given byte slice with 1 to n bytes, where
// n is the block size. The size of the result is x times n, where x
// is at least 1.
func pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}

	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))

	return pb, nil
}

// pkcs7Unpad validates and unpads data from the given bytes slice.
// The returned value will be 1 to n bytes smaller depending on the
// amount of padding, where n is the block size.
func pkcs7Unpad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	if len(b)%blocksize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}

	c := b[len(b)-1]
	n := int(c)
	if n == 0 || n > len(b) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if b[len(b)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}

	return b[:len(b)-n], nil
}
