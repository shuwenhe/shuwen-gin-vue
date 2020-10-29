<template>
  <div id="app">
    <b-container>
      <navbar/>
    </b-container>
  </div>
</template>

<script>
import { component } from 'vuedraggable'
import Navbar from "./views/layout/Navbar.vue"

export default {
  data(){
    component:{Navbar} // 声明组件
  },
}
</script>

<style lang="scss">

</style>
