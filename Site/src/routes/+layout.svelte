<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>

<div class="flex h-16 items-center bg-gray-800">
    <div class="hidden sm:ml-6 sm:block pr-80">
        <div class="flex space-x-4">
            <a href="/" class="rounded-md px-3 py-2 text-sm font-medium text-gray-300 hover:bg-gray-700 hover:text-white">Dashboard</a>
            <a href="/about" class="rounded-md px-3 py-2 text-sm font-medium text-gray-300 hover:bg-gray-700 hover:text-white">About</a>
            <a href="/login" class="rounded-md px-3 py-2 text-sm font-medium text-gray-300 hover:bg-gray-700 hover:text-white">Login</a>
        </div>
    </div>
    <span class="flex text-3xl font-bold text-blue-300 pl-60">Cloud-native Guard Map</span>
</div>
{@render children()}
