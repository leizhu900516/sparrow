package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sparrow/utils"
	"strconv"
	"strings"
)
var (
    msg string
    code int
)
// 获取图书分类接口
func GetBookCate(c *gin.Context){
	var id int
	var catename string
	var datalist []interface{}
	rows,err := Db.Query("SELECT id,cate_name FROM sp_book_cate")
	if err != nil{
		utils.Logger.Info("图书馆功能分类读取失败")
		code = 1
		msg = "读取分类失败"
	}else {
		for rows.Next(){
			var data = make(map[string]interface{})
			rows.Scan(&id,&catename)
			data["id"] = id
			data["catename"] = catename
			datalist = append(datalist, data)
		}
		defer rows.Close()
		code = 0
		msg = "success"
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":datalist,
		"msg":msg,
	})
}

func GetBookList(c *gin.Context){
	var (
		id int
		name string
		fileurl string
		avatorurl string
		download int
		avatorMd5 string
		datas []interface{}
		sql string = "SELECT id, b_name, b_url, b_avator_url, download,avator_md5 FROM sparrow.sp_book  "
	)
	cate := c.DefaultQuery("flag","all")
	kw := c.Query("kw")
	rank := c.Query("rank")
	offset := c.DefaultQuery("page","1")
	size := c.DefaultQuery("size","10")
	intOffset,_ := strconv.Atoi(offset)
	intSize,_ := strconv.Atoi(size)
	switch cate {
	case "1":
		sql += fmt.Sprintf(" where b_cate_id = %d ",1)
	case "2":
		sql += fmt.Sprintf(" where b_cate_id = %d ",2)
	default:

	}
	if kw != "" {
		if !strings.Contains(sql,"where"){
			sql  =  sql + " where b_name like '%"+kw+"%'"
		}else {
			sql  =  sql + " and b_name like '%"+kw+"%'"
		}
	}
	if rank == "hot"{
		sql  =  sql + " order by download desc "
	}
	sql += fmt.Sprintf(" limit %d,%d",(intOffset-1)*intSize ,intSize)
	fmt.Println(sql)
	rows,err := Db.Query(sql)
	if err != nil{
		utils.Logger.Error("获取图书列表错误",zap.String("error",err.Error()))
		code = 1
		msg = "读取分类失败"
	}else{
		for rows.Next(){
			var tmpMap = make(map[string]interface{})
			rows.Scan(&id,&name,&fileurl,&avatorurl,&download,&avatorMd5)
			tmpMap["id"] = id
			tmpMap["name"] = name
			tmpMap["fileurl"] = fileurl
			tmpMap["avatorurl"] = "api/v1/img/"+ avatorMd5
			tmpMap["download"] = download
			datas = append(datas,tmpMap)
		}
		defer rows.Close()
		code = 0
		msg = "success"
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":datas,
		"msg":msg,
	})
}
