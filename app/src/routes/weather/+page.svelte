<script lang="ts">
	import { options } from '$lib/stores/WeatherStore';
	import type { EChartOption } from 'echarts';
	import { Breadcrumb, BreadcrumbItem, Button, Input, Skeleton } from 'flowbite-svelte';
	import Chart from '../../components/chart/Chart.svelte';
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
		$options = {
			xAxis: {
				type: 'category',
				data: data?.weather?.Forecasts?.map((f) => f.Date),
				name: 'Date'
			},
			yAxis: {
				type: 'value',
				name: 'Temperature',
				axisLabel: {
					formatter: '{value} °C'
				}
			},
			series: [
				{
					data: data?.weather?.Forecasts?.map((f) => f.Temperature.replace('°C', '')),
					type: 'line'
				}
			]
		} as EChartOption;
		loading = false;
	}
</script>

<div class="container flex flex-col space-y-4">
	<Breadcrumb aria-label="Solid background breadcrumb example" solid>
		<BreadcrumbItem href="/" home>Home</BreadcrumbItem>
		<BreadcrumbItem>Weather</BreadcrumbItem>
	</Breadcrumb>
	<div class="flex flex-row space-x-4">
		<Input inline={true} id="city" placeholder="Input city" />
		<Button on:click={handleWeatherRequest}>Check</Button>
	</div>
	{#if loading}
		<Skeleton />
	{:else}
		<Chart options={$options} />
	{/if}
</div>
