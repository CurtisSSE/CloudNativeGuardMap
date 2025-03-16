package main

import (
	"CloudNativeGuardMap/Azure/Advisor"
	"CloudNativeGuardMap/Azure/Auth"
	"CloudNativeGuardMap/Azure/Compute"
	"CloudNativeGuardMap/Azure/Subscription"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
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
		type jsonResponse struct {
			SubidData string `json:"subscriptionid"`
		}
		var requestData jsonResponse
		if !Auth.CurrentAuthState.LoggedIn {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
			return
		}
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request invalid"})
			return
		}
		subiddataTransfer := requestData.SubidData
		Auth.BuildArmAdvisorClientFactory(subiddataTransfer, Auth.SharedCreds)
		c.JSON(http.StatusOK, gin.H{"SubidData": subiddataTransfer})
	})

	siteRouter.POST("/advisor-request", func(c *gin.Context) {
		if Auth.CurrentAdvisorFactoryState.AdvisorFactoryRecommendationsClient == nil {
			log.Fatal("AdvisorFactoryRecommendationsClient not properly initialized (nil)")
		} else {
			var recommendations []string
			var format string
			currentrecommendations := Advisor.GetAllAdvisories(Auth.CurrentAdvisorFactoryState.AdvisorFactoryRecommendationsClient)
			for _, value := range currentrecommendations {
				format = value.RecName + " | " + value.RecID + " | " + value.ShortDescription + " | " + value.ShortSolution + " | " + value.Description + " | " +
					value.ActionsRepo.Description + " | " + value.ActionsRepo.ActionType + " | " + value.ActionsRepo.Caption + " | " +
					value.ActionsRepo.Link + " | " + value.ActionsRepo.Metadata.ID + " | " + string(value.Category) + " | " + string(value.Impact) + " | " + value.ImpactedField + " | " + value.PotentialBenefits
				recommendations = append(recommendations, format)
			}
			fmt.Println(recommendations)
			c.JSON(http.StatusOK, gin.H{"output": recommendations})
		}
	})

	siteRouter.POST("/resources-auth", func(c *gin.Context) {
		type jsonResponse struct {
			SubidData string `json:"subscriptionid"`
		}
		var requestData jsonResponse
		if !Auth.CurrentAuthState.LoggedIn {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
			return
		}
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request invalid"})
			return
		}
		subiddataTransfer := requestData.SubidData
		Auth.CurrentResourceGraphFactoryState.ResourceFactoryCurrentSubscriptionID = subiddataTransfer
		Auth.BuildArmResourceGraphClientFactory(subiddataTransfer, Auth.SharedCreds)
		c.JSON(http.StatusOK, gin.H{"SubidData": subiddataTransfer})
	})

	siteRouter.POST("/resources-request", func(c *gin.Context) {
		//var returnedVMOutputString string
		if Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient == nil {
			log.Fatal("ResourceFactoryClient not properly initialized (nil)")
		} else {
			currentresources := Compute.GetAllResourceGroups(Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient, &Auth.CurrentResourceGraphFactoryState.ResourceFactoryCurrentSubscriptionID)
			var currentreturnedVMNames []string
			var currentreturnedResGroups []string
			var currentreturnedOperatingSystems []string
			var currentreturnedAdminUsernames []string
			var currentreturnedNetworkInterfaces []string
			for i := range currentresources {
				currentreturnedVMNames = append(currentreturnedVMNames, currentresources[i][0])
				currentreturnedResGroups = append(currentreturnedResGroups, currentresources[i][1])
				currentreturnedOperatingSystems = append(currentreturnedOperatingSystems, currentresources[i][2])
				currentreturnedAdminUsernames = append(currentreturnedAdminUsernames, currentresources[i][3])
				currentreturnedNetworkInterfaces = append(currentreturnedNetworkInterfaces, currentresources[i][4])
			}
			var buildStringOutput []string
			for i := range currentreturnedVMNames {
				buildStringOutput = append(buildStringOutput, currentreturnedVMNames[i]+" "+currentreturnedResGroups[i]+" "+currentreturnedOperatingSystems[i]+" "+currentreturnedAdminUsernames[i]+" "+currentreturnedNetworkInterfaces[i])
			}
			fmt.Println(buildStringOutput)
			c.JSON(http.StatusOK, gin.H{"output": buildStringOutput})
		}
	})

	// Error handling, start the server.
	if err := siteRouter.Run(":5000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
