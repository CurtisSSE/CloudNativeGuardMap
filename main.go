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
		if Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient == nil {
			log.Fatal("ResourceFactoryClient not properly initialized (nil)")
		} else {
			currentresources := Compute.GetAllVirtualMachines(Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient, &Auth.CurrentResourceGraphFactoryState.ResourceFactoryCurrentSubscriptionID)
			var currentreturnedVMNames []string
			var currentreturnedResGroups []string
			var currentreturnedOperatingSystems []string
			var currentreturnedAdminUsernames []string
			var currentreturnedNetworkInterfaces []string
			var currentreturnedOsDisks []string
			var currentreturnedDataDisks []string

			for i := range currentresources {
				currentreturnedVMNames = append(currentreturnedVMNames, currentresources[i][0])
				currentreturnedResGroups = append(currentreturnedResGroups, currentresources[i][1])
				currentreturnedOperatingSystems = append(currentreturnedOperatingSystems, currentresources[i][2])
				currentreturnedAdminUsernames = append(currentreturnedAdminUsernames, currentresources[i][3])
				currentreturnedNetworkInterfaces = append(currentreturnedNetworkInterfaces, currentresources[i][4])
				currentreturnedOsDisks = append(currentreturnedOsDisks, currentresources[i][5])
				currentreturnedDataDisks = append(currentreturnedDataDisks, currentresources[i][6])
			}

			var vmstringOutput []string
			for n := range currentreturnedVMNames {
				vmstringOutput = append(vmstringOutput, currentreturnedVMNames[n]+" "+currentreturnedResGroups[n]+" "+currentreturnedOperatingSystems[n]+" "+currentreturnedAdminUsernames[n]+" "+currentreturnedNetworkInterfaces[n]+" "+currentreturnedOsDisks[n]+" "+currentreturnedDataDisks[n])
			}

			fmt.Println(vmstringOutput)
			c.JSON(http.StatusOK, gin.H{"output": vmstringOutput})
		}
	})

	siteRouter.POST("/vn-request", func(c *gin.Context) {
		if Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient == nil {
			log.Fatal("ResourceFactoryClient not properly initialized (nil)")
		} else {
			currentresourcesNetworks := Compute.GetAllVirtualNetworks(Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient, &Auth.CurrentResourceGraphFactoryState.ResourceFactoryCurrentSubscriptionID)
			var currentreturnedVNNames []string
			var currentreturnedVNResGroups []string
			var currentreturnedIPAddresses []string
			var currentreturnedAddressPrefixes []string

			for j := range currentresourcesNetworks {
				currentreturnedVNNames = append(currentreturnedVNNames, currentresourcesNetworks[j][0])
				currentreturnedVNResGroups = append(currentreturnedVNResGroups, currentresourcesNetworks[j][1])
				currentreturnedIPAddresses = append(currentreturnedIPAddresses, currentresourcesNetworks[j][2])
				currentreturnedAddressPrefixes = append(currentreturnedAddressPrefixes, currentresourcesNetworks[j][3])
			}

			var vnstringOutput []string
			for o := range currentreturnedVNNames {
				vnstringOutput = append(vnstringOutput, currentreturnedVNNames[o]+" "+currentreturnedVNResGroups[o]+" "+currentreturnedIPAddresses[o]+" "+currentreturnedAddressPrefixes[o])
			}

			fmt.Println(vnstringOutput)
			c.JSON(http.StatusOK, gin.H{"output": vnstringOutput})
		}
	})

	siteRouter.POST("/vni-request", func(c *gin.Context) {
		if Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient == nil {
			log.Fatal("ResourceFactoryClient not properly initialized (nil)")
		} else {
			currentresourcesNetworkInterfaces := Compute.GetAllVirtualNetworkInterfaces(Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient, &Auth.CurrentResourceGraphFactoryState.ResourceFactoryCurrentSubscriptionID)
			var currentreturnedVNINames []string
			var currentreturnedPrivateIPs []string
			var currentreturnedPublicIPIDs []string

			for k := range currentresourcesNetworkInterfaces {
				currentreturnedVNINames = append(currentreturnedVNINames, currentresourcesNetworkInterfaces[k][0])
				currentreturnedPrivateIPs = append(currentreturnedPrivateIPs, currentresourcesNetworkInterfaces[k][1])
				currentreturnedPublicIPIDs = append(currentreturnedPublicIPIDs, currentresourcesNetworkInterfaces[k][2])
			}

			var nistringOutput []string
			for p := range currentreturnedVNINames {
				nistringOutput = append(nistringOutput, currentreturnedVNINames[p]+" "+currentreturnedPrivateIPs[p]+" "+currentreturnedPublicIPIDs[p])
			}

			fmt.Println(nistringOutput)
			c.JSON(http.StatusOK, gin.H{"output": nistringOutput})
		}
	})

	siteRouter.POST("/pip-request", func(c *gin.Context) {
		if Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient == nil {
			log.Fatal("ResourceFactoryClient not properly initialized (nil)")
		} else {
			currentresourcesPublicIPs := Compute.GetAllPublicIPs(Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient, &Auth.CurrentResourceGraphFactoryState.ResourceFactoryCurrentSubscriptionID)
			var currentreturnedPIPIDs []string
			var currentreturnedPIPResGroups []string
			var currentreturnedActualPublicIPs []string

			for l := range currentresourcesPublicIPs {
				currentreturnedPIPIDs = append(currentreturnedPIPIDs, currentresourcesPublicIPs[l][0])
				currentreturnedPIPResGroups = append(currentreturnedPIPResGroups, currentresourcesPublicIPs[l][1])
				currentreturnedActualPublicIPs = append(currentreturnedActualPublicIPs, currentresourcesPublicIPs[l][2])
			}

			var pipstringOutput []string
			for q := range currentreturnedPIPIDs {
				pipstringOutput = append(pipstringOutput, currentreturnedPIPIDs[q]+" "+currentreturnedPIPResGroups[q]+" "+currentreturnedActualPublicIPs[q])
			}

			fmt.Println(pipstringOutput)
			c.JSON(http.StatusOK, gin.H{"output": pipstringOutput})
		}
	})

	siteRouter.POST("/nsg-request", func(c *gin.Context) {
		if Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient == nil {
			log.Fatal("ResourceFactoryClient not properly initialized (nil)")
		} else {
			currentresourcesNSGs := Compute.GetAllNSGs(Auth.CurrentResourceGraphFactoryState.ResourceFactoryClient, &Auth.CurrentResourceGraphFactoryState.ResourceFactoryCurrentSubscriptionID)
			var currentreturnedNSGNames []string
			var currentreturnedNSGResGroups []string
			var currentreturnedAttachedNIs []string

			for m := range currentresourcesNSGs {
				currentreturnedNSGNames = append(currentreturnedNSGNames, currentresourcesNSGs[m][0])
				currentreturnedNSGResGroups = append(currentreturnedNSGResGroups, currentresourcesNSGs[m][1])
				currentreturnedAttachedNIs = append(currentreturnedAttachedNIs, currentresourcesNSGs[m][2])
			}

			var nsgstringOutput []string
			for r := range currentreturnedNSGNames {
				nsgstringOutput = append(nsgstringOutput, currentreturnedNSGNames[r]+" "+currentreturnedNSGResGroups[r]+" "+currentreturnedAttachedNIs[r])
			}

			fmt.Println(nsgstringOutput)
			c.JSON(http.StatusOK, gin.H{"output": nsgstringOutput})
		}
	})

	// Error handling, start the server.
	if err := siteRouter.Run(":5000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
