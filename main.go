package main

import (
	"CloudNativeGuardMap/Azure/Auth"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Global variable to store authentication data (in-memory for simplicity)
var authData azidentity.AuthenticationRecord
var userToken azcore.TokenCredential

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	// Endpoint to trigger authentication on 'Authenticate' click.
	r.POST("/auth-login", func(c *gin.Context) {
		authData, _ = Auth.AuthN() // Authenticate and store the result globally.
		c.JSON(http.StatusOK, gin.H{"result": authData})
	})

	r.POST("/auth-logout", func(c *gin.Context) {
		Auth.AuthNLogOut()
		c.JSON(http.StatusOK, gin.H{"result": nil})
	})

	// Endpoint to provide authData and userToken after authentication.
	r.GET("/auth-data", func(c *gin.Context) {
		sendData := gin.H{
			"userupn":   authData.Username,
			"usertoken": userToken,
		}
		c.JSON(http.StatusOK, sendData)
	})

	// Start the server
	r.Run(":8080")
}
