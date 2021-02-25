<template>
  <div id="myChart" :style="{width: '100%', height: '300px'}"></div>
</template>

<script>
import * as echarts from 'echarts';
export default {
  name: "Echart",
  data(){
    return {
      goods:{}
    }
  },
  methods: {
    drawLine(){
      // 基于准备好的dom，初始化echarts实例
      let myChart = echarts.init(document.getElementById("myChart"));
      // 绘制图表
      myChart.setOption({
        title: {},
        tooltip: {},
        xAxis: {
          data: []
        },
        yAxis: {},
        series: [{
          name: '销量',
          type: 'bar',
          data: []
        }]
      });
      this.$axios.get("http://picture.nj-jay.com/dat.json") .then((res) => {
        const data = res.data;
        const list = data.series.map(good=>{
          let list =  good.data;
          return [...list]
        })
        console.log(list);
        console.log(Array.from(...list));
        myChart.setOption({
          title: data.title,
          xAxis: [{
            data: data.xAxis.data
          }],
          series: [{
            name: '销量',
            type: 'bar',
            data: Array.from(...list) //[5, 20, 36, 10, 10, 20]
          }]
        });
      })
    }
  },
  mounted() {
    this.drawLine();
  },
  created() {
    this.$axios.get("http://picture.nj-jay.com/dat.json")
      .then(function (response) {
        const data = response.data;
        this.goods = data;
        console.log(response)
    })
  }
}
</script>

<style scoped>

</style>
