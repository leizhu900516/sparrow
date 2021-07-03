<template>
  <el-container class="body-container">
    <Header></Header>
    <el-header >
      <div class="wiki-for-me main-containers-custom">
        <span v-bind:class="{ active: isActive === 1 }" @click="changeIsActie(1)">推荐</span>
        <span v-bind:class="{ active: isActive === 2 }" @click="changeIsActie(2)">关注</span>
      </div>
    </el-header>
    <el-main :body-style="{ padding:'10px' }">
      <div class="main-containers-custom">
        <el-row :gutter="12" v-if="isActive === 1">
          <el-col :span="18">
            <span v-if="wikiloading === true">正在加载...</span>
            <ul v-loading="wikiloading">
              <li class="wiki-card" v-for="article in articlelist" v-bind:key="article.id">
                <el-row :gutter="12">
                  <el-col :span="2" >
                    <div class="goodicon">
                      <span class="iconfont iconzan goodicon-can-click" @click="goodHandle(article)" v-if="article.isGood === false"></span>
                      <span class="iconfont iconzan goodicon-not-click" v-if="article.isGood === true"></span>
                      <span>{{article.liked}}</span>
                    </div>
                  </el-col>
                  <el-col :span="18">
                    <div >
                      <h4 class="title" @click="gotoArticleDesc(article.ar_unique_code)"> {{article.title}}</h4>
                      <p class="article-desc">{{article.desc}}</p>
                      <div class="article-info">
                        <el-popover
                          placement="top-start"
                          width="250"
                          trigger="hover"
                          content="这是一段内容,这是一段内容,这是一段内容,这是一段内容。">
                          <div class="userinfo-popover">
                            <div class="userinfo-popover-avator-username">
                              <img :src="article.user.avatarurl" class="samll-avator">
                              <div>
                                <span class="username">{{article.user.username}}</span>
                                <p class="user-desc">{{article.user.profile}}</p>
                              </div>
                            </div>
                            <div class="userinfo-follow-info">
                              <div>
                                <span>关注者</span>
                                <span class="follow-number">{{article.user.beFollowCount}}</span>
                              </div>
                              <div>
                                <span>关注了</span>
                                <span class="follow-number">{{article.user.followCount}}</span>
                              </div>
                              <el-button type="primary" size="mini" @click="followAuthor(article)" v-if="article.isFollow === false">关注</el-button>
                              <el-button type="info" size="mini" plain disabled v-if="article.isFollow === true">已关注</el-button>
                            </div>
                          </div>
                          <span class="username" slot="reference">{{article.username}}</span>
                        </el-popover>
                        <span>发布于</span>
                        <span>{{article.catename}}</span>
                        <span>{{article.createtime}}</span>
                      </div>
                    </div>
                  </el-col>
<!--                  <el-col :span="4">这是图片区域</el-col>-->
                </el-row>
              </li>
            </ul>
          </el-col>
          <el-col :span="6">
            <el-card shadow="never">
              <div slot="header" class="clearfix">
                <span>热门文章</span>
<!--                <span class="more-btn" type="text">更多</span>-->
              </div>
              <div v-for="(hot,index) in hotArticlelist" :key="hot.id" class="text item" v-loading="hotArticleLoading">
                <div class="hot-article-title" @click="gotoArticleDesc(hot.ar_unique_code)">
                  <span>{{index+1}}</span>
                  <span>{{hot.title }}</span>
                </div>
              </div>
            </el-card>
            <el-card shadow="never" style="margin-top: 20px">
              <div slot="header" class="clearfix">
                <span>热门知识库</span>
