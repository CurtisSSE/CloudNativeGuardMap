<script lang="ts">
    import { goto } from '$app/navigation';
    import { loggedInUser } from "../../stores/logindata.js";
    
    async function azureLogout() {
        const response = await fetch('http://localhost:5000/auth-logout', { method: 'POST' });
        if (response.ok) {
            await response.json()
            goto('/login')
            loggedInUser.update(() => 'abstract random penguin ceaser')
            loggedInUser.set('')
            
        }
    }
</script>

<div class="text-center">
    <button class="rounded-lg border bg-blue-300 justify-center text-lg" on:click={azureLogout}>Logout from Azure</button>
</div>