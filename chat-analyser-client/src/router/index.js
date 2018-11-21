import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import topList from '@/components/topList/topListPage'
import userList from '@/components/topList/userListPage'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'topListPage',
      component: topList
    },
    {
      path: '/user',
      name: 'userListPage',
      component: userList
    }
  ]
})
