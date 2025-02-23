<script lang="ts">
	import '../app.css';
    import { loggedInUser, selectedSubscriptionName } from '../stores/logindata.js';
 
    let { children } = $props();
</script>

<div class="flex h-16 items-center justify-between bg-gray-800 px-4">
    <!-- Left Section: Buttons -->
    <div class="flex space-x-4">
        <a href="/" class="rounded-md px-3 py-2 text-sm font-medium text-gray-300 hover:bg-gray-700 hover:text-white">Dashboard</a>
        <a href="/about" class="rounded-md px-3 py-2 text-sm font-medium text-gray-300 hover:bg-gray-700 hover:text-white">About</a>
        {#if $loggedInUser != ''}
            <a href="/logout" class="rounded-md px-3 py-2 text-sm font-medium text-gray-300 hover:bg-gray-700 hover:text-white">Logout</a>
        {:else}
            <a href="/login" class="rounded-md px-3 py-2 text-sm font-medium text-gray-300 hover:bg-gray-700 hover:text-white">Login</a>
        {/if}
    </div>

    <!-- Center Section: Title -->
    <span class="text-3xl font-bold text-blue-300 text-center absolute left-1/2 transform -translate-x-1/2">
        Cloud-native Guard Map
    </span>

    <!-- Right Section: Logged-In Info -->
    <div class="flex items-center justify-end ">
        {#if $loggedInUser != ''}
            <p class="text-sm font-medium text-gray-300 p-4">Logged in as: {$loggedInUser}</p>
            <p class="text-sm font-medium text-gray-300 p-4">Azure Subscription: {$selectedSubscriptionName}</p>
        {/if}
    </div>
</div>


{@render children()}
