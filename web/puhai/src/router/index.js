import Vue from 'vue'
import Router from 'vue-router'
import page1 from '../components/page1'
import page2 from '../components/page2'
import element from "../components/element"
import home from "../components/home"
import Dashboard from "../../view/dashboard/Dashboard"
import Echart from "../components/Echart";

Vue.use(Router)


export default new Router ({
  routes: [
    {
      path: '/page1',
      name: 'page1',
      component: page1
    },
    {
      path: '/page2',
      name: 'page2',
      component: page2
    },
    {
      path: '/element',
      name: 'element',
      component: element
    },
    {
      path: '/home',
      name: 'home',
      component: home
    },
    {
      path: '/Echart',
      name: 'Echart',
      component: Echart
    },
    {
      path: '/Dashboard',
      name: 'Dashboard',
      component: Dashboard
    },
    {
      path: '',
      redirect: Dashboard
    },
  ]
  }
)
