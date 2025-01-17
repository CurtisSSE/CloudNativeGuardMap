<script lang="ts">
    import { goto } from "$app/navigation";
    import { loggedInUser } from "../../stores/logindata.js";
    
    async function azureLogout() {
        const response = await fetch('http://localhost:8080/auth-logout', { method: 'POST' });
        if (response.ok) {
            await response.json()
            loggedInUser.set('')
        }
        goto('/login')
    }
</script>

<div class="text-center">
    <button class="rounded-lg border bg-blue-300 justify-center text-lg" on:click={azureLogout}>Logout from Azure</button>
</div>