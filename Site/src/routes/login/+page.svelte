<script lang="ts">
    import { goto } from '$app/navigation';
    import { loggedInUser } from "../../stores/persistentsession.js";


    async function azureLogin() {
        const response = await fetch('http://localhost:5000/auth-login', { method: 'POST' });
        if (response.ok) {
            let result = await response.json();
            let strresult = JSON.stringify(result.user.username);
            loggedInUser.set(strresult.replace(/"/g, ''));
            goto('/');
        }
    }
</script>

<div class="text-center">
    <button class="rounded-lg border-2 border-black bg-blue-300 justify-center text-lg" on:click={azureLogin}>Authenticate to Azure</button>
</div>