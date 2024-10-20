package main

import (
	"gpg-gin-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize the GPGController with the paths to your keys
	gpgController := controllers.NewGPGController("keys/public.key", "keys/private.key")

	// Define routes
	r.POST("/encrypt", gpgController.EncryptData)
	r.POST("/decrypt", gpgController.DecryptData)

	// Start the Gin server
	r.Run(":8080")
}
