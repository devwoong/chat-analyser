import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import topList from '@/components/topList/topListPage'
import userList from '@/components/userList/userListPage'
import dateList from '@/components/dateList/dateListPage'
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
    },
    {
      path: '/date',
      name: 'dateListPage',
      component: dateList
    }
  ]
})
