<script lang="ts">
    // Import all persistent state functions.
    import * as FunctionPersist from "../stores/persistentfunctions.js"
    // Import individual identity states.
    import { loggedInUser } from "../stores/persistentsession.js";
    // Import individual subscription states.
    import { selectedSubscriptionName, selectedSubscriptionID, subscriptionIsSelectedState, subscriptionButtonState, existingSubscriptions } from "../stores/persistentsession.js";
    // Import individual recommendation states.
    import { advisorRecommendationsGeneratedState, recExpandButton, zeroRecs, recName, recID, shortDesc, impactedField, impactfromAlert, impactedValue, expandedAdvisorButtonIdx, existingRecommendations} from "../stores/persistentsession.js";
    // Import individual resources states.
    import { resourcesGeneratedState, existingResources } from "../stores/persistentsession.js";
    // Import individual threat model states.
    import { threatModelGeneratedState } from "../stores/persistentsession.js";
    // Helper variables.
    const splitVar = " | "

</script>

<div>
    {#snippet selectedSubscriptionSnippet()}
        <p class="font-bold">Current selected subscription:</p>{$selectedSubscriptionName}       
    {/snippet}

    {#snippet availableAzureSubscriptionsSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 justify-center text-lg" onclick={FunctionPersist.azureStartSubscriptionsAuth}>Show available Azure Subscriptions</button>
    {/snippet}

    {#snippet eachAvailableAzureSubscriptionsSnippet()}
        <p class="font-bold">Choose an Azure subscription:</p>
        {#each $existingSubscriptions as subscriptions}
        <br/><button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {FunctionPersist.azureSetSubscription(subscriptions.split(splitVar)[1], subscriptions.split(splitVar)[0]); FunctionPersist.azureStartAdvisorAuths($selectedSubscriptionID); FunctionPersist.azureStartResourcesAuth($selectedSubscriptionID); FunctionPersist.azureToggleSubscriptionState()}}>{subscriptions}</button><br/>
        {/each}
    {/snippet}

    {#snippet changeSelectedSubscriptionSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {FunctionPersist.azureReleaseSubscription(); FunctionPersist.azureToggleThreatModelState(); FunctionPersist.azureReleaseAdvisorRecommendations(); FunctionPersist.azureToggleSubscriptionState()}}>Change Subscription</button>
    {/snippet}

    {#snippet startAdvisorRecommendationsSnippet()}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => { FunctionPersist.azureToggleAdvisorState(); FunctionPersist.azureSetAdvisorRecommendations();}}>Check Azure Advisor security recommendations</button>
    {/snippet}
    {#snippet startThreatModelSnippet()}
    <a href="/threatmodel" class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => { FunctionPersist.azureToggleThreatModelState();}}>Start Threat Modeling</a>
    {/snippet}


    {#snippet displayAdvisorRecommendationsSnippet()}
        <center><br/>
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {FunctionPersist.azureReleaseAdvisorRecommendations(); FunctionPersist.azureToggleAdvisorState() }}>Cancel and return to menu</button><br/><br/>
        <p class="font-semibold">Azure Security Posture Recommendations:</p><br/>
        {#if $recExpandButton == true && $zeroRecs == true}
        <p>No security recommendations available for this subscription.</p>
        {:else if $recExpandButton == true && $zeroRecs == false}
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {FunctionPersist.azureRecButtonHandler()}}>Close Recommendations</button><br/>
        {#each $recName as _, i}
        <br/>
        <div class="my-3">
            <button class="p-4 rounded-lg border border-gray-300 bg-white shadow-sm cursor-pointer hover: bg-gray-50 transition-colors flex justify-between items-center" onclick={() => FunctionPersist.toggleAlertExpand(i)}>
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
        <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {FunctionPersist.azureSetAdvisorRecommendations(); FunctionPersist.azureRecButtonHandler()}}>Expand Recommendations</button>
        {/if}
        </center>
    {/snippet}

    {#snippet displayResourcesSnippet()}
    <center><br/>
    <button class="rounded-lg border-2 border-black bg-blue-300 font-semibold" onclick={() => {FunctionPersist.azureReleaseResources(); FunctionPersist.azureToggleResourcesState() }}>Cancel and return to menu</button><br/><br/>
    <p class="font-semibold">Azure Resources:</p><br/>
    {$existingResources}
    </center>
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
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == true && $advisorRecommendationsGeneratedState == false && $resourcesGeneratedState == false && $threatModelGeneratedState == false}
    <center>
    {@render selectedSubscriptionSnippet()}<br/><br/>
    {@render changeSelectedSubscriptionSnippet()}<br/><br/>
    {@render startAdvisorRecommendationsSnippet()}<br/><br/>
    {@render startThreatModelSnippet()}
    </center>

    <!-- If the user is logged in, has selected a subscription and has selected to generate security advisor recommendations. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == true && $advisorRecommendationsGeneratedState == true && $resourcesGeneratedState == false && $threatModelGeneratedState == false}
    {@render displayAdvisorRecommendationsSnippet()}


    <!-- If the user is logged in, has selected a subscription and has selected to display resources in their subscription. -->
    {:else if $loggedInUser != '' && $subscriptionButtonState == true && $subscriptionIsSelectedState == true && $advisorRecommendationsGeneratedState == false && $resourcesGeneratedState == true && $threatModelGeneratedState == false}
    {@render displayResourcesSnippet()}

    <!-- If the user is not logged in. -->
    {:else if $loggedInUser == ''}
    <div class="text-center">Not logged in. Please log in above to continue.</div>
    {/if}
</div>
