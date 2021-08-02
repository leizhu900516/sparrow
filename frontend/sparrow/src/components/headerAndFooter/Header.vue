<template>
  <el-header>
    <div class="logo" @click="gotohome">
      <img src="../../assets/logo.png">
      <span>麻雀 WIKI</span>
    </div>
    <div class="search" v-if="displaysearch">
      <el-input
        placeholder="搜索"
        v-model="globalSearchKeyword"
        @change="gotoSearch"
        prefix-icon="el-icon-search">
      </el-input>
    </div>
    <div class="header-items">
      <el-menu :default-active="activeIndexMenu" class="el-menu-demo" mode="horizontal" @select="handleSelect" router>
        <el-menu-item index="/workspace">工作台</el-menu-item>
        <el-menu-item index="/wiki">知识</el-menu-item>
        <el-menu-item index="/book">图书馆</el-menu-item>
        <el-submenu index="4">
          <template slot="title"><i class="el-icon-circle-plus"></i>新建</template>
          <el-menu-item @click="dialogVisible = true">新建文档</el-menu-item>
          <el-menu-item index="/new">新建知识库</el-menu-item>
          <el-menu-item index="/add/team">新建团队</el-menu-item>
        </el-submenu>
        <el-submenu index="5">
          <template slot="title" v-if="username !== ''"><i class="el-icon-user-solid" ></i>{{username}}</template>
          <template slot="title" v-if="username === ''"><i class="el-icon-user-solid" ></i>未登录</template>
          <el-menu-item :index="username">个人主页</el-menu-item>
          <el-menu-item index="/setting">账户设置</el-menu-item>
          <el-menu-item index="/sys/setting">系统设置</el-menu-item>
          <el-menu-item @click="logout">退出</el-menu-item>
        </el-submenu>
      </el-menu>
    </div>
    <!--新建文档弹出框-->
    <el-dialog title="新建文档" :visible.sync="dialogVisible" width="30%" >
      <div v-if="repository">
        <span>点击选择一个最近参与的知识库</span>
        <ul class="repository_list"  v-for="rep in repository" :key="rep.id">
          <li @click="selectRepository($event)"  v-bind:repid="rep.repo_unique_code" v-bind:repname="rep.name" :aaa="rep.name">
            <i class="el-icon-document"></i>
            <span>{{ rep.name }}</span>
          </li>
        </ul>
      </div>
      <div>
        <el-button type="primary" v-if="!repository" @click="createRepo">还未有知识库，请先创建知识库</el-button>
      </div>
    </el-dialog>
  </el-header>
</template>

<script>
    export default {
      name: 'Header',
      // provide() {
      //   return {
      //     updateGlobalSearchKeyword:this.updateGlobalSearchKeyword
      //   }
      // },
      data() {
        return {
          activeIndexMenu: '/workspace',
          globalSearchKeyword:'',
          dialogVisible: false,
          form: {
            name:'',
            region:''
          },
          displaysearch:true,
          repository:[],
          username:''
        }
      },
      created() {
        this.Init()
        this.getRepository()
        this.username = window.localStorage.getItem('username')
      },
      methods:{
        Init() {
          var localpath = this.$route.path
          if (localpath === '/') {
            this.displaysearch = true
          }
          this.activeIndexMenu = localpath
        },
        openNewBox() {
          var that = this
          that.$msgbox('这是一段内容', '标题名称', {
            confirmButtonText: '确定'
          })
        },
        selectRepository(event) {
          var repid = event.currentTarget.getAttribute('repid')
          var repname = event.currentTarget.getAttribute('repname')
          this.$router.push({ name: 'wikiedit', params: { docCode:'luckid' ,repid: repid, repname:repname } })
        },
        getRepository() {
          var that = this
          that.$http.get('/api/v1/repository?type=doc').then(function(response) {
              if (response.status === 200 && response.data.code === 0) {
                that.repository = response.data.data
            }
          })
        },
        createRepo() {
          this.$router.push('/new')
        },
        // 头部菜单栏
        handleSelect(key, keyPath) {
          this.activeIndexMenu = key
        },
        gotoSearch(value) {
          console.log('搜索词：',value)
          this.$router.push({ path: '/search', query: { kw: value } })
        },
        gotohome() {
          this.$router.push('/')
        },
        logout() {
          // 清除token
          this.$store.commit('settoken','')
          this.$router.push('/login')
        }
      }
    }
</script>

<style>
  .el-header{
    background-color: #fff;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #dcdfe6;
  }
  .logo{
    height: 100%;
    display: flex;
    align-items: center;
    font-size: 18px;
    text-space: 10px;
    cursor: pointer;
  }
  .logo span{
    margin-left: 10px;
  }
  .logo img{
    height: 30px;
  }
  .el-dialog{
    padding: 10px 20px 10px;
  }
  .el-dialog__header {
    padding: 10px 20px 10px;
  }
  .repository_list li{
    cursor: pointer;
    display: flex;
    height: 50px;
    flex-direction: row;
    align-content: center;
    border-bottom: 1px solid #e8e8e8;
    line-height: 24px;
    align-items: center;
  }
  .repository_list li:hover{
    background-color:#f5f5f5;
  }
  .repository_list li span{
    margin-left: 10px;
  }
</style>
