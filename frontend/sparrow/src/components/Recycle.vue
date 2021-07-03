<!--回收站-->
<template>
  <div>
    <el-breadcrumb separator=" ">
      <el-breadcrumb-item :to="{ path: '/' }">回收站</el-breadcrumb-item>
    </el-breadcrumb>
    <el-table
      ref="multipleTable"
      :data="tableData"
      tooltip-effect="dark"
      style="width: 100%">
      <el-table-column
        prop="title"
        label="标题"
        width="250">
        <template slot-scope="scope">
          <span class="go-to-desc-title" @click="gotodoArtilceDesc(scope.row.id,scope.row.catename,scope.row.title)">{{ scope.row.title }}</span>
        </template>
      </el-table-column>
      <el-table-column
        prop="catename"
        label="分类"
        width="120">
      </el-table-column>
      <el-table-column
        prop="deletetime"
        label="删除时间"
        sortable
        show-overflow-tooltip>
        <template slot-scope="scope">{{ scope.row.deletetime }}</template>
      </el-table-column>
      <el-table-column label="操作" fixed="right" width='140' align="center">
        <template slot-scope="scope">
          <span @click="recover(scope.row.id,scope.row.resource_unique_code)" class="recover-delete">恢复</span>
          <span @click="deletesure(scope.row.id)" class="recover-delete">彻底删除</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
  export default {
    name: 'Recycle',
    created() {
      this.getArticleList()
    },
    data() {
      return {
        activeName: 'first',
        multipleSelection: [],
        tableData: []
      }
    },
    methods: {
      getArticleList() {
        var that = this
        that.$http.get('api/v1/recycle').then(function (respoonse) {
          if (respoonse.data.code === 0) {
            that.tableData = respoonse.data.data
          }
        })
      },
      recover(rid,aid) {
        var that = this
        that.$confirm('是否恢复数据?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          that.$http.put('api/v1/article/' + aid + '?rid=' + rid).then(function (respoonse) {
            if (respoonse.data.code === 0) {
              that.$message.success('恢复成功')
              that.tableData.forEach((item, index, arr) => {
                if (item.resource_unique_code === aid) {
                  that.tableData.splice(item,1)
                }
              })
            }
          })
        }).catch(() => {
          // 取消的时候，什么也不做
        })
      },
      deletesure(rid) {
        var that = this
        that.$confirm('是否彻底删除数据?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          that.$http.delete('api/v1/recycle/' + rid).then(function (respoonse) {
            if (respoonse.data.code === 0) {
              that.$message.success('删除成功!')
              for (var i = 0; i < that.tableData.length; i++) {
                if (that.tableData[i].id === rid) {
                  that.tableData.splice(that.tableData[i],1)
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
  .recover-delete{
    cursor: pointer;
    color: #409EFF;
    margin-right: 10px;
  }
</style>
