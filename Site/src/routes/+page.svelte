<script lang="ts">
    // Svelte tick functionality for dom re-rendering.
    import { tick } from "svelte";
    // Svelte stores for logged in user related variables.
    import { loggedInUser, threatModelButtonState } from "../stores/persistentsession.js";
    // Svelte stores for threat model related variables.
    import { threatModelGeneratedState } from "../stores/persistentsession.js";
    // Svelte stores for Resources.
    import { existingResources, resourcesGeneratedState } from "../stores/persistentsession.js";
    // Svelte stores for Subscriptions.
    import { existingSubscriptions, subscriptionButtonState, selectedSubscriptionName, selectedSubscriptionID, subscriptionIsSelectedState } from "../stores/persistentsession.js";
    // Svelte stores for Advisor.
    import { recExpandButton, zeroRecs, existingRecommendations, advisorRecommendationsGeneratedState, recName, recID, shortDesc, shortSol, longDesc, actionsDesc, actionsType, actionsCaption, actionsLink, actionsMetaID, impactfromAlert, impactedField, impactedValue, potentialBenefits } from "../stores/persistentsession.js";
    // Svelte stores for button states.
    import { expandedAdvisorButtonIdx } from "../stores/persistentsession.js"


    // Helper variables variables.
    const splitVar = " | "
    const bogusText = "abstract random penguin caesar"

    // Recommendations alert toggle states.
    function toggleAlertExpand(idx: any) {
        $expandedAdvisorButtonIdx = expandedAdvisorButtonIdx === idx ? null : idx;
    }

    // Primary Azure SubscriptionsFactoryClient auth handler and returns list of subscriptions for logged in user's tenant.
    async function azureStartSubscriptionsAuth() {
        await fetch('http://localhost:5000/subscriptions-auth', { method: 'POST' })
        .then(async (response) => {
            if (response.ok) {
            let subs = await response.json();
            let subscriptionsfromGin = subs.output
            existingSubscriptions.update(($existingSubscriptions) => [...$existingSubscriptions, ...subscriptionsfromGin]);
            };
        subscriptionButtonState.set(true);
        })
    }

    // Authentication handler and primary triggering function for additional auths required for diagram, such as ARMAdvisor Client Factory.
    async function azureStartAdvisorAuths(subscriptionid: string) {
        await fetch("http://localhost:5000/advisor-auth", {method: "POST", headers: { "Content-Type": "application/json"}, 
        body: JSON.stringify({ subscriptionid })
        }).then(async (response) => {
            if (response.ok) {
            let subid = await response.json();
            return subid.output;
        };
        })
        await tick();
    }

    async function azureAdvisorRequest() {
        await fetch("http://localhost:5000/advisor-request", {method: "POST"}).then(async (response) => {
            if (response.ok) {
            let recs = await response.json();
            let recommendationsFromGin = recs.output;
            existingRecommendations.update(($existingRecommendations) => [...$existingRecommendations, ...recommendationsFromGin]);
            };
        })
        await tick();
    }

    async function azureStartResourcesAuth(subscriptionid: string) {
        await fetch("http://localhost:5000/resources-auth", {method: "POST", headers: { "Content-Type": "application/json"},
        body: JSON.stringify({ subscriptionid })
        }).then(async (response) => {
            if (response.ok) {
                let subid = await response.json();
                return subid.output;
            };
        })
        await tick();
    }

    async function azureResourcesRequest() {
        azureReleaseResources();
        await fetch("http://localhost:5000/resources-request", {method: "POST"}).then(async (response) => {
            if (response.ok) {
            let resources = await response.json();
            let resourcesFromGin = resources.output;
            existingResources.update(($existingResources) => [...$existingResources, ...resourcesFromGin])
            };
        })
        await tick();
    }

    async function azureReleaseResources() {
        existingResources.set([]);
    }

    // Sets the Azure subscription state amongst Svelte stores and other components.
    async function azureSetSubscription(subscriptionname: string, subscriptionid: string) {
        azureReleaseSubscription();
        selectedSubscriptionName.set(subscriptionname);
        selectedSubscriptionID.set(subscriptionid);
        await tick();
        return selectedSubscriptionName;
    }

    // Unsets the Azure subscription state amongst Svelte stores and other components.
    async function azureReleaseSubscription() {
        selectedSubscriptionName.set(bogusText);
        selectedSubscriptionID.set(bogusText);
        selectedSubscriptionName.set('');
        selectedSubscriptionID.set('');
        await tick();
    }

    async function azureToggleSubscriptionState() {
        if ($subscriptionIsSelectedState == true) {
            subscriptionIsSelectedState.set(false);
        } else {
            subscriptionIsSelectedState.set(true);
        }
        await tick();
    }

    async function azureToggleThreatModelState() {
        if ($threatModelGeneratedState == true) {
            threatModelGeneratedState.set(false);
        } else {
            threatModelGeneratedState.set(true);
        }
        await tick();
    }

    async function azureSetAdvisorRecommendations(recommendations: string[]) {
        await azureReleaseAdvisorRecommendations();
        await azureAdvisorRequest();
        if ($existingRecommendations.length > 0) {
            $zeroRecs = false;
            for (let i = 0; i < recommendations.length; i++) {
                const advisorsections = recommendations[i].split(splitVar);
                //if (advisorsections[10] == 'Security') {
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
                //} else {
                    //continue
                //}
            }
        } else {
            $zeroRecs = true;
        }
        await tick();
    }

    async function azureReleaseAdvisorRecommendations() {
        $zeroRecs = true;
        $existingRecommendations = [];
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
        await tick();
    }

    async function azureToggleAdvisorState() {
        if ($advisorRecommendationsGeneratedState == true) {
            advisorRecommendationsGeneratedState.set(false);
        } else {
            advisorRecommendationsGeneratedState.set(true);
        }
        await tick();
    }

    async function azureRecButtonHandler() {
        if ($recExpandButton == false) {
            $recExpandButton = true
        } else {
            $recExpandButton = false
        }
        await tick();
    }

    async function azureToggleResourcesState() {
        if ($resourcesGeneratedState == true) {
            resourcesGeneratedState.set(false);
        } else {
            resourcesGeneratedState.set(true);
        }
        await tick();
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
        <br/><button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {azureSetSubscription(subscriptions.split(splitVar)[1], subscriptions.split(splitVar)[0]); azureStartAdvisorAuths($selectedSubscriptionID); azureStartResourcesAuth($selectedSubscriptionID); azureToggleSubscriptionState()}}>{subscriptions}</button><br/>
        {/each}
    {/snippet}

    {#snippet changeSelectedSubscriptionSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {azureReleaseSubscription(); azureToggleThreatModelState(); azureReleaseAdvisorRecommendations(); azureToggleSubscriptionState()}}>Change Subscription</button>
    {/snippet}

    {#snippet startAdvisorRecommendationsSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => { azureToggleAdvisorState(); azureSetAdvisorRecommendations($existingRecommendations);}}>Check Azure Advisor security recommendations</button>
    {/snippet}

    {#snippet startResourcesSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => { azureToggleResourcesState(); azureResourcesRequest();}}>Check Azure Resources</button>
    {/snippet}

    {#snippet displayAdvisorRecommendationsSnippet()}
        <center><br/>
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {azureReleaseAdvisorRecommendations(); azureToggleAdvisorState() }}>Cancel and return to menu</button><br/><br/>
        <p class="font-semibold">Azure Security Posture Recommendations:</p><br/>
        {#if $recExpandButton == true && $zeroRecs == true}
        <p>No security recommendations available for this subscription.</p>
        {:else if $recExpandButton == true && $zeroRecs == false}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {azureRecButtonHandler()}}>Close Recommendations</button><br/>
        {#each $recName as _, i}
        <br/>
        <div class="my-3">
            <button class="p-4 rounded-lg border border-gray-300 bg-white shadow-sm cursor-pointer hover: bg-gray-50 transition-colors flex justify-between items-center" onclick={() => toggleAlertExpand(i)}>
                <div class="flex items-center space-x-3">
                    <span class="bg-blue-100 text-blue-800 px-2 py-1 rounded-md font-medium">Alert #{i+1}</span>
                    <span class="text-gray-700 font-medium">{$shortDesc[i]} | Severity: </span>
                </div>
                <span class="px-2 py-1 rounded-md text-sm font-medium"
                    class:bg-red-100={$impactedField[i] === 'High'}
                    class:text-red-800={$impactedField[i] === 'High'}
                    class:bg-yellow-100={$impactedField[i] === 'Medium'}
                    class:text-yellow-800={$impactedField[i] === 'Medium'}
                    class:bg-green-100={$impactedField[i] === 'Low'}
                    class:text-green-800={$impactedField[i] === 'Low'}>
                    {$impactedField[i]}
                </span>
            </button>

            {#if $expandedAdvisorButtonIdx === i}
            <div class="mt-1 p-5 rounded-b-lg border border-t-0 border-gray-300 bg-white shadow-sm animate-expandVertical">
                <div class="grid grid-cols-1 gap-3">
                    <div class="flex">
                        <p class="font-bold w-40">Alert ID:</p>
                        <span class="text-gray-700 font-mono">{$recName[i]}</span>
                    </div>
                    <div class="flex">
                        <p class="font-bold w-40">Resource:</p>
                        <span class="text-gray-700">{$recID[i]}</span>
                    </div>
                    <div class="flex">
                        <p class="font-bold w-40">Risk Type:</p>
                        <span class="text-gray-700">{$impactfromAlert[i]}</span>
                    </div>
                    <div class="flex">
                        <p class="font-bold w-40">Resource Type:</p>
                        <span class="text-gray-700">{$impactedValue[i]}</span>
                    </div>
                    <!--<div class="flex flex-col">
                        <p class="font-bold mb-1">Long Description:</p>
                        <p class="text-gray-700 bg-gray-50 p-3 rounded">{$longDesc[i]}</p>
                    </div>
                    <div class="flex flex-col">
                        <p class="font-bold mb-1">Potential Benefits:</p>
                        <p class="text-gray-700 bg-blue-50 p-3 rounded">{$potentialBenefits[i]}</p>
                    </div> -->
                </div>
                <br/><br/>
            </div>
            {/if}
        </div>
        {/each}
        {:else if $recExpandButton == false}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {azureSetAdvisorRecommendations($existingRecommendations); azureRecButtonHandler()}}>Expand Recommendations</button>
        {/if}
        </center>
    {/snippet}

    {#snippet displayResourcesSnippet()}
    <center><br/>
    <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {azureReleaseResources(); azureToggleResourcesState() }}>Cancel and return to menu</button><br/><br/>
    <p class="font-semibold">Azure Resources:</p><br/>
    {$existingResources}
    </center>
    {/snippet}

    {#snippet generateThreatModelSnippet()}
    
    {/snippet}


    <!-- If the user is logged in and a subscription is not selected. -->
    {#if $loggedInUser != '' && $subscriptionButtonState == false}
    <center>
    {@render availableAzureSubscriptionsSnippet()}
    </center>
    <!-- If the user is logged in, they clicked choose a subscription and a subscription is not yet selected. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == false}
    <center>
    {@render eachAvailableAzureSubscriptionsSnippet()}
    </center>
    <!-- If the user is logged in and has selected a subscription. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == true && $advisorRecommendationsGeneratedState == false && $resourcesGeneratedState == false}
    <center>
    {@render selectedSubscriptionSnippet()}<br/><br/>
    {@render changeSelectedSubscriptionSnippet()}<br/><br/>
    {@render startAdvisorRecommendationsSnippet()}<br/><br/>
    {@render startResourcesSnippet()}
    </center>

    <!-- If the user is logged in, has selected a subscription and has selected to generate security advisor recommendations. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == true && $advisorRecommendationsGeneratedState == true && $resourcesGeneratedState == false}
    {@render displayAdvisorRecommendationsSnippet()}


    <!-- If the user is logged in, has selected a subscription and has selected to display resources in their subscription. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == true && $advisorRecommendationsGeneratedState == false && $resourcesGeneratedState == true}
    {@render displayResourcesSnippet()}

    <!-- If the user is not logged in. -->
    {:else if $loggedInUser == ''}
    <div class="text-center">Not logged in. Please log in above to continue.</div>
    {/if}
</div>
