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
          xAxis: [{
            name:'日期',
            data: ["1", "2", "3", "4", "5"]
          }],
          yAxis: {
            name:'电力负荷/MW',
            type: 'value'
          },
          series: [{
            name: '销量',
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
