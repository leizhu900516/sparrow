<template>
  <el-container class="body-container">
    <Header></Header>
    <el-header style="height: 120px">
      <div class="team-info main-containers-custom">
        <span class="iconfont iconzu font-size-30" v-if="group_avator_url_id === ''"></span>
        <img :src="group_avator_url_id" v-if="group_avator_url_id !== ''" class="smallavatar">
        <div class="team-desc-info">
          <h2>{{teamname}}</h2>
          <span class="teamdesc">{{teamdesc}}</span>
        </div>
      </div>
    </el-header>
    <el-main  class="main-containers-custom" style="width: 100%">
      <el-tabs v-model="activeName" @tab-click="handleClick">
        <el-tab-pane label="首页" name="first">
          <div>
            <h4>常用功能</h4>
            <el-card shadow="never">
              <el-row :gutter="12">
                <el-col :span="8">
                  <el-card shadow="never" @click.native="addRepository" class="common-function">
                    <span>新建知识库</span>
                    <h6>沉淀团队知识</h6>
                  </el-card>
                </el-col>
                <el-col :span="8">
                  <el-card shadow="never" @click.native="addGroupMember" class="common-function">
                    <span>添加成员</span>
                    <h6>和成员一起创作知识</h6>
                  </el-card>
                </el-col>
                <el-col :span="8">
                </el-col>
              </el-row>
            </el-card>
          </div>
          <div>
            <h4>知识库</h4>
            <div>
              <el-row :gutter="12">
                <el-col :span="8">
                  <el-card shadow="never">
                    <div slot="header" class="clearfix">
                      <span>知识库名称</span>
                    </div>
                    <div v-for="repo in respositoryList" :key="repo.id" class="repository-name" @click="gotoRepoWorkHome(repo.respuniquecode)">
                      <span class="iconfont iconwenjianjia"></span>{{repo.respname }}
                    </div>
                  </el-card>
                </el-col>
                <el-col :span="8">
                </el-col>
                <el-col :span="8">
                </el-col>
              </el-row>
            </div>
          </div>
        </el-tab-pane>
        <el-tab-pane label="知识库" name="second">
          <div>
            <div class="repository-and-search-create">
              <h4>知识库</h4>
              <div class="search-create">
<!--                <el-input-->
<!--                  placeholder="搜索"-->
<!--                  v-model="globalSearchKeyword"-->
<!--                  @change="gotoSearch"-->
<!--                  prefix-icon="el-icon-search">-->
<!--                </el-input>-->
                <el-button type="primary" @click="addRepository">新建知识库</el-button>
              </div>
            </div>

            <el-table
              :data="respositoryList"
              style="width: 100%">
              <el-table-column
                prop="respname"
                label="知识库"
                width="280">
                <template slot-scope="scope">
                  <span class="iconfont iconwenjianjia"></span>
                  <span class="repository-name" @click="gotoRepoWorkHome(scope.row.respuniquecode)">{{ scope.row.respname }}</span>
                </template>
              </el-table-column>
              <el-table-column
                prop="resp_desc"
                label="简介">
              </el-table-column>
              <el-table-column
                prop="createtime"
                label="日期"
                width="180">
              </el-table-column>
            </el-table>
          </div>
        </el-tab-pane>
        <el-tab-pane label="成员" name="third">
          <div>
            <div class="repository-and-search-create">
              <h4>成员列表</h4>
              <div class="search-create">
                <el-input
                  placeholder="搜索成员"
                  v-model="globalSearchUsername"
                  @change="gotoSearchUsername"
                  prefix-icon="el-icon-search">
                </el-input>
                <el-button type="primary" @click="dialogVisible = true" class="invitation-btn">邀请新成员</el-button>
              </div>
            </div>
            <el-table
              :data="userlist"
              style="width: 100%">
              <el-table-column
                prop="username"
                label="昵称"
                width="180">
              </el-table-column>
              <el-table-column
                prop="username"
                label="用户名">
              </el-table-column>
              <el-table-column
                prop="usertype"
                label="角色"
                width="180">
              </el-table-column>
              <el-table-column
                prop="handle"
                label="操作"
                width="180">
                <template slot-scope="scope">
                  <span class="exit-team-btn" @click="deleteGroupMember(scope.row.uid)">删除</span>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-tab-pane>
      </el-tabs>
    </el-main>
    <Footer></Footer>
    <!--新建文档弹出框-->
    <el-dialog title="邀请新成员" :visible.sync="dialogVisible" width="30%" >
      <div class="invitation-dialog">
        <span>通过链接进行邀请，默认3天后失效</span>
        <div class="invitation-value-btn">
          <el-input
            placeholder="invitationUrl"
            v-model="invitationUrl"
            :disabled="true">
          </el-input>
          <el-button type="primary" size="small" :data-clipboard-text="invitationUrl" @click.native="copyUrl" class="copy_btn">复制链接</el-button>
        </div>
      </div>
    </el-dialog>
  </el-container>
</template>

