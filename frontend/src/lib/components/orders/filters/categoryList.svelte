<script lang="ts">
	import { removeFilter } from '$lib/api/filters';

	interface FiltersProps {
		filters?: string[];
	}

	let { filters = $bindable() }: FiltersProps = $props();
</script>

{#each filters! as filter}
	<div class="flex flex-row items-center gap-2">
		<button
			onclick={() => {
				if (!filters) console.error('filters is undefined while removing a filter');
				filters = filters!.filter((f) => f != filter);

				if (!localStorage) console.error('localStorage is undefined');
				if (!localStorage.getItem('sid'))
					console.error('sid not found in localStorage while removing filter');
				if (!localStorage.getItem('canteen'))
					console.error('canteen not found in localStorage while removing filter');
				removeFilter(filter);
			}}>X</button
		>
		<p>{filter}</p>
	</div>
{/each}
