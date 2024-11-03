<script lang="ts">
	import '../app.css';

	import { afterNavigate, beforeNavigate, goto } from '$app/navigation';
	import { navigating } from '$app/stores';
	import { onMount } from 'svelte';
	import { pageLoading } from '$lib/stores/page.svelte';

	import RouteLoader from '$lib/components/layout/RouteLoader.svelte';

	let { children } = $props();

	let loggedIn: boolean = $state(false);
	// let pageLoading: boolean = $state(false);

	function refreshLoginStatus() {
		localStorage.getItem('sid') ? (loggedIn = true) : (loggedIn = false);
	}

	onMount(() => {
		if (localStorage) refreshLoginStatus();
	});

	$effect(() => {
		if ($navigating) refreshLoginStatus();
	});

	beforeNavigate(() => {
		pageLoading.value = true;
	});

	afterNavigate(() => {
		pageLoading.value = false;
	});
</script>

<div hidden={!pageLoading.value}>
	<RouteLoader />
</div>

<div
	class="sticky flex flex-row bg-ctp-mantle shadow shadow-ctp-mantle justify-center md:justify-between items-center gap-2 top-0 left-0 h-[50px] w-full mb-5 px-20 md:px-10 lg:px-20 xl:px-40"
>
	<p class="font-bold text-2xl">Stravule</p>
	{#if loggedIn}
		<button
			onclick={() => {
				localStorage.clear();
				goto('/');
			}}>Logout</button
		>
	{/if}
</div>

<div class="mb-5 mx-5 md:mx-10 lg:mx-20 xl:mx-40">
	{@render children()}
</div>
