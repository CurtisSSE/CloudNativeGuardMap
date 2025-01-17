package main

import (
	"CloudNativeGuardMap/Azure/Auth"
	"log"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type AuthState struct {
	AuthData    azidentity.AuthenticationRecord
	AccessToken azcore.AccessToken
}

var (
	authState AuthState
)

func main() {

	// Router definition, CORS configuration.
	siteRouter := gin.Default()
	siteCorsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}
	siteRouter.Use(cors.New(siteCorsConfig))

	// Authentication POST requests.
	siteRouter.POST("/auth-login", func(c *gin.Context) {
		userData, accessToken := Auth.AuthN()
		authState.AuthData = userData
		authState.AccessToken = accessToken
		c.JSON(http.StatusOK, gin.H{"user": authState.AuthData})
	})

	siteRouter.POST("/auth-logout", func(c *gin.Context) {
		userData, accessToken := Auth.AuthNLogOut()
		authState.AuthData = userData
		authState.AccessToken = accessToken
		c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
	})

	siteRouter.GET("/auth-userdata", func(c *gin.Context) {
		if authState.AuthData.Username != "" {
			c.JSON(http.StatusOK, gin.H{"userupn": authState.AuthData.Username})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		}
	})

	// Error handling, start the server.
	if err := siteRouter.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
