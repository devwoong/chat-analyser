import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import topList from '@/components/topList/topListPage'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'topList',
      component: topList
    }
  ]
})
