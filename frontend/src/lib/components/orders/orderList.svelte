<script lang="ts">
	import type { Order } from '$lib/api/orders';
	import { flyAndScale } from '$lib/utils/flyAndScale';
	import { formatDate } from '$lib/utils/dates';
	import Checkbox from '$lib/components/ui/Checkbox.svelte';

	interface OrderListProps {
		orders: Order[][];
		selected: boolean[][];
	}

	let { orders = $bindable(), selected = $bindable() }: OrderListProps = $props();

	const checkIfAfterNow = (timeString: string) => {
		const givenTime = new Date(timeString);
		const now = new Date();
		return givenTime > now;
	};
</script>

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
									className={`size-[28px] rounded-md bg-selected-order-checkbox data-[state=unchecked]:bg-surface1 data-[state=unchecked]:disabled:bg-surface01 data-[state=checked]:hover:bg-selected-order-checkbox-hover ${checkIfAfterNow(order.casKonec) ? 'data-[state=unchecked]:hover:bg-surface2' : 'data-[state=checked]:hover:bg-selected-order-checkbox'}`}
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
									disabled={!checkIfAfterNow(order.casKonec)}
								/>
							{:else}
								<div class="ml-[28px]"></div>
							{/if}
							<div
								class={`flex flex-wrap flex-row break ${(!checkIfAfterNow(order.casKonec) || order.omezeni.endsWith('B')) && !selected[orderTableIndex][orderIndex] ? 'text-subtext1' : ''} ${selected[orderTableIndex][orderIndex] ? 'text-selected-order-text font-medium' : ''}`}
							>
								{order.id + 1}. {order.nazev}
							</div>
						</div>
					{/each}
				</div>
				<div>
					<p class="text-text text-xl font-bold">
						{formatDate(orderTable[0].datum)}
					</p>
				</div>
			</div>
		{/if}
	{/each}
</div>
