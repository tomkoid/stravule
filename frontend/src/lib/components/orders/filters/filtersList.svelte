<script lang="ts">
	import type { Filters } from '$lib/api/filters';
	import Icon from '@iconify/svelte';
	import CategoryList from './categoryList.svelte';

	interface FiltersProps {
		filters?: Filters;
	}

	let { filters = $bindable() }: FiltersProps = $props();
</script>

{#if filters}
	<div
		class="flex flex-col sm:flex-row bg-ctp-mantle border border-ctp-surface0 gap-8 sm:gap-4 p-3 items-center msm:items-start pl-8 sm:pl-6 md:pl-4 rounded-xl"
	>
		<div class="flex-1">
			<p class="font-bold flex flex-row items-center gap-1">
				<Icon icon="mdi:tick" />
				Chci
			</p>
			<p class="text-ctp-subtext0 text-sm">
				Jídla obsahující zadané slovo v názvu, která bych si přál/a
			</p>
			<div class="rounded-xl mt-4">
				<CategoryList bind:filters={filters.include} category="include" />
			</div>
		</div>
		<div class="flex-1">
			<p class="font-bold flex flex-row items-center gap-1">
				<Icon icon="maki:cross" />
				Nechci
			</p>
			<p class="text-ctp-subtext0 text-sm">
				Jídla obsahující zadané slovo v názvu, která bych si nepřál/a
			</p>
			<div class="rounded-xl mt-4">
				<CategoryList bind:filters={filters.exclude} category="exclude" />
			</div>
		</div>
	</div>
{/if}
