package controllers

import (
    "io/ioutil"
    "gpg-gin-api/utils"
    "net/http"

    "github.com/gin-gonic/gin"
)

// GPGController handles encryption and decryption requests.
type GPGController struct {
    publicKeyPath  string
    privateKeyPath string
}

// NewGPGController creates a new GPGController with paths to keys.
func NewGPGController(pubKeyPath, privKeyPath string) *GPGController {
    return &GPGController{
        publicKeyPath:  pubKeyPath,
        privateKeyPath: privKeyPath,
    }
}

// EncryptData encrypts the incoming data using GPG.
func (ctrl *GPGController) EncryptData(c *gin.Context) {
    data, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
        return
    }

    encryptedData, err := utils.Encrypt(ctrl.publicKeyPath, data)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption failed"})
        return
    }

    c.Data(http.StatusOK, "text/plain", encryptedData)
}

// DecryptData decrypts the incoming encrypted data using GPG.
func (ctrl *GPGController) DecryptData(c *gin.Context) {
    data, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
        return
    }

    decryptedData, err := utils.Decrypt(ctrl.privateKeyPath, data)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Decryption failed"})
        return
    }

    c.Data(http.StatusOK, "text/plain", decryptedData)
}