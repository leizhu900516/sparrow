<template>
  <el-container   class="body-container">
    <Header></Header>
    <el-container  class="top-left-right-bottom-container">
      <el-aside width="200px">
        <el-menu
          default-active="2"
          class="el-menu-vertical-demo"
          @open="handleOpen"
          @close="handleClose" router>
          <el-menu-item index="collect">
            <i class="el-icon-star-on"></i>
            <span slot="title">收藏</span>
          </el-menu-item>
          <el-submenu index="2">
            <template slot="title">
              <i class="el-icon-location"></i>
              <span>个人知识库</span>
            </template>
            <el-menu-item-group>
              <el-menu-item  v-for="rep in docRepositoryList"  v-bind:key="rep.id" @click="gotoRepoWorkHome(rep.repo_unique_code)">
                <span v-bind:class="rep.cateid == 1 ? 'iconfont iconwendang' :'iconfont iconwenjianjia'"> {{rep.name}}</span>
              </el-menu-item>
            </el-menu-item-group>
          </el-submenu>
          <el-submenu index="3">
            <template slot="title">
              <i class="el-icon-share"></i>
              <span>协作知识库</span>
            </template>
            <el-menu-item-group>
              <el-menu-item v-for="rep in fileRepositoryList " v-bind:key="rep.id" @click="gotoRepoWorkHome(rep.repo_unique_code)">
                <span v-bind:class="rep.cateid == 1 ? 'iconfont iconwendang' :'iconfont iconwenjianjia'"> {{rep.name}}</span>
              </el-menu-item>
            </el-menu-item-group>
          </el-submenu>
          <el-submenu index="4">
            <template slot="title">
              <i class="el-icon-s-check"></i>
              <span>团队</span>
            </template>
            <el-menu-item-group>
              <el-menu-item  v-for="team in teams" v-bind:key="team.id" @click="gotoTeamHome(team.flag)">{{team.name}}</el-menu-item>
            </el-menu-item-group>
          </el-submenu>
          <el-menu-item index="/recycle">
            <i class="el-icon-delete-solid"></i>
            <span slot="title">回收站</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-container>
        <el-main>
          <router-view></router-view>
        </el-main>
        <el-footer>© Copyright by Sparrow</el-footer>
      </el-container>
    </el-container>
  </el-container>
</template>

<script>
  import Header from './headerAndFooter/Header'
  export default {
    name: 'Home',
    components: {
      Header
    },
    created() {
      this.getRepository()
      this.getTeam()
    },
    data() {
      return {
        docRepositoryList:[],
        fileRepositoryList:[],
        teams:[]
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
        // var path = event.target.getAttribute('index')
        // this.$router.push(path)
      },
      handleOpen(key, keyPath) {
        console.log(key, keyPath)
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath)
      },
      getTeam() {
        var that = this
        this.$http.get('/api/v1/team').then(function (response) {
            var datas = response.data.data
            if (response.data.code === 0) {
              if (datas != null) {
                that.teams = response.data.data
              }
            }
        })
      },
      getRepository() {
        var that = this
        this.$http.get('/api/v1/repository').then(function (response) {
          var datas = response.data.data
          console.log(datas)
          if (response.data.code === 0) {
            if (datas != null) {
              for (var i = 0; i < datas.length; i++) {
                if (datas[i].groupid === 9999) { // 9999是个人知识库
                  that.docRepositoryList.push(datas[i])
                } else {
                  that.fileRepositoryList.push(datas[i])
                }
              }
            }
          }
        })
      },
      gotoRepoWorkHome(respUniqueCode) {
        this.$router.push({ name:'repoflag', params: { username:'chenhuachao',repoflag:respUniqueCode } })
      },
      gotoTeamHome(flag) {
        this.$router.push({ name:'teamflag',params:{ teamflag:flag } })
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
  .el-aside {
    line-height: 200px;
  }
  .el-menu{
    height: 100%;
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
