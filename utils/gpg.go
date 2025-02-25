package utils

import (
	"gpg-gin-api/library"
)

// Encrypt encrypts the given input data using the public key located at `publicKeyPath`.
// It returns the encrypted data or an error if encryption fails.
func Encrypt(publicKeyPath string, inputData []byte) ([]byte, error) {
	return library.Encrypt(publicKeyPath, inputData)
}

// Decrypt decrypts the given encrypted data using the private key located at `privateKeyPath`.
// It returns the decrypted data or an error if decryption fails.
func Decrypt(privateKeyPath string, encryptedData []byte) ([]byte, error) {
	return library.Decrypt(privateKeyPath, encryptedData)
}
