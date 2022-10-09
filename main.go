package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"sparrow/handlers"
	"sparrow/utils"
)

func main() {
	fmt.Println("start>>>>>>>")
	var (
		configFile string
		port       string
	)
	flag.StringVar(&configFile, "c", "./config/config.ini", "please input  config path ,defalt path is ./config/config.ini")
	flag.StringVar(&port, "port", "12345", "please input server port")
	flag.Parse()
	utils.ParseConfig(configFile)
	//models.DbInit()
	handlers.DbConnectInit()
	//utils.CreateFilePath()
	r := gin.Default()
	//r.Static("/static","./static")
	r.Use(utils.Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/api/v1/login", handlers.Login)
	r.POST("/api/v1/register", handlers.Register)
	// api路由组
	//api := r.Group("/api/v1",utils.JwtAuthMiddleware)
	api := r.Group("/api/v1")
	api.GET("/img/:md5", handlers.ImgShow)               // 图片预览接口
	api.GET("/book/download/:id", handlers.BookDownload) // 图书下载接口
	api.Use(utils.JwtAuthMiddleware)
	{
		api.POST("/upload", handlers.Upload)          // 上传接口 通用上传接口
		api.POST("/book/upload", handlers.BookUpload) // 图书上传接口
		api.GET("/download/:md5", handlers.Download)  // 下载接口

		// 知识库
		api.GET("/repository", handlers.Getrespositorys)                  //获取知识库列表
		api.POST("/repository", handlers.Addrepository)                   //新增知识库
		api.GET("/repository/:repoflag", handlers.GetRidrespositorys)     //获取指定知识库列表
		api.DELETE("/repository/:repoflag", handlers.Delrespositorys)     //删除指定知识库
		api.GET("/repositoryinfo/:repoflag", handlers.GetrespositoryInfo) //获取知识库详情信息

		// 图书
		api.GET("/book/cate", handlers.GetBookCate)         //获取图书分类
		api.POST("/book/cate", handlers.AddBookCate)        //添加图书分类
		api.GET("/books", handlers.GetBookList)             //获取图书列表
		api.DELETE("/book/cate/:cid", handlers.DelBookCate) //删除图书分类
		api.DELETE("/book/del/:bookid", handlers.DelBook)   //删除图书
		// 文章
		api.POST("/article", handlers.AddArticle)           // 更新和新增是否一起
		api.GET("/article", handlers.GetArticle)            // 获取全部文章
		api.GET("/self/article", handlers.GetMySelfArticle) // 获取自己的文章
		api.POST("/good", handlers.GoodArticle)             //点赞
		api.GET("/follow/article", handlers.GetArticleFollow)
		api.GET("/article/:id", handlers.GetArticleDesc)
		api.DELETE("/article/:ar_unique_code", handlers.DelArticle)
		api.PUT("/article/:id", handlers.UpdateArticle)
		// 我的关注动态
		api.POST("/follow", handlers.AddFollow)
		api.GET("/follow/moments", handlers.FollowMoments)

		// 团队 组
		api.GET("/team", handlers.GetTeam)
		api.POST("/team", handlers.AddTeam)
		api.GET("/team/repository", handlers.GetTeamRepository)
		api.GET("/team/info/:flag", handlers.GetTeamInfo)
		api.GET("/member/:teamflag", handlers.GetTeamMember)
		api.DELETE("/member/:teamflag/:uid", handlers.DelTeamMember)

		// 回收站
		api.GET("/recycle", handlers.Recycle)
		api.DELETE("/recycle/:id", handlers.DelRecycle)
		// 收藏接口
		api.GET("/collect", handlers.Collect)
		api.POST("/collect", handlers.AddCollect)
		api.DELETE("/collect/:id", handlers.DelCollect)
		// 用户接口
		api.GET("/user", handlers.GetUserinfo)
		api.POST("/user", handlers.UpdateUserinfo)
		api.POST("/user/passwd", handlers.UpdateUserPasswd)
		// 文件目录=文件分类
		api.POST("/file/dir", handlers.AddFileCate)
		// 文件=资源列表
		api.GET("/file", handlers.GetFiles)
		api.DELETE("/file/:id", handlers.DelFiles) // 删除文件
		// 邀请入团队
		api.GET("/generate/invitation/encode", handlers.GenerateInvitationEncode)
		api.GET("/invitation/:encode", handlers.Invitation)
		// 搜索todo
		api.GET("/search", handlers.Search)

	}
	port, _ = utils.Config.GetValue(utils.ConfigName, "serverPort")
	r.Run(fmt.Sprintf("0.0.0.0:%s", port)) // 监听并在 0.0.0.0:8080 上启动服务
}
