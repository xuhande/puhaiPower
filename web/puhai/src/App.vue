<template>
  <div id="app">
    <router-view></router-view>
  </div>

</template>

<script>
import HelloWorld from './components/HelloWorld'
import Echart from "./components/Echart"
import home from "./components/home"
import Dashboard from "../view/dashboard/Dashboard";
export default {
  name: 'App',
  components: {
    HelloWorld,
    Echart,
    home,
    Dashboard
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: start;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
