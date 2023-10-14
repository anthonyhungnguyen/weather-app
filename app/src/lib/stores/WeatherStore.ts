import type { EChartOption } from 'echarts'
import { writable } from 'svelte/store'

const options = writable({} as EChartOption)

export {
  options
}
