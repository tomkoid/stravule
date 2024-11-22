<script lang="ts">
	import { onMount } from 'svelte';

	import { loadOrders } from '$lib/api/orders';
	import { sendOrders } from '$lib/api/orders';
	import { loadFilters } from '$lib/api/filters';
	import { pickOrders } from '$lib/stores/page.svelte';
	import type { Order } from '$lib/api/orders';
	import type { Filters } from '$lib/api/filters';
	import Checkbox from '$lib/components/ui/Checkbox.svelte';
	import { Collapsible } from 'bits-ui';
	import * as FiltersList from '$lib/components/orders/filters/index';
	import { slide } from 'svelte/transition';
	import { flyAndScale } from '$lib/utils/flyAndScale';
	import Icon from '@iconify/svelte';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';

	let orders: Order[][] | undefined = $state();
	let filters: Filters | undefined = $state();
	let selected: any | boolean[][] | undefined = $state();
	// let pickOrders: any = $state(false);

	let filterTabOpened = $state(false);

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
			pickOrders.value
		);
		filters = await loadFilters();
		renderOrders(orders);
		mountOrdersLoaded = true;
	});

	$effect(() => {
		if (!mountOrdersLoaded) return;
		loadOrders(
			localStorage.getItem('sid')!,
			localStorage.getItem('canteen')!,
			pickOrders.value
		).then((ordersFetched) => {
			orders = ordersFetched;
			renderOrders(ordersFetched);
		});
	});

	const checkIfAfterNow = (timeString: string) => {
		const givenTime = new Date(timeString);
		const now = new Date();
		return givenTime > now;
	};
</script>

<div class="flex flex-col flex-nowrap gap-3">
	<div class="flex gap-2">
		<Checkbox
			className="size-[25px] rounded-md border border-surface1 bg-surface0 data-[state=unchecked]:bg-surface0 data-[state=unchecked]:hover:bg-surface1 data-[state=checked]:hover:bg-mantle"
			bind:checked={pickOrders.value}
			label="Zobrazit vybrané obědy od Stravule"
		/>
	</div>

	<Collapsible.Root bind:open={filterTabOpened}>
		<Collapsible.Trigger class="flex flex-row items-center gap-1">
			<div>
				{#if filterTabOpened}
					<Icon class="text-xl rotate-180" icon="mdi:caret" />
				{:else}
					<Icon class="text-xl" icon="mdi:caret" />
				{/if}
			</div>
			<p>Filtrování</p>
		</Collapsible.Trigger>
		<Collapsible.Content class="mt-2" transition={slide}>
			<FiltersList.Root {filters} />
		</Collapsible.Content>
	</Collapsible.Root>

	{#if pickOrders.value}
		<button
			onclick={async () => {
				await sendOrders(localStorage.getItem('sid')!, localStorage.getItem('canteen')!);
				pickOrders.value = false;
			}}
			class="flex flex-row items-center gap-1 px-2 py-1 w-fit rounded bg-ctp-green text-base transition-all hover:rounded-xl"
		>
			<Icon color="inherit" icon="mdi:check" />
			Potrvdit změny od Stravule a nastavit obědy</button
		>
	{/if}

	{#if orders && selected}
		<div transition:flyAndScale class="flex flex-col flex-nowrap gap-4 mt-3">
			{#each orders as orderTable, orderTableIndex}
				{#if orderTable}
					<div
						class="flex flex-col-reverse gap-2 justify-between bg-surface0 border border-surface1 rounded-xl p-3"
					>
						<div class="flex flex-col">
							{#each orderTable as order, orderIndex}
								<div
									class={`flex flex-nowrap flex-row items-center transition-all gap-3 p-[6px] ${selected[orderTableIndex][orderIndex] ? 'bg-selected-order rounded-xl' : ''}`}
								>
									<!-- <input type="radio" name={orderTableIndex.toString()} bind:group={sel} /> -->
									<!-- <input type="radio" bind:group={sel} /> -->
									{#if order.omezeni.endsWith('E')}
										<Checkbox
											className={`size-[28px] rounded-md bg-selected-order-checkbox data-[state=unchecked]:bg-surface1 data-[state=unchecked]:disabled:bg-surface01 data-[state=checked]:hover:bg-selected-order-checkbox-hover ${checkIfAfterNow(order.casKonec) ? 'data-[state=unchecked]:hover:bg-surface2' : ''}`}
											onclick={() => {
												if (selected) {
													for (let item in selected[orderTableIndex]) {
														if (
															selected[orderTableIndex][item] !=
															selected[orderTableIndex][orderIndex]
														) {
															selected[orderTableIndex][item] = false;
														}
													}
												} else {
													console.log('something really bad happened');
												}
											}}
											bind:checked={selected[orderTableIndex][orderIndex]}
											disabled={!checkIfAfterNow(order.casKonec)}
										/>
									{:else}
										<div class="ml-[28px]"></div>
									{/if}
									<div
										class={`flex flex-wrap flex-row break-all ${(!checkIfAfterNow(order.casKonec) || order.omezeni.endsWith('B')) && !selected[orderTableIndex][orderIndex] ? 'text-subtext1' : ''} ${selected[orderTableIndex][orderIndex] ? 'text-selected-order-text' : ''}`}
									>
										{order.id + 1}. {order.nazev}
									</div>
								</div>
							{/each}
						</div>
						<div>
							<p class="text-text text-xl font-bold">
								{orderTable[0].datum}
							</p>
						</div>
					</div>
				{/if}
			{/each}
		</div>
	{:else}
		<p>Načítání obědů..</p>
	{/if}
</div>

<svelte:head>
	<title>Stravule - Objednávky</title>
	<link rel="preconnect" href={PUBLIC_BACKEND_URL} />
</svelte:head>
