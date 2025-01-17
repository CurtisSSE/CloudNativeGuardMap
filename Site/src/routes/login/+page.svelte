<script lang="ts">
    import { loggedInUser } from "../../stores/logindata.js";

    async function azureLogin() {
        const response = await fetch('http://localhost:8080/auth-login', { method: 'POST' });
        if (response.ok) {
            let result = await response.json()
            let strresult = JSON.stringify(result.user.username)
            let trimmedresult = 
            loggedInUser.set(strresult.replace(/"/g, ''))
        }
    }
</script>

<div class="text-center">
    <button class="rounded-lg border bg-blue-300 justify-center text-lg" on:click={azureLogin}>Authenticate to Azure</button>
    <button class="rounded-lg border bg-yellow-300 justify-center text-lg">Authenticate to AWS</button>
    <button class="rounded-lg border bg-green-300 justify-center text-lg">Authenticate to GCP</button>
</div>