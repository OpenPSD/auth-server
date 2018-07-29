import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/Login'
import Indexx from '@/components/Index'

Vue.use(Router)

export default new Router({
  // mode: 'history',
  routes: [
    {
      path: '/login/',
      name: 'Login',
      component: Login
    },
    {
      path: '/',
      name: 'Index',
      component: Indexx
    }
  ]
})
