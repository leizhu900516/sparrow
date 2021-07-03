package handlers

import "github.com/gin-gonic/gin"


// 通用返回http json结构体
func ReturnData(code int,data interface{},msg string) gin.H{
	return gin.H{
		"code":code,
		"msg":msg,
		"data":data,
	}
}
// 带返回数据大小的json 结构体
func ReturnDataTwo(code int,data interface{},msg string,count int) gin.H{
	return gin.H{
		"code":code,
		"msg":msg,
		"data":data,
		"count":count,
	}
}
