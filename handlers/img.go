package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sparrow/utils"
)


//图片接口 配置文件根目录 + /20121223/d41d8cd98f00b204e9800998ecf8427e
func ImgShow(c *gin.Context) {
	//fullpath := "files/"+c.Param("date")+"/"+c.Param("md5")
	//fmt.Println(fullpath)
	var (
		fileUrl string
	)
	md5Str := c.Param("md5")
	fmt.Println(md5Str)
	//date:= c.Param("date")
	row:=Db.QueryRow("select fileurl from sp_files where md5=?",md5Str)
	err := row.Scan(&fileUrl)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"msg":"图片不存在",
		})
	}else{
		fmt.Println("fileUrl",fileUrl)
		exits := utils.Isfile(fileUrl)
		if !exits{
			c.JSON(http.StatusOK,gin.H{
				"code":1,
				"msg":"图片不存在",
			})
		}else{
			c.File(fileUrl)
		}
	}
}
