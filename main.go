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
			// VM sep
			currentresources := Compute.GetAllVirtualMachines(Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient, &Auth.CurrentResourceGraphFactoryState.ResourceFactoryCurrentSubscriptionID)
			var currentreturnedVMNames []string
			var currentreturnedResGroups []string
			var currentreturnedOperatingSystems []string
			var currentreturnedAdminUsernames []string
			var currentreturnedNetworkInterfaces []string
			var currentreturnedOsDisks []string
			var currentreturnedDataDisks []string
			// VN sep
			currentresourcesNetworks := Compute.GetAllVirtualNetworks(Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient, &Auth.CurrentResourceGraphFactoryState.ResourceFactoryCurrentSubscriptionID)
			var currentreturnedVNNames []string
			var currentreturnedVNResGroups []string
			var currentreturnedIPAddresses []string
			var currentreturnedAddressPrefixes []string
			// VNI sep
			currentresourcesNetworkInterfaces := Compute.GetAllVirtualNetworkInterfaces(Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient, &Auth.CurrentResourceGraphFactoryState.ResourceFactoryCurrentSubscriptionID)
			var currentreturnedVNINames []string
			var currentreturnedPrivateIPs []string
			var currentreturnedPublicIPIDs []string
			// PIP sep
			currentresourcesPublicIPs := Compute.GetAllPublicIPs(Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient, &Auth.CurrentResourceGraphFactoryState.ResourceFactoryCurrentSubscriptionID)
			var currentreturnedPIPIDs []string
			var currentreturnedPIPResGroups []string
			var currentreturnedActualPublicIPs []string
			// NSG sep
			currentresourcesNSGs := Compute.GetAllNSGs(Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient, &Auth.CurrentResourceGraphFactoryState.ResourceFactoryCurrentSubscriptionID)
			var currentreturnedNSGNames []string
			var currentreturnedNSGResGroups []string
			var currentreturnedAttachedNIs []string

			for i := range currentresources {
				currentreturnedVMNames = append(currentreturnedVMNames, currentresources[i][0])
				currentreturnedResGroups = append(currentreturnedResGroups, currentresources[i][1])
				currentreturnedOperatingSystems = append(currentreturnedOperatingSystems, currentresources[i][2])
				currentreturnedAdminUsernames = append(currentreturnedAdminUsernames, currentresources[i][3])
				currentreturnedNetworkInterfaces = append(currentreturnedNetworkInterfaces, currentresources[i][4])
				currentreturnedOsDisks = append(currentreturnedOsDisks, currentresources[i][5])
				currentreturnedDataDisks = append(currentreturnedDataDisks, currentresources[i][6])
			}

			for i := range currentresourcesNetworks {
				currentreturnedVNNames = append(currentreturnedVNNames, currentresourcesNetworks[i][0])
				currentreturnedVNResGroups = append(currentreturnedVNResGroups, currentresourcesNetworks[i][1])
				currentreturnedIPAddresses = append(currentreturnedIPAddresses, currentresourcesNetworks[i][2])
				currentreturnedAddressPrefixes = append(currentreturnedAddressPrefixes, currentresourcesNetworks[i][3])
			}

			for i := range currentresourcesNetworkInterfaces {
				currentreturnedVNINames = append(currentreturnedVNINames, currentresourcesNetworkInterfaces[i][0])
				currentreturnedPrivateIPs = append(currentreturnedPrivateIPs, currentresourcesNetworkInterfaces[i][1])
				currentreturnedPublicIPIDs = append(currentreturnedPublicIPIDs, currentresourcesNetworkInterfaces[i][2])
			}

			for i := range currentresourcesPublicIPs {
				currentreturnedPIPIDs = append(currentreturnedPIPIDs, currentresourcesPublicIPs[i][0])
				currentreturnedPIPResGroups = append(currentreturnedPIPResGroups, currentresourcesPublicIPs[i][1])
				currentreturnedActualPublicIPs = append(currentreturnedActualPublicIPs, currentresourcesPublicIPs[i][2])
			}

			for i := range currentresourcesNSGs {
				currentreturnedNSGNames = append(currentreturnedNSGNames, currentresourcesNSGs[i][0])
				currentreturnedNSGResGroups = append(currentreturnedNSGResGroups, currentresourcesNSGs[i][1])
				currentreturnedAttachedNIs = append(currentreturnedAttachedNIs, currentresourcesNSGs[i][2])
			}

			var vmstringOutput []string
			for i := range currentreturnedVMNames {
				vmstringOutput = append(vmstringOutput, currentreturnedVMNames[i]+" "+currentreturnedResGroups[i]+" "+currentreturnedOperatingSystems[i]+" "+currentreturnedAdminUsernames[i]+" "+currentreturnedNetworkInterfaces[i]+" "+currentreturnedOsDisks[i]+" "+currentreturnedDataDisks[i])
			}

			var vnstringOutput []string
			for i := range currentreturnedVNNames {
				vnstringOutput = append(vnstringOutput, currentreturnedVNNames[i]+" "+currentreturnedVNResGroups[i]+" "+currentreturnedIPAddresses[i]+" "+currentreturnedAddressPrefixes[i])
			}

			var nistringOutput []string
			for i := range currentreturnedVNINames {
				nistringOutput = append(nistringOutput, currentreturnedVNINames[i]+" "+currentreturnedPrivateIPs[i]+" "+currentreturnedPublicIPIDs[i])
			}

			var pipstringOutput []string
			for i := range currentreturnedPIPIDs {
				pipstringOutput = append(pipstringOutput, currentreturnedPIPIDs[i]+" "+currentreturnedPIPResGroups[i]+" "+currentreturnedActualPublicIPs[i])
			}

			var nsgstringOutput []string
			for i := range currentreturnedNSGNames {
				nsgstringOutput = append(nsgstringOutput, currentreturnedNSGNames[i]+" "+currentreturnedNSGResGroups[i]+" "+currentreturnedAttachedNIs[i])
			}

			var firstStringOutput []string
			var actualStringOutput []string
			firstStringOutput = append(firstStringOutput, vmstringOutput...)
			vnstringOutput = append(firstStringOutput, vnstringOutput...)
			nistringOutput = append(vnstringOutput, nistringOutput...)
			pipstringOutput = append(nistringOutput, pipstringOutput...)
			nsgstringOutput = append(pipstringOutput, nsgstringOutput...)
			actualStringOutput = append(actualStringOutput, nsgstringOutput...)

			fmt.Println(actualStringOutput)
			c.JSON(http.StatusOK, gin.H{"output": actualStringOutput})
		}
	})

	// Error handling, start the server.
	if err := siteRouter.Run(":5000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
