package Subscription

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

func GetAllSubscriptions(subscriptionsclientvariable *armsubscriptions.Client) (subscriptionoutput map[string]string) {
	currentctx := context.Background()
	currentpager := subscriptionsclientvariable.NewListPager(nil)
	subscriptionoutput = make(map[string]string)
	for currentpager.More() {
		currentpage, err := currentpager.NextPage(currentctx)
		if err != nil {
			fmt.Println("Unable to retrieve subscriptions. Check authentication.", err)
		}
		for _, sub := range currentpage.Value {
			if sub.DisplayName != nil && sub.SubscriptionID != nil {
				subscriptionoutput[*sub.SubscriptionID] = *sub.DisplayName
			} else if sub.DisplayName == nil {
				subscriptionoutput[*sub.SubscriptionID] = "No display name"
			} else {
				continue
			}
		}
	}
	return subscriptionoutput
}
