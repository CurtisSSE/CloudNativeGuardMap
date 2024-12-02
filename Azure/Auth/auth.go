package Auth

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

var authData azidentity.AuthenticationRecord
var userToken azcore.TokenCredential

// AuthN returns an azidentity.AuthenticationRecord and azcore.AccessToken and returns as variables to main.
func AuthN() (record azidentity.AuthenticationRecord, usertoken azcore.AccessToken) {
	cred, err := azidentity.NewInteractiveBrowserCredential(nil)
	if err != nil {
		fmt.Println("Could not login. Please try again.")
	}
	ctx := context.TODO()
	record, err = cred.Authenticate(ctx, nil)
	usertoken, err = cred.GetToken(ctx, policy.TokenRequestOptions{})
	return record, usertoken
}

// AuthNLogOut clears all azidentity.AuthenticationRecord and azcore.AccessToken variable data (authData and userToken).
func AuthNLogOut() {
	authData.Authority = ""
	authData.ClientID = ""
	authData.HomeAccountID = ""
	authData.TenantID = ""
	authData.Username = ""
	authData.Version = ""
}

func GetUserData() {
	userdata, usertoken := AuthN()
	fmt.Println(userdata.Authority, userdata.ClientID, userdata.HomeAccountID, userdata.TenantID, userdata.Username, userdata.Version)
	fmt.Println(usertoken)
}

/*func azureGraphConnectAuth() {
	record := AuthN()
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{"Files.Read"})
	if err != nil {
		fmt.Println("Could not connect to Graph, try again.")
	}
	result, err := client.Me().Drive().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Failure to get drive: %v\n", err)
	}
	fmt.Println("Found Drive: $v\n", *result.GetId())
	return
}*/
