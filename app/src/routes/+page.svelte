<script lang="ts">
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		Label,
		Input,
		Button,
		Skeleton
	} from 'flowbite-svelte';
	let loading: boolean = false;
	let data: {
		weather: {
			City: string;
			Forecasts: [
				{
					Date: string;
					Temperature: string;
				}
			];
		};
	};
	async function handleWeatherRequest(event: MouseEvent) {
		loading = true;
		const resp = await fetch(
			'/?' +
				new URLSearchParams({
					city: (document.getElementById('city') as HTMLInputElement).value
				}),
			{
				method: 'GET',
				headers: {
					'Content-Type': 'application/json'
				}
			}
		);
		data = await resp.json();
		loading = false;
	}
</script>

<div class="container mx-auto space-y-4">
	<p class="text-xl font-medium dark:text-white">The Weather App</p>
	<Input inline={true} id="city" placeholder="Input city" />
	<Button on:click={handleWeatherRequest}>Check</Button>
	<Table hoverable={true} shadow divClass="h-screen overflow-auto">
		<TableHead>
			<TableHeadCell>Date</TableHeadCell>
			<TableHeadCell>Temperatures</TableHeadCell>
		</TableHead>
		{#if loading}
			<Skeleton size="xxl" />
		{:else}
			<TableBody>
				{#each data?.weather?.Forecasts ?? [] as temp}
					<TableBodyRow>
						<TableBodyCell>{temp.Date}</TableBodyCell>
						<TableBodyCell>{temp.Temperature}</TableBodyCell>
					</TableBodyRow>
				{/each}
			</TableBody>
		{/if}
	</Table>
</div>
