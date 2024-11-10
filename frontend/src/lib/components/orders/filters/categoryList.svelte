<script lang="ts">
	import { removeFilter } from '$lib/api/filters';
	import Icon from '@iconify/svelte';

	interface FiltersProps {
		filters?: string[];
	}

	let { filters = $bindable() }: FiltersProps = $props();
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
