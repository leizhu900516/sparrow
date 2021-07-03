<template>
  <el-container class="body-container">
    <Header></Header>
    <router-view></router-view>
    <Footer></Footer>
  </el-container>
</template>

<script>
    import Header from '../headerAndFooter/Header'
    import Footer from '../headerAndFooter/Footer'
    export default {
      name: 'RespositoryHome',
      components: {
        Header,
        Footer
      },
      data() {
        return {
        }
      },
      created() {
      }
    }
</script>

<style scoped>

</style>
