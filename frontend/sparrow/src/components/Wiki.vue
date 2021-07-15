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
            <div class="no-data-show" v-if="articlelist.length === 0">
              <svg t="1626330460185" class="icon" viewBox="0 0 1576 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2510" width="128" height="128"><path d="M1374.43519999 856.3904h-143.79733333l10.0128-16.74346667c6.73066667-11.81866667 10.01386667-25.28 10.01386666-39.56053333V145.94133333c0-21.0112-8.37226667-40.38186667-22.6528-55.48373333-14.2816-14.28053333-34.47253333-22.6528-55.48373333-22.6528H634.76586666c-21.0112 0-40.3808 8.37226667-55.48266667 22.6528-15.10186667 15.10186667-22.6528 34.47253333-22.6528 55.48373333v34.47146667h-78.13653333c-21.0112 0-40.38186667 8.37226667-55.48373334 22.65386667-15.10186667 15.10186667-23.47306667 34.47146667-23.47306666 55.48266666v653.16053334c0 13.46133333 3.2832 26.92053333 10.0128 39.56053333l10.01386666 16.74346667H298.58133332c-5.90933333 0-10.99733333 5.088-10.99733333 10.99733333 0 2.464 0.82133333 5.9104 3.44746667 7.552 1.6416 1.6416 5.088 3.2832 7.55093333 3.2832h569.11466667l57.9456 57.9456c5.90933333 5.90933333 15.10186667 10.0128 23.47306666 10.0128 8.37226667 0 17.728-3.2832 23.47413334-10.0128 10.99733333-10.9984 12.64-27.74186667 5.088-41.20213333l-10.0128-17.728h137.888c5.90933333 0 10.99733333-5.08906667 10.99733333-10.9984s-5.088-10.9984-10.99733333-10.9984h-32.01066667l10.01386667-16.74346667c6.73066667-11.81866667 10.01386667-25.28 10.01386666-39.56053333v-34.47253334h280.69973334c5.90933333 0 10.99733333-5.088 10.99733333-10.99733333 0.16426667-4.43306667-4.92373333-8.7008-10.83306667-8.7008z m-302.53226667 55.48373333c0 31.02506667-25.28 56.304-56.304 56.304h-75.67466666l-78.13546667-78.9568c-8.37226667-8.37226667-21.01226667-11.81866667-32.01066667-8.37226666l-6.7296 1.64266666-32.82986666-32.832 5.90933333-7.54986666c45.46986667-61.39306667 36.11306667-147.9008-21.01226667-198.45973334-57.12426667-50.39466667-144.6176-47.11146667-198.45866666 7.55093334-53.84213333 53.84106667-56.30506667 141.17013333-6.73066667 199.28 49.57333333 57.12533333 136.24533333 67.3024 196.65386667 21.83253333l7.552-5.90933333 32.82986666 32.82986666-1.64266666 6.73066667c-3.28213333 11.81866667 0 24.45866667 8.37333333 32.82986667l31.024 31.02506666H477.50933332c-31.02506667 0-56.30506667-25.28-56.30506666-56.304V257.89333333c0-30.2048 24.45866667-55.48373333 54.66346666-55.48373333h541.3728c30.20373333 0 54.6624 25.28 54.6624 55.48373333v653.9808z m-305.1584-69.76426666c-24.45866667 24.45866667-55.48373333 36.11306667-87.49333333 36.11306666s-63.0336-11.81866667-87.49226667-36.11306666c-23.30986667-23.14666667-36.27733333-54.66346667-36.11413333-87.49333334 0-32.83093333 12.64-64.67626667 36.11413333-87.49333333 23.47306667-23.47306667 54.6624-36.11306667 87.49226667-36.11306667s63.85493333 12.64 87.49333333 36.11306667c23.47413333 23.47413333 36.11306667 54.6624 36.11306667 87.49333333s-12.64 64.0192-36.11306667 87.49333334z m462.25066667-42.84373334c0.00106667 31.024-25.27786667 56.304-56.30293333 56.304h-78.9568V258.7136c0-21.0112-8.37226667-40.3808-22.65386667-55.48373333-14.28053333-14.28053333-34.47146667-22.6528-55.48266667-22.6528H578.46186666v-33.6512c0-31.024 25.28-56.304 56.304-56.304h538.09066666c31.024 0 56.304 25.28 56.304 56.304v652.34026666zM208.46506666 968.9984H74.02346666c-5.90933333 0-10.99733333 5.08906667-10.99733334 10.9984 0 3.2832 0.82026667 5.90933333 3.28213334 7.552 1.6416 1.64053333 5.088 3.28213333 7.552 3.28213333H208.29866666c5.90933333 0 10.99733333-5.088 10.99733333-10.99733333s-4.1024-10.8352-10.83306667-10.8352z m0 0" p-id="2511" fill="#e6e6e6"></path><path d="M544.81066666 315.01866667h246.39253333c5.90933333 0 10.99733333-5.09013333 10.99733333-10.99946667s-5.088-10.99733333-10.99733333-10.99733333H544.81066666c-5.90933333 0-10.99733333 5.088-10.99733334 10.99733333 0 2.46293333 0.82133333 5.90933333 3.2832 7.552 1.8048 2.62613333 5.088 3.4464 7.71413334 3.4464z m359.00053333 90.77546666H544.81066666c-5.90933333 0-10.99733333 5.088-10.99733334 10.99733334 0 3.2832 0.82133333 5.90933333 3.2832 7.552 1.64053333 1.64053333 5.088 3.28213333 7.55093334 3.28213333h358.99946666c5.90933333 0 10.9984-5.088 10.9984-10.99733333 0-5.74613333-5.088-10.83413333-10.83413333-10.83413334z m-201.9072 112.608H544.64746666c-5.90933333 0-10.9984 5.088-10.9984 10.99733334 0 3.2832 0.82133333 5.90933333 3.2832 7.552 1.6416 1.6416 5.088 3.28213333 7.55093333 3.28213333h157.2576c5.90933333 0 10.99733333-5.088 10.99733333-10.99733333s-4.92373333-10.83413333-10.83413333-10.83413334zM197.62986666 720.96533333h22.6528c5.90933333 0 10.9984 5.088 10.9984 10.99733334s-5.088 10.99946667-10.99733334 10.99946666h-22.65386666v22.6528c0 5.90933333-5.088 10.99733333-10.99733334 10.99733334-3.28426667 0-5.90933333-0.82026667-7.552-3.28213334-2.46186667-1.64266667-3.2832-5.08906667-3.2832-7.552V743.12533333h-22.816c-5.90933333 0-10.9984-5.088-10.9984-10.99733333s5.088-10.9984 10.99733334-10.9984h22.6528v-22.6528c0-5.90933333 5.08906667-10.9984 10.9984-10.9984s10.9984 5.088 10.9984 10.99733333v22.4896z m1479.5008-90.77546666v-22.65386667c0-3.4464-0.82133333-5.90933333-3.2832-7.55093333-2.46186667-1.6416-5.088-3.2832-7.55093334-3.2832-5.90933333 0-10.99733333 5.08906667-10.99733333 10.9984v22.6528h-22.65386667c-3.28213333 0-5.90933333 0.82133333-7.55093333 3.44746666-1.6416 2.46293333-3.4464 5.088-3.4464 7.552 0 5.90933333 5.088 10.99733333 10.99733333 10.99733334h22.65386667v22.6528c0 5.90933333 5.088 10.99733333 10.99733333 10.99733333s10.9984-5.088 10.9984-10.99733333V652.02133333h22.6528c5.90933333 0 10.9984-5.088 10.9984-10.99733333s-5.088-10.9984-10.99733333-10.9984h-22.81813333zM315.32799999 16.58986667h33.6512c9.19253333 0 16.74346667 7.55093333 16.74346667 16.74346666s-7.552 16.74346667-16.74346667 16.74346667h-33.6512v33.6512c0 9.19253333-7.552 16.74346667-16.74346667 16.74346667-4.26773333 0-8.37226667-1.64266667-11.81866666-5.088-3.2832-2.464-5.08906667-6.73066667-5.08906667-11.81973334V50.24h-33.6512c-4.26773333 0-8.37226667-1.64266667-11.81866667-5.08906667-3.2832-2.46186667-5.088-6.73066667-5.088-11.81866666 0-9.19253333 7.54986667-16.74346667 16.7424-16.74346667h33.6512V-17.06133333c0-9.19253333 7.552-16.74346667 16.74346667-16.74346667s16.74346667 7.55093333 16.74346667 16.74346667v33.6512zM96.67626666 258.7136c-24.45866667 0-46.29013333 12.64-57.9456 33.6512-11.81866667 21.0112-11.81866667 47.11146667 0 67.3024 11.81866667 21.0112 34.47253333 33.6512 57.9456 33.6512 36.9344 0 67.3024-30.20373333 67.3024-67.3024s-30.368-67.3024-67.3024-67.3024z m29.38346666 84.86613333c-5.90933333 10.01386667-16.74346667 16.74346667-29.38346666 16.74346667-18.54933333 0-33.6512-15.10186667-33.6512-33.6512s15.10186667-33.65013333 33.6512-33.65013333c11.81866667 0 22.6528 6.73066667 29.38346666 16.74346666 5.90933333 10.34133333 5.90933333 22.816 0 33.8144z m1360.3264-219.47093333c-24.45866667 0-46.29013333 12.64-57.9456 33.6512-11.81866667 21.0112-11.81866667 47.11146667 0 67.3024 11.81866667 21.0112 34.47253333 33.6512 57.9456 33.6512 36.9344 0 67.3024-30.20373333 67.3024-67.3024-0.16426667-36.9344-29.54666667-67.3024-67.3024-67.3024z m29.38346667 84.04586667c-5.90933333 10.01386667-16.74346667 16.74346667-29.38346667 16.74346666-18.54933333 0-33.6512-15.10186667-33.6512-33.6512s15.10186667-33.6512 33.6512-33.6512c11.81866667 0 23.47413333 6.73066667 29.38346667 16.74346667 5.90933333 10.3424 5.90933333 23.80266667 0 33.81546667z m0 0" p-id="2512" fill="#e6e6e6"></path></svg>
              <span class="no-data-span">没有任何知识哦</span>
            </div>
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
              <div class="no-hot-data" v-if="hotArticlelist.length === 0">
                <svg t="1626331028665" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3789" width="64" height="64"><path d="M170.54 849.055c0 6.082-4.044 10.125-10.128 10.125H62.129c-6.073 0-10.129-4.044-10.129-10.126 0-6.07 4.056-10.124 10.129-10.124h98.284c6.083 0 10.127 4.055 10.127 10.125z m744.706-87.115c6.083 0 10.139 4.055 10.139 10.135 0 6.071-4.056 10.126-10.14 10.126H744.008v26.332c0 28.37-22.273 50.647-50.64 50.647h-30.405c0 6.082-3.036 12.163-7.091 17.226-10.129 10.125-25.33 10.125-35.459 0l-16.22-17.226h-393.12c-6.082 0-10.127-4.043-10.127-10.126 0-6.07 4.045-10.124 10.129-10.124h44.587a52.071 52.071 0 0 1-10.139-30.397V302.068c0-28.359 22.295-50.648 50.661-50.648h32.422V225.09c0-28.37 22.294-50.648 50.66-50.648h398.192c28.368 0 50.661 22.288 50.661 50.648v506.474c0 11.135-4.055 22.278-10.138 30.377h97.266z m-160.093 0h21.286c17.23 0 30.394-13.161 30.394-30.387V225.078c0-17.215-13.164-30.376-30.394-30.376H378.247c-17.219 0-30.394 13.16-30.394 30.387v26.331h346.513c28.377 0 50.671 22.29 50.671 50.648V761.95h10.128l-0.012-0.01z m-212.238-87.62a88.016 88.016 0 0 0-25.84-62.296c-16.212-16.196-38.505-25.324-61.808-25.324-23.303 0-45.596 9.128-61.808 25.324a88.027 88.027 0 0 0-25.84 62.3 88.027 88.027 0 0 0 25.84 62.302 87.419 87.419 0 0 0 61.808 25.315c22.294 0 44.588-9.119 61.807-25.325a88.016 88.016 0 0 0 25.84-62.296z m-132.236-30.89c2.027 1.008 3.047 4.055 2.027 7.09-7.091 12.152-8.1 28.359-3.035 41.53 1.008 3.046 0 5.063-3.037 7.09h-2.028a5.563 5.563 0 0 1-5.064-3.036c-6.083-16.207-5.063-35.45 3.037-50.648a7.879 7.879 0 0 1 8.1-2.026z m314.102 165.103V302.068c0-17.225-13.165-30.386-30.394-30.386H295.175c-17.23 0-30.405 13.16-30.405 30.386v506.475c0 17.215 13.175 30.387 30.405 30.387h288.76l-28.377-28.37a25.51 25.51 0 0 1-6.073-26.322l-25.33-26.342a106.996 106.996 0 0 1-68.899 25.324c-29.386 0-55.734-11.145-77.01-31.405a107.596 107.596 0 0 1-31.912-76.476c0-28.729 11.49-56.265 31.913-76.476 20.276-20.272 47.624-31.415 77.01-31.415 29.374 0 55.734 11.143 76.999 31.415 39.513 39.504 41.54 103.313 7.091 145.862l26.35 26.332c9.12-3.036 19.247-2.017 26.339 6.081l62.815 62.8h39.525c16.2-1.008 30.394-14.18 30.394-31.405h0.011z m-366.79-432.521c-6.073 0-10.128-4.054-10.128-10.125 0-6.082 4.055-10.136 10.128-10.136h177.312c6.083 0 10.138 4.055 10.138 10.136 0 6.07-4.055 10.126-10.138 10.126H357.98l0.01-0.001z m236.083 44.578c6.083 0 10.128 4.043 10.128 10.125 0 6.07-4.045 10.125-10.128 10.125H361.038c-6.094 0-10.14-4.055-10.14-10.125 0-6.082 4.046-10.136 10.13-10.136h233.024l0.022 0.01z m-103.36 74.942c0 6.07-4.056 10.125-10.127 10.125H355.942c-6.073 0-10.117-4.055-10.117-10.126 0-6.08 4.044-10.136 10.128-10.136h124.623c6.083 0 10.138 4.055 10.138 10.137zM113.818 282.826c0 11.144-9.13 20.25-20.277 20.25-11.172-0.035-20.22-9.08-20.256-20.25 0-11.144 9.13-20.271 20.267-20.271 11.146 0 20.266 9.127 20.266 20.271z m-20.266-40.523c-22.36 0.065-40.469 18.17-40.533 40.523 0 22.288 18.238 40.511 40.533 40.511 22.294 0 40.522-18.234 40.522-40.51 0-22.29-18.228-40.524-40.522-40.524z m848.054 77.999c6.072 0 10.128 4.055 10.128 10.126 0 6.081-4.056 10.135-10.128 10.135-6.083 0-10.138-4.055-10.138-10.135-0.001-6.071 4.055-10.126 10.138-10.126zM911.2 330.428c0 17.225 13.174 30.386 30.405 30.386 17.22 0 30.394-13.161 30.394-30.386 0-17.216-13.164-30.387-30.394-30.387-17.23 0-30.405 13.16-30.405 30.387zM200.946 194.702c0 6.071 4.044 10.126 10.127 10.126 6.083 0 10.128-4.055 10.128-10.126v-12.163h12.155c6.083 0 10.14-4.055 10.14-10.125 0-6.081-4.057-10.136-10.14-10.136h-12.155v-12.152c0-6.082-4.045-10.126-10.128-10.126s-10.128 4.044-10.128 10.126v12.152H188.78c-6.073 0-10.128 4.055-10.128 10.136 0 6.071 4.055 10.126 10.128 10.126h12.166v12.162z m-36.478 271.46c0-6.08-4.056-10.124-10.138-10.124-6.073 0-10.129 4.043-10.129 10.126v12.151h-12.155c-6.083 0-10.139 4.055-10.139 10.137 0 6.081 4.056 10.125 10.14 10.125h12.155v12.152c0 6.082 4.056 10.136 10.128 10.136 6.083 0 10.138-4.055 10.138-10.136v-12.152h12.155c6.083 0 10.129-4.044 10.129-10.126s-4.046-10.136-10.129-10.136h-12.155v-12.152z" p-id="3790" fill="#e6e6e6"></path></svg>                <span class="no-data-span">空</span>
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
              <div class="no-hot-data" v-if="hotRepositoryList.length === 0">
                <svg t="1626331028665" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3789" width="64" height="64"><path d="M170.54 849.055c0 6.082-4.044 10.125-10.128 10.125H62.129c-6.073 0-10.129-4.044-10.129-10.126 0-6.07 4.056-10.124 10.129-10.124h98.284c6.083 0 10.127 4.055 10.127 10.125z m744.706-87.115c6.083 0 10.139 4.055 10.139 10.135 0 6.071-4.056 10.126-10.14 10.126H744.008v26.332c0 28.37-22.273 50.647-50.64 50.647h-30.405c0 6.082-3.036 12.163-7.091 17.226-10.129 10.125-25.33 10.125-35.459 0l-16.22-17.226h-393.12c-6.082 0-10.127-4.043-10.127-10.126 0-6.07 4.045-10.124 10.129-10.124h44.587a52.071 52.071 0 0 1-10.139-30.397V302.068c0-28.359 22.295-50.648 50.661-50.648h32.422V225.09c0-28.37 22.294-50.648 50.66-50.648h398.192c28.368 0 50.661 22.288 50.661 50.648v506.474c0 11.135-4.055 22.278-10.138 30.377h97.266z m-160.093 0h21.286c17.23 0 30.394-13.161 30.394-30.387V225.078c0-17.215-13.164-30.376-30.394-30.376H378.247c-17.219 0-30.394 13.16-30.394 30.387v26.331h346.513c28.377 0 50.671 22.29 50.671 50.648V761.95h10.128l-0.012-0.01z m-212.238-87.62a88.016 88.016 0 0 0-25.84-62.296c-16.212-16.196-38.505-25.324-61.808-25.324-23.303 0-45.596 9.128-61.808 25.324a88.027 88.027 0 0 0-25.84 62.3 88.027 88.027 0 0 0 25.84 62.302 87.419 87.419 0 0 0 61.808 25.315c22.294 0 44.588-9.119 61.807-25.325a88.016 88.016 0 0 0 25.84-62.296z m-132.236-30.89c2.027 1.008 3.047 4.055 2.027 7.09-7.091 12.152-8.1 28.359-3.035 41.53 1.008 3.046 0 5.063-3.037 7.09h-2.028a5.563 5.563 0 0 1-5.064-3.036c-6.083-16.207-5.063-35.45 3.037-50.648a7.879 7.879 0 0 1 8.1-2.026z m314.102 165.103V302.068c0-17.225-13.165-30.386-30.394-30.386H295.175c-17.23 0-30.405 13.16-30.405 30.386v506.475c0 17.215 13.175 30.387 30.405 30.387h288.76l-28.377-28.37a25.51 25.51 0 0 1-6.073-26.322l-25.33-26.342a106.996 106.996 0 0 1-68.899 25.324c-29.386 0-55.734-11.145-77.01-31.405a107.596 107.596 0 0 1-31.912-76.476c0-28.729 11.49-56.265 31.913-76.476 20.276-20.272 47.624-31.415 77.01-31.415 29.374 0 55.734 11.143 76.999 31.415 39.513 39.504 41.54 103.313 7.091 145.862l26.35 26.332c9.12-3.036 19.247-2.017 26.339 6.081l62.815 62.8h39.525c16.2-1.008 30.394-14.18 30.394-31.405h0.011z m-366.79-432.521c-6.073 0-10.128-4.054-10.128-10.125 0-6.082 4.055-10.136 10.128-10.136h177.312c6.083 0 10.138 4.055 10.138 10.136 0 6.07-4.055 10.126-10.138 10.126H357.98l0.01-0.001z m236.083 44.578c6.083 0 10.128 4.043 10.128 10.125 0 6.07-4.045 10.125-10.128 10.125H361.038c-6.094 0-10.14-4.055-10.14-10.125 0-6.082 4.046-10.136 10.13-10.136h233.024l0.022 0.01z m-103.36 74.942c0 6.07-4.056 10.125-10.127 10.125H355.942c-6.073 0-10.117-4.055-10.117-10.126 0-6.08 4.044-10.136 10.128-10.136h124.623c6.083 0 10.138 4.055 10.138 10.137zM113.818 282.826c0 11.144-9.13 20.25-20.277 20.25-11.172-0.035-20.22-9.08-20.256-20.25 0-11.144 9.13-20.271 20.267-20.271 11.146 0 20.266 9.127 20.266 20.271z m-20.266-40.523c-22.36 0.065-40.469 18.17-40.533 40.523 0 22.288 18.238 40.511 40.533 40.511 22.294 0 40.522-18.234 40.522-40.51 0-22.29-18.228-40.524-40.522-40.524z m848.054 77.999c6.072 0 10.128 4.055 10.128 10.126 0 6.081-4.056 10.135-10.128 10.135-6.083 0-10.138-4.055-10.138-10.135-0.001-6.071 4.055-10.126 10.138-10.126zM911.2 330.428c0 17.225 13.174 30.386 30.405 30.386 17.22 0 30.394-13.161 30.394-30.386 0-17.216-13.164-30.387-30.394-30.387-17.23 0-30.405 13.16-30.405 30.387zM200.946 194.702c0 6.071 4.044 10.126 10.127 10.126 6.083 0 10.128-4.055 10.128-10.126v-12.163h12.155c6.083 0 10.14-4.055 10.14-10.125 0-6.081-4.057-10.136-10.14-10.136h-12.155v-12.152c0-6.082-4.045-10.126-10.128-10.126s-10.128 4.044-10.128 10.126v12.152H188.78c-6.073 0-10.128 4.055-10.128 10.136 0 6.071 4.055 10.126 10.128 10.126h12.166v12.162z m-36.478 271.46c0-6.08-4.056-10.124-10.138-10.124-6.073 0-10.129 4.043-10.129 10.126v12.151h-12.155c-6.083 0-10.139 4.055-10.139 10.137 0 6.081 4.056 10.125 10.14 10.125h12.155v12.152c0 6.082 4.056 10.136 10.128 10.136 6.083 0 10.138-4.055 10.138-10.136v-12.152h12.155c6.083 0 10.129-4.044 10.129-10.126s-4.046-10.136-10.129-10.136h-12.155v-12.152z" p-id="3790" fill="#e6e6e6"></path></svg>                <span class="no-data-span">空</span>
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
                that.isMore = false
                that.$message.warning('还没有任何知识，赶快去创建吧！')
              } else {
                for (var i = 0; i < response.data.data.length; i++) {
                  var article = response.data.data[i]
                  if (article.user.avatarurl.indexOf('null') === -1) {
                    article.user.avatarurl = that.$http.defaults.baseURL + article.user.avatarurl
                  } else {
                    article.user.avatarurl = require('@/assets/imgs/defaultAvator.png')
                  }
                  that.articlelist.push(article)
                }
              }
              that.wikiloading = false
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
  .no-data-show{
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    margin-top: 44px;
  }
.no-data-span{
  color: #C0C4CC;
}
  .no-hot-data{
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>
