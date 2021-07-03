<!--收藏夹-->
<template>
  <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="收藏夹" name="first">
      <el-table
        ref="multipleTable"
        :data="tableData"
        tooltip-effect="dark"
        style="width: 100%">
        <el-table-column
          prop="resourcename"
          label="标题"
          width="250">
          <template slot-scope="scope">
            <span class="go-to-desc-title" @click="gotodoArtilceDesc(scope.row.id,scope.row.catename,scope.row.title)">{{ scope.row.resourcename }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="catename"
          label="分类"
          width="120">
        </el-table-column>
        <el-table-column
          prop="collecttime"
          label="收藏时间"
          sortable
          show-overflow-tooltip>
          <template slot-scope="scope">{{ scope.row.collecttime }}</template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" width='140' align="center">
          <template slot-scope="scope">
            <span @click="cancelCollectSure(scope.row.id)" class="recover-delete">取消收藏</span>
          </template>
        </el-table-column>
      </el-table>
    </el-tab-pane>
    <el-tab-pane label="文档" name="second">
      <el-table
        ref="multipleTable"
        :data="tableData"
        tooltip-effect="dark"
        style="width: 100%">
        <el-table-column
          prop="resourcename"
          label="标题"
          width="250">
          <template slot-scope="scope">
            <span class="go-to-desc-title" @click="gotodoArtilceDesc(scope.row.id,scope.row.catename,scope.row.title)">{{ scope.row.resourcename }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="catename"
          label="分类"
          width="120">
        </el-table-column>
        <el-table-column
          prop="collecttime"
          label="收藏时间"
          sortable
          show-overflow-tooltip>
          <template slot-scope="scope">{{ scope.row.collecttime }}</template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" width='140' align="center">
          <template slot-scope="scope">
            <span @click="cancelCollectSure(scope.row.id)" class="recover-delete">取消收藏</span>
          </template>
        </el-table-column>
      </el-table>
    </el-tab-pane>
    <el-tab-pane label="资源" name="third">
      <el-table
        ref="multipleTable"
        :data="tableData"
        tooltip-effect="dark"
        style="width: 100%">
        <el-table-column
          prop="resourcename"
          label="标题"
          width="250">
          <template slot-scope="scope">
            <span class="go-to-desc-title" @click="gotodoArtilceDesc(scope.row.id,scope.row.catename,scope.row.title)">{{ scope.row.resourcename }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="catename"
          label="分类"
          width="120">
        </el-table-column>
        <el-table-column
          prop="collecttime"
          label="收藏时间"
          sortable
          show-overflow-tooltip>
          <template slot-scope="scope">{{ scope.row.collecttime }}</template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" width='140' align="center">
          <template slot-scope="scope">
            <span @click="cancelCollectSure(scope.row.id)" class="recover-delete">取消收藏</span>
          </template>
        </el-table-column>
      </el-table>
    </el-tab-pane>
    <el-tab-pane label="团队" name="fourth">
      <el-table
        ref="multipleTable"
        :data="tableData"
        tooltip-effect="dark"
        style="width: 100%">
        <el-table-column
          prop="resourcename"
          label="标题"
          width="250">
          <template slot-scope="scope">
            <span class="go-to-desc-title" @click="gotodoArtilceDesc(scope.row.id,scope.row.catename,scope.row.title)">{{ scope.row.resourcename }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="catename"
          label="分类"
          width="120">
        </el-table-column>
        <el-table-column
          prop="collecttime"
          label="收藏时间"
          sortable
          show-overflow-tooltip>
          <template slot-scope="scope">{{ scope.row.collecttime }}</template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" width='140' align="center">
          <template slot-scope="scope">
            <span @click="cancelCollectSure(scope.row.id)" class="recover-delete">取消收藏</span>
          </template>
        </el-table-column>
      </el-table>
    </el-tab-pane>
    <el-tab-pane label="知识库" name="fifth">
      <el-table
        ref="multipleTable"
        :data="tableData"
        tooltip-effect="dark"
        style="width: 100%">
        <el-table-column
          prop="resourcename"
          label="标题"
          width="250">
          <template slot-scope="scope">
            <span class="go-to-desc-title" @click="gotodoArtilceDesc(scope.row.id,scope.row.catename,scope.row.title)">{{ scope.row.resourcename }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="catename"
          label="分类"
          width="120">
        </el-table-column>
        <el-table-column
          prop="collecttime"
          label="收藏时间"
          sortable
          show-overflow-tooltip>
          <template slot-scope="scope">{{ scope.row.collecttime }}</template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" width='140' align="center">
          <template slot-scope="scope">
            <span @click="cancelCollectSure(scope.row.id)" class="recover-delete">取消收藏</span>
          </template>
        </el-table-column>
      </el-table>
    </el-tab-pane>
  </el-tabs>
</template>

<script>
  export default {
    name: 'Collect',
    created() {
      this.getCollectList()
    },
    data() {
      return {
        activeName: 'first',
        tableData: [],
        collectType:{
          first: 'all',
          second: 'doc',
          third: 'file',
          fourth: 'team',
          fifth: 'repository'
        }
      }
    },
    methods: {
      cancelCollectSure(id) {
        var that = this
        that.$confirm('是否取消收藏?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          that.$http.delete('api/v1/collect/' + id).then(function (respoonse) {
            if (respoonse.data.code === 0) {
              that.$message.success(respoonse.data.msg)
              for (var i = 0; i < that.tableData.length; i++) {
                console.log('>>>>>',that.tableData[i].id,id)
                if (that.tableData[i].id === id) {
                  that.tableData.splice(that.tableData[i],1)
                }
              }
            } else {
              that.$message.error(respoonse.data.msg)
            }
          })
        }).catch(() => {
          // 取消的时候，什么也不做
        })
      },
      getCollectList() {
        var that = this
        that.$http.get('api/v1/collect').then(function (respoonse) {
          if (respoonse.data.code === 0) {
            that.tableData = respoonse.data.data
          }
        })
      },
      handleClick(tab, event) {
        var that = this
        that.activeName = tab.$props.name
        that.$http.get('api/v1/collect?flag=' + that.collectType[that.activeName]).then(function (respoonse) {
          if (respoonse.data.code === 0) {
            that.tableData = respoonse.data.data
          }
        })
      },
      gotodoArtilceDesc(aid,catename,title) {
        console.log('去详情页了')
        // this.$router.push('/article')
        this.$router.push({ name: 'article', params: { aid:aid,title: title, repname:catename } })
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
  .collect-type{
    display: flex;
    flex-direction: row;
    font-size: 16px;
  }
  .collect-type li{
    margin-right: 15px;
  }
  .isActive{
    font-weight: bold;
  }
</style>
