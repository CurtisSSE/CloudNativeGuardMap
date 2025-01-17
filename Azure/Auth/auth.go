package Auth

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

var authData azidentity.AuthenticationRecord
var accessToken azcore.AccessToken

// AuthN returns an azidentity.AuthenticationRecord and azcore.AccessToken and returns as variables to main.
func AuthN() (authData azidentity.AuthenticationRecord, accessToken azcore.AccessToken) {
	cred, err := azidentity.NewInteractiveBrowserCredential(nil)
	if err != nil {
		fmt.Println("Could not login. Please try again.")
	}
	ctx := context.TODO()
	authData, err = cred.Authenticate(ctx, nil)
	accessToken, err = cred.GetToken(ctx, policy.TokenRequestOptions{})
	return authData, accessToken
}

// AuthNLogOut clears all azidentity.AuthenticationRecord and azcore.AccessToken variable data and returns them as empty (authData and userToken).
func AuthNLogOut() (authData azidentity.AuthenticationRecord, accessToken azcore.AccessToken) {
	authData.Authority = ""
	authData.ClientID = ""
	authData.HomeAccountID = ""
	authData.TenantID = ""
	authData.Username = ""
	authData.Version = ""
	accessToken.Token = ""
	accessToken.ExpiresOn = time.Now()
	return authData, accessToken
}
