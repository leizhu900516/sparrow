<template>
  <el-container class="body-container">
    <Header></Header>
    <el-main  class="main-containers-custom" style="width: 100%">
      <div class="header-title">
        <h3>新建团队</h3>
        <span>和团队成员一起编写文档、交流想法、沉淀经验</span>
      </div>
      <el-card shadow="never">
        <el-row>
          <el-col :span="12">
            <div class="team-item">
              <span>名称</span>
              <el-input v-model="teamname" placeholder="团队名称"></el-input>
            </div>
            <div class="team-item">
              <span>简介</span>
              <el-input
                type="textarea"
                :rows="2"
                placeholder="请输入团队简介"
                v-model="desc">
              </el-input>
            </div>
            <div class="team-item">
              <span>可见范围</span>
              <el-select v-model="auth" placeholder="请选择" >
                <el-option
                  v-for="item in options"
                  :key="item.id"
                  :label="item.name"
                  :value="item.id">
                </el-option>
              </el-select>
            </div>
            <div class="team-item-img">
              <span>头像</span>
              <el-upload
                class="upload-demo"
                :action="uploadApi"
                :data="uploadAvatorParams"
                :headers="headers"
                :on-change="handleChange"
                :on-success="uploadAvatorHandle"
                drag>
                <i class="el-icon-upload"></i>
                <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
              </el-upload>
              <el-button type="primary" @click="addTeam">新建团队</el-button>
            </div>
          </el-col>
        </el-row>
      </el-card>
    </el-main>
    <Footer></Footer>
  </el-container>
</template>

<script>
  import Header from '../headerAndFooter/Header'
  import Footer from '../headerAndFooter/Footer'

  export default {
    name: 'NewTeam',
    components: {
      Header,
      Footer
    },
    data() {
      return {
        desc: '',
        teamname:'',
        auth:1,
        avatorUrl : require('@/assets/imgs/defaultAvator.png'),
        avatorMd5: '',
        avator:'',
        options: [{ id:1,name:'仅团队成员可见' },
          { id:2,name:'所有人可见' }],
        uploadApi: '',
        uploadAvatorParams: { repo_unique_code:'avator',file_dir_level:0 },
        headers: { authorization: '' }
      }
    },
    created() {
      this.uploadApi = this.$http.defaults.baseURL + 'api/v1/upload'
      this.headers = { authorization: window.sessionStorage.getItem('token') }
    },
    methods: {
      addTeam() {
        var that = this
        var data = { name:that.teamname,desc:that.desc,auth:that.auth,avator: that.avatorMd5 }
        that.$http.post('api/v1/team',data).then(function (response) {
              console.log(response.data)
              that.$message.success(response.data.msg)
              if (response.data.code === 0) {
                setTimeout(function () {
                  that.$router.push('/team/' + response.data.data)
                },2000)
              }
        })
      },
      handleChange(file, fileList) {
        this.fileList = fileList.slice(-3)
      },
      uploadAvatorHandle(response,file,fileList) {
        if (response.code === 0) {
          this.avatorUrl = this.$http.defaults.baseURL + response.data.url
          this.avatorMd5 = response.data.md5
          this.$message.success('success')
        } else {
          this.$message.error(response.msg)
        }
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
  .header-title{
    display: flex;
    flex-direction: row;
    height: 60px;
    align-items: center;
  }
  .header-title span{
    font-size: 16px;
    color: #8c8c8c;
    margin-left: 10px;
  }
  .team-item{
    display: flex;
    flex-direction: column;
    line-height: 35px;
    font-size: 14px;
    margin-bottom: 5px;
    color: #8c8c8c;
  }
  .team-item-img{
    display: flex;
    flex-direction: column;
    line-height: 35px;
    font-size: 14px;
    margin-bottom: 5px;
    color: #8c8c8c;
  }
</style>
