<template>
  <div id="app">
    <b-navbar toggleable="md" type="dark" variant="info">
      <b-navbar-brand href="#/">카카오 채팅 분석</b-navbar-brand>
        <b-collapse is-nav id="nav_collapse">
          <b-navbar-nav>
            <b-nav-item href="#/" @click="changeTitle('키워드')">키워드</b-nav-item>
            <b-nav-item href="#/user" @click="changeTitle('사용자')">사용자</b-nav-item>
            <b-nav-item href="#" @click="changeTitle('날짜별')">날짜별</b-nav-item>
          </b-navbar-nav>
        </b-collapse>
        <b-navbar-brand href="#" right>{{title}}</b-navbar-brand>
    </b-navbar>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      title : "키워드"
    }
  },
  methods : {
    changeTitle(title) {
      this.title = title;
    }
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
