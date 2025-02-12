package main

import (
	"CloudNativeGuardMap/Azure/Auth"
	"CloudNativeGuardMap/Azure/Subscription"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Router definition, CORS configuration.
	siteRouter := gin.Default()
	siteCorsConfig := cors.DefaultConfig()
	siteCorsConfig.AllowAllOrigins = true
	siteRouter.Use(cors.New(siteCorsConfig))

	// Authentication POST requests.
	siteRouter.POST("/auth-login", func(c *gin.Context) {
		Auth.AuthN()
		c.JSON(http.StatusOK, gin.H{"user": Auth.CurrentAuthState.AuthData})
	})

	siteRouter.POST("/auth-logout", func(c *gin.Context) {
		Auth.AuthNLogOut()
		c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
	})

	siteRouter.POST("/get-subscriptions", func(c *gin.Context) {
		if !Auth.CurrentAuthState.LoggedIn {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		} else {
			Auth.BuildArmSubscriptionsClientFactory(Auth.SharedCreds)
			if Auth.CurrentSubscriptionsFactoryState.SubscriptionsFactoryClient == nil {
				log.Fatal("ARMSubscriptionsFactoryClient not properly initialized (nil)")
			} else {
				var output string
				currentsubscriptions := Subscription.GetAllSubscriptions(Auth.CurrentSubscriptionsFactoryState.SubscriptionsFactoryClient)
				for key, value := range currentsubscriptions {
					output += "Subscription ID: " + key + "| Subscription Name: " + value
				}
				fmt.Println(output)
				c.JSON(http.StatusOK, gin.H{"output": output})
			}
		}
	})

	// Error handling, start the server.
	if err := siteRouter.Run(":5000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
