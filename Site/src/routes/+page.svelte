<script>
    import { onMount } from 'svelte';
    let data = null;

    // Function to fetch user data
    async function loginData() {
        try {
            const res = await fetch('http://localhost:8080/auth-data');
            if (res.ok) {
                data = await res.json();
            } else {
                console.error("Failed to fetch data.");
            }
        } catch (error) {
            console.error("Error fetching data:", error);
        }
    }

    // Function to logout
    async function clearData() {
        try {
            const res = await fetch('http://localhost:8080/auth-logout', { method: 'POST' });
            if (res.ok) {
                data = null; // Clear the data in frontend
                console.log("Logged out successfully.");
            } else {
                console.error("Failed to log out.");
            }
        } catch (error) {
            console.error("Error logging out:", error);
        }
    }

    // Fetch data on component mount
    onMount(loginData);
</script>

<div class="text-center">
    {#if data}
        Logged in account: {data.userupn}
    {:else}
        Not logged in. Please authenticate your Azure account.
    {/if}
</div>
