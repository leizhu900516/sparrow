<template>
  <el-container class="body-container">
    <Header></Header>
    <el-main :body-style="{ padding:'10px' }">
      <router-view></router-view>
    </el-main>
    <el-footer>© Copyright by Sparrow</el-footer>
  </el-container>
</template>

<script>
  import Header from './headerAndFooter/Header'
  export default {
    name: 'SecondHome',
    components: {
      Header
    },
    created() {
      this.getRepository()
    },
    data() {
      return {
        dialogVisible: false,
        avatarurl: 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png',
        respositorys:[{ name:'知识库1', id:1 }, { name:'知识库2', id:1 }]
      }
    },
    methods:{
      openNewBox() {
        var that = this
        that.$msgbox('这是一段内容', '标题名称', {
          confirmButtonText: '确定'
        })
      },
      selectRepository(event) {
        var path = event.target.getAttribute('index')
        this.$router.push(path)
      },
      getRepository() {
        this.$http.get('/api/v1/repository')
      }
    }
  }
</script>

<style scoped>
  .el-container{
    height: 100%;
  }
  .el-footer{
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: row;
    color: #595959;
    font-size: 14px;
  }
  .el-header{
    display: flex;
    justify-content: space-between;
    flex-direction: row;
    align-items: center;
    background-color: #fff;
    border-bottom: 1px solid #dcdfe6;
  }
  .logo{
    height: 100%;
    display: flex;
    align-items: center;
    font-size: 20px;
    text-space: 10px;
  }
  .logo span{
    margin-left: 10px;
  }
  .logo img{
    height: 45px;
  }

  .el-menu{
    height: 100%;
  }

</style>
