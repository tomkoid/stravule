<script lang="ts">
	import { goto } from '$app/navigation';
	import { login } from '$lib/api/login';
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';

	let username = $state('');
	let password = $state('');
	let canteen = $state('');
	let errorMessage = $state('');
	onMount(() => {
		if (localStorage) {
			if (localStorage.getItem('sid')) {
				goto('/orders');
			}
			localStorage.getItem('username') ? (username = localStorage.getItem('username')!) : null;
			localStorage.getItem('password') ? (password = localStorage.getItem('password')!) : null;
			localStorage.getItem('canteen') ? (canteen = localStorage.getItem('canteen')!) : null;
		}
	});
</script>

<div class="flex flex-col mt-5 gap-2">
	<h1 class="flex flex-row justify-start items-center gap-2 text-2xl font-extrabold mb-1">
		<Icon class="text-2xl" icon="tabler:login" />
		Přihlášení
	</h1>
	<input
		class="login-input"
		type="username"
		placeholder="Uživatelské jméno"
		bind:value={username}
	/>
	<input class="login-input" type="password" placeholder="Heslo" bind:value={password} />
	<input class="login-input" placeholder="Jídelna" bind:value={canteen} />
	<button
		class="flex flex-row flex-nowrap w-full justify-center items-center gap-2 bg-blue-300 hover:bg-blue-200 text-base shadow-surface0 rounded-xl hover:rounded-2xl mt-5 p-2 transition-all"
		onclick={() => {
			errorMessage = '';
			login(username, password, canteen)
				.then(() => {
					goto('/orders');
				})
				.catch((e) => {
					if (typeof e === 'string') {
						errorMessage = e;
					} else if (e instanceof Error) {
						errorMessage = e.message;
					}
				});
		}}
	>
		<Icon class="min-w-[16px] min-h-[16px]" color="inherit" icon="mdi:login" />
		Přihlásit se
	</button>
	<p class="text-ctp-red">{errorMessage}</p>
</div>

<svelte:head>
	<title>Stravule - Přihlášení</title>
</svelte:head>
