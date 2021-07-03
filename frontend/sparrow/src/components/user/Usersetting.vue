<!--个人主页-->
<template>
  <el-container class="body-container">
    <Header></Header>
    <el-main  class="main-containers-custom" style="width: 100%">
      <el-row :gutter="20">
        <el-col :span="6">
          <ul class="usersetting-item">
            <li :class="{isActive:settingFlag == 1}" @click="changeSetting(1)">个人信息</li>
            <li :class="{isActive:settingFlag == 2}" @click="changeSetting(2)">账户设置</li>
          </ul>
        </el-col>
        <el-col :span="18">
          <el-card shadow="never" v-if="settingFlag==1" >
            <div slot="header">
              <span>个人信息</span>
            </div>
            <div class="userinfo-setting">
              <span>昵称</span>
              <el-input v-model="username" :placeholder="userinfo.username" :disabled="disabled"></el-input>
              <span>简介</span>
              <el-input type="textarea"
                        :rows="2"
                        :placeholder="userinfo.desc"
                        v-model="userdesc"></el-input>
              <span>头像</span>
              <div class="upload-avator">
                <img class="smallavatar" :src="userinfo.avatarMd5 !== 'null' ? userinfo.avatarurl:avatorUrl">
                <el-upload
                  class="upload-demo"
                  :action="uploadApi"
                  :data="uploadAvatorParams"
                  :headers="headers"
                  :on-change="handleChange"
                  :on-success="uploadAvatorHandle">
                  <el-button size="small" type="primary">点击上传</el-button>
                </el-upload>
              </div>
              <el-button  size="small" type="primary" @click="updateUserinfo()" style="width: 50%">更新信息</el-button>

            </div>
          </el-card>
          <el-card shadow="never" v-if="settingFlag==2" >
            <div slot="header">
              <span>账户设置</span>
            </div>
            <div class="userinfo-setting">
              <span>更改密码</span>
              <el-input v-model="oldpasswd" placeholder="请输原密码" show-password></el-input>
              <span>新密码</span>
              <el-input v-model="newpasswd1" placeholder="请输入新密码" show-password></el-input>
              <el-input v-model="newpasswd2" placeholder="请再次输入新密码" show-password></el-input>
              <el-button  size="small" type="primary" @click="updateUserPasswd()" style="width: 50%">更新信息</el-button>
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
    name: 'Usersetting',
    components: {
      Header,
      Footer
    },
    data() {
      return {
        username:'',
        userdesc:'',
        settingFlag:1,
        oldpasswd:'',
        newpasswd1:'',
        newpasswd2:'',
        userinfo:'',
        avatorUrl : require('@/assets/imgs/defaultAvator.png'),
        avatorMd5: '',
        uploadApi: '',
        uploadAvatorParams: { repo_unique_code:'avator',file_dir_level:0 },
        headers: { authorization: '' }
      }
    },
    created() {
      this.getUserinfo()
      this.uploadApi = this.$http.defaults.baseURL + 'api/v1/upload'
      this.headers = { authorization: window.localStorage.getItem('token') }
    },
    methods: {
      changeSetting(flag) {
        this.settingFlag = flag
      },
      updateUserinfo() {
        var that = this
        // 更新用户信息
        that.$http.post('api/v1/user',{ userdesc:that.userdesc,avatarurl:that.avatorMd5 }).then(
          function (response) {
            if (response.data.code === 0) {
              that.$message.success(response.data.msg)
            } else {
              that.$message.error(response.data.msg)
            }
          }
        )
      },
      updateUserPasswd() {
        var that = this
        if (that.newpasswd2 === '' || that.newpasswd1 === '' || that.oldpasswd === '') {
          that.$message.error('密码不能为空')
          return false
        }
        if (that.newpasswd1 !== that.newpasswd2) {
          that.$message.error('两次密码不相同')
          return false
        }
        that.$http.post('api/v1/user/passwd',{ old_passwd:that.oldpasswd,new_passwd_one:that.newpasswd1,new_passwd_two:that.newpasswd2 })
        .then(function (response) {
          if (response.data.code === 0) {
            that.$message.success(response.data.msg)
          } else {
            that.$message.error(response.data.msg)
          }
        })
      },
      getUserinfo() {
        var that = this
        that.$http.get('api/v1/user').then(function (response) {
          if (response.data.code === 0) {
            if (response.data.data != null) {
              var data = response.data.data
              that.userinfo = data
              that.avatorMd5 = data.avatarMd5
              that.userinfo.avatarurl = that.$http.defaults.baseURL + that.userinfo.avatarurl
            }
          }
        })
      },
      handleChange(file, fileList) {
        this.fileList = fileList.slice(-3)
      },
      uploadAvatorHandle(response,file,fileList) {
        if (response.code === 0) {
          this.avatorUrl = window.location.origin + '/' + response.data.url
          this.avatorMd5 = response.data.md5
        }
      }
    }
  }
</script>

<style scoped>
  >>> .el-input{
    width: 50%;
  }
  >>> .el-textarea{
    width: 50%;
  }
  .main-containers-custom{
    margin-left: auto;
    margin-right: auto;
    max-width: 1056px;
  }
  .usersetting-item{
    border: 1px solid #dcdfe6;
  }
  .usersetting-item li{
    border-bottom: 1px solid #dcdfe6;
    padding: 10px;
  }
  .usersetting-item li:nth-last-child(1){
    border-bottom: none;
  }
  .isActive {
    border-left: 3px solid #409EFF;
  }
  .smallavatar{
    height: 40px;
    width: 40px;
    border-radius: 20px;
    border-color: rgba(0, 0, 0, 0.06);
    border-width: 1px;
    border-style: solid;
  }
  .userinfo-setting{
    line-height: 50px;
    display: flex;
    flex-direction: column;
  }
  .upload-avator{
    display: flex;
    flex-direction: row;
  }
</style>
