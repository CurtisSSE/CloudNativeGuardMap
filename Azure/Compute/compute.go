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
	DataDisksName    string
}

type VirtualNetwork struct {
	VNName        string
	ResGroup      string
	IPAddresses   string
	AddressPrefix string
}

type VirtualNetworkInterface struct {
	VNIName    string
	ResGroup   string
	PrivateIP  string
	PublicIPID string
}

type PublicIPAttributes struct {
	IPID           string
	ResGroup       string
	ActualPublicIP string
}

type NetworkSecurityGroups struct {
	NSGName                  string
	ResGroup                 string
	AttachedNetworkInterface string
}

var VMquery armresourcegraph.QueryRequest
var VMqueryoptions *armresourcegraph.ClientResourcesOptions

//var dataDiskCycle string

func GetAllVirtualMachines(resourceclientvariable *armresourcegraph.Client, subscriptionid *string) (currentresourceFields map[int][]string) {
	var currentResource VirtualMachine
	VMquerytext := "Resources | where type =~ 'Microsoft.Compute/virtualMachines' | project name, resourceGroup, operatingSystem = (properties['extended']['instanceView']['osName']), operatingVersion = (properties['extended']['instanceView']['osVersion']), adminUsername = (properties['osProfile']['adminUsername']), networkProfile = (properties['networkProfile']['networkInterfaces']), osDisk = (properties['storageProfile']['osDisk']['name']), dataDisks = (properties['storageProfile']['dataDisks'])"
	VMquery.Query = &VMquerytext
	VMquery.Subscriptions = append(VMquery.Subscriptions, subscriptionid)
	currentctx := context.Background()
	queryresponse, err := resourceclientvariable.Resources(currentctx, VMquery, VMqueryoptions)
	if err != nil {
		fmt.Println("Querying virtual machines failed. Check authentication or other issues", err)
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
			//dataDiskCycle = string(rune(i))
			currentResource.VMName = resourceumIter["name"].(string)
			currentResource.ResGroup = resourceumIter["resourceGroup"].(string)
			currentResource.OperatingSystem = resourceumIter["operatingSystem"].(string) + "-" + resourceumIter["operatingVersion"].(string)
			currentResource.AdminUsername = resourceumIter["adminUsername"].(string)
			networkjsonProcessing, _ := json.Marshal(resourceumIter["networkProfile"])
			var networkumData []map[string]interface{}
			json.Unmarshal(networkjsonProcessing, &networkumData)
			currentResource.NetworkInterface = networkumData[0]["id"].(string)
			currentResource.OsDiskName = resourceumIter["osDisk"].(string)
			datadisksjsonProcessing, _ := json.Marshal(resourceumIter["dataDisks"])
			var datadisksumData []map[string]interface{}
			json.Unmarshal(datadisksjsonProcessing, &datadisksumData)
			if len(datadisksumData) != 0 {
				currentResource.DataDisksName = datadisksumData[0]["name"].(string)
			} else {
				currentResource.DataDisksName = "No-Data-Disks"
			}

			currentresourceFields[i] = []string{currentResource.VMName, currentResource.ResGroup, currentResource.OperatingSystem, currentResource.AdminUsername, currentResource.NetworkInterface, currentResource.OsDiskName, currentResource.DataDisksName}
		}
	}
	return currentresourceFields
}

var NIquery armresourcegraph.QueryRequest
var NIqueryoptions *armresourcegraph.ClientResourcesOptions

func GetAllVirtualNetworks(resourceclientvariable *armresourcegraph.Client, subscriptionid *string) (currentresourceVNFields map[int][]string) {
	var currentResourceVirtualNetwork VirtualNetwork
	NIquerytext := "Resources | where type =~ 'Microsoft.Network/virtualNetworks' | project name, resourceGroup, ipAddress1 = (properties['subnets'][0]['properties']['ipConfigurations'][0]['id']), ipAddress2 = (properties['subnets'][0]['properties']['ipConfigurations'][1]['id']), addressPrefix = (properties['subnets'][0]['properties']['addressPrefix'])"
	NIquery.Query = &NIquerytext
	NIquery.Subscriptions = append(NIquery.Subscriptions, subscriptionid)
	currentctx := context.Background()
	queryresponse, err := resourceclientvariable.Resources(currentctx, NIquery, NIqueryoptions)
	if err != nil {
		fmt.Println("Querying virtual networks failed. Check authentication or other issues.", err)
	}
	resourceoutputd, _ := json.Marshal(queryresponse.Data)
	var resourceumData []map[string]interface{}
	resourceerr := json.Unmarshal([]byte(resourceoutputd), &resourceumData)
	if resourceerr != nil {
		fmt.Println("Error unmarshaling resource data.", resourceerr)
		return
	} else {
		currentresourceVNFields := make(map[int][]string)
		for i, resourceumIter := range resourceumData {
			currentResourceVirtualNetwork.VNName = resourceumIter["name"].(string)
			currentResourceVirtualNetwork.ResGroup = resourceumIter["resourceGroup"].(string)
			currentResourceVirtualNetwork.IPAddresses = resourceumIter["ipAddress1"].(string) + "~" + resourceumIter["ipAddress2"].(string)
			currentResourceVirtualNetwork.AddressPrefix = resourceumIter["addressPrefix"].(string)
			currentresourceVNFields[i] = []string{currentResourceVirtualNetwork.VNName, currentResourceVirtualNetwork.ResGroup, currentResourceVirtualNetwork.IPAddresses, currentResourceVirtualNetwork.AddressPrefix}
		}
	}
	return currentresourceVNFields
}

var VINquery armresourcegraph.QueryRequest
var VINqueryoptions *armresourcegraph.ClientResourcesOptions

