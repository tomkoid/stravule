<script lang="ts">
	import { errors } from '$lib/stores/errors.svelte';

	$effect(() => {
		if (errors.length > 5) {
			errors.clean();
		}
	});
</script>

<!---->
<!-- <div class="errors"> -->
<!-- 	<button -->
<!-- 		onclick={() => { -->
<!-- 			errors.removeOne(); -->
<!-- 		}} -->
<!-- 		class="error m-auto mb-2 pl-4 pr-4 p-2 rounded-full w-fit bg-ctp-red text-ctp-base" -->
<!-- 	> -->
<!-- 		<p>test</p> -->
<!-- 	</button> -->
<!-- </div> -->

<div class="errors">
	{#each errors.errors as error}
		<button
			onclick={() => {
				errors.removeOne();
			}}
			class="error mb-2 pl-4 pr-4 p-2 rounded-md w-fit bg-ctp-red"
		>
			<p class="text-ctp-base">{error}</p>
		</button>
	{/each}
</div>

<style lang="scss">
	// create an animation for error bubbles that are on the top of the screen that the first they are not seen by the user and then after 0.5s slides
	// down to the top of the screen
	@keyframes slideDown {
		0% {
			transform: translateY(200%);
		}
		100% {
			transform: translateY(0%);
		}
	}

	.errors {
		position: fixed;
		bottom: 15px;
		right: 70px;
		max-width: 300px;
		margin-left: -35px; /* Negative half of width. */
		z-index: 1000;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.error {
		animation: slideDown 0.3s ease-in-out;
		// end animation
		animation-direction: alternate;
	}
</style>
