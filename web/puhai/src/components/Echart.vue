<template>
  <div id="myChart" :style="{width: '100%', height: '300px'}"></div>
</template>

<script>
import * as echarts from 'echarts';

export default {
  name: "Echart",
  data(){
    return {
    }
  },
  methods: {
    drawLine(){
      let myChart = echarts.init(document.getElementById("myChart"));
      let option;
      option = {
          title: {
            "text": "电力负荷",
          },
          tooltip: {
            trigger: 'axis',
            axisPointer: {
              type: 'cross'
            }
          },
          toolbox: {
            show: true,
            feature: {
              saveAsImage: {}
            }
          },
          xAxis: [{
            type: 'category',
            boundaryGap: false,
            name:'日期',
            data: ['00:00', '01:15', '02:30', '03:45', '05:00', '06:15', '07:30', '08:45', '10:00', '11:15', '12:30', '13:45', '15:00', '16:15', '17:30', '18:45', '20:00', '21:15', '22:30', '23:45']
          }],
          yAxis: {
            name:'电力负荷/MW',
            type: 'value'
          },
          series: [{
            name: '用电量',
            type: 'line',
            smooth:'true',
            data: [10, 20, 30, 40, 50]
          }]
        };
      let ws = new WebSocket("ws://localhost:8081/ws")
      ws.onopen = function () {
        console.log("链接成功")
      }
      ws.onmessage = function (msg) {
        let pl = JSON.parse(msg.data)
        console.log(pl)
        option.series[0].data = pl
        option && myChart.setOption(option);
      }

    }
  },
  mounted() {
    this.drawLine();
  }
}
</script>

<style scoped>
</style>
