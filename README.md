# GPG Encryption/Decryption API with Gin Framework

This repository contains a **Go (Golang) web application** built with the **Gin framework**, which provides endpoints for **GPG encryption and decryption**. It leverages the `gpg` command-line tool to encrypt and decrypt data using public and private keys.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Project Structure](#project-structure)
- [Usage](#usage)
  - [Start the Server](#start-the-server)
  - [Encrypt Data](#encrypt-data)
  - [Decrypt Data](#decrypt-data)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

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
