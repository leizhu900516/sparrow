<!--个人主页-->
<template>
  <el-container class="body-container">
    <Header></Header>
    <el-main  class="main-containers-custom" style="width: 100%">
      <el-row :gutter="20">
        <el-col :span="4">
          <ul class="syssetting-item">
            <li :class="{isActive:settingFlag == 1}" @click="changeSetting(1)">图书分类</li>
            <li :class="{isActive:settingFlag == 3}" @click="changeSetting(3)">图书管理</li>
            <li :class="{isActive:settingFlag == 2}" @click="changeSetting(2)">用户管理</li>
          </ul>
        </el-col>
        <el-col :span="20">
          <el-card shadow="never" v-if="settingFlag==1" >
            <div slot="header" class="book-header">
              <span>图书分类管理</span>
              <el-button type="primary" icon="el-icon-circle-plus-outline" size="small" @click="addBookCate">新建分类</el-button>
            </div>
            <div class="">
              <template>
                <el-table
                  :data="bookcate"
                  border
                  style="width: 100%">
                  <el-table-column
                    prop="id"
                    label="序号"
                    style="width: 10%">
                  </el-table-column>
                  <el-table-column
                    prop="catename"
                    label="分类名称"
                    style="width: 45%">
                  </el-table-column>
                  <el-table-column
                    prop="id"
                    label="操作"
                    style="width: 45%">
                    <template slot-scope="scope">
                      <el-button type="primary" icon="el-icon-edit" size="small" @click="editBookCate(scope.row)"></el-button>
                      <el-button type="danger" icon="el-icon-delete" size="small" @click="delBookCate(scope.row.id)"></el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </template>
            </div>
          </el-card>
          <el-card shadow="never" v-if="settingFlag==2" >
            <div slot="header">
              <span>用户管理</span>
            </div>
            <div class="">
              <span>等待开发...</span>
            </div>
          </el-card>
          <el-card shadow="never" v-if="settingFlag==3" >
            <div slot="header">
              <span>图书管理</span>
            </div>
            <div class="">
              <span>等待开发...</span>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </el-main>
    <Footer></Footer>

    <el-dialog title="添加图书分类"
               :visible.sync="dialogFormWithAddBookCate"
               :before-close="closeDialogForm">
      <el-form :model="bookCateForm">
        <el-form-item label="分类名称" :label-width="formLabelWidth">
          <el-input v-model="bookCateForm.name" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancelBookCateForm">取 消</el-button>
        <el-button type="primary" @click="submitAddBookCateForm">确 定</el-button>
      </div>
    </el-dialog>
  </el-container>
</template>

<script>
  import Header from './headerAndFooter/Header'
  import Footer from './headerAndFooter/Footer'

  export default {
    name: 'Syssetting',
    components: {
      Header,
      Footer
    },
    data() {
      return {
        disabled:true,
        bookcate: [],
        updateBookCateId: '',
        formLabelWidth: '120px',
        bookCateForm: {
          name: ''
        },
        dialogFormWithAddBookCate: false,
        settingFlag:1
      }
    },
    created() {
      this.getBookCateList()
    },
    methods: {
      changeSetting(flag) {
        this.settingFlag = flag
      },
      cancelBookCateForm() {
        this.dialogFormWithAddBookCate = false
        // 讲更新的分类id标识清空
        this.updateBookCateId = ''
      },
      addBookCate() {
        this.dialogFormWithAddBookCate = true
      },
      closeDialogForm() {
        this.updateBookCateId = ''
        this.dialogFormWithAddBookCate = false
      },
      submitAddBookCateForm() {
        var that = this
        var data = {}
        if (that.updateBookCateId !== '') {
          data = {
            name:that.bookCateForm.name,
            id:that.updateBookCateId
          }
        } else {
          data = {
            name:that.bookCateForm.name
          }
        }
        that.$http.post('api/v1/book/cate/',data).then(function (respoonse) {
          if (respoonse.data.code === 0) {
            that.$message.success(respoonse.data.msg)
            that.dialogFormWithAddBookCate = false
            that.getBookCateList()
            // 讲更新的分类id标识清空
            that.updateBookCateId = ''
          } else {
            that.$message.error(respoonse.data.msg)
          }
        })
      },
      editBookCate(row) {
        // 添加更新的分类id标识
        this.updateBookCateId = row.id
        this.bookCateForm.name = row.catename
        this.dialogFormWithAddBookCate = true
      },
      delBookCate(id) {
        var that = this
        that.$confirm('删除不可恢复，是否确定删除?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          that.$http.delete('api/v1/book/cate/' + id).then(function (respoonse) {
            if (respoonse.data.code === 0) {
              that.$message.success(respoonse.data.msg)
            } else {
              that.$message.error(respoonse.data.msg)
            }
          })
        }).catch(() => {
          // 取消的时候，什么也不做
        })
      },
      getBookCateList() {
        var that = this
        that.$http.get('api/v1/book/cate')
          .then(function (response) {
            if (response.data.code === 0) {
              that.bookcate = response.data.data
            } else {
              that.$message.error(response.data.msg)
            }
          })
      }
    }
  }
</script>

<style scoped>
  >>> .el-input{
    width: 90%;
  }
  >>> .el-textarea{
    width: 50%;
  }
  .main-containers-custom{
    margin-left: auto;
    margin-right: auto;
    max-width: 1056px;
  }
  .syssetting-item{
    border: 1px solid #dcdfe6;
  }
  .syssetting-item li{
    border-bottom: 1px solid #dcdfe6;
    padding: 10px;
    font-size: 15px;
  }
  .syssetting-item li:nth-last-child(1){
    border-bottom: none;
  }
  .isActive {
    background-color: #F2F6FC;
    color: #409EFF;
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
  .book-header{
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }
</style>
