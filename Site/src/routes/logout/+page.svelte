<script lang="ts">
    import { goto } from '$app/navigation';
    import { existingSubscriptions, loggedInUser, selectedSubscriptionID, selectedSubscriptionName, subscriptionButtonState, subscriptionIsSelectedState } from "../../stores/logindata.js";
    
    async function azureLogout() {
        const response = await fetch('http://localhost:5000/auth-logout', { method: 'POST' });
        if (response.ok) {
            await response.json();
            loggedInUser.set('abstract random penguin caesar');
            loggedInUser.set('');
            existingSubscriptions.update(() => ['abstract', 'random', 'penguin' , 'caesar']);
            existingSubscriptions.set(['']);
            selectedSubscriptionName.set('abstract random penguin caesar');
            selectedSubscriptionName.set('');
            selectedSubscriptionID.set('');
            selectedSubscriptionID.set('');
            subscriptionButtonState.set(false);
            subscriptionIsSelectedState.set(false);
            goto('/login');            
        }
    }
</script>

<div class="text-center">
    <button class="rounded-lg border-2 border-black bg-blue-300 justify-center text-lg" on:click={azureLogout}>Logout from Azure</button>
</div>