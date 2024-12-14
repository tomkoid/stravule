<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Calendar } from 'bits-ui';
	import { CalendarDate, type DateValue } from '@internationalized/date';
	import { onMount } from 'svelte';
	import { listNoOrderDays, setNoOrderDay } from '$lib/api/orders';
	import { pickOrders } from '$lib/stores/page.svelte';

	let originalCalendarValue: DateValue[] = $state([]);
	let calendarValue: DateValue[] = $state([]);
	let loaded = $state(false);

	onMount(async () => {
		loaded = false;
		const dayList = await listNoOrderDays();

		for (let day of dayList) {
			const daySplit = day.split('-');
			if (daySplit.length != 3) {
				console.error('day split doesnt have the length of 3');
			}

			const year = parseInt(daySplit[0]);
			const month = parseInt(daySplit[1]);
			const dayDay = parseInt(daySplit[2]);

			const newDate = new CalendarDate(year, month, dayDay);
			calendarValue.push(newDate);
			originalCalendarValue.push(newDate);
		}

		loaded = true;
	});

	async function toggleDate(date: DateValue) {
		if (originalCalendarValue && originalCalendarValue.toString().includes(date.toString())) {
			console.log(`calendar: removing ${date.toString()}`);
			originalCalendarValue = calendarValue;
			await setNoOrderDay(false, date.toString());
		} else {
			console.log(`calendar: adding ${date.toString()}`);
			originalCalendarValue = calendarValue;
			await setNoOrderDay(true, date.toString());
		}

		originalCalendarValue = calendarValue;
		changes = true;

		if (pickOrders.value == true) {
			pickOrders.value = false;
			pickOrders.value = true;
		}
	}

	let changes = $state(false);
</script>

{#if loaded}
	<div>
		<div class="flex flex-col">
			<div class="flex flex-row items-center gap-2">
				<Icon icon="mdi:calendar" class="text-2xl" />
				<p class="text-lg font-bold">Kdy odhlásit obědy?</p>
			</div>
			<p class="text-sm text-subtext0">Vyber si dny, kdy automaticky odhlásit oběd.</p>
		</div>
		<Calendar.Root
			class="mt-4 rounded-[15px] border border-surface1 bg-mantle/30 p-[22px] shadow-sm shadow-surface0 max-w-fit"
			let:months
			let:weekdays
			weekdayFormat="short"
			weekStartsOn={1}
			fixedWeeks={true}
			multiple={true}
			bind:value={calendarValue}
		>
			<Calendar.Header class="flex items-center justify-between">
				<Calendar.PrevButton
					class="inline-flex size-10 items-center justify-center rounded-9px bg-background-alt hover:bg-muted active:scale-98 active:transition-all"
				>
					<Icon icon="mdi:chevron-left" class="size-6" />
				</Calendar.PrevButton>
				<Calendar.Heading class="text-[15px] font-medium" />
				<Calendar.NextButton
					class="inline-flex size-10 items-center justify-center rounded-9px bg-background-alt hover:bg-muted active:scale-98 active:transition-all"
				>
					<Icon icon="mdi:chevron-right" class="size-6" />
				</Calendar.NextButton>
			</Calendar.Header>
			<div class="flex flex-col space-y-4 pt-4 sm:flex-row sm:space-x-4 sm:space-y-0">
				{#each months as month, i (i)}
					<Calendar.Grid class="w-full border-collapse select-none space-y-1">
						<Calendar.GridHead>
							<Calendar.GridRow class="mb-3 flex w-full justify-between">
								{#each weekdays as day}
									<Calendar.HeadCell class="w-10 rounded-md text-xs !font-normal text-subtext0">
										<div class="text-inherit">{day.slice(0, 2)}</div>
									</Calendar.HeadCell>
								{/each}
							</Calendar.GridRow>
						</Calendar.GridHead>
						<Calendar.GridBody>
							{#each month.weeks as weekDates}
								<Calendar.GridRow class="flex w-full">
									{#each weekDates as date}
										<Calendar.Cell {date} class="relative size-10 !p-0 text-center text-sm">
											<Calendar.Day
												{date}
												month={month.value}
												class="group transition-all relative inline-flex size-8 items-center justify-center whitespace-nowrap rounded-full border border-transparent bg-transparent p-0 text-sm font-normal text-foreground hover:border-surface2 data-[disabled]:pointer-events-none data-[outside-month]:pointer-events-none data-[selected]:bg-text data-[selected]:font-medium data-[disabled]:text-text/30 data-[selected]:text-base data-[unavailable]:text-text/20 data-[unavailable]:line-through"
												onclick={async () => await toggleDate(date)}
											>
												<div
													class="absolute top-[5px] hidden size-1 rounded-full bg-foreground group-data-[today]:block group-data-[selected]:bg-background"
												></div>
												{date.day}
											</Calendar.Day>
										</Calendar.Cell>
									{/each}
								</Calendar.GridRow>
							{/each}
						</Calendar.GridBody>
					</Calendar.Grid>
				{/each}
			</div>
		</Calendar.Root>
	</div>
{:else}
	<p>Načítání..</p>
{/if}
