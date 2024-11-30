<script lang="ts">
	import { addFilter, removeFilter, updateFilterWeight } from '$lib/api/filters';
	import { errors } from '$lib/stores/errors.svelte';
	import { pickOrders } from '$lib/stores/page.svelte';
	import { flyAndScale } from '$lib/utils/flyAndScale';
	import Icon from '@iconify/svelte';
	import { Slider, DropdownMenu } from 'bits-ui';
	import { onMount } from 'svelte';

	interface FiltersProps {
		filters?: Filter[];
		category: string;
	}

	interface Filter {
		value: string;
		weight: number;
		category: string;
		created_at: string;
	}

	let { filters = $bindable(), category }: FiltersProps = $props();

	let newFilter: string = $state('');
	let confirmButton: HTMLButtonElement | undefined = $state();

	let transactionHappening = false;
	let dropdownShown: boolean[] = $state(new Array(filters!.length).fill(false));
	let editFilter: Filter = $state({ value: '', weight: 0, category: '', created_at: '' }); // what filter to remove state
	let editWeightConfirm = $state(new Array(filters!.length).fill(false));

	let sliderVal = $state(new Array(filters!.length));

	onMount(() => {
		for (let i = 0; i < filters!.length; i++) {
			sliderVal[i] = new Array(1).fill(filters![i].weight);
		}
	});
</script>

