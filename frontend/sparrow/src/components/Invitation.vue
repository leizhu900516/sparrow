<template>
  <el-container class="body-container">
    <Header></Header>
    <el-main  class="main-containers-custom" style="width: 100%">
      <el-card class="box-card" shadow="never">
        <div slot="header" class="clearfix">
          <span>一起加入团队来创作吧！</span>
        </div>
        <el-row :gutter="12" v-if="noTeamFlag === false">
          <el-col :span="12">
            <el-card shadow="never">
              <div class="team-desc">
                <h3>团队名称：{{teamname}}</h3>
                <p class="team-desc-content">团队描述</p>
                <span class="groupdesc">{{teamdesc}}</span>
                <el-divider></el-divider>
                <span style="margin-bottom: 10px;">团队知识库</span>
                <div>
                  <el-col :span="12" v-for="resp in resplist" v-bind:key="resp.respname">
                    <el-card shadow="never">
                      <div class="resp-list-desc">
                        <span class="repo-title">{{resp.respname}}</span>
                        <span class="repo-desc">{{resp.respdesc}}</span>
                      </div>
                    </el-card>
                  </el-col>

                </div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="12" class="team-join-div">
              <div >
                <span class="alert-join-title">是否加入该团队</span>
                <div class="team-join-btn">
                  <el-button @click="cancelJoin">暂不加入</el-button>
                  <el-button type="primary" @click="joinTeam">加入团队</el-button>
                </div>
              </div>
          </el-col>
        </el-row>
        <el-row :gutter="12" v-if="noTeamFlag === true">
          <el-col :span="24" class="no-team-info">
              <span><i class="el-icon-error"></i>{{errmsg}}</span>
          </el-col>

        </el-row>
      </el-card>
    </el-main>
    <Footer></Footer>
  </el-container>
</template>

<script>
  import Header from './headerAndFooter/Header'
  import Footer from './headerAndFooter/Footer'
  export default {
    name: 'Invitation',
    components: {
      Header,
      Footer
    },
    data() {
      return {
        teamname: '',
        group_avator_url_id: '', // 团队图片
        teamdesc: '',
        teamid: '',
        teamflag: '',
        encode: '',
        noTeamFlag: false,
        errmsg: '',
        resplist: []
      }
    },
    created() {
      var encode = this.$route.params.encode
      this.encode = encode
      this.decodeCode(encode)
      // this.getTeamInfo(teamflag)
      // this.getTeamMember(teamflag)
    },
    methods: {
      decodeCode(code) {
        var that = this
        that.$http.get('api/v1/invitation/' + code).then(function (response) {
          console.log(response)
          if (response.data.code === 0) {
            that.teamname = response.data.data.groupname
            that.teamdesc = response.data.data.groupdesc
            that.teamid = response.data.data.gid
            that.resplist = response.data.data.resplist
            that.teamflag = response.data.data.groupUniqueCode
          } else {
            that.noTeamFlag = true
            that.errmsg = response.data.msg
          }
        })
      },
      getTeamRepo(teamid) {
        var that = this
        that.$http.get('api/v1/team/repository?groupid=' + teamid).then(function (response) {
          if (response.data.code === 0) {
            that.respositoryList = response.data.data
          }
        })
      },
      cancelJoin() {
        setTimeout(this.$router.push('/'),1000)
      },
      joinTeam() {
        var that = this
        console.log(that.$route.path)
        that.$http.get('api/v1/invitation/' + that.encode + '?join=true').then(function (response) {
          if (response.data.code === 0 || response.data.code === 2) {
            that.$message.success(response.data.msg)
            setTimeout(function () {
              that.$router.push({ name:'teamflag',params:{ teamflag:that.teamflag } })
            },2000)
          } else {
            that.$message.success(response.data.msg)
          }
        })
      }
    }
  }
</script>

<style scoped>
  .team-desc{
    width: 100%;
    display: flex;
    flex-direction: column;
    height: 100%;
  }
  .team-join-btn{
    margin-top: 10px;
  }
  .alert-join-title{
    color: #409EFF;
    margin-left: 20px;
  }
  .team-join-div{
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    margin-top: 50px;
  }
  .groupdesc{
    font-size: 13px;
    color: #909399;
  }
  .no-team-info{
    height: 400px;
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    font-size: 18px;
    color: #F56C6C;
  }
  .resp-list-desc{
    display: flex;
    flex-direction: column;
  }
  .repo-title,.team-desc-content{
    font-size: 16px;
    color: #303133;
  }
  .resp-list-desc .repo-desc{
    font-size: 14px;
    color: #909399;
  }
</style>
