<script lang="ts">
	import { addFilter, removeFilter } from '$lib/api/filters';
	import { errors } from '$lib/stores/errors.svelte';
	import { pickOrders } from '$lib/stores/page.svelte';
	import { flyAndScale } from '$lib/utils/flyAndScale';
	import Icon from '@iconify/svelte';
	import { Dialog } from 'bits-ui';
	import { fade } from 'svelte/transition';

	interface FiltersProps {
		filters?: Filter[];
		category: string;
	}

	interface Filter {
		value: string;
	}

	let { filters = $bindable(), category }: FiltersProps = $props();

	let newFilter: string = $state('');
	let confirmButton: HTMLButtonElement;

	let transactionHappening = false;
	let dialogShown: boolean[] = $state(new Array(filters!.length).fill(false));
	let editFilter: Filter = $state({ value: '' }); // what filter to remove state
</script>

{#if filters}
	{#each filters as filter, i}
		<div class="flex flex-row items-center gap-1">
			<p class="text">{filter.value}</p>

			<Dialog.Root bind:open={dialogShown[i]}>
				<Dialog.Trigger onclick={() => (editFilter = filter)}>
					<Icon icon="material-symbols:delete" />
				</Dialog.Trigger>
				<Dialog.Portal>
					<Dialog.Overlay
						transition={fade}
						transitionConfig={{ duration: 150 }}
						class={`fixed inset-0 z-50 bg-black/70`}
					/>
					<Dialog.Content
						transition={flyAndScale}
						class="fixed left-[50%] top-[50%] z-50 w-full max-w-[94%] translate-x-[-50%] translate-y-[-50%] rounded-card-lg rounded-xl border border-mantle bg-base p-5 shadow-popover outline-none sm:max-w-[490px] md:w-full"
					>
						<Dialog.Title
							class="flex w-full items-center justify-center text-lg font-semibold tracking-tight"
							>Jste si jisti?</Dialog.Title
						>
						<Dialog.Description class="text-sm text-subtext1 mb-4">
							Tímto potvrzujete, že chcete smazat filtr "{editFilter.value}".
						</Dialog.Description>
						<div class="flex flex-row items-center justify-center gap-2 h-12">
							<button
								class="bg-surface0 hover:bg-surface1 transition-all rounded hover:rounded-lg flex flex-row items-center justify-center gap-1 px-2 h-full w-full"
								onclick={() => (dialogShown[i] = false)}
							>
								Ne
							</button>
							<button
								class="flex flex-row items-center justify-center text-center gap-1 px-2 h-full w-full rounded hover:rounded-lg text-base bg-ctp-red transition disabled:bg-crust hover:bg-ctp-maroon"
								onclick={async () => {
									if (!filters) console.error('filters is undefined while removing a filter');
									filters = filters!.filter((f) => f != editFilter);

									if (!localStorage) console.error('localStorage is undefined');
									if (!localStorage.getItem('sid'))
										console.error('sid not found in localStorage while removing filter');
									if (!localStorage.getItem('canteen'))
										console.error('canteen not found in localStorage while removing filter');

									dialogShown[i] = false;

									await removeFilter(editFilter.value);

									if (pickOrders.value == true) {
										pickOrders.value = false;
										pickOrders.value = true;
									}
								}}
							>
								Ano
							</button>
						</div>
						<Dialog.Close
							class="absolute right-5 top-5 rounded-md focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-foreground focus-visible:ring-offset-2 focus-visible:ring-offset-background active:scale-98"
						>
							<div>
								<Icon icon="material-symbols:close" />
								<span class="sr-only">Close</span>
							</div>
						</Dialog.Close>
					</Dialog.Content>
				</Dialog.Portal>
			</Dialog.Root>
		</div>
	{/each}
{/if}

<div class="flex flex-row items-center gap-1 bg-surface0 border border-surface0 mt-2 rounded w-fit">
	<input
		placeholder="brambory, ..."
		class="border-none bg-surface0 pl-1 w-28 h-6"
		type="text"
		onkeypresscapture={(e) => {
			if (e.key == 'Enter') {
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
				filters = [...filters!, { value: newFilter }];
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
