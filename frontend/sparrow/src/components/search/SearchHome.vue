<template>
  <el-container class="body-container">
    <Header ref="Header"></Header>
    <div class="header-second" >
      <div class="main-container">
        <ul class="doc-items">
          <li :class="{isActive:active==1}" @click="changeDocType(1)">内容</li>
          <li :class="{isActive:active==2}" @click="changeDocType(2)">知识库</li>
          <li :class="{isActive:active==3}" @click="changeDocType(3)">团队</li>
          <li :class="{isActive:active==4}" @click="changeDocType(4)">用户</li>
        </ul>
      </div>
    </div>
    <el-main :body-style="{ padding:'10px' }">
      <div class="main-container" v-if="active === 1">
        <span class="result-count">找到 {{dataCount}} 个结果</span>
        <ul class="resource-item" >
          <li v-for="doc in dataList" v-bind:key="doc.uniqueCode">
              <i class="el-icon-document resource-icon"></i>
              <div>
                <span class="h3title" @click="gotoArticleDesc(doc.uniqueCode)">{{doc.title}}</span>
                <p class="resource-desc">{{doc.desc}}</p>
                <div class="resource-info">
                  <span>{{doc.repoName}} </span>
                  <span> / </span>
                  <span>{{doc.createtime}}</span>
                </div>
              </div>
          </li>
        </ul>
      </div>
      <div class="main-container" v-if="active === 2">
        <span class="result-count">找到 {{dataCount}} 个结果</span>
        <ul class="resource-item" >
          <li v-for="doc in dataList" v-bind:key="doc.uniqueCode">
              <i class="el-icon-document resource-icon"></i>
              <div>
                <span class="h3title" @click="gotoRepoWorkHome(doc.uniqueCode)">{{doc.title}}</span>
                <p class="resource-desc">{{doc.desc}}</p>
                <div class="resource-info">
                  <span>{{doc.createtime}}</span>
                </div>
              </div>
          </li>
        </ul>
      </div>
      <div class="main-container" v-if="active === 3">
        <span class="result-count">找到 {{dataCount}} 个结果</span>
        <ul class="resource-item" >
          <li v-for="doc in dataList" v-bind:key="doc.uniqueCode">
              <i class="el-icon-document resource-icon"></i>
              <div>
                <span class="h3title">{{doc.title}}</span>
                <p class="resource-desc">{{doc.desc}}</p>
                <div class="resource-info">
                  <span>{{doc.createtime}}</span>
                </div>
              </div>
          </li>
        </ul>
      </div>
      <div class="main-container" v-if="active === 4">
        <span class="result-count">找到 {{dataCount}} 个结果</span>
        <ul class="resource-item" >
          <li v-for="doc in dataList" v-bind:key="doc.uniqueCode">
            <div>
              <div class="user-info">
                <img :src="$http.defaults.baseURL + 'api/v1/img/' + doc.avatarmd5" class="smallavatar">
                <span>{{doc.username}}</span>
              </div>
              <span>{{doc.profile}}</span>
            </div>
          </li>
        </ul>
      </div>
      <div class="page-div" v-if="dataCount > 0">
        <el-pagination
          background
          layout="prev, pager, next"
          :total="dataCount">
        </el-pagination>
      </div>
    </el-main>
    <Footer></Footer>
  </el-container>
</template>

<script>
    import Header from '../headerAndFooter/Header'
    import Footer from '../headerAndFooter/Footer'
    export default {
      name: 'SearchHome',
      components: {
        Header,
        Footer
      },

      created() {
        const kw = this.$route.query.kw
        if (kw !== '') {
          this.kw = kw
          this.getSearchData(kw, 1)
        } else {
          this.$message.error('请输入要搜索的内容')
        }
      },
      watch:{
        $route(to,from) {
          this.$router.go(0)
        }
      },
      data() {
        return {
          activeName: 'second',
          kw: '',
          active:1,
          dataList: [],
          dataCount: 0
        }
      },
      methods: {
        handleClick(tab, event) {
          console.log(tab, event)
        },
        gotoRepoWorkHome(respUniqueCode) {
          this.$router.push({ name:'repoflag', params: { username:'chenhuachao',repoflag:respUniqueCode } })
        },
        changeDocType(id) {
          this.active = id
          this.dataList = []
          this.getSearchData(this.kw, id)
        },
        getSearchData(kw, id) {
          const that = this
          that.$http.get(`/api/v1/search?flag=${id}&kw=${kw}`).then(function(response) {
            console.log(response)
            if (response.data.code === 0) {
              if ('doc' in response.data.data) {
                that.dataList = response.data.data.doc
                that.dataCount = response.data.count
              }
              if ('user' in response.data.data) {
                console.log('user')
                that.dataList = response.data.data.user
                that.dataCount = response.data.count
              }
            }
          })
        },
        gotoArticleDesc(uniquecode) {
          this.gotodoArtilceDescFun(uniquecode)
        }
      }

    }
</script>

<style scoped>
  .body-container{
    background: #fafafa;
    height: 100%;
  }
  .header-second{
    border-bottom: 1px solid #dcdfe6;
    height: 50px;
  }
  .main-container{
    margin-left: auto;
    margin-right: auto;
    max-width: 1056px;
    width: 100%;
    height: 100%;
  }
  .doc-items{
    height: 50px;
    display: flex;
    flex-direction: row;
    align-items: center;
  }
  .doc-items li{
    margin-right: 20px;
    padding: 14px;
  }
  .isActive{
    border-bottom: 2px solid #409EFF;
  }
  .result-count{
    font-size: 14px;
    color: #8c8c8c;
  }
  .resource-item{
    margin-top: 10px;
  }
  .resource-item li{
    display: flex;
    flex-direction: row;
    padding: 15px;
    border-bottom: 1px solid #f0f0f0;
    background: #fff;
  }
  .h3title{
    font-size: 18px;
    font-weight: bold;
    cursor: pointer;
  }
  .resource-icon{
    margin-right: 15px;
    color: #409EFF;
    font-size: 30px;
  }
  .resource-desc{
    display: -webkit-box;
    max-height: 36px;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
    line-height: 18px;
    font-size: 14px;
    margin-bottom: 8px;
  }
  .resource-info{
    line-height: 18px;
    font-size: 12px;
    color: #8c8c8c;
  }
  .user-info{
    display: flex;
    flex-direction: row;
    align-items: center;
  }
  .user-info span{
    margin-left: 10px;
  }
  .page-div{
    text-align: center;
  }
</style>
