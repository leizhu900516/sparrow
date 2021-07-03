package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strings"
)
func Sha1String(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}
//api认证函数
func ApiAuthCheck(c *gin.Context)  {

	fmt.Println("权限检查")
	apiKey := c.PostForm("api_key")
	timestrap := c.PostForm("timestrap")
	sign := c.PostForm("sign")
	paramsSlice :=[]string{apiKey,timestrap,sign}
	sort.Strings(paramsSlice)
	fmt.Println(paramsSlice)
	sortStr := strings.Join(paramsSlice,"")
	newSign := Sha1String(sortStr)
	fmt.Println(sign,newSign)
	fmt.Println("check auth")
	if sign != newSign {
		c.Abort()
		c.JSON(200,gin.H{
			"code":1,
			"msg":"auth fail",
		})
		fmt.Println("auth fail")
		return
	}
	//
	c.Next()
	fmt.Println("aaa")
}

/**
	是否登陆检测，如果登陆则在context中写入登陆的用户id
	如果未登陆，则进入不登陆处理逻辑
	authflag  是否需要认证-为获取个人行程数据逻辑 1为需要认证 如果没有该参数，则是获取全部到列表页
*/
func JwtAuthMiddleware (c *gin.Context){
	var (
		token string
	)
	tokenString := c.Request.Header.Get("Authorization")

	isAuthFlag := c.DefaultQuery("authflag","1") //强制认证才能继续，例如个人的列表，如果没有登陆，则跳转到登陆。
	//fmt.Println("isAuthFlag=",isAuthFlag,"tokenString=",tokenString)
	if isAuthFlag =="1"{
		//fmt.Println("请求的token=",tokenString)
		if  tokenString != ""{
			//_token := strings.Split(tokenString," ")
			//if len(_token) > 1{
			token =tokenString
			msg,status:=ParseJwt(token)
			if status{
				userid := int64(msg.(jwt.MapClaims)["id"].(float64))
				c.Set("userid",userid)
				c.Next()

			}else{
				c.Abort()
				c.JSON(http.StatusForbidden,gin.H{
					"code":1,
					"msg":msg,
				})
			}
		}else {
			c.Abort()
			c.JSON(http.StatusForbidden,gin.H{
				"code":1,
				"msg":"未认证",
			})

		}
	}else{
		c.Next()
	}


}

/*
历史专用认证处理逻辑
登陆则记录，没有登陆则不处理
*/
func JwtAuthMiddleware2 (c *gin.Context){
	tokenString := c.Request.Header.Get("Authorization")
	fmt.Println("token==",tokenString)
	if  tokenString != ""{
		token := strings.Split(tokenString," ")[1]
		msg,status:=ParseJwt(token)

		//fmt.Println("msg=",msg)
		//fmt.Println("status=",status)
		if status{
			userid := int64(msg.(jwt.MapClaims)["id"].(float64))
			c.Set("userid",userid)


		}
	}
	c.Next()
}

/**
解决跨域请求中间件
*/
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")		//请求头部
		var headerKeys []string								// 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")		// 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")		//服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//				允许跨域设置																										可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")		// 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")		// 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")		//	跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")		// 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()        //  处理请求
	}
}
