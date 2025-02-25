# GPG Encryption/Decryption API with Gin Framework

This repository contains a **Go (Golang) web application** built with the **Gin framework**, which provides endpoints for **GPG encryption and decryption**. It leverages the `gpg` command-line tool to encrypt and decrypt data using public and private keys.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Library Usage](#library-usage)
https://weaspire.atlassian.net/wiki/x/5wCMOg
---

## Features

- **Encrypt Data:** Encrypts incoming data using a **public key**.
- **Decrypt Data:** Decrypts encrypted data using a **private key**.
- Simple and easy-to-use **REST API** built with the **Gin framework**.
- **GPG key management** using public and private keys.

---

## Prerequisites

Before you can run this project, ensure you have the following installed:

1. **Golang (Go) 1.19+**: [Download here](https://golang.org/dl/).
2. **GPG (GNU Privacy Guard)**: Install via your package manager:
   - Ubuntu/Debian: `sudo apt-get install gnupg`
   - macOS: `brew install gpg`
   - Windows: [Download GPG](https://gnupg.org/download/index.html)
3. **Public and Private GPG Keys**:
   - Generate keys:
     ```bash
     gpg --gen-key
     ```
   - Export public key:
     ```bash
     gpg --export -a "Your Name" > keys/public.key
     ```
   - Export private key:
     ```bash
     gpg --export-secret-key -a "Your Name" > keys/private.key
     ```

---

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/gauravaspire/gpg-gin-api.git
   cd gpg-gin-api
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Run the application:**

   ```bash
   go run main.go
   ```

---

## Usage

Once the application is running, you can use the following endpoints:

1. **Encrypt Data:**

   - **Endpoint:** `/encrypt`
   - **Method:** `POST`
   - **Description:** Encrypts the incoming data using the public key.
   - **Request Body:** Raw data to be encrypted.
   - **Response:** Encrypted data.

   ```bash
   curl -X POST http://localhost:8080/encrypt -d "Your data to encrypt"
   ```

2. **Decrypt Data:**

   - **Endpoint:** `/decrypt`
   - **Method:** `POST`
   - **Description:** Decrypts the incoming encrypted data using the private key.
   - **Request Body:** Raw encrypted data.
   - **Response:** Decrypted data.

   ```bash
   curl -X POST http://localhost:8080/decrypt -d "Your encrypted data"
   ```

---

## Library Usage

The core GPG encryption and decryption logic has been refactored into a common library. You can use this library in your own Go projects.

1. **Import the library:**

   ```go
   import "gpg-gin-api/library"
   ```

2. **Encrypt Data:**

   ```go
   publicKeyPath := "path/to/public.key"
   data := []byte("Your data to encrypt")

   encryptedData, err := library.Encrypt(publicKeyPath, data)
   if err != nil {
       // Handle error
   }

   // Use encryptedData
   ```

3. **Decrypt Data:**

   ```go
   privateKeyPath := "path/to/private.key"
   encryptedData := []byte("Your encrypted data")

   decryptedData, err := library.Decrypt(privateKeyPath, encryptedData)
   if err != nil {
       // Handle error
   }

   // Use decryptedData
   ```

---
