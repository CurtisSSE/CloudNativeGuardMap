package Auth

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity/cache"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

type AuthState struct {
	AuthData    azidentity.AuthenticationRecord
	AccessToken azcore.AccessToken
	LoggedIn    bool
}

type SubscriptionsClientFactoryState struct {
	SubscriptionsFactoryClient         *armsubscriptions.Client
	SubscriptionsFactoryClientLoggedIn bool
}

// Shared Azure identity states across packages and main.
var (
	SharedCreds                      azcore.TokenCredential
	BrowserCreds                     *azidentity.InteractiveBrowserCredential
	CurrentAuthState                 AuthState
	CurrentSubscriptionsFactoryState SubscriptionsClientFactoryState
)

// CurrentAuthState internal variables.
var (
	authData    azidentity.AuthenticationRecord
	accessToken azcore.AccessToken
)

// CurrentARMSubscriptionsFactoryState internal variables.

/* Authentication using the azidentity library (AuthenticationRecord) only returns non-secret account information as strings,
so byte manipulation, zeroing and so forth need not be considered for this particular data.*/

func AuthN() azcore.TokenCredential {
	cachevar, _ := cache.New(nil)
	BrowserCreds, _ = azidentity.NewInteractiveBrowserCredential(&azidentity.InteractiveBrowserCredentialOptions{Cache: cachevar})
	ctx := context.TODO()
	authData, _ = BrowserCreds.Authenticate(ctx, nil)
	accessToken, _ = BrowserCreds.GetToken(ctx, policy.TokenRequestOptions{})
	CurrentAuthState.AuthData = authData
	CurrentAuthState.AccessToken = accessToken
	CurrentAuthState.LoggedIn = true
	SharedCreds = BrowserCreds
	return SharedCreds
}

func AuthNLogOut() {
	SharedCreds = nil
	BrowserCreds = nil
	authData.Authority = ""
	authData.ClientID = ""
	authData.HomeAccountID = ""
	authData.TenantID = ""
	authData.Username = ""
	authData.Version = ""
	accessToken.Token = ""
	accessToken.ExpiresOn = time.Now()
	CurrentAuthState.AuthData = authData
	CurrentAuthState.AccessToken = accessToken
	CurrentAuthState.LoggedIn = false
}

func BuildArmSubscriptionsClientFactory(tokencreds azcore.TokenCredential) {
	clientFactory, err := armsubscriptions.NewClientFactory(tokencreds, nil)
	if err != nil {
		fmt.Println("Could not login. Please try again.")
		CurrentSubscriptionsFactoryState.SubscriptionsFactoryClientLoggedIn = false
	}
	CurrentSubscriptionsFactoryState.SubscriptionsFactoryClient = clientFactory.NewClient()
	CurrentSubscriptionsFactoryState.SubscriptionsFactoryClientLoggedIn = true
}
