<script lang="ts">
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
			}}>X</button
		>
		<p>{filter}</p>
	</div>
{/each}
