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
	OsDiskName       string
	OsDiskSize       string
	OsDiskType       string
	DataDisksName    string
}

type NetworkComponents struct {
	PublicIP string
}

var VMquery armresourcegraph.QueryRequest
var VMqueryoptions *armresourcegraph.ClientResourcesOptions

func GetAllResourceGroups(resourceclientvariable *armresourcegraph.Client, subscriptionid *string) (currentresourceFields map[int][]string) {
	var currentResource VirtualMachine
	VMquerytext := "Resources | where type =~ 'Microsoft.Compute/virtualMachines' | project name, resourceGroup, operatingSystem = (properties['extended']['instanceView']['osName']), operatingVersion = (properties['extended']['instanceView']['osVersion']), adminUsername = (properties['osProfile']['adminUsername']), networkProfile = (properties['networkProfile']['networkInterfaces']), osDisk = (properties['storageProfile']['osDisk']), osDiskType = (properties['storageProfile']['osDisk']['managedDisk'], dataDisks = (properties['storageProfile']['dataDisks']), dataDisksType = (properties['storageProfile']['dataDisks']['managedDisk'])"
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
			currentResource.OperatingSystem = resourceumIter["operatingSystem"].(string) + "-" + resourceumIter["operatingVersion"].(string)
			currentResource.AdminUsername = resourceumIter["adminUsername"].(string)
			networkjsonProcessing, _ := json.Marshal(resourceumIter["networkProfile"])
			var networkumData []map[string]interface{}
			json.Unmarshal(networkjsonProcessing, &networkumData)
			currentResource.NetworkInterface = networkumData[0]["id"].(string)
			osdiskjsonProcessing, _ := json.Marshal(resourceumIter["osDisk"])
			var osdiskumData []map[string]interface{}
			json.Unmarshal(osdiskjsonProcessing, &osdiskumData)
			currentResource.OsDiskName = osdiskumData[0]["name"].(string)
			currentResource.OsDiskSize = osdiskumData[0]["diskSizeGB"].(string)
			osdisktypejsonProcessing, _ := json.Marshal(resourceumIter["osDiskType"])
			var osdisktypeumData []map[string]interface{}
			json.Unmarshal(osdisktypejsonProcessing, &osdisktypeumData)
			currentResource.OsDiskType = osdisktypeumData[0]["storageAccountType"].(string)
			datadisksjsonProcessing, _ := json.Marshal(resourceumIter["dataDisks"])
			var datadisksumData []map[string]interface{}
			json.Unmarshal(osdisk)
			currentresourceFields[i] = []string{currentResource.VMName, currentResource.ResGroup, currentResource.OperatingSystem, currentResource.AdminUsername, currentResource.NetworkInterface, currentResource.OsDiskName, currentResource.OsDiskSize, currentResource.OsDiskType, currentResource.DataDisks}
		}
	}
	return currentresourceFields
}
