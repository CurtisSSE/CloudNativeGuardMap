package Compute

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

type VirtualMachine struct {
	VMName           string
	ResGroup         string
	OperatingSystem  string
	AdminUsername    string
	NetworkInterface string
}

type VMtoModel struct {
	ReturnedVMs [][]string
}

var VMquery armresourcegraph.QueryRequest
var VMqueryoptions *armresourcegraph.ClientResourcesOptions

func GetAllResourceGroups(resourceclientvariable *armresourcegraph.Client, subscriptionid *string) (currentresourceFields map[int][]string) {
	var currentResource VirtualMachine
	VMquerytext := "Resources | where type =~ 'Microsoft.Compute/virtualMachines' | project name, resourceGroup, operatingSystem = (properties['extended']['instanceView']['osName']), adminUsername = (properties['osProfile']['adminUsername']), networkProfile = (properties['networkProfile']['networkInterfaces'])"
	VMquery.Query = &VMquerytext
	VMquery.Subscriptions = append(VMquery.Subscriptions, subscriptionid)
	currentctx := context.Background()
	queryresponse, err := resourceclientvariable.Resources(currentctx, VMquery, VMqueryoptions)
	if err != nil {
		fmt.Println("Querying resources failed. Check authentication or other issues", err)
	}
	resourceoutputd, _ := json.Marshal(queryresponse.Data)
	var resourceumData []map[string]interface{}
	resourceerr := json.Unmarshal([]byte(resourceoutputd), &resourceumData)
	if resourceerr != nil {
		fmt.Println("Error unmarshaling resource data.", resourceerr)
		return
	} else {
		currentresourceFields = make(map[int][]string)
		for i, resourceumIter := range resourceumData {
			currentResource.VMName = resourceumIter["name"].(string)
			currentResource.ResGroup = resourceumIter["resourceGroup"].(string)
			currentResource.OperatingSystem = resourceumIter["operatingSystem"].(string)
			currentResource.AdminUsername = resourceumIter["adminUsername"].(string)
			networkjsonProcessing, _ := json.Marshal(resourceumIter["networkProfile"])
			var networkumData []map[string]interface{}
			json.Unmarshal(networkjsonProcessing, &networkumData)
			currentResource.NetworkInterface = networkumData[0]["id"].(string)
			currentresourceFields[i] = []string{currentResource.VMName, currentResource.ResGroup, currentResource.OperatingSystem, currentResource.AdminUsername, currentResource.NetworkInterface}
		}
	}
	return currentresourceFields
}
