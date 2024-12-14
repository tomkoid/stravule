<script lang="ts">
	import Icon from '@iconify/svelte';
	import CategoryList from './categoryList.svelte';
	import type { Filter } from '$lib/api/filters';

	interface FiltersProps {
		filters?: Filter[];
		categoryTitle: string;
		categorySubtitle: string;
		category: string;
	}

	let { filters = $bindable(), categoryTitle, categorySubtitle, category }: FiltersProps = $props();
</script>

{#if filters}
	<div class="flex-1">
		<p class="font-bold flex flex-row items-center gap-1 text-lg">
			<Icon icon="mdi:tick" />
			{categoryTitle}
		</p>
		<p class="text-subtext0 text-sm">{categorySubtitle}</p>
		<div class="rounded-xl mt-4">
			<CategoryList bind:filters {category} />
		</div>
	</div>
{/if}
