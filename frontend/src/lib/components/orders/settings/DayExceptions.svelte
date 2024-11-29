<script lang="ts">
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';
	import { listOrderDayExceptions, setOrderDayException } from '$lib/api/orders';
	import { pickOrders } from '$lib/stores/page.svelte';

	const days = ['Po', 'Út', 'St', 'Čt', 'Pá', 'So', 'Ne'];
	const daysValue = ['1', '2', '3', '4', '5', '6', '0'];
	let daysChecked: boolean[] = $state([]);

	let daysFetched = $state(false);

	onMount(async () => {
		const dayList = await listOrderDayExceptions();

		daysChecked = [false, false, false, false, false, false, false];

		for (let i = 0; i < dayList.length; i++) {
			daysChecked[daysValue.indexOf(dayList[i].toString())] = true;
		}

		daysFetched = true;
	});
</script>

<div class="flex flex-col gap-2">
	<div>
		<div class="flex gap-2 items-center">
			<Icon width={24} icon="lets-icons:date-fill" />
			<h1 class="text-lg font-bold">Vybrat dny automatického odhlášení</h1>
		</div>

		<p class="text-subtext0 text-sm">Dny, které se automaticky odhlásí.</p>
	</div>
	{#if daysChecked.length == days.length}
		<div class="flex flex-row gap-2">
			{#each days as day, i}
				<div class="flex flex-row items-center gap-4">
					<button
						onclick={async () => {
							daysChecked[i] = !daysChecked[i];
							await setOrderDayException(daysChecked[i], Number(daysValue[i]));

							if (pickOrders.value == true) {
								pickOrders.value = false;
								pickOrders.value = true;
							}
						}}
						class={`p-1 w-8 h-8 text-center rounded-full transition ${daysChecked[i] ? 'bg-white hover:bg-text text-base' : 'bg-surface0 hover:bg-surface1 text-text'}`}
					>
						{day}
					</button>
				</div>
			{/each}
		</div>
	{:else}
		<p>Načítání..</p>
	{/if}
</div>
