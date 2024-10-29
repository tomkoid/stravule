<script lang="ts">
	import { login } from '$lib/api/login';
	let username = $state('test');
	let password = $state('testpass');
	let canteen = $state('0000');
</script>

<div class="flex flex-col mt-5">
	<input class="login-input" type="username" bind:value={username} />
	<input class="login-input" type="password" bind:value={password} />
	<input class="login-input" bind:value={canteen} />
	<button
		class="bg-lime-400 border p-2"
		onclick={() => {
			login(username, password, canteen);
		}}>Login</button
	>
</div>

