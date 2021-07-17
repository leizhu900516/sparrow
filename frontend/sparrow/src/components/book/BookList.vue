<template>
  <el-container class="main-container">
    <el-row class="bookcate" :gutter="20">
      <el-col :span="5">
        <el-select v-model="bookcate"
                   @change="selectBookCate($event)"
                   placeholder="请选择分类">
          <el-option
            v-for="item in options"
            :key="item.id"
            :label="item.catename"
            :value="item.id">
          </el-option>
        </el-select>
      </el-col>
      <el-col :span="5">
        <el-input v-model="searchkeyword" placeholder="搜索" prefix-icon="el-icon-search" @blur="searchKwBook"></el-input>
      </el-col>
      <el-col :span="5">
        <span class="downloadhot" @click="downloadRankHandle">下载最多</span>
      </el-col>
      <el-col :span="5"></el-col>
      <el-col :span="4" class="upload_btn">
        <el-button type="primary" @click="gotoUpload()"><i class="el-icon-upload"></i>上传图书</el-button>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="4" v-for="book in booklist" v-bind:key="book.id">
        <el-card :body-style="{ padding: '0px' }">
          <img :src="book.avatorurl" class="image">
          <div class="book-item">
            <el-tooltip class="item" effect="dark" :content="book.name" placement="top-start">
              <span class="book-name">{{ book.name }}</span>
            </el-tooltip>
            <div class="bottom clearfix">
              <i class="el-icon-download">{{ book.download }}</i>
              <el-button type="text" class="button" @click="downloadBook(book.id)"><i class="el-icon-download"></i>下载</el-button>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </el-container>
</template>

<script>
  export default {
    name: 'BookList',
    created() {
      this.getBookCate()
      this.getBookList('all','')
    },
    data() {
      return {
        searchkeyword: '',
        bookcate: '', // book分类
        booklist: [], // 图书列表
        options: []
      }
    },
    methods:{
      gotoUpload() {
        this.$router.push('/book/upload')
      },
      downloadBook(id) {
        window.location = window.location.origin + '/api/v1/book/download/' + id
      },
      getBookCate() {
        var that = this
        that.$http.get('/api/v1/book/cate').then(function (response) {
          if (response.data.code === 0) {
            that.options = response.data.data
          }
        })
      },
      getBookList(flag,kw,hot) {
        // 获取图书列表
        var that = this
        that.$http.get('/api/v1/books?flag=' + flag + '&kw=' + kw + '&rank=' + hot).then(function (response) {
          console.log('>>>>',response.data.data)
          that.booklist = []

          if (response.data.code === 0) {
            if (response.data.data === null) {
              that.$message.warning('没有找到相关图书')
              return
            }
            that.booklist = response.data.data
            if (that.booklist !== null) {
              that.booklist.forEach((item,index,arr) =>{
                item.avatorurl = window.location.origin + '/' + item.avatorurl
              })
            }
          } else {
            that.$message.error(response.data.msg)
          }
        })
      },
      selectBookCate(event) {
        this.bookcate = event
        // 获取图书列表
        this.getBookList(event,'','')
      },
      searchKwBook() {
        this.getBookList('all',this.searchkeyword,'')
      },
      downloadRankHandle() {
        this.getBookList('all',this.searchkeyword,'hot')
      }
    }
  }
</script>

<style scoped>
  .main-container{
    margin-left: auto;
    margin-right: auto;
    padding: 24px 16px 32px;
    max-width: 1056px;
    display: flex;
    flex-direction: column;
  }
  .bookcate{
    margin-bottom: 10px;
    height: 41px;
    display: flex;
    flex-direction: row;
    align-content: center;
    align-items: center;
    color: #595959;
  }
  /*卡片样式*/
  .time {
    font-size: 13px;
    color: #999;
  }
  .bottom {
    margin-top: 13px;
    line-height: 12px;
  }

  .button {
    padding: 0;
    float: right;
  }
  .image {
    width: 100%;
    display: block;
    height: 200px;
  }
  .clearfix:before,
  .clearfix:after {
    display: table;
    content: "";
  }
  .clearfix:after {
    clear: both
  }
  .downloadhot{
    cursor: pointer;
    color:#909399;
  }
  .downloadhot:hover{
    color: #262626;
  }
  .upload_btn{
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
  }
  .book-name{
    overflow: hidden;
    text-overflow:ellipsis;
    white-space: nowrap;
  }
  .book-item{
    padding: 10px;font-size: 14px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
</style>
