<div class="app">
	<slot />
</div>

<style>
	* {
		box-sizing: border-box;
	}

	.app {
		width: 100%;
		display: flex;
		flex-direction: column;
		min-height: 100vh;
	}
</style>
