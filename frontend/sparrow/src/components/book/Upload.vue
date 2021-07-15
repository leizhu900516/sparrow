<template>
  <div class="main-containers">
    <div class="newTitle">
      <h3>上传图书</h3>
      <span>创作、管理各种类型的知识</span>
    </div>
    <el-card class="box-card" shadow="never">
      <el-row>
        <el-col :span="12">
          <div class="left_new">
            <span class="belong">属于</span>
            <el-select v-model="bookcate"
                       @change="selectBookCate($event)"
                       placeholder="请选择分类"
                       class="teamlist">
              <el-option
                v-for="item in options"
                :key="item.id"
                :label="item.catename"
                :value="item.id">
              </el-option>
            </el-select>
            <el-divider></el-divider>
            <el-input placeholder="请输入图书名称"
                      class="repository_name"
                      @change="changefilename"
                      v-model="fileName"></el-input>
            <span class="belong">上传封面</span>
            <el-upload
              action=""
              class="avatar-uploader"
              :show-file-list="false"
              :before-upload="beforeAvatarUpload"
              :auto-upload="false"
              :on-change="imgBroadcastChange" >
              <img v-if="imageUrl" :src="imageUrl" class="avatar">
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
            </el-upload>
            <span class="belong">上传图书</span>
            <el-upload
              class="file-uploader"
              ref="upload"
              action=""
              :file-list="fileList"
              :on-change="fileBroadcastChange"
              :auto-upload="false">
              <el-button slot="trigger" size="small" type="success">选取图书文件</el-button>
            </el-upload>
            <el-button type="primary" @click="startUpload()">提交</el-button>
          </div>

        </el-col>
        <el-col :span="12"><div class="grid-content bg-purple-light"></div></el-col>
      </el-row>
    </el-card>
  </div>
</template>

<script>
  import { Loading } from 'element-ui'
    export default {
      name: 'Upload',
      created() {
        this.getBookCate()
      },
      data() {
          return {
            imageUrl: '',
            bookcate: '',
            options:[],
            uploadapi:this.$http.defaults.baseURL + 'api/v1/upload',
            fileName:'',
            avator:'',
            file:'',
            fileList:[]
          }
      },
      methods: {
        getBookCate() {
          var that = this
          that.$http.get('/api/v1/book/cate').then(function (response) {
            console.log('>>>>',response.data.data)
            if (response.data.code === 0) {
              that.options = response.data.data
            }
          })
        },
      selectBookCate(event) {
        this.bookcate = event
      },
        handleAvatarSuccess(res, file) {
          this.imageUrl = URL.createObjectURL(file.raw)
        },
        beforeAvatarUpload(file) {
          const isJPG = file.type === 'image/jpeg' || file.type === 'image/png'
          const isLt2M = file.size / 1024 / 1024 < 2

          if (!isJPG) {
            this.$message.error('上传头像图片只能是 JPG 或者PNG 格式!')
          }
          if (!isLt2M) {
            this.$message.error('上传头像图片大小不能超过 2MB!')
          }
          return isJPG && isLt2M
        },
        // 图书封面信息获取
        imgBroadcastChange(file, fileList) {
          this.imageUrl = URL.createObjectURL(file.raw)
          this.avator = file.raw
          if (!this.fileName) {
            this.fileName = file.name
          }
        },
        // 文件信息获取
        fileBroadcastChange(file, fileList) {
          this.file = file.raw
          if (!this.fileName) {
            this.fileName = file.name
          }
        },
        uploadSuccess(res, file) {
        },
        startUpload() {
          const formData = new FormData()
          formData.append('filename', this.fileName)
          formData.append('filecate', this.bookcate)
          formData.append('avator', this.avator)
          formData.append('file', this.file)
          // console.log(formData)
          const loadingInstance = Loading.service({ fullscreen: true })
          this.$http.post(window.location.origin + '/' + 'api/v1/book/upload', formData)
            .then(su => {
              if (su.data.code === 0) {
                this.$message.success(su.data.msg)
                setTimeout(this.$router.push('/book'), 1500)
              } else if (su.data.code === 403) {
                this.$router.push('/login')
                this.$message.error(su.data.msg)
              } else {
                this.$message.error(su.data.msg)
              }
            })
            .catch(err => {
              console.log(err)
            })
          loadingInstance.close()
        },
        handleChange(file, fileList) {
          this.fileList = fileList.slice(-3)
        },
        changefilename(value) {
          this.fileName = value
        }

      }
    }
</script>

<style scoped>
  .main-containers{
    margin-left: auto;
    margin-right: auto;
    padding: 1px 16px 32px;
    max-width: 1056px;
    display: flex;
    flex-direction: column;
  }
  .newTitle{
    display: flex;
    flex-direction: row;
    align-content: center;
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
  .avatar-uploader{
    margin-bottom: 10px;
    margin-top: 10px;
  }
  .repository_desc,.repository_auth,.file-uploader{
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
  /*上传图片*/
  .avatar-uploader .el-upload {  /* 不起作用，放到global.css中控制 */
    /*border: 1px dashed #d9d9d9!important;*/
    /*border-radius: 6px;*/
    /*cursor: pointer;*/
    /*position: relative;*/
    /*overflow: hidden;*/
    /*margin-bottom: 10px;*/
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 100px;
    height: 100px;
    line-height: 100px;
    text-align: center;
  }
  .avatar {
    width: 100px;
    height: 100px;
    display: block;
  }
</style>
