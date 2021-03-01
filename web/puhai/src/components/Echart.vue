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
      this.$axios.get("http://localhost:8081/getData") .then((res) => {
        const data = res.data;
        // const list = data.series.map(good=>{
        //   let list =  good.data;
        //   return [...list]
        // })
        console.log(data)
        option = {
          title: data.title,
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
            data: data.series.data
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
