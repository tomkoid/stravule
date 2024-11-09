<script lang="ts">
	import { onMount } from 'svelte';

	import { loadOrders } from '$lib/api/orders';
	import { loadFilters } from '$lib/api/filters';
	import type { Order } from '$lib/api/orders';
	import type { Filters } from '$lib/api/filters';
	import Checkbox from '$lib/components/ui/Checkbox.svelte';
	import { Collapsible } from 'bits-ui';
	import * as FiltersList from '$lib/components/orders/filters/index';
	import { fly, slide } from 'svelte/transition';
	import { quadInOut } from 'svelte/easing';

	let orders: Order[][] | undefined = $state();
	let filters: Filters | undefined = $state();
	let selected: any | boolean[][] | undefined = $state();
	let pickOrders: any = $state(false);

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

	let mountOrdersLoaded = $state(false);
	onMount(async () => {
		orders = await loadOrders(
			localStorage.getItem('sid')!,
			localStorage.getItem('canteen')!,
			pickOrders
		);
		filters = await loadFilters();
		renderOrders(orders);
		mountOrdersLoaded = true;
	});

	$effect(() => {
		if (!mountOrdersLoaded) return;
		loadOrders(localStorage.getItem('sid')!, localStorage.getItem('canteen')!, pickOrders).then(
			(ordersFetched) => {
				orders = ordersFetched;
				renderOrders(ordersFetched);
			}
		);
	});
</script>

<div class="flex flex-col flex-nowrap gap-3">
	<Collapsible.Root>
		<Collapsible.Trigger>> Filtry</Collapsible.Trigger>
		<Collapsible.Content class="mt-2" transition={slide}>
			<FiltersList.Root {filters} />
		</Collapsible.Content>
	</Collapsible.Root>

	<div class="flex gap-2">
		<Checkbox
			className="size-[25px] rounded-md shadow shadow-ctp-mantle bg-ctp-surface0 data-[state=unchecked]:bg-ctp-surface0 data-[state=unchecked]:hover:bg-ctp-surface1 data-[state=checked]:hover:bg-ctp-mantle"
			bind:checked={pickOrders}
			label="Zobrazit vybrané obědy od Stravule"
		/>
	</div>
	{#if orders && selected}
		<div
			in:fly={{ y: 200, delay: 0, easing: quadInOut }}
			out:fly={{ x: 200, duration: 300 }}
			class="flex flex-col flex-nowrap gap-3"
		>
			{#each orders as orderTable, orderTableIndex}
				{#if orderTable}
					<div class="bg-ctp-surface0 shadow shadow-ctp-surface0 rounded-xl p-3">
						{#each orderTable as order, orderIndex}
							<div class="flex flex-nowrap flex-row gap-2">
								<!-- <input type="radio" name={orderTableIndex.toString()} bind:group={sel} /> -->
								<!-- <input type="radio" bind:group={sel} /> -->
								{#if order.omezeni.endsWith('E')}
									<Checkbox
										className="size-[20px] rounded-md shadow shadow-ctp-surface0 bg-ctp-surface1 data-[state=unchecked]:bg-ctp-surface1 data-[state=unchecked]:hover:bg-ctp-surface2 data-[state=checked]:hover:bg-ctp-mantle"
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
									<div class="ml-[20px] msm:ml-[28px] md:ml-[32px]"></div>
								{/if}
								<div class="flex flex-wrap flex-row break-all">
									{order.id + 1}. {order.nazev}
								</div>
							</div>
						{/each}
					</div>
				{/if}
			{/each}
		</div>
	{:else}
		<p>Načítání obědů..</p>
	{/if}
	<button
		class="bg-ctp-blue text-ctp-base border rounded-xl p-2 mb-2"
		onclick={() => {
			console.log($state.snapshot(selected));
		}}>print order list object to console</button
	>
</div>

<svelte:head>
	<title>Stravule - Objednávky</title>
</svelte:head>
