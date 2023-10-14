<script lang="ts">
	export let options: EChartOption = {};
	import { chartable } from '$lib/Chart';
	import type { EChartOption } from 'echarts';
</script>

<div id="chart" use:chartable={options} />

<style>
	#chart {
		width: 100%;
		height: 500px;
	}
</style>