func GetAllVirtualNetworkInterfaces(resourceclientvariable *armresourcegraph.Client, subscriptionid *string) (currentresourceVNIFields map[int][]string) {
	var currentResourceVirtualNetworkInterface VirtualNetworkInterface
	VINquerytext := "Resources | where type =~ 'Microsoft.Network/networkInterfaces' | project name, privateip = (properties['ipConfigurations'][0]['properties']['privateIPAddress']), resourceGroup, publicip = (properties['ipConfigurations'][0]['properties']['publicIPAddress']['id'])"
	VINquery.Query = &VINquerytext
	VINquery.Subscriptions = append(VINquery.Subscriptions, subscriptionid)
	currentctx := context.Background()
	queryresponse, err := resourceclientvariable.Resources(currentctx, VINquery, VINqueryoptions)
	if err != nil {
		fmt.Println("Querying virtual network interfaces failed. Check authentication or other issues.", err)
	}
	resourceoutputd, _ := json.Marshal(queryresponse.Data)
	var resourceumData []map[string]interface{}
	resourceerr := json.Unmarshal([]byte(resourceoutputd), &resourceumData)
	if resourceerr != nil {
		fmt.Println("Error unmarshaling resource data.", resourceerr)
		return
	} else {
		currentresourceVNIFields := make(map[int][]string)
		for i, resourceumIter := range resourceumData {
			currentResourceVirtualNetworkInterface.VNIName = resourceumIter["name"].(string)
			currentResourceVirtualNetworkInterface.PrivateIP = resourceumIter["privateip"].(string)
			currentResourceVirtualNetworkInterface.PublicIPID = resourceumIter["publicip"].(string)
			currentResourceVirtualNetworkInterface.ResGroup = resourceumIter["resourceGroup"].(string)
			currentresourceVNIFields[i] = []string{currentResourceVirtualNetworkInterface.VNIName, currentResourceVirtualNetworkInterface.PrivateIP, currentResourceVirtualNetworkInterface.ResGroup, currentResourceVirtualNetworkInterface.PublicIPID}
		}
	}
	return currentresourceVNIFields
}

var PIPquery armresourcegraph.QueryRequest
var PIPqueryoptions *armresourcegraph.ClientResourcesOptions

func GetAllPublicIPs(resourceclientvariable *armresourcegraph.Client, subscriptionid *string) (currentresourcePIPFields map[int][]string) {
	var currentResourcePublicIP PublicIPAttributes
	PIPquerytext := "Resources | where type =~ 'Microsoft.Network/publicIPAddresses' | project id, resourceGroup, subscriptionId, actualip = (properties['ipAddress'])"
	PIPquery.Query = &PIPquerytext
	PIPquery.Subscriptions = append(PIPquery.Subscriptions, subscriptionid)
	currentctx := context.Background()
	queryresponse, err := resourceclientvariable.Resources(currentctx, PIPquery, PIPqueryoptions)
	if err != nil {
		fmt.Println("Querying public IPs failed. Check authentication or other issues.", err)
	}
	resourceoutputd, _ := json.Marshal(queryresponse.Data)
	var resourceumData []map[string]interface{}
	resourceerr := json.Unmarshal([]byte(resourceoutputd), &resourceumData)
	if resourceerr != nil {
		fmt.Println("Error unmarshaling resource data.", resourceerr)
		return
	} else {
		currentresourcePIPFields := make(map[int][]string)
		for i, resourceumIter := range resourceumData {
			currentResourcePublicIP.IPID = resourceumIter["id"].(string)
			currentResourcePublicIP.ResGroup = resourceumIter["resourceGroup"].(string)
			currentResourcePublicIP.ActualPublicIP = resourceumIter["actualip"].(string)
			currentresourcePIPFields[i] = []string{currentResourcePublicIP.IPID, currentResourcePublicIP.ResGroup, currentResourcePublicIP.ActualPublicIP}
		}
	}
	return currentresourcePIPFields
}

var NSGquery armresourcegraph.QueryRequest
var NSGqueryoptions *armresourcegraph.ClientResourcesOptions

func GetAllNSGs(resourceclientvariable *armresourcegraph.Client, subscriptionid *string) (currentresourceNSGFields map[int][]string) {
	var currentResourceNSG NetworkSecurityGroups
	NSGquerytext := "where type =~ 'Microsoft.Network/networkSecurityGroups' | project name, resourceGroup, attachedinterface = (properties['networkInterfaces'][0]['id'])"
	NSGquery.Query = &NSGquerytext
	NSGquery.Subscriptions = append(NSGquery.Subscriptions, subscriptionid)
	currentctx := context.Background()
	queryresponse, err := resourceclientvariable.Resources(currentctx, NSGquery, NSGqueryoptions)
	if err != nil {
		fmt.Println("Querying network security groups failed. Check authentication or other issues.", err)
	}
	resourceoutputd, _ := json.Marshal(queryresponse.Data)
	var resourceumData []map[string]interface{}
	resourceerr := json.Unmarshal([]byte(resourceoutputd), &resourceumData)
	if resourceerr != nil {
		fmt.Println("Error unmarshaling resource data.", resourceerr)
		return
	} else {
		currentresourceNSGFields := make(map[int][]string)
		for i, resourceumIter := range resourceumData {
			currentResourceNSG.NSGName = resourceumIter["name"].(string)
			currentResourceNSG.ResGroup = resourceumIter["resourceGroup"].(string)
			currentResourceNSG.AttachedNetworkInterface = resourceumIter["attachedinterface"].(string)
			currentresourceNSGFields[i] = []string{currentResourceNSG.NSGName, currentResourceNSG.ResGroup, currentResourceNSG.AttachedNetworkInterface}
		}
	}
	return currentresourceNSGFields
}