<!--                <span class="more-btn" type="text">更多</span>-->
              </div>
              <div v-for="(repo,index) in hotRepositoryList" :key="repo.id" class="text item" v-loading="hotRepoLoading">
                <div class="hot-article-title" @click="gotoRepoWorkHome(repo.repo_unique_code)">
                  <span>{{index+1}}</span>
                  <span>{{repo.name}}</span>
                </div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="24">
            <div class="load-more" v-if="isMore === true">
              <el-button @click="loadMore">加载更多...</el-button>
            </div>
          </el-col>
        </el-row >
        <el-row :gutter="12" v-if="isActive === 2">
          <el-col :span="18">
            <el-card shadow="never" class="change-card">
              <div slot="header">
                <span>关注动态</span>
              </div>
              <span v-if="moments.length === 0">暂时还没有动态</span>
               <div v-for="m in moments" v-bind:key="m.unique_code" class="follow-moments">
                 <div class="author-info" v-if="m.flag === 'user'">
                   <img :src="m.avator !== '' ? $http.defaults.baseURL+m.avator:avatorUrl" class="smallavatar">
                   <span>{{m.username}}</span>
                 </div>
                 <div class="author-info" v-if="m.flag === 'repo'">
                   <span>{{m.username}}</span>
                   <span>{{m.avator}}</span>
                 </div>
                 <div>
                   <span class="follow_title_createtime">{{m.createtime}} 发布了</span>
                   <span class="follow_title" @click="gotoArticleDesc(m.unique_code)">{{m.title}}</span>
                 </div>
               </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="never">
              <div slot="header" class="clearfix">
                <span>关注的人</span>
                <span class="more-btn" type="text"></span>
              </div>
              <div v-for="(hot,index) in hotArticlelist" :key="hot.id" class="text item">
                <div class="hot-article-title">
                  <span>{{index+1}}</span>
                  <span>{{hot.title }}</span>
                </div>
              </div>
            </el-card>
            <el-card shadow="never" style="margin-top: 20px">
              <div slot="header" class="clearfix">
                <span>关注的知识库</span>
                <span class="more-btn" type="text"></span>
              </div>
              <div v-for="(repo,index) in hotRepositoryList" :key="repo.id" class="text item">
                <div class="hot-article-title">
                  <span>{{index+1}}</span>
                  <span>{{repo.name}}</span>
                </div>
              </div>
            </el-card>
          </el-col>
        </el-row >
      </div>
    </el-main>
    <Footer></Footer>
  </el-container>
</template>

<script>
    import Header from './headerAndFooter/Header'
    import Footer from './headerAndFooter/Footer'
    export default {
      name: 'Wiki',
      data() {
        return {
          avatorUrl : require('@/assets/imgs/defaultAvator.png'),
          articlelist: [],
          hotArticlelist: [],
          hotRepositoryList: [],
          isActive: 1,
          page: 1,
          isMore: true,
          moments: [],
          wikiloading: true,
          hotArticleLoading: true,
          hotRepoLoading: true
        }
      },
      components: {
        Header,
        Footer
      },
      created() {
        this.getArticleList(1)
        this.getArticleHot()
        this.getRepositoryList()
      },
      methods: {
        getArticleList(page) {
          var that = this
          that.$http.get('api/v1/article?page=' + page).then(function (response) {
            if (response.data.code === 0) {
              if (response.data.data === null) {
                console.log('no data')
                that.isMore = false
                that.$message.warning('没有更多数据了')
              } else {
                for (var i = 0; i < response.data.data.length; i++) {
                  var article = response.data.data[i]
                  if (article.user.avatarurl.indexOf('null') === -1) {
                    console.log('存在',article.user.avatarurl)
                    article.user.avatarurl = that.$http.defaults.baseURL + article.user.avatarurl
                  } else {
                    console.log('不存在')
                    article.user.avatarurl = require('@/assets/imgs/defaultAvator.png')
                  }
                  that.articlelist.push(article)
                }
                that.wikiloading = false
              }
            }
          })
        },
        getArticleHot() {
          var that = this
          that.$http.get('api/v1/article?flag=hot&page=1&limit=7').then(function (response) {
            if (response.data.code === 0) {
              that.hotArticlelist = response.data.data
              that.hotArticleLoading = false
            }
          })
        },
        getRepositoryList() {
          var that = this
          that.$http.get('api/v1/repository?flag=hot&page=1&limit=7').then(function (response) {
            if (response.data.code === 0) {
              that.hotRepositoryList = response.data.data
              that.hotRepoLoading = false
            }
          })
        },
        gotoArticleDesc(uniquecode) {
          this.gotodoArtilceDescFun(uniquecode)
        },
        followAuthor(article) {
          var that = this
          that.$http.post('api/v1/follow',{ flag:'user',id:article.userid }).then(function (response) {
            if (response.data.code === 0) {
              article.isFollow = true
              that.$message.success(response.data.msg)
            } else {
              that.$message.error(response.data.msg)
            }
          })
        },
        changeIsActie(status) {
          this.isActive = status
          if (status === 2) {
            this.getFollow()
          }
        },
        getFollow() {
          var that = this
          that.$http.get('api/v1/follow/moments').then(function (response) {
            console.log(response)
            if (response.data.code === 0) {
              that.moments = response.data.data
            } else {
            }
          })
        },
        loadMore() {
          this.page += 1
          console.log(this.page)
          this.getArticleList(this.page)
        },
        gotoRepoWorkHome(respUniqueCode) {
          this.$router.push({ name:'repoflag', params: { username:'chenhuachao',repoflag:respUniqueCode } })
        },
        goodHandle(article) {
          var that = this
          that.$http.post('api/v1/good',{ uniquecode:article.ar_unique_code }).then(function (response) {
              if (response.data.code === 0) {
                that.$message.success('sccess')
                article.liked += 1
                article.isGood = true
              } else if (response.data.code === 1) {
                that.$message.error(response.data.msg)
              }
          })
        }
      }
    }
