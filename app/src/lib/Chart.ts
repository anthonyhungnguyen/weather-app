
import * as charts from 'echarts'

import type { EChartOption } from 'echarts'

export function chartable(node: HTMLElement, options: EChartOption) {
  const chart = charts.init(node)
  chart.setOption(options)

  return {
    destroy() {
      chart.dispose()
    },
    update(newOptions: EChartOption) {
      chart.setOption(newOptions)
    }
  }
} 