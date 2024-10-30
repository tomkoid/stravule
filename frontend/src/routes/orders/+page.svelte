<script lang="ts">
	import { onMount } from 'svelte';
	import { loadOrders } from '$lib/api/orders';
	import type { Order } from '$lib/api/orders';

	let orders: Order[][] | undefined = $state();
	let selected: boolean[][] | undefined = $state();

	onMount(async () => {
		orders = await loadOrders(localStorage.getItem('sid')!, localStorage.getItem('canteen')!);

		selected = new Array(orders.length);
		for (let i = 0; i < orders.length; i++) {
			selected[i] = new Array(orders[i].length).fill(false);
			for (let j = 0; j < orders[i].length; j++) {
				console.log(i, j);
				selected[i][j] = orders[i][j].pocet == 1;
			}
		}
	});
</script>

<div class="flex flex-col flex-wrap gap-3">
	<button
		class="bg-blue-200 border rounded-xl p-2 mb-2"
		onclick={() => {
			console.log(selected);
		}}>print order list object to console</button
	>
	{#if orders && selected}
		{#each orders as orderTable, orderTableIndex}
			{#if orderTable}
				<div class="bg-lime-200 rounded-xl border p-2">
					{#each orderTable as order, orderIndex}
						<div class="flex flex-wrap flex-row gap-2">
							<!-- <input type="radio" name={orderTableIndex.toString()} bind:group={sel} /> -->
							<!-- <input type="radio" bind:group={sel} /> -->
							{#if order.omezeni.endsWith('E')}
								<input
									type="checkbox"
									onclick={() => {
										if (selected) {
											for (let item in selected[orderTableIndex]) {
												if (
													selected[orderTableIndex][item] != selected[orderTableIndex][orderIndex]
												) {
													selected[orderTableIndex][item] = false;
												}
											}
										} else {
											console.log('something really bad happened');
										}
									}}
									bind:checked={selected[orderTableIndex][orderIndex]}
								/>
							{:else}
								<div class="ml-3"></div>
							{/if}
							<div class="flex flex-row">
								{order.id + 1}. {order.nazev}
							</div>
						</div>
					{/each}
				</div>
			{/if}
		{/each}
	{/if}
</div>
