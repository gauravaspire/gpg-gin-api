package library

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Encrypt encrypts the given input data using the public key located at `publicKeyPath`.
// It returns the encrypted data or an error if encryption fails.
func Encrypt(publicKeyPath string, inputData []byte) ([]byte, error) {
	// Step 1: Import the public key into the GPG keyring.
	err := importKey(publicKeyPath)
	if err != nil {
		fmt.Println("failed to import public key: %w", err)
		return nil, fmt.Errorf("failed to import public key: %w", err)
	}

	// Step 2: Set up the GPG command for encryption.
	cmd := exec.Command(
		"gpg", "--encrypt", "--armor", "--recipient", "cis", "--trust-model", "always",
	)

	// Step 3: Set up input and output buffers.
	cmd.Stdin = bytes.NewReader(inputData) // Pass input data to stdin.
	var encryptedData bytes.Buffer
	cmd.Stdout = &encryptedData // Capture encrypted data in stdout.

	// Step 4: Run the command and handle errors.
	if err := cmd.Run(); err != nil {
		fmt.Println("encryption failed: %w", err)
		return nil, fmt.Errorf("encryption failed: %w", err)
	}

	// Step 5: Return the encrypted data.
	return encryptedData.Bytes(), nil
}

// Decrypt decrypts the given encrypted data using the private key located at `privateKeyPath`.
// It returns the decrypted data or an error if decryption fails.
func Decrypt(privateKeyPath string, encryptedData []byte) ([]byte, error) {
	// Step 1: Import the private key into the GPG keyring.
	err := importKey(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to import private key: %w", err)
	}

	// Step 2: Set up the GPG command for decryption.
	cmd := exec.Command("gpg", "--decrypt")

	// Step 3: Set up input and output buffers.
	cmd.Stdin = bytes.NewReader(encryptedData) // Pass encrypted data to stdin.
	var decryptedData bytes.Buffer
	cmd.Stdout = &decryptedData // Capture decrypted data in stdout.

	// Step 4: Run the command and handle errors.
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}

	// Step 5: Return the decrypted data.
	return decryptedData.Bytes(), nil
}

// importKey imports a GPG key (public or private) from the given file path.
// It returns an error if the import fails.
func importKey(keyPath string) error {
	// Step 1: Set up the GPG command to import the key.
	cmd := exec.Command("gpg", "--import", keyPath)

	// Step 2: Run the command and handle errors.
	if err := cmd.Run(); err != nil {
		fmt.Println("key import failed: %w", err)

		return fmt.Errorf("key import failed: %w", err)
	}

	return nil
}
