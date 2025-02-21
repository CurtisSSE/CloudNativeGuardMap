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

	siteRouter.POST("/subscriptions-auth", func(c *gin.Context) {
		if !Auth.CurrentAuthState.LoggedIn {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		} else {
			Auth.BuildArmSubscriptionsClientFactory(Auth.SharedCreds)
			if Auth.CurrentSubscriptionsFactoryState.SubscriptionsFactoryClient == nil {
				log.Fatal("ARMSubscriptionsFactoryClient not properly initialized (nil)")
			} else {
				var subscriptions []string
				var format string
				currentsubscriptions := Subscription.GetAllSubscriptions(Auth.CurrentSubscriptionsFactoryState.SubscriptionsFactoryClient)
				for key, value := range currentsubscriptions {
					format = key + " | " + value
					subscriptions = append(subscriptions, format)
				}
				fmt.Println(subscriptions)
				c.JSON(http.StatusOK, gin.H{"output": subscriptions})
			}
		}
	})

	siteRouter.POST("/advisor-auth", func(c *gin.Context) {
		if !Auth.CurrentAuthState.LoggedIn {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
			return
		}

		var requestData struct {
			subidData string `json:"subscriptionid"`
		}

		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request invalid"})
			return
		}

		subiddataTransfer := requestData.subidData
		c.JSON(http.StatusOK, gin.H{"subid": subiddataTransfer})
		Auth.BuildArmAdvisorClientFactory(subiddataTransfer, Auth.SharedCreds)
	})

	// Error handling, start the server.
	if err := siteRouter.Run(":5000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
