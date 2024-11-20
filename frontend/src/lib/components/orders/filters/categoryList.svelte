<script lang="ts">
	import { addFilter, removeFilter } from '$lib/api/filters';
	import Icon from '@iconify/svelte';

	interface FiltersProps {
		filters?: string[];
		category: string;
	}

	let { filters = $bindable(), category }: FiltersProps = $props();

	let newFilter: string = $state('');
	let confirmButton: HTMLButtonElement;
</script>

{#each filters! as filter}
	<div class="flex flex-row items-center gap-1">
		<p class="text text-sm">{filter}</p>
		<button
			class="transition rounded-xl text-ctp-red hover:bg-ctp-surface1"
			onclick={() => {
				if (!filters) console.error('filters is undefined while removing a filter');
				filters = filters!.filter((f) => f != filter);

				if (!localStorage) console.error('localStorage is undefined');
				if (!localStorage.getItem('sid'))
					console.error('sid not found in localStorage while removing filter');
				if (!localStorage.getItem('canteen'))
					console.error('canteen not found in localStorage while removing filter');
				removeFilter(filter);
			}}
		>
			<Icon icon="material-symbols:delete" />
		</button>
	</div>
{/each}

<div
	class="flex flex-row items-center gap-1 bg-ctp-surface0 border border-ctp-surface0 mt-2 rounded w-fit text-sm"
>
	<input
		placeholder="brambory, ..."
		class="border-none bg-ctp-surface0 pl-1 w-28 h-6"
		type="text"
		onkeypresscapture={(e) => {
			if (e.key == 'Enter') {
				e.preventDefault();
				confirmButton.click();
			}
		}}
		bind:value={newFilter}
	/>
	<button
		class="flex flex-row items-center gap-1 px-2 h-6 bg-ctp-mantle transition disabled:bg-ctp-crust hover:bg-ctp-crust"
		bind:this={confirmButton}
		disabled={!newFilter}
		onclick={async () => {
			await addFilter(newFilter, category);

			// add filter to the list
			filters = [...filters!, newFilter];
			newFilter = '';
		}}
	>
		<Icon icon="material-symbols:add" />
		Přidat
	</button>
</div>
