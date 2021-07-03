<template>
  <el-container >
    <el-header>
      <div>
        <el-breadcrumb separator="/">
          <el-breadcrumb-item :to="{ path: '/' }">{{ reponame }}</el-breadcrumb-item>
          <el-breadcrumb-item>
            <span>{{title}}</span>
          </el-breadcrumb-item>
        </el-breadcrumb>
      </div>
      <span class="editstatus">{{ title }}</span>
      <div class="edit-icon-btn">
        <i class="el-icon-star-off collect-icon" @click="collectSubmit(docCode)" v-if="isCollect === false && isauthor === false"></i>
        <el-tooltip class="item" effect="dark" content="已经收藏过" placement="top" v-if="isCollect === true">
          <i class="el-icon-star-on collect-icon" ></i>
        </el-tooltip>
        <el-button type="primary" @click="modifyArticle(docCode)" size="mini" v-if="isauthor == true">继续编辑</el-button>
      </div>
    </el-header>
    <el-main class="main-containers">
      <div class="ArticleDetail">
        <div ref="editor" style="text-align:left" v-html="content"></div>
      </div>
    </el-main>
    <Footer></Footer>
  </el-container>
</template>

<script>
  import Footer from '../headerAndFooter/Footer'
  export default {
    name: 'ArticleDesc',
    components:{
      Footer
    },
    created() {
      var docCode = this.$route.params.docCode
      this.docCode = docCode
      this.getArticleDesc(docCode)
      console.log('>>>',docCode)
    },
    data: function () {
      return {
        reponame:'',
        repocode:'',
        content:'',
        title:'',
        docCode:'',
        isCollect:false,
        isauthor: false
      }
    },
    methods:{
      getArticleDesc(docCode) {
        var that = this
        if (that.docCode) {
          that.$http.get('api/v1/article/' + docCode).then(function (response) {
            console.log(response)
            if (response.data.code === 0) {
              that.content = response.data.data.content
              that.title = response.data.data.title
              that.reponame = response.data.data.reponame
              that.repocode = response.data.data.repocode
              that.isauthor = response.data.data.isauthor
              that.isCollect = response.data.data.iscollect
            }
          })
        }
      },
      modifyArticle(docCode) {
        console.log('修改文章')
        this.$router.push({ name: 'wikiedit', params: { docCode:docCode } })
      },
      // 收藏按钮
      collectSubmit(aid) {
        var that = this
        that.$http.post('api/v1/collect',{ docCode:that.docCode,resourcetype:1 }).then(function (response) {
            that.$message.success(response.data.msg)
          if (response.data.code === 0) {
            that.isCollect = true
          }
        })
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
  .el-header{
    background-color: #fff;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #dcdfe6;
    align-content: center;
  }
  .el-footer{
    justify-content: center;
    display: flex;
    align-items: center;
    flex-direction: row;
    color: #595959;
    font-size: 14px;
    width: 100%;
  }
  .editstatus{
    font-size: 14px;
    color: #595959;
  }
  .el-breadcrumb{
    display: flex;
    flex-direction: row;
    align-items: center;
  }
  .collect-icon{
    font-size: 30px;
    color: #409EFF;
    margin-right: 10px;
    cursor: pointer;
  }
  .edit-icon-btn{
    display: flex;
    flex-direction: row;
    align-items: center;
  }
</style>
