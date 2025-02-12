<script lang="ts">
    import { goto } from "$app/navigation";
    import { loggedInUser, existingSubscriptions } from "../stores/logindata.js";

    async function azureGetSubscriptions() {
        const response = await fetch('http://localhost:5000/get-subscriptions', { method: 'POST' });
        if (response.ok) {
            let result = await response.json()
            let strresult = JSON.stringify(result.output)
            existingSubscriptions.set(strresult.replace(/"/g, '').replace(/%/g, '/nl'))
            goto('/')
        }
    }
</script>

<div>
    {#if $loggedInUser != ''}
        <div class="text-center">Logged in as: {$loggedInUser}</div>
        <button class="rounded-lg border bg-blue-300 justify-center text-lg" on:click={azureGetSubscriptions}>List Subs</button>
        <div class="text-center">{$existingSubscriptions}</div>
    {:else}
    <div class="text-center">Not logged in. Please authenticate your Azure account.</div>
    {/if}
</div>
