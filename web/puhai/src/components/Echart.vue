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
      this.$axios.get("http://picture.nj-jay.com/dat.json") .then((res) => {
        const data = res.data;
        const list = data.series.map(good=>{
          let list =  good.data;
          return [...list]
        })
        option = {
          title: data.title,
          xAxis: [{
            name:'日期',
            data: data.xAxis.data
          }],
          yAxis: {
            type: 'value'
          },
          series: [{
            name: '销量',
            type: 'line',
            data: Array.from(...list)
          }]
        };
        option && myChart.setOption(option);
      })
    }
  },
  mounted() {
    this.drawLine();
  },
}
</script>

<style scoped>
</style>
