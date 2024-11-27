<script lang="ts">
	import { Avatar, DropdownMenu } from 'bits-ui';
	import { flyAndScale } from '$lib/utils/flyAndScale';
	import Icon from '@iconify/svelte';

	let dropdownItemClass = `flex h-10 select-none items-center rounded-button py-2 pl-2 pr-1.5 text-sm font-medium !ring-0 !ring-transparent data-[highlighted]:bg-surface0 rounded transition`;

	interface Props {
		loggedIn: boolean;
	}

	let { loggedIn = $bindable(false) }: Props = $props();
</script>

<div
	class="flex mbl:flex-row flex-col bg-crust bg-opacity-75 backdrop-blur-sm shadow shadow-crust justify-center mbl:justify-between items-center gap-6 mbl:gap-2 top-0 left-0 min-h-[50px] pt-2 pb-2 w-full mb-5 px-6 sm:px-8 md:px-10 lg:px-20 xl:px-40"
>
	<p class="font-bold text-2xl">Stravule</p>

	{#if loggedIn && localStorage.getItem('jmeno')}
		<DropdownMenu.Root>
			<DropdownMenu.Trigger>
				<button
					class="p-1 pr-2 bg-base rounded-xl flex flex-row items-center text-subtext2 gap-1 transition hover:bg-surface0"
					onclick={() => {
						// localStorage.removeItem('sid');
						// goto('/');
					}}
				>
					<Avatar.Root>
						<Avatar.Image asChild={true} alt="Účet"
							><Icon width={24} icon="mdi:user" /></Avatar.Image
						>
					</Avatar.Root>
					<p class="text-inherit">{localStorage.getItem('jmeno')}</p>
				</button>
			</DropdownMenu.Trigger>
			<DropdownMenu.Content
				class="w-full max-w-[229px] rounded-xl border border-surface0 bg-mantle px-1 py-1.5 shadow shadow-mantle"
				transition={flyAndScale}
				sideOffset={8}
			>
				<DropdownMenu.Item class={dropdownItemClass}>
					<button
						onclick={() => {
							localStorage.removeItem('sid');
							window.location.href = '/';
						}}
						disabled={true}
						class="flex gap-2 h-10 w-full items-center text-subtext0"
					>
						<Icon class="text-xl" icon="mdi:settings" />
						Nastavení
					</button>
				</DropdownMenu.Item>
				<DropdownMenu.Item class={dropdownItemClass}>
					<button
						onclick={() => {
							localStorage.removeItem('sid');
							window.location.href = '/';
						}}
						class="flex gap-2 h-10 items-center w-full"
					>
						<Icon class="text-xl" icon="mdi:logout" />
						Odhlásit se
					</button>
				</DropdownMenu.Item>
			</DropdownMenu.Content>
		</DropdownMenu.Root>
	{/if}

	<!-- {#if loggedIn} -->
	<!-- 	<button -->
	<!-- 		class="hover:text-blue transition flex items-center justify-center text-center gap-2" -->
	<!-- 		onclick={() => { -->
	<!-- 			localStorage.removeItem('sid'); -->
	<!-- 			goto('/'); -->
	<!-- 		}} -->
	<!-- 	> -->
	<!-- 		<Icon icon="mdi:logout" /> -->
	<!-- 		Odhlásit se</button -->
	<!-- 	> -->
	<!-- {/if} -->
</div>
