<template>
  <img alt="logo" src="./assets/logo.svg" class="logo">
  <DbView />
</template>

<script>
import DbView from './components/DbView.vue'

export default {
  name: 'App',
  data() {
    return {};
  },
  components: {
    DbView
  }
}
</script>

<style lang="scss">
.logo {
  width: 200px;
}
body {
  background: #0d3049;
  #app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #F2F2F2;
    margin-top: 60px;
  }
}
</style>
