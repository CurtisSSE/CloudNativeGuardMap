<script>
    async function azureLogin() {
        const response = await fetch('http://localhost:8080/auth-login', { method: 'POST' });
        if (response.ok) {
            const result = await response.json()
        }
    }
    async function azureLogout() {
        const response = await fetch('http://localhost:8080/auth-logout', { method: 'POST' });
        if (response.ok) {
            const result = await response.json()
        }
    }
</script>

<div class="text-center">
    <button class="rounded-lg border bg-blue-300 justify-center text-lg" on:click={azureLogin}>Authenticate to Azure</button>
    <button class="rounded-lg border bg-blue-300 justify-center text-lg" on:click={azureLogout}>Logout</button>
</div>