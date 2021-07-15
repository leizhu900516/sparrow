<!--个人主页-->
<template>
    <el-container class="body-container">
      <Header></Header>
      <el-main  class="main-containers-custom" style="width: 100%">
        <el-row :gutter="20">
          <el-col :span="6">
            <el-card shadow="never" class="userinfo">
              <img :src="avatarurl !== 'null' ? avatarurl:require('@/assets/imgs/defaultAvator.png')" class="img-avatar">
              <h1>{{username}}</h1>
              <div>
                <div>
                  <span>{{beFollowCount}}</span>
                  <span>关注者</span>
                </div>
                <div>
                  <span>{{followCount}}</span>
                  <span>关注了</span>
                </div>
              </div>
              <el-button ><router-link :to="{name:'usersetting'}">编辑资料</router-link></el-button>
              <el-divider></el-divider>
              <span>共{{docCount}} 篇公开文档</span>
            </el-card>
            <el-card shadow="never" class="user-team">
              <div slot="header" class="teams-item">
                <span >团队</span>
              </div>
              <ul class="team-ul-list">
                <li v-for="team in teams" v-bind:key="team.id" @click="gotoTeamHome(team.flag)">{{team.name}}</li>
              </ul>
            </el-card>
          </el-col>
          <el-col :span="18">
            <el-card shadow="never">
              <div slot="header">
                <span>知识库</span>
              </div>
              <div>
                <ul>
                  <li class="repo-li"  v-for="resp in RepositoryList" v-bind:key="resp.id">
                    <div class="resp-name-and-del-btn">
                      <span class="repo-name" @click="gotoRepoWorkHome(resp.repo_unique_code)">{{resp.name}}</span>
                      <el-button type="danger" size="mini" @click="deleteResp(resp.repo_unique_code)">删除</el-button>
                    </div>
                    <p class="repo-desc">可用于学习笔记、周报、项目文档等场景</p>
                    <div class="repo-info">
                      <span>文档 {{resp.followsum}}</span>
                      <span>关注 {{resp.docsum}}</span>
                    </div>
                  </li>
                </ul>
              </div>
              <div class="no-repository-warring" v-if="RepositoryList.length === 0 ">
                <span>还未创建知识库</span>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </el-main>
      <Footer></Footer>
    </el-container>
</template>

<script>
    import Header from '../headerAndFooter/Header'
    import Footer from '../headerAndFooter/Footer'

    export default {
        name: 'Userhome',
          components: {
            Header,
              Footer
          },
      data() {
          return {
            username:'',
            followCount:'',
            beFollowCount:'',
            docCount:'',
            avatarurl:'',
            RepositoryList:[],
            teams:[]
          }
      },
      created() {
        this.getUserinfo()
        this.getRepositoryList()
        this.getTeam()
      },
      methods:{
          getUserinfo() {
            var that = this
            that.$http.get('api/v1/user').then(function (response) {
              if (response.data.code === 0) {
                if (response.data.data != null) {
                  var data = response.data.data
                  console.log(data)
                  that.followCount = data.followCount
                  that.beFollowCount = data.beFollowCount
                  that.docCount = data.docCount
                  that.avatarurl = window.location.origin + '/' + data.avatarurl
                  that.username = data.username
                }
              }
            })
          },
        gotoRepoWorkHome(respUniqueCode) {
            console.log(respUniqueCode)
          this.$router.push({ name:'repoflag', params: { username:'chenhuachao',repoflag:respUniqueCode } })
        },
        getRepositoryList() {
          var that = this
          that.$http.get('api/v1/repository').then(function (response) {
            if (response.data.code === 0) {
              that.RepositoryList = response.data.data
            }
          })
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
        gotoTeamHome(flag) {
          this.$router.push({ name:'teamflag',params:{ teamflag:flag } })
        },
        deleteResp(respflag) {
          var that = this
          that.$confirm('此操作将清空知识库, 是否继续?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            that.$http.delete('api/v1/repository/' + respflag).then(function (response) {
              if (response.data.code === 0) {
                that.$message.success({
                  type: 'success',
                  message: '删除成功!'
                })
                for (var i = 0; i < that.RepositoryList.length; i++) {
                  if (that.RepositoryList[i].repo_unique_code === respflag) {
                    that.RepositoryList.splice(that.RepositoryList[i],1)
                  }
                }
              }
            })
          }).catch(() => {
            // 取消的时候，什么也不做
          })
        }
      }
    }
</script>

<style scoped>
  .main-containers-custom{
    margin-left: auto;
    margin-right: auto;
    max-width: 1056px;
  }

  .el-card >>> .el-card__body{
    padding: 0px!important;
  }
  .img-avatar{
    width: 160px;
    min-width: 160px;
    height: 160px;
    border-radius: 80px;
    border-color: rgba(0, 0, 0, 0.06);
    border-width: 1px;
    border-style: solid;
  }
  .user-team{
    margin-top: 10px;
  }
  .teams-item{
    display: flex;
    flex-direction: column;
  }
  .repo-li{
    border-bottom: 1px solid #dcdfe6;
    padding: 20px;
  }
  .repo-name{
    font-size: 16px;
    font-weight: bold;
    cursor: pointer;
  }
  .repo-desc{
    font-size: 14px;
    color: #595959;
  }
  .repo-info{
    font-size: 12px;
    color: #595959;
  }
  .repo-info span:nth-child(2){
    margin-left: 10px;
  }
  .userinfo{
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    line-height: 40px;
    text-align: center;
  }
  .team-ul-list{
    padding: 20px;
  }
  .team-ul-list li{
    cursor: pointer;
    margin-bottom: 10px;
  }
  .no-repository-warring{
    height: 200px;
    width: 100%;
    text-align: center;
    margin-top: 50px;
  }
  .resp-name-and-del-btn{
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }
</style>
