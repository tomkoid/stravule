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
	import * as Settings from '$lib/components/orders/settings/index';
	import OrderList from '$lib/components/orders/orderList.svelte';
	import { slide } from 'svelte/transition';
	import Icon from '@iconify/svelte';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';

	let orders: Order[][] | undefined = $state();
	let filters: Filters | undefined = $state();
	let selected: any | boolean[][] | undefined = $state();

	let filterTabOpened = $state(false);

	async function renderOrders(orders: Order[][]) {
		selected = new Array(orders.length);
		for (let i = 0; i < orders.length; i++) {
			selected[i] = new Array(orders[i].length).fill(false);
			for (let j = 0; j < orders[i].length; j++) {
				selected[i][j] = orders[i][j].pocet == 1;
			}
		}
	}

	let mountOrdersLoaded = $state(false);
	let originalOrders = $state<Order[][]>([]);
	onMount(async () => {
		orders = await loadOrders(
			localStorage.getItem('sid')!,
			localStorage.getItem('canteen')!,
			pickOrders.value
		);
		originalOrders = orders;
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

	function getDifference(ordersA: Order[][], ordersB: Order[][]): number {
		if (!ordersA || !ordersB) {
			return -1;
		}

		let difference = 0;
		for (let i = 0; i < ordersA.length; i++) {
			for (let j = 0; j < ordersA[i].length; j++) {
				if (ordersA[i][j].pocet != ordersB[i][j].pocet) {
					difference++;
				}
			}
		}
		return difference;
	}

	let areChanges = $derived(
		pickOrders.value && orders && getDifference(orders, originalOrders) > 0
	);
	let areNoChanges = $derived(
		pickOrders.value && orders && getDifference(orders, originalOrders) == 0
	);
</script>

<div class="flex flex-col flex-nowrap gap-3">
	<div class="flex flex-wrap gap-2 flex-row justify-between">
		<h1 class="text-2xl font-extrabold">Objednávky</h1>
		{#if areChanges}
			<div>
				<button
					onclick={async () => {
						await sendOrders(localStorage.getItem('sid')!, localStorage.getItem('canteen')!);
						originalOrders = orders!;
						pickOrders.value = false;
					}}
					class="flex flex-row items-center w-fit gap-1 px-2 py-1 rounded-xl bg-blue-300 text-base transition-all hover:rounded-2xl"
				>
					<Icon class="min-w-[16px] min-h-[16px]" color="inherit" icon="mdi:check" />
					Potrvdit změny od Stravule a nastavit objednávky (změn: {getDifference(
						orders!,
						originalOrders
					)})</button
				>
			</div>
		{:else if areNoChanges}
			<div class="flex flex-row items-center gap-1">
				<p class="text-subtext0">--- Žádné změny od Stravule ---</p>
			</div>
		{/if}
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
			<p>Nastavení objednávek</p>
		</Collapsible.Trigger>
		<Collapsible.Content class="mt-2" transition={slide}>
			<Settings.Root>
				<div class="flex gap-2">
					<Checkbox
						className="size-[25px] rounded-md border border-surface1 bg-surface0 data-[state=unchecked]:bg-surface0 data-[state=unchecked]:hover:bg-surface1 data-[state=checked]:hover:bg-mantle"
						bind:checked={pickOrders.value}
						label="Zobrazit vybrané objednávky od Stravule"
					/>
				</div>
				<FiltersList.Root {filters} />
				<div class="flex flex-row flex-wrap gap-6 w-full">
					<div class="flex-1 mt-6">
						<Settings.DayExceptions />
					</div>

					<div class="flex-1 mt-6">
						<Settings.Calendar />
					</div>
				</div>
			</Settings.Root>
		</Collapsible.Content>
	</Collapsible.Root>

	{#if orders && selected}
		<OrderList {orders} {selected} />
	{:else}
		<p>Načítání objednávek..</p>
	{/if}
</div>

<svelte:head>
	<title>Stravule - Objednávky</title>
	<link rel="preconnect" href={PUBLIC_BACKEND_URL} />
</svelte:head>
