<script lang="ts">
	import type { Filters } from '$lib/api/filters';
	import CategoryList from './categoryList.svelte';

	interface FiltersProps {
		filters?: Filters;
	}

	let { filters = $bindable() }: FiltersProps = $props();
</script>

{#if filters}
	<div class="flex flex-col gap-1 bg-ctp-surface1 p-3 rounded-xl">
		<p class="font-bold">Chci</p>
		<div class="bg-ctp-surface2 rounded-xl p-2">
			<CategoryList bind:filters={filters.include} />
		</div>
		<p class="font-bold">Nechci</p>
		<div class="bg-ctp-surface2 rounded-xl p-2">
			<CategoryList bind:filters={filters.exclude} />
		</div>
	</div>
{/if}
