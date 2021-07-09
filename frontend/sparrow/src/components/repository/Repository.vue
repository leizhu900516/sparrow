<template>
    <div>
      <div v-if="cateid === '1' ">
        <el-header >
          <div class="wiki-for-me main-containers-custom">
            <div class="user-repo-follow">
              <span>{{username}}</span>
              <span>/</span>
              <span>{{respname}}</span>
              <div v-if="isSelf === false">
                <el-button size="mini"  type="primary" v-if="followFlag === true" :disabled="true">已关注</el-button>
                <el-button size="mini"  type="primary" v-if="followFlag === false" @click="handleFollow">关注</el-button>
              </div>

            </div>
            <div>
              <el-button @click="gotoRepository">新建</el-button>
            </div>
          </div>
        </el-header>
        <el-main :body-style="{ padding:'10px' }" class="main-repo-container" >
          <div class="main-containers-custom">
            <div class="repo-list">
              <div class="title-and-search">
                <span class="resource_title">资源</span>
                <el-input v-model="searchKw" placeholder="搜索"  @keyup.enter.native="searchRepo"></el-input>
              </div>
              <el-table
                :data="doclist"
                style="width: 100%"
                :default-sort = "{prop: 'date', order: 'descending'}"
              >
                <el-table-column
                  prop="title"
                  label="名称"
                  sortable
                  width="300">
                  <template slot-scope="scope">
                    <span class="go-to-desc-title" @click="gotodoArtilceDesc(scope.row.ar_unique_code)">{{ scope.row.title }}</span>
                  </template>
                </el-table-column>
                <el-table-column
                  prop="createtime"
                  label="日期"
                  sortable
                  width="200">
                </el-table-column>
                <el-table-column
                  prop="username"
                  label="作者"
                  width="180">
                </el-table-column>
                <el-table-column
                  prop="address"
                  fixed="right"
                  label="操作">
                  <template slot-scope="scope">
                    <el-dropdown  trigger="click">
                          <span class="el-dropdown-link el-button--lightblue dropbutton">
                                  <i class="el-icon-more"></i>
                          </span>
                      <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item @click.native="modifyArticle(scope.row.ar_unique_code)">修改</el-dropdown-item>
                        <el-dropdown-item @click.native="delArticle(scope.row.ar_unique_code)">删除</el-dropdown-item>
                      </el-dropdown-menu>
                    </el-dropdown>
                  </template>
                </el-table-column>
              </el-table>
            </div>
          </div>
        </el-main>
      </div>
      <div v-if="cateid === '2' ">
        <el-header >
          <div class="wiki-for-me main-containers-custom">
            <div>
              <span>{{username}}</span>
              <span>/</span>
              <span>{{respname}}</span>
            </div>
            <div class="mkdir-upload">
              <el-button @click="mkdirdir = true">新建文件夹</el-button>
              <el-upload
                class="upload-demo"
                :action="uploadUrl"
                :on-success="uploadSuccess"
                multiple
                :headers="headers"
                :data="formData"
                :limit="3"
                :show-file-list="false">
                <el-button>上传</el-button>
              </el-upload>
            </div>
          </div>
        </el-header>
        <el-main :body-style="{ padding:'10px' }" class="main-repo-container">
          <div class="main-containers-custom">
            <div class="repo-list">
              <div class="title-and-search">
                <span class="resource_title">资源</span>
                <el-input v-model="searchKw" placeholder="搜索" @keyup.enter.native="searchRepo"></el-input>
              </div>
              <el-table
                :data="doclist"
                style="width: 100%"
                :default-sort = "{prop: 'date', order: 'descending'}"
              >
                <el-table-column
                  prop="filename"
                  label="名称"
                  sortable
                  width="300">
                  <template slot-scope="scope">
                    <span v-if="scope.row.filetype === 'dir'" class="filedir-stype" @click="getRepositoryDirList(respUniqueCode,scope.row.id)">
                      <span class="iconfont iconwenjianjia"></span>{{ scope.row.filename }}
                    </span>
                    <span v-if="scope.row.filetype === 'file'" class="file-stype" >
                      {{ scope.row.filename }}
                    </span>
                  </template>
                </el-table-column>
                <el-table-column
                  prop="fileext"
                  label="类型"
                  sortable
                  width="180">
                </el-table-column>
                <el-table-column
                  prop="size"
                  label="大小"
                  sortable
                  width="100">
                </el-table-column>
                <el-table-column
                  prop="createtime"
                  label="日期"
                  sortable
                  width="180">
                </el-table-column>
                <el-table-column
                  prop="username"
                  label="上传人"
                  sortable
                  width="180">
                </el-table-column>
                <el-table-column
                  prop="address"
                  fixed="right"
                  label="操作">
                  <template slot-scope="scope" v-if="scope.row.filetype === 'file'">
                    <el-dropdown  trigger="click">
                          <span class="el-dropdown-link el-button--lightblue dropbutton">
                                  <i class="el-icon-more"></i>
                          </span>

                      <el-dropdown-menu slot="dropdown" >
                        <el-dropdown-item @click.native="downloadFile(scope.row.md5Str)" >下载</el-dropdown-item>
                        <el-dropdown-item @click.native="delFile(scope.row.id)" >删除</el-dropdown-item>
                      </el-dropdown-menu>
                    </el-dropdown>
                  </template>
                </el-table-column>
              </el-table>
            </div>
          </div>
        </el-main>

      </div>

      <div class="pagination-block" v-if="total >0">
        <el-pagination
          background
          layout="prev, pager, next"
          :page-size="10"
          @current-change="changePage"
          :total="total">
        </el-pagination>
      </div>
      <!--新建文档弹出框-->
      <el-dialog title="新建文件夹" :visible.sync="mkdirdir" width="33%" >
        <div class="demo-input-suffix">
          文件夹：
          <el-input
            placeholder="请输入文件夹名称"
            v-model="dirname">
          </el-input>
          <el-button type="primary" @click="mkdirdirbtn">确定</el-button>
        </div>
      </el-dialog>
    </div>
