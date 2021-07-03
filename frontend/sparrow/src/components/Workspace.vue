<template>
  <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="文档" name="first">
      <template>
        <el-table
          ref="multipleTable"
          :data="tableData"
          tooltip-effect="dark"
          style="width: 100%">
          <el-table-column
            type="index"
            :index='(index)=>{return (index+1) + (page-1)*pageSize }'
            width="50">
          </el-table-column>
          <el-table-column
            prop="title"
            label="标题"
            show-overflow-tooltip
            width="250">
            <template slot-scope="scope">
               <span class="go-to-desc-title" @click="gotodoArtilceDesc(scope.row.ar_unique_code)">{{ scope.row.title }}</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="catename"
            label="分类"
            width="130">
          </el-table-column>
          <el-table-column
            prop="createtime"
            label="时间"
            sortable
            show-overflow-tooltip>
            <template slot-scope="scope">{{ scope.row.createtime }}</template>
          </el-table-column>
          <el-table-column label="操作" fixed="right" width='140' align="center">
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
      </template>
      <div class="pagination-block" v-if="total >0">
        <el-pagination
          background
          layout="prev, pager, next"
          :page-size="15"
          @current-change="changePage"
          :total="total">
        </el-pagination>
      </div>
    </el-tab-pane>
    <el-tab-pane label="资源" name="second">
      <template>
        <el-table
          ref="multipleTable"
          :data="fileTableData"
          tooltip-effect="dark"
          style="width: 100%">
          <el-table-column
            prop="filename"
            label="名称"
            width="300">
          </el-table-column>
          <el-table-column
            prop="repositoryname"
            label="分类"
            width="120">
          </el-table-column>
          <el-table-column
            prop="filetype"
            label="类型"
            width="120">
          </el-table-column>
          <el-table-column
            prop="createtime"
            label="时间"
            sortable
            show-overflow-tooltip>
            <template slot-scope="scope">{{ scope.row.createtime }}</template>
          </el-table-column>
          <el-table-column label="操作" fixed="right" width='140' align="center">
            <template slot-scope="scope">
              <el-dropdown  trigger="click">
                          <span class="el-dropdown-link el-button--lightblue dropbutton">
                                  <i class="el-icon-more"></i>
                          </span>

                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item @click.native="downloadFile(scope)">下载</el-dropdown-item>
                  <el-dropdown-item @click.native="delFile(scope.row.id)">删除</el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </template>
          </el-table-column>
        </el-table>
      </template>
    </el-tab-pane>
  </el-tabs>
</template>
<script>
    export default {
        name: 'Workspace',
      created() {
          this.getArticleList(this.page, this.pageSize)
          this.$store.commit('increment')
          console.log(this.$store.state.count)
      },
      data() {
          return {
            activeName: 'first',
            tableData: [],
            multipleSelection: [],
            total:0,
            fileTableData:[],
            page: 1,
            pageSize: 15
          }
      },
      methods: {
        handleClick(tab, event) {
          // console.log('>>>',tab, event)
          // console.log('>>>',tab.name)
          if (tab.name === 'second') {
            this.getFileResource()
          }
        },
        toggleSelection(rows) {
          if (rows) {
            rows.forEach(row => {
              this.$refs.multipleTable.toggleRowSelection(row)
            })
          } else {
            this.$refs.multipleTable.clearSelection()
          }
        },
        changePage(val) {
          this.page = val
          this.getArticleList(this.page, this.pageSize)
        },
        getArticleList(page, pageSize) {
          var that = this
          that.$http.get(`api/v1/self/article?page=${page}&size=${pageSize}`).then(function (respoonse) {
              if (respoonse.data.code === 0) {
                if (respoonse.data.data != null) {
                  that.tableData = respoonse.data.data
                  that.total = respoonse.data.count
                } else {
                  that.$message.warning('没有更多数据了')
                }
              }
          })
        },
        getFileResource() {
          var that = this
          that.$http.get('api/v1/file').then(function (respoonse) {
            if (respoonse.data.code === 0) {
              if (respoonse.data.data != null) {
                that.fileTableData = respoonse.data.data
                that.total = respoonse.data.count
              }
            }
          })
        },
        gotodoArtilceDesc(aid) {
          this.gotodoArtilceDescFun(aid)
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
                for (var i = 0; i < that.tableData.length; i++) {
                  if (that.tableData[i].ar_unique_code === uniqueCode) {
                    that.tableData.splice(that.tableData[i],1)
                  }
                }
              }
            })
          }).catch(() => {
            // 取消的时候，什么也不做
          })
        },
        downloadFile(scope) {
          // 下载文件
          var that = this
          that.$http.get(this.$http.defaults.baseURL + 'api/v1/download/' + scope.row.md5str, { responseType:'blob' }).then(function (response) {
            const { data,headers } = response
            const blob = new Blob([data], { type: headers['content-type'] })
            const dom = document.createElement('a')
            const url = window.URL.createObjectURL(blob)
            dom.href = url
            // dom.download = decodeURI(fileName)
            dom.download = decodeURI(scope.row.filename)
            dom.style.display = 'none'
            document.body.appendChild(dom)
            dom.click()
            dom.parentNode.removeChild(dom)
            window.URL.revokeObjectURL(url)
          })
          // window.location = this.$http.defaults.baseURL + 'api/v1/download/' + md5str
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
                    that.tableData.pop(that.tableData[i])
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
.display-file{
  cursor: pointer;
  color: #409EFF;
}
  .pagination-block{
    text-align: center;
    margin-top: 20px;
  }
</style>
