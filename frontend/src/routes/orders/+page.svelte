<script lang="ts">
	import { onMount } from 'svelte';
	import { loadOrders } from '$lib/api/orders';
	import { loadFilters } from '$lib/api/filters';
	import type { Order } from '$lib/api/orders';
	import type { Filters } from '$lib/api/filters';

	let orders: Order[][] | undefined = $state();
	let filters: Filters | undefined = $state();
	let selected: boolean[][] | undefined = $state();
	let pickOrders: boolean = $state(false);

	async function renderOrders(orders: Order[][]) {
		selected = new Array(orders.length);
		for (let i = 0; i < orders.length; i++) {
			selected[i] = new Array(orders[i].length).fill(false);
			for (let j = 0; j < orders[i].length; j++) {
				console.log(i, j);
				selected[i][j] = orders[i][j].pocet == 1;
			}
		}
	}

	onMount(async () => {
		orders = await loadOrders(
			localStorage.getItem('sid')!,
			localStorage.getItem('canteen')!,
			pickOrders
		);
		filters = await loadFilters(localStorage.getItem('sid')!, localStorage.getItem('canteen')!)!;
		renderOrders(orders);
	});

	$effect(() => {
		loadOrders(localStorage.getItem('sid')!, localStorage.getItem('canteen')!, pickOrders).then(
			(ordersFetched) => {
				orders = ordersFetched;
				renderOrders(ordersFetched);
			}
		);
	});
</script>

<div class="flex flex-col flex-nowrap gap-3">
	<div class="flex flex-wrap flex-row md:justify-between gap-2">
		<div class="flex flex-wrap gap-2">
			{#if filters}
				{#each filters.include as filter}
					<div class="bg-ctp-green text-ctp-base shadow shadow-ctp-green rounded-xl p-2">
						{filter}
					</div>
				{/each}
			{/if}
		</div>
		<div class="flex flex-wrap gap-2">
			{#if filters}
				{#each filters.exclude as filter}
					<div class="bg-ctp-red text-ctp-base shadow shadow-ctp-red rounded-xl p-2">
						{filter}
					</div>
				{/each}
			{/if}
		</div>
	</div>
	<div class="flex gap-2">
		<input class="outline-none" type="checkbox" bind:checked={pickOrders} /> Show picked orders by Stravule
	</div>
	{#if orders && selected}
		{#each orders as orderTable, orderTableIndex}
			{#if orderTable}
				<div class="bg-ctp-surface0 shadow shadow-ctp-surface0 rounded-xl p-3">
					{#each orderTable as order, orderIndex}
						<div class="flex flex-nowrap flex-row gap-2">
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
								<div class="ml-[13px]"></div>
							{/if}
							<div class="flex flex-wrap flex-row break-all">
								{order.id + 1}. {order.nazev}
							</div>
						</div>
					{/each}
				</div>
			{/if}
		{/each}
	{/if}
	<button
		class="bg-ctp-blue text-ctp-base border rounded-xl p-2 mb-2"
		onclick={() => {
			console.log(selected);
		}}>print order list object to console</button
	>
</div>