</template>

<script>
    export default {
      name: 'Repository',
      data() {
          return {
            followFlag: true,
            isSelf: true,
            total:0,
            mkdirdir: false,
            dirname:'',
            respname:'',
            searchKw:'',
            catename:'',
            respUniqueCode:'',
            rid:'',
            username:'',
            cateid:'', // 知识库类型 1文档 2 文件
            doclist:[],
            file_dir_level:0,
            uploadUrl:'api/v1/upload',
            formData:{ file_dir_level:this.file_dir_level, repo_unique_code:this.respUniqueCode },
            headers: { authorization: '' },
            page: 1,
            size: 15,
            fromPage: '' // 上一个页面
          }
      },
      watch:{
        cateid(id) {
        }
      },
      created() {
        var repoflag = this.$route.params.repoflag
        this.respUniqueCode = repoflag
        this.formData = { file_dir_level:this.file_dir_level, repo_unique_code:this.respUniqueCode }
        this.getRepositoryInfo(repoflag)
        this.headers = { authorization: window.localStorage.getItem('token') }
      },
      methods:{
        handleFollow() {
          var that = this
          that.$http.post('api/v1/follow',{ flag:'repo',id:that.respUniqueCode }).then(function (response) {
            if (response.data.code === 0) {
              that.followFlag = true
              that.$message.success(response.data.msg)
            } else {
              that.$message.error(response.data.msg)
            }
          })
        },
        modifyArticle(aid) {
          this.modifyArticleFun(aid)
        },
        delArticle(uniqueCode) {
          var that = this
          that.$confirm('此操作将永久删除, 是否继续?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            that.$http.delete('api/v1/article/' + uniqueCode).then(function (response) {
              if (response.data.code === 0) {
                that.$message.success({
                  type: 'success',
                  message: '删除成功!'
                })
                for (var i = 0; i < that.doclist.length; i++) {
                  if (that.doclist[i].ar_unique_code === uniqueCode) {
                    that.doclist.splice(that.doclist[i],1)
                  }
                }
              }
            })
          }).catch(() => {
            // 取消的时候，什么也不做
          })
        },
        changePage(val) {
          this.page = val
          this.getRepositoryList(this.respUniqueCode,this.file_dir_level,this.page,this.size,'')
        },
        // 搜索接口
        searchRepo() {
          this.getRepositoryList(this.respUniqueCode,this.file_dir_level,this.page,this.size,this.searchKw)
        },
        getRepositoryInfo(repoflag) {
          var that = this
          that.$http.get('api/v1/repositoryinfo/' + repoflag).then(function (response) {
            if (response.data.code === 0) {
              that.respname = response.data.data.resp_name
              that.rid = response.data.data.id
              that.cateid = response.data.data.resp_cate
              that.username = response.data.data.username
              that.isSelf = response.data.data.isAuthorOrMember
              that.followFlag = response.data.data.follow
              that.getRepositoryList(that.respUniqueCode,that.file_dir_level,that.page,that.size,'')
            } else if (response.data.code === 2) {
              that.$message.error(response.data.msg)
              setTimeout(function() {
                that.$router.push(that.fromPage)
              },2000)
            } else {
              that.$message.error(response.data.msg)
            }
          })
        },
        getRepositoryDirList(repoflag,id) {
          // 进入文件夹,查看文件列表
          this.file_dir_level = id
          this.getRepositoryList(repoflag,id,this.page,this.size,'')
          this.formData = { file_dir_level:this.file_dir_level, repo_unique_code:this.respUniqueCode }
        },
        getRepositoryList(repoflag,pathid,page,size,kw) {
          var that = this
          that.file_dir_level = pathid
          that.$http.get(`api/v1/repository/${repoflag}?file_dir_level=${pathid}&page=${page}&size=${size}&kw=${kw}`).then(function (response) {
              if (response.data.code === 0) {
                that.doclist = response.data.data
                that.total = response.data.count
              }
          })
        },
        gotodoArtilceDesc(aid) {
          this.gotodoArtilceDescFun(aid)
        },
        downloadFile(md5str) {
          // 下载文件
          window.location = this.$http.defaults.baseURL + 'api/v1/download/' + md5str
        },
        delFile(fileid) {
          //    删除文件
          var that = this
          that.$confirm('此操作将永久删除, 是否继续?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            that.$http.delete('api/v1/file/' + fileid).then(function (response) {
              if (response.data.code === 0) {
                that.$message.success({
                  type: 'success',
                  message: '删除成功!'
                })
                for (var i = 0; i < that.tableData.length; i++) {
                  if (that.tableData[i].id === fileid) {
                    that.tableData.splice(that.tableData[i],1)
                  }
                }
              }
            })
          }).catch(() => {
            // 取消的时候，什么也不做
          })
        },
        mkdirdirbtn() {
          var that = this
          var Data = new Date()
          var nowdata = Data.getFullYear() + '-' + (Data.getMonth() + 1) + '-' + Data.getDate() + ' ' + Data.getHours() + ':' + Data.getMinutes()
          var data = { filedir:that.dirname,repositoryid:that.respUniqueCode,parentid:that.file_dir_level }
          that.$http.post('api/v1/file/dir',data)
          .then(function (response) {
              if (response.data.code === 0) {
                that.$message.success(response.data.msg)
                that.doclist.push({
                  filename: that.dirname,
                  createtime: nowdata,
                  filetype: 'dir'
                })
              } else {
                that.$message.error(response.data.msg)
              }
          })
          that.mkdirdir = false
        },
        uploadSuccess(response,file,fileList) {
          if (response.code === 0) {
            var Data = new Date()
            var nowdata = Data.getFullYear() + '-' + (Data.getMonth() + 1) + '-' + Data.getDate() + ' ' + Data.getHours() + ':' + Data.getMinutes()
            this.$message.success('上传成功')
            this.doclist.push({
              filename: file.name,
              createtime: nowdata,
              filetype: 'file'
            })
          } else {
            this.$message.warning(response.msg)
          }
        },
        gotoRepository() {
          this.$router.push({ name: 'wikiedit', params: { docCode:'luckid' ,repid: this.respUniqueCode, repname:this.respname } })
        }
      },
      beforeRouteEnter (to, from, next) {
        next(vm => {
          // beforeRouteEnter不能通过this访问组件实例，但是可以通过 vm 访问组件实例
          vm.fromPage = from.fullPath
        })
      }
    }
</script>

<style scoped>
  .demo-input-suffix{
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
  }
  >>> .el-main{
   padding: 0px!important;
  }
  .wiki-for-me{
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }
  .wiki-for-me span{
    margin-right: 10px;
    font-size: 14px;
    color: #606266;
  }
  .main-containers-custom{
    margin-left: auto;
    padding: 0px!important;
    margin-right: auto;
    max-width: 1056px;
  }
  .user-repo-follow{
    display: flex;
    flex-direction: row;
    align-items: center;
  }
.repo-header{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  border-bottom: 1px solid #dcdfe6;
  height: 70px;
  align-items: center;
}
  .title-and-search{
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }
  .resource_title{
    color: #606266;
    font-size: 14px;
  }
  .repo-list{
    margin-top: 10px;
  }
  .el-input{
    width: 200px;
  }
  .filedir-stype{
    color: #409EFF;
    cursor: pointer;
  }
  .mkdir-upload{
    display: flex;
    flex-direction: row;
  }
  .mkdir-upload button:nth-child(1){
    margin-right: 5px;
  }
</style>
