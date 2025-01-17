<script lang="ts">
    import { loggedInUser } from "../stores/logindata.js";
</script>

<div class="text-center">
    {#if $loggedInUser != ''}
        Logged in as: {$loggedInUser}
    {:else}
        Not logged in. Please authenticate your Azure account.
    {/if}
</div>
