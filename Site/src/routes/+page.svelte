<script lang="ts">
    import { goto } from "$app/navigation";
    import { loggedInUser, existingSubscriptions, subscriptionButtonState, selectedSubscription, subscriptionIsSelectedState, threatModelGeneratedState, subidfromGin} from "../stores/logindata.js";

    // Primary Azure SubscriptionsFactoryClient auth handler and returns list of subscriptions for logged in user's tenant.
    async function azureStartSubscriptionsAuth() {
        const response = await fetch('http://localhost:5000/subscriptions-auth', { method: 'POST' });
        if (response.ok) {
            let subs = await response.json();
            let subscriptionsfromGin = subs.output
            existingSubscriptions.update(($existingSubscriptions) => [...$existingSubscriptions, ...subscriptionsfromGin]);
        }
        goto('/');
        $subscriptionButtonState = true;
    }

    // Authentication handler for additional auths required for diagram, such as ARMAdvisor Client Factory.
    // NEEDS WORK
    async function azureStartAdditionalAuths(subscriptionid:string) {
        const response = await fetch("http://localhost:5000/advisor-auth", {method: "POST", headers: { "Content-Type": "application/json"}, 
        body: JSON.stringify({ subidData: subscriptionid })
        });
        if (response.ok) {
            let responsetext = await response.json();
            let responsetextfromGin = responsetext.subid;
            subidfromGin.set(responsetextfromGin);
        }
    }

    // Sets the Azure subscription state amongst Svelte stores and other components.
    function azureSetSubscription(subscriptionid:string) {
        selectedSubscription.set(subscriptionid);
        $subscriptionIsSelectedState = true;
        return selectedSubscription;
    }

    // Unsets the Azure subscription state amongst Svelte stores and other components.
    function azureReleaseSubscription() {
        selectedSubscription.set('abstract random penguin caesar');
        selectedSubscription.set('');
        $subscriptionIsSelectedState = false;
    }

    function azureSetThreatModel() {
        $threatModelGeneratedState = true;
    }

    function azureReleaseThreatModel() {
        $threatModelGeneratedState = false;
    }

</script>

<div>
    {#snippet selectedSubscriptionSnippet()}
        <p class="font-bold">Current selected subscription:</p>{$selectedSubscription}       
    {/snippet}

    {#snippet availableAzureSubscriptionsSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 justify-center text-lg" on:click={azureStartSubscriptionsAuth}>Show available Azure Subscriptions</button>
    {/snippet}

    {#snippet eachAvailableAzureSubscriptionsSnippet()}
        {#each $existingSubscriptions as subscriptions}
        <br/><button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" on:click={() => azureSetSubscription(subscriptions.split(" | ")[1])}>{subscriptions}</button><br/>
        {/each}
    {/snippet}

    {#snippet changeSelectedSubscriptionSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" on:click={() => azureReleaseSubscription()}>Change Subscription</button>
    {/snippet}

    {#snippet startRenderThreatModelSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" on:click={() => { azureStartAdditionalAuths($selectedSubscription); azureSetThreatModel(); }}>Start Threat Model generation</button>
    {/snippet}

    {#snippet displayThreatModelSnippet()}
        <p>Under construction</p>
        <p>{$subidfromGin}</p>
    {/snippet}

    <!-- If the user is logged in and a subscription is not selected. -->
    {#if $loggedInUser != '' && $subscriptionButtonState == false}
    {@render availableAzureSubscriptionsSnippet()}

    <!-- If the user is logged in, they clicked choose a subscription and a subscription is not yet selected. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == false}
    <center>
    {@render eachAvailableAzureSubscriptionsSnippet()}
    </center>

    <!-- If the user is logged in and has selected a subscription. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == true && $threatModelGeneratedState == false}
    <center>
    {@render selectedSubscriptionSnippet()}<br/>
    {@render changeSelectedSubscriptionSnippet()}<br/><br/>
    {@render startRenderThreatModelSnippet()}
    </center>

    <!-- If the user is logged in, has selected a subscription and has selected to start threat model generation. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == true && $threatModelGeneratedState == true}
    {@render displayThreatModelSnippet()}

    <!-- If the user is not logged in. -->
    {:else if $loggedInUser == ''}
    <div class="text-center">Not logged in. Please log in above to continue.</div>
    {/if}
</div>
