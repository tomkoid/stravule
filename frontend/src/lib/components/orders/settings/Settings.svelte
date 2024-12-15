<script lang="ts">
	interface Props {
		children: () => any;
	}
	let { children }: Props = $props();
</script>

<div
	class="flex flex-col bg-mantle border border-surface0 gap-6 sm:gap-4 p-3 items-center msm:items-start px-4 sm:px-6 rounded-xl"
>
	{@render children()}
</div>