{#if filters}
	{#each filters as filter, i}
		<div class="flex flex-row items-center gap-4">
			<p class="text">{filter.value}</p>

			<DropdownMenu.Root
				preventScroll={false}
				closeOnItemClick={false}
				bind:open={dropdownShown[i]}
			>
				<DropdownMenu.Trigger onclick={() => (editFilter = filter)}>
					<button
						class={`p-1 bg-base rounded-xl flex flex-row items-center text-subtext2 gap-1 transition hover:bg-surface0 ${dropdownShown[i] ? 'bg-surface1 hover:bg-surface1' : ''}`}
					>
						<Icon width={16} icon="tabler:dots" />
					</button>
				</DropdownMenu.Trigger>
				<DropdownMenu.Content
					class="w-full max-w-[229px] rounded-xl border border-surface1 bg-mantle px-1 py-1.5 shadow shadow-base"
					transition={flyAndScale}
					sideOffset={8}
				>
					<div class="flex flex-col text-center justify-center">
						<div class="flex flex-row justify-center items-center gap-2">
							<Icon width={16} icon="material-symbols:scale" />
							<h1 class="font-bold text-lg">Váha filtru</h1>
						</div>
						<p class="text-sm text-subtext0">Mění váhu filtru {editFilter.value}</p>
					</div>
					<!-- <DropdownMenu.Item -->
					<!-- 	class={`${dropdownItemClass} flex self-center justify-center text-center`} -->
					<!-- > -->
					<Slider.Root
						class="relative flex touch-none select-none items-center max-w-[12rem] m-5"
						min={1}
						max={category == 'include' ? 10 : 20}
						step={1}
						bind:value={sliderVal[i]}
						onValueChange={(value) => {
							editWeightConfirm[i] = true;
						}}
						let:ticks
						let:thumbs
					>
						<span class="relative h-2 w-full grow overflow-hidden rounded-full bg-surface1">
							<Slider.Range class="absolute h-full bg-text" />
						</span>

						{#each thumbs as thumb}
							<Slider.Thumb
								class="block size-[25px] cursor-pointer rounded-full border border-base bg-white shadow  transition-colors hover:border-dark-40 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-foreground focus-visible:ring-offset-2 active:scale-98 disabled:pointer-events-none disabled:opacity-50 dark:bg-foreground dark:shadow-card"
								{thumb}
							/>
						{/each}

						{#each ticks as tick}
							<Slider.Tick {tick} />
						{/each}
					</Slider.Root>
					{#if category == 'include'}
						<div class="flex flex-row justify-between w-[12rem] m-5">
							{#each { length: 10 } as _, i}
								<p class="text-subtext0 text-sm mt-[-4px] w-[19px]">{i + 1}</p>
							{/each}
						</div>
					{:else}
						<p class="text-subtext0 text-sm mt-[-13px] mb-2 text-center">Váha: {sliderVal[i][0]}</p>
					{/if}
					<div>
						<!-- </DropdownMenu.Item> -->
						<button
							class={`flex flex-row gap-1 justify-center text-center items-center w-full bg-blue-300 hover:bg-blue-200 text-base mb-2 transition-all rounded hover:rounded-xl ${editWeightConfirm[i] ? '' : 'disabled hidden'}`}
							onclick={async () => {
								editWeightConfirm[i] = false;

								// update weight
								await updateFilterWeight(editFilter.value, sliderVal[i][0]);

								if (pickOrders.value == true) {
									pickOrders.value = false;
									pickOrders.value = true;
								}
							}}
						>
							<Icon class="text-inherit" width={20} icon="material-symbols:check" />
							<p class="text-inherit">Potvrdit váhu</p>
						</button>
						<!-- <button -->
						<!-- 	onclick={() => { -->
						<!-- 		localStorage.removeItem('sid'); -->
						<!-- 		window.location.href = '/'; -->
						<!-- 	}} -->
						<!-- 	class="flex gap-2 h-10 items-center w-full" -->
						<!-- > -->
						<!-- 	<Icon class="text-xl" icon="mdi:logout" /> -->
						<!-- 	Odhlásit se -->
						<!-- </button> -->
						<button
							class="flex flex-row items-center justify-center text-center gap-1 h-full w-full rounded hover:rounded-xl text-base bg-ctp-red transition-all disabled:bg-crust hover:bg-ctp-maroon"
							onclick={async () => {
								if (!filters) console.error('filters is undefined while removing a filter');
								filters = filters!.filter((f) => f != editFilter);

								if (!localStorage) console.error('localStorage is undefined');
								if (!localStorage.getItem('sid'))
									console.error('sid not found in localStorage while removing filter');
								if (!localStorage.getItem('canteen'))
									console.error('canteen not found in localStorage while removing filter');

								// dialogShown[i] = false;

								await removeFilter(editFilter.value);

								if (pickOrders.value == true) {
									pickOrders.value = false;
									pickOrders.value = true;
								}

								dropdownShown[i] = false;
							}}
						>
							<Icon class="text-base" icon="material-symbols:delete" />
							Smazat filtr
						</button>
					</div></DropdownMenu.Content
				>
			</DropdownMenu.Root>
		</div>
	{/each}
{/if}

{#if filters}
	<div
		class="flex flex-row items-center gap-1 bg-surface0 border border-surface0 mt-2 rounded w-fit"
	>
		<input
			placeholder="brambory, ..."
			class="border-none bg-surface0 pl-1 w-28 h-6"
			type="text"
			onkeypresscapture={(e) => {
				if (e.key == 'Enter' && confirmButton) {
					e.preventDefault();
					confirmButton.click();
				}
			}}
			bind:value={newFilter}
		/>
		<button
			class="flex flex-row items-center gap-1 px-2 h-6 bg-base transition disabled:bg-crust hover:bg-crust"
			bind:this={confirmButton}
			disabled={!newFilter}
			onclick={async () => {
				if (transactionHappening) return;

				transactionHappening = true;
				try {
					await addFilter(newFilter, category);

					// add filter to the list
					filters!.push({ value: newFilter, weight: 1, category: category, created_at: '' });
					dropdownShown.push(false);
					sliderVal.push(category == 'include' ? [5] : [10]);

					newFilter = '';

					if (pickOrders.value == true) {
						pickOrders.value = false;
						pickOrders.value = true;
					}
				} catch (e: any) {
					errors.add(e);
					transactionHappening = false;
				}
				transactionHappening = false;
			}}
		>
			<Icon icon="material-symbols:add" />
			Přidat
		</button>
	</div>
{/if}
