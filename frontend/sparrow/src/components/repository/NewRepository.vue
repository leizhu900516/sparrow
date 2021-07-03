<template>
  <!--选择知识库分类-->
  <div>
    <v-header></v-header>
    <el-main class="main-container">
        <div class="newTitle">
          <h3>新建知识库</h3>
          <span>创作、管理各种类型的知识</span>
        </div>
        <el-card class="box-card">
          <el-row>
            <el-col :span="12">
              <div class="left_new">
                <span class="belong">属于</span>
                <el-select v-model="team" placeholder="不选择默认团队是自己" class="teamlist" @change="selectTeam($event)">
                  <el-option
                    v-for="item in teams"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"></el-option>
                </el-select>
                <el-divider></el-divider>
                <el-input placeholder="请输入名称" class="repository_name" @change="inputTeamName" v-model="reponame"></el-input>
                <div class="new_type">
                  <div @click="changeSelectStatus(1)"  v-bind:class="{ active: repotype === 1 }" class="typeitem">
                    <i class="el-icon-document big-icon-40"></i><span>文档知识库</span>
                  </div>
                  <div @click="changeSelectStatus(2)"  v-bind:class="{ active: repotype === 2  }" class="typeitem">
                    <i class="el-icon-folder big-icon-40"></i><span>资源知识库</span>
                  </div>
                </div>
                <span class="rep_desc">简介</span>
                <el-input type="textarea" v-model="desc" class="repository_desc" @change="inputTeamDesc"></el-input>
                <label class="rep_auth">可见范围</label>
                <el-select v-model="auth"  class="repository_auth"  @change="selectAuth($event)">
                  <el-option
                  v-for="item in auths"
                             :key="item.id"
                             :label="item.name"
                             :value="item.id">
                  </el-option>
                </el-select>
                <el-button type="primary" @click="submitBtn">提交</el-button>
              </div>

            </el-col>
            <el-col :span="12"><div class="grid-content bg-purple-light"></div></el-col>
          </el-row>
        </el-card>
    </el-main>
  </div>
</template>

<script>
    import Header from '../headerAndFooter/Header'
    export default {
      name: 'NewRepository',
      created() {
        var groupid = this.$route.params.groupid
        this.getTeam()
        if (groupid) {
          this.team = parseInt(groupid)
        }
      },
      mounted() {
        // this.getTeam()
        // console.log('>>>',this.teams)
      },
      data() {
        return {
          teams:[
            { id:9999,name:'默认团队是自己' }
          ],
          team: 9999,
          desc:'',
          groupid:'',
          reponame:'',
          repotype: 1,
          auth:1,
          options:'',
          auths:[
            { id:0,name:'知识库仅自己和知识库成员可见' },
            { id:1,name:'所有人可见' }
          ]
        }
      },
      methods: {
        changeSelectStatus(sid) {
          this.repotype = sid
        },
        inputTeamName(value) {
          this.reponame = value
        },
        inputTeamDesc(value) {
          this.desc = value
        },
        submitBtn() {
          var that = this
          if (that.team === undefined) {
            that.team = 9999
          }
          const data = { reponame:that.reponame,desc:that.desc,team:that.team,auth:that.auth,repotype:that.repotype }
          console.log(data)
          that.$http.post('api/v1/repository',data).then(function (response) {
            if (response.data.code === 0) {
              that.$message.success('添加成功')
              setTimeout(function () {
                  that.$router.push('/workspace')
              },2000)
            }
          })
        },
        selectTeam(event) {
          this.team = event
        },
        getTeam() {
          var that = this
          that.$http.get('api/v1/team').then(function (response) {
              if (response.data.code === 0) {
                var teams = response.data.data
                that.teams = teams
              }
          })
        },
        selectAuth(event) {
          this.auth = event
        }
      },
      components: {
        'v-header':Header
      }
    }
</script>

<style scoped>
.newTitle{
  display: flex;
  flex-direction: row;
  align-items: center;
  height: 30px;
  margin-bottom: 10px;
}
.newTitle span:nth-child(1){
  margin-right: 16px;
  font-size: 16px;
  line-height: 24px;
  color: #262626;
}
.newTitle span:nth-child(2){
  margin-left: 3px;
  color: #8c8c8c;
  font-size: 14px;
}
  .left_new{
    display: flex;
    flex-direction: column;
    line-height: 15px;
  }
  .new_type{
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    margin-bottom: 10px;
  }
  .new_type_item{
    cursor: pointer;
  }
  .belong{
    padding: 0;
    margin:0;
  }
  .teamlist{
    margin-top: 10px;
  }
  .repository_name{
    margin-bottom: 10px;
  }
  .repository_desc,.repository_auth{
    margin-top: 10px;
    margin-bottom: 10px;
  }
  .rep_auth,.rep_desc{
    color: rgba(0,0,0,.85);
    font-size: 14px;
  }
  .main-container{
    margin-left: auto;
    margin-right: auto;
    padding: 24px 16px 32px;
    max-width: 1056px;
  }

  .active{
    background-color: #faf9fc;
  }
  .typeitem{
    height: 50px;
    width: 45%;
    display: flex;
    flex-direction: row;
    align-items: center;
    padding: 10px;
    border: 1px solid #DCDFE6;
  }
  .big-icon-40{
    font-size: 40px;
  }
</style>
