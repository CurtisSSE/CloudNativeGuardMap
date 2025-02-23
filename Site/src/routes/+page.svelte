<script lang="ts">
    import { goto } from "$app/navigation";
    // Svelte stores for logged in user related variables.
    import { loggedInUser } from "../stores/logindata.js";
    // Svelte stores for threat model related variables.
    import { threatModelGeneratedState } from "../stores/logindata.js";
    // Svelte stores for Subscriptions.
    import { existingSubscriptions, subscriptionButtonState, selectedSubscriptionName, selectedSubscriptionID, subscriptionIsSelectedState } from "../stores/logindata.js";
    // Svelte stores for Advisor.
    import { existingRecommendations, advisorRecommendationsGeneratedState, recName, recID, shortDesc, shortSol, longDesc, actionsDesc, actionsType, actionsCaption, actionsLink, actionsMetaID, impactfromAlert, impactedField, impactedValue, potentialBenefits } from "../stores/logindata.js";
   
    // General variables.
    const splitVar = " | "
    const bogusText = "abstract random penguin caesar"

    // General states.
    let tickState = $state(false);
    let recExpandButton = $state(false);

    // Timeouts and async handlers.
    function resolveAfterTick() {
        tickState = true;
        return new Promise((resolve) => {
            setTimeout(() => {
                resolve("resolved");
            }, 2000);
            goto("/");
            tickState = false;
        });
    }

    // Primary Azure SubscriptionsFactoryClient auth handler and returns list of subscriptions for logged in user's tenant.
    async function azureStartSubscriptionsAuth() {
        const response = await fetch('http://localhost:5000/subscriptions-auth', { method: 'POST' });
        if (response.ok) {
            let subs = await response.json();
            let subscriptionsfromGin = subs.output
            existingSubscriptions.update(($existingSubscriptions) => [...$existingSubscriptions, ...subscriptionsfromGin]);
        }
        $subscriptionButtonState = true;
    }

    // Authentication handler and primary triggering function for additional auths required for diagram, such as ARMAdvisor Client Factory.
    async function azureStartAdvisorAuths(subscriptionid: string) {
        const response = await fetch("http://localhost:5000/advisor-auth", {method: "POST", headers: { "Content-Type": "application/json"}, 
        body: JSON.stringify({ subscriptionid })
        });
        if (response.ok) {
            let subid = await response.json();
            return subid.output;
        }
    }

    async function azureAdvisorRequest() {
        const response = await fetch("http://localhost:5000/advisor-request", {method: "POST"});
        if (response.ok) {
            let recs = await response.json();
            let recommendationsFromGin = recs.output;
            existingRecommendations.update(($existingRecommendations) => [...$existingRecommendations, ...recommendationsFromGin]);
        }
    }

    // Sets the Azure subscription state amongst Svelte stores and other components.
    async function azureSetSubscription(subscriptionname: string, subscriptionid: string) {
        azureReleaseSubscription();
        selectedSubscriptionName.set(subscriptionname);
        selectedSubscriptionID.set(subscriptionid);
        return selectedSubscriptionName;
    }

    // Unsets the Azure subscription state amongst Svelte stores and other components.
    async function azureReleaseSubscription() {
        selectedSubscriptionName.set(bogusText);
        selectedSubscriptionID.set(bogusText);
        selectedSubscriptionName.set('');
        selectedSubscriptionID.set('');
    }

    function azureToggleSubscriptionState() {
        if ($subscriptionIsSelectedState == true) {
            subscriptionIsSelectedState.set(false);
        } else {
            subscriptionIsSelectedState.set(true);
        }
    }

    function azureSetThreatModel() {
    }

    function azureReleaseThreatModel() {
    }

    async function azureSetAdvisorRecommendations(recommendations: string[]) {
        const resolve = await resolveAfterTick();
        azureReleaseAdvisorRecommendations();
        azureAdvisorRequest();
        for (let i = 0; i < recommendations.length; i++) {
            const advisorsections = recommendations[i].split(splitVar);
            if (advisorsections[10] == 'Security') {
                recName.update(($recName) => [...$recName, advisorsections[0]]);
                recID.update(($recID) => [...$recID, advisorsections[1].split('/providers/Microsoft.Advisor/rec')[0]]);
                shortDesc.update(($shortDesc) => [...$shortDesc, advisorsections[2]]);
                shortSol.update(($shortSol) => [...$shortSol, advisorsections[3]]);
                longDesc.update(($longDesc) => [...$longDesc, advisorsections[4]]);
                actionsDesc.update(($actionsDesc) => [...$actionsDesc, advisorsections[5]]);
                actionsType.update(($actionsType) => [...$actionsType, advisorsections[6]]);
                actionsCaption.update(($actionsCaption) => [...$actionsCaption, advisorsections[7]]);
                actionsLink.update(($actionsLink) => [...$actionsLink, advisorsections[8]]);
                actionsMetaID.update(($actionsMetaID) => [...$actionsMetaID, advisorsections[9]]);
                impactfromAlert.update(($impactfromAlert) => [...$impactfromAlert, advisorsections[10]]);
                impactedField.update(($impactedField) => [...$impactedField, advisorsections[11]]);
                impactedValue.update(($impactedValue) => [...$impactedValue, advisorsections[12]]);
                potentialBenefits.update(($potentialBenefits) => [...$potentialBenefits, advisorsections[13]]);
            } else {
                continue
            }
        }
        console.log(resolve);
    }

    async function azureReleaseAdvisorRecommendations() {
        existingRecommendations.set([]);
        recName.set([]);
        recID.set([]);
        shortDesc.set([]);
        shortSol.set([]);
        longDesc.set([]);
        actionsDesc.set([]);
        actionsType.set([]);
        actionsCaption.set([]);
        actionsLink.set([]);
        actionsMetaID.set([]);
        impactfromAlert.set([]);
        impactedField.set([]);
        impactedValue.set([]);
        potentialBenefits.set([]);
    }

    function azureToggleAdvisorState() {
        if ($advisorRecommendationsGeneratedState == true) {
            advisorRecommendationsGeneratedState.set(false);
        } else {
            advisorRecommendationsGeneratedState.set(true);
        }
    }

    function azureRecButtonHandler() {
        if (recExpandButton == false) {
            recExpandButton = true
        } else {
            recExpandButton = false
        }
    }


