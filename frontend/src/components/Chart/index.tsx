import React, { useEffect } from 'react'
import Apexcharts from 'apexcharts'

type IProps = {
  style?: React.CSSProperties
  options: any
}

const Chart: React.FC<IProps> = ({ style, options }) => {
  useEffect(() => {
    const chart = new Apexcharts(document.querySelector('#chart'), options)

    chart.render()

    return () => {
      chart.destroy()
    }
  }, [])

  return (
    <div
      id="chart"
      style={style}
    />
  )
}

export default Chart
