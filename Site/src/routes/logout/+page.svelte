<script lang="ts">
    import { goto } from '$app/navigation';
    import { actionsCaption, actionsDesc, actionsLink, actionsMetaID, actionsType, existingRecommendations, existingResources, existingSubscriptions, expandedAdvisorButtonIdx, impactedField, impactedValue, impactfromAlert, loggedInUser, longDesc, potentialBenefits, recExpandButton, recID, recName, selectedSubscriptionID, selectedSubscriptionName, shortDesc, shortSol, subscriptionButtonState, subscriptionIsSelectedState, threatModelButtonState, threatModelDrawing, threatModelGeneratedState, zeroRecs } from "../../stores/persistentsession.js";
    
    async function azureLogout() {
        const response = await fetch('http://localhost:5000/auth-logout', { method: 'POST' });
        if (response.ok) {
            await response.json();
            loggedInUser.set('');
            existingSubscriptions.set([]);
            selectedSubscriptionName.set('');
            selectedSubscriptionID.set('');
            selectedSubscriptionID.set('');
            subscriptionButtonState.set(false);
            subscriptionIsSelectedState.set(false);
            threatModelButtonState.set(false);
            threatModelGeneratedState.set(false);
            recExpandButton.set(false);
            expandedAdvisorButtonIdx.set(null);
            zeroRecs.set(true);
            existingResources.set([]);
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
            goto('/login');            
        }
    }
</script>

<div class="text-center">
    <button class="rounded-lg border-2 border-black bg-blue-300 justify-center text-lg" on:click={azureLogout}>Logout from Azure</button>
</div>