</script>

<div>
    {#snippet selectedSubscriptionSnippet()}
        <p class="font-bold">Current selected subscription:</p>{$selectedSubscriptionName}       
    {/snippet}

    {#snippet availableAzureSubscriptionsSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 justify-center text-lg" onclick={azureStartSubscriptionsAuth}>Show available Azure Subscriptions</button>
    {/snippet}

    {#snippet eachAvailableAzureSubscriptionsSnippet()}
        <p class="font-bold">Choose an Azure subscription:</p>
        {#each $existingSubscriptions as subscriptions}
        <br/><button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {azureSetSubscription(subscriptions.split(splitVar)[1], subscriptions.split(splitVar)[0]); azureStartAdvisorAuths($selectedSubscriptionID); azureToggleSubscriptionState()}}>{subscriptions}</button><br/>
        {/each}
    {/snippet}

    {#snippet changeSelectedSubscriptionSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {azureReleaseSubscription(); azureReleaseThreatModel(); azureReleaseAdvisorRecommendations(); azureToggleSubscriptionState()}}>Change Subscription</button>
    {/snippet}

    {#snippet startAdvisorRecommendationsSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {azureToggleAdvisorState()}}>Check Azure Advisor security recommendations</button>
    {/snippet}

    {#snippet displayAdvisorRecommendationsSnippet()}
        <center><br/>
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {azureReleaseAdvisorRecommendations(); azureToggleAdvisorState() }}>Cancel and return to menu</button><br/><br/>
        <p class="font-semibold">Azure Security Posture Recommendations:</p><br/>
        {#if recExpandButton == true}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {azureRecButtonHandler()}}>Close Alerts</button><br/>
        {#each $recName as _, i}
        <br/>
        <p class="font-bold">Alert Number:</p> {i+1}
        <p class="font-bold">Summary:</p> {$shortDesc[i]}
        <p class="font-bold">Alert ID:</p> {$recName[i]}
        <p class="font-bold">Short Description:</p> {$shortDesc[i]}
        <p class="font-bold">Long Description:</p> {$longDesc[i]}
        <p class="font-bold">Risk Type:</p> {$impactfromAlert[i]}
        <p class="font-bold">Severity:</p> {$impactedField[i]}
        <p class="font-bold">Resource Type:</p> {$impactedValue[i]}
        <p class="font-bold">Resource:</p> {$recID[i]}
        <p class="font-bold">Potential Benefits:</p> {$potentialBenefits[i]}
        <br/><br/>
        {/each}
        {:else if recExpandButton == false}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {azureSetAdvisorRecommendations($existingRecommendations); azureRecButtonHandler()}}>Expand Alerts</button>
        {/if}
        
        </center>
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
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == true && $advisorRecommendationsGeneratedState == false}
    <center>
    {@render selectedSubscriptionSnippet()}<br/><br/>
    {@render changeSelectedSubscriptionSnippet()}<br/><br/>
    {@render startAdvisorRecommendationsSnippet()}
    </center>

    <!-- If the user is logged in, has selected a subscription and has selected to generate security advisor recommendations. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == true && $advisorRecommendationsGeneratedState == true}
    {@render displayAdvisorRecommendationsSnippet()}

    <!-- If the user is not logged in. -->
    {:else if $loggedInUser == ''}
    <div class="text-center">Not logged in. Please log in above to continue.</div>
    {/if}
</div>
