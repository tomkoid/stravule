<script lang="ts">
	import type { Filters } from '$lib/api/filters';
	import CategoryList from './categoryList.svelte';

	interface FiltersProps {
		filters?: Filters;
	}

	let { filters = $bindable() }: FiltersProps = $props();
</script>

{#if filters}
	<div class="flex flex-col gap-3 bg-ctp-surface1 p-3 rounded-xl">
		<p class="font-bold">Pozitivni</p>
		<div class="bg-ctp-surface2">
			<CategoryList bind:filters={filters.include} />
		</div>
		<hr class="border-ctp-mantle" />
		<p class="font-bold">Negativni</p>
		<div class="bg-ctp-surface2">
			<CategoryList bind:filters={filters.exclude} />
		</div>
	</div>
{/if}
