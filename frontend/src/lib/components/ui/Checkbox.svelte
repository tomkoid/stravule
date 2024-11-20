<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Checkbox, Label, type CustomEventHandler } from 'bits-ui';

	interface Props {
		checked: boolean | 'indeterminate' | undefined;
		label?: string;
		className?: string;
		onclick?: (e: CustomEventHandler<MouseEvent, HTMLButtonElement>) => void;
		[props: string]: any;
	}

	let {
		checked = $bindable(false),
		onclick,
		label = '',
		className = '',
		...props
	}: Props = $props();
</script>

<div
	class="flex flex-col msm:flex-row items-center justify-center text-center md:justify-start md:text-left"
>
	<Checkbox.Root
		bind:checked
		{...props}
		on:click={onclick!}
		id="terms"
		aria-labelledby="terms-label"
		class={`peer inline-flex items-center justify-center transition-all duration-150 ease-in-out active:scale-98 ${className}`}
	>
		<Checkbox.Indicator
			let:isChecked
			let:isIndeterminate
			class="inline-flex items-center justify-center text-background"
		>
			{#if isChecked}
				<Icon width="20" icon="material-symbols:check-rounded" />
			{:else if isIndeterminate}
				-
			{/if}
		</Checkbox.Indicator>
	</Checkbox.Root>
	{#if label !== ''}
		<Label.Root
			id="label"
			for="label"
			class="ml-2 leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
		>
			{label}
		</Label.Root>
	{/if}
</div>
