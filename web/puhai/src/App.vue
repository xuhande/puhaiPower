<template>
  <div id="app">
    <img src="./assets/logo.png">
    <HelloWorld/>
    <Echart/>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld'
import Echart from "./components/Echart";

export default {
  name: 'App',
  components: {
    HelloWorld,
    Echart
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
