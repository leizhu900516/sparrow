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



// 添加图书分类接口
type BookCate struct {
	Name string `json:"name"`
	Id int   	`json:"id"`
}
func AddBookCate(c *gin.Context) {
	var rJson map[string]interface{}
	var sql string
	form := &BookCate{}
	if err := c.ShouldBind(form);err != nil{
		utils.Logger.Error("图书分类表单提交错误" + err.Error())
		rJson = ReturnData(1,"",err.Error())
	}else{
		name := form.Name
		id := form.Id
		if id == 0 { // 添加新分类
			sql = fmt.Sprintf("insert into sp_book_cate (cate_name) values ('%s')",name )
		} else { // 更新分类信息
			sql = fmt.Sprintf("update  sp_book_cate set cate_name =' %s' where id= %d",name,id)
		}
		result,err := Db.Exec(sql)
		if err !=nil{
			utils.Logger.Error("图书分类表单写入错误")
			rJson = ReturnData(1,"",err.Error())
		}else {
			insertID,_ := result.RowsAffected()
			rJson = ReturnData(0,insertID,"success")
		}
	}
	c.JSON(http.StatusOK,rJson)
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
	case "all":
	default:
		sql += fmt.Sprintf(" where b_cate_id = %s ",cate)
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


// 删除图书分类接口
func DelBookCate(c *gin.Context) {
	var cid = c.Param("cid")
	var rJson  map[string]interface{}
	result,err :=Db.Exec("delete from sp_book_cate where id =?",cid)
	if err != nil{
		rJson = ReturnData(1,"",err.Error())
		c.JSON(http.StatusOK,rJson)
	}
	affectId,_ := result.RowsAffected()
	rJson = ReturnData(0,affectId,"success")
	c.JSON(http.StatusOK,rJson)
}
