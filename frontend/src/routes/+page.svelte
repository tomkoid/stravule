<script lang="ts">
	import { goto } from '$app/navigation';
	import { login } from '$lib/api/login';
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
	<input
		class="login-input"
		type="username"
		placeholder="Uživatelské jméno"
		bind:value={username}
	/>
	<input class="login-input" type="password" placeholder="Heslo" bind:value={password} />
	<input class="login-input" placeholder="Jídelna" bind:value={canteen} />
	<button
		class="bg-ctp-blue text-ctp-base shadow-ctp-blue rounded-xl mt-5 p-2"
		onclick={() => {
			errorMessage = '';
			login(username, password, canteen)
				.then(() => {
					goto('/orders');
				})
				.catch((e) => {
					console.log(typeof e);
					if (typeof e === 'string') {
						errorMessage = e;
					} else if (e instanceof Error) {
						errorMessage = e.message;
					}
				});
		}}>Login</button
	>
	<p class="text-ctp-red">{errorMessage}</p>
</div>

<svelte:head>
	<title>Stravule - Přihlášení</title>
</svelte:head>