</script>

<style scoped>
.wiki-for-me{
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 100%;
}
.wiki-for-me span{
  margin-right: 10px;
  font-size: 16px;
  color: #606266;
}
.main-containers-custom{
  margin-left: auto;
  margin-right: auto;
  padding: 20px 16px 32px;
  max-width: 1056px;
}
  .wiki-card{
    border: 1px solid #EBEEF5;
    padding: 5px;
  }
  .goodicon{
    margin-top: 20px;
    margin-left: 13px;
    width: 38px;
    display: flex;
    flex-direction: column;
    align-items: center;
    background-color: #F2F6FC;
    padding: 5px 0px 5px 0px;
  }
  .goodicon-can-click{
    cursor: pointer;
    font-size: 25px;
    font-weight: 100;
  }
  .goodicon-not-click{
    cursor:not-allowed;
    color: #C0C4CC;
    font-size: 25px;
  }
  .article-desc{
    font-size: 14px;
    color: #8c8c8c;
  }
  .article-info{
    font-size: 12px;
    color: #8c8c8c;
  }
  .article-info span{
    margin-right: 10px;
  }
  .title{
    cursor: pointer;
    font-size: 18px;
    color: #262626;
  }
  .active{
    font-weight: bold;
  }
  .load-more{
    text-align: center;
    margin-top: 20px;
  }
  .more-btn{
    cursor: pointer;
    color: #409EFF;
    font-size: 14px;
    float: right; padding: 3px 0
  }
  .hot-article-title{
    cursor: pointer;
    line-height: 27px;
  }
  .hot-article-title span:nth-child(1){
    font-weight: bold;
    margin-right: 5px;
  }
  .change-card{
    height: 500px;
  }
  .username{
    color: #303133;
    font-size: 16px;
  }
  .userinfo-popover{
    display: flex;
    flex-direction: column;
  }
  .userinfo-popover-avator-username{
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    border-bottom: 1px solid #EBEEF5;
  }
  .userinfo-follow-info{
    margin-top: 10px;
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    align-items: center;
  }
.userinfo-follow-info span{
  margin-left: 3px;
}
  .user-desc{
    white-space:nowrap;
    overflow:hidden;
    text-overflow:ellipsis;
  }
  .follow-number{
    color: #409EFF;
  }
  .samll-avator{
    height: 40px;
    width: 40px;
    border-radius: 20px;
    border-color: rgba(0, 0, 0, 0.06);
    border-width: 1px;
    border-style: solid;
  }
  .follow-moments{
    border-radius: 2px
  }
  .author-info{
    display: flex;
    flex-direction: row;
    align-content: center;
    align-items: center;
    margin-bottom: 5px;
  }
  .author-info span{
    margin-left: 5px;
  }
  .follow_title_createtime{
    color: #909399;
    font-size: 13px;
    cursor: pointer;
  }
  .follow_title{
    color: #303133;
    font-size: 14px;
    margin-left: 10px;
  }
</style>
