<script lang="ts">
	import type { Filters } from '$lib/api/filters';
	import FilterCategory from './filterCategory.svelte';

	interface FiltersProps {
		filters?: Filters;
	}

	let { filters = $bindable() }: FiltersProps = $props();
</script>

{#if filters}
	<div
		class="flex flex-col w-full sm:flex-row gap-8 sm:gap-4 items-center msm:items-start rounded-xl"
	>
		<FilterCategory
			categoryTitle={'Chci'}
			categorySubtitle={'Jídla obsahující zadané slovo v názvu, která bych si přál/a'}
			category="include"
			bind:filters={filters.include}
		/>

		<FilterCategory
			categoryTitle={'Nechci'}
			categorySubtitle={'Jídla obsahující zadané slovo v názvu, která bych si nepřál/a'}
			category="exclude"
			bind:filters={filters.exclude}
		/>
	</div>
{/if}
