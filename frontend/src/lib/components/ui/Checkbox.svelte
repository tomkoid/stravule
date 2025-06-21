<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Checkbox, Label } from 'bits-ui';

	interface Props {
		checked: boolean | undefined;
		label?: string;
		className?: string;
		onclick?: (e: any) => void;
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
		onclick={onclick!}
		id="terms"
		aria-labelledby="terms-label"
		class={`peer inline-flex items-center justify-center transition-all duration-150 ease-in-out active:scale-98 ${className}`}
	>
		{#snippet children({ checked, indeterminate })}
			{#if checked}
				<Icon width="20" icon="material-symbols:check-rounded" />
			{:else if indeterminate}
				-
			{/if}
		{/snippet}
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
