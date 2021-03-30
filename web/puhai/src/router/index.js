import Vue from 'vue'
import Router from 'vue-router'
import upload from '../components/upload'
import login from '../view/login/login'
import element from "../components/element"
import home from "../components/home"
import Dashboard from "../view/dashboard/Dashboard"
import Echart from "../components/Echart";

Vue.use(Router)


export default new Router ({
  routes: [
    {
      path: '/upload',
      name: 'upload',
      component: upload
    },
    {
      path: '/login',
      name: 'login',
      component: login
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
