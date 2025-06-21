<script lang="ts">
	import '../app.css';

	import { afterNavigate, beforeNavigate } from '$app/navigation';
	import { navigating } from '$app/stores';
	import { onMount } from 'svelte';
	import { pageLoading } from '$lib/stores/page.svelte';

	import { getUserInfo } from '$lib/api/user_info';

	import RouteLoader from '$lib/components/layout/RouteLoader.svelte';
	import Navbar from '$lib/components/layout/Navbar.svelte';
	import Errors from '$lib/components/layout/Errors.svelte';

	let { children } = $props();

	let loggedIn: boolean = $state(false);

	function refreshLoginStatus() {
		localStorage.getItem('sid') ? (loggedIn = true) : (loggedIn = false);
	}

	function isLoggedIn() {
		return localStorage.getItem('sid') ? true : false;
	}

	onMount(async () => {
		if (localStorage) {
			refreshLoginStatus();

			if (!localStorage.getItem('jmeno') && loggedIn) {
				await getUserInfo(localStorage.getItem('sid')!, localStorage.getItem('canteen')!);
				console.log('reloading web page after getting user info');

				window.location.reload();
			}
		}
	});

	$effect(() => {
		if ($navigating) {
			if (localStorage) {
				refreshLoginStatus();
			}
		}
		if (!isLoggedIn() && window.location.pathname != '/') window.location.href = '/';
	});

	beforeNavigate(() => {
		pageLoading.value = true;
	});

	afterNavigate(() => {
		pageLoading.value = false;
	});
</script>

<Errors />

<div hidden={!pageLoading.value}>
	<RouteLoader />
</div>

<Navbar bind:loggedIn />

<div class="mb-5 mx-5 md:mx-10 lg:mx-20 xl:mx-40">
	{@render children()}
</div>