<script>
    import Header from '../headerAndFooter/Header'
    import Footer from '../headerAndFooter/Footer'
    import Clipboard from 'clipboard'

    export default {
        name: 'Team',
      components: {
        Header,
        Footer
      },
      data() {
          return {
            teamname:'',
            group_avator_url_id:'', // 团队图片
            teamdesc:'',
            teamid:'',
            teamflag:'',
            activeName: 'first',
            // globalSearchKeyword:'',
            globalSearchUsername:'',
            invitationUrl:'未绑定',
            dialogVisible:false,
            tableData: [],
            userlist:[],
            respositoryList:[]
          }
      },
      created() {
        var teamflag = this.$route.params.teamflag
        this.teamflag = teamflag
        this.getTeamInfo(teamflag)
        this.getTeamMember(teamflag)
        this.invitationUrl = window.location.origin + '/invitation/:encode'
      },
      watch:{
        dialogVisible:function(val) {
          var that = this
          if (val === true) {
            if (that.teamflag === '') {
              this.$message.error('团队标识为空')
              return
            }
            that.$http.get('api/v1/generate/invitation/encode?expired=3&flag=' + that.teamflag).then(function (response) {
                if (response.data.code === 0) {
                  that.invitationUrl = window.location.origin + '#/invitation/' + response.data.data
                }
            })
          }
          console.log('dialogVisible',val)
          // 获取邀请链接
        }
      },
      methods: {
        handleClick(tab, event) {
          console.log(tab, event)
        },
        getTeamInfo(teamflag) {
          var that = this
          that.$http.get('api/v1/team/info/' + teamflag).then(function (response) {
              if (response.data.code === 0) {
                that.teamname = response.data.data.group_name
                that.teamdesc = response.data.data.group_desc
                that.teamid = response.data.data.id
                if (response.data.data.group_avator_url_id) {
                  that.group_avator_url_id = that.$http.defaults.baseURL + 'api/v1/img/' + response.data.data.group_avator_url_id
                }
                that.getTeamRepo(response.data.data.id)
              }
          })
        },
        // gotoSearch(value) {
        //   console.log('搜索词：',value)
        // },
        gotoSearchUsername(value) {
          console.log('搜索词：',value)
        },
        getTeamRepo(teamid) {
          var that = this
          that.$http.get('api/v1/team/repository?groupid=' + teamid).then(function (response) {
                if (response.data.code === 0) {
                  that.respositoryList = response.data.data
                }
          })
        },
        deleteGroupMember(uid) {
          var that = this
          that.$confirm('确认删除, 是否继续?' , '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            that.$http.delete('api/v1/member/' + that.teamflag + '/' + uid).then(function (response) {
              if (response.data.code === 0) {
                that.$message.success('删除成功')
              } else {
                that.$message.error(response.data.msg)
              }
            })
          }).catch(() => {
          })
        },
        addGroupMember() {
          // 弹出框 进行站内搜索或者链接邀请
          this.activeName = 'third'
        },
        copyUrl() {
          console.log('开始复制')
          var that = this
          var clipboard = new Clipboard('.copy_btn')
          clipboard.on('success', e => {
            that.$message.success('复制成功')
            // 释放内存
            clipboard.destroy()
          })
          clipboard.on('error', e => {
            // 不支持复制
            that.$message.error('该浏览器不支持自动复制')
            // 释放内存
            clipboard.destroy()
          })
        },
        addRepository() {
          this.$router.push({ name:'New',params:{ groupid:this.teamid } })
        },
        getTeamMember(flag) {
          var that = this
          that.$http.get('api/v1/member/' + flag).then(function (response) {
            if (response.data.code === 0) {
              that.userlist = response.data.data
            }
          })
        },
        gotoRepoWorkHome(respUniqueCode) {
          this.$router.push({ name:'repoflag', params: { username:'chenhuachao',repoflag:respUniqueCode } })
        }
      }
    }
</script>

<style scoped>
  .team-info{
    width: 100%;
    display: flex;
    flex-direction: row;
    height: 100%;
    align-items: center;
  }
  .main-containers-custom{
    margin-left: auto;
    margin-right: auto;
    max-width: 1056px;
  }
  .teamdesc{
    color: #606266;
    font-size: 14px;
  }
  .big-icon-40{
    font-size: 40px;
  }
  .team-desc-info{
    display: flex;
    flex-direction: column;
    margin-left: 20px;
  }
  .repository-and-search-create {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    height: 40px;
  }
  .search-create{
    display: flex;
    flex-direction: row;
  }
  .exit-team-btn{
    cursor: pointer;
    color: #409EFF;
  }
  .common-function{
    cursor: pointer;
    background-color: #409EFF;
    color: #fff;
    /*text-align: center;*/
  }
  .repository-name{
    cursor: pointer;
    margin-bottom: 5px;
  }
  >>> .invitation-btn{
    margin-left: 10px;
  }
  .invitation-value-btn{
    display: flex;
    flex-direction: row;
  }
  .invitation-dialog{
    line-height: 40px;
  }
</style>
