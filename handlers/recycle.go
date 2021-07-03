package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)


// 回收站接口
func Recycle(c *gin.Context) {
	var (
		id int
		respCate int
		count  int = 0
		createtime *time.Time
		resourceUniqueCode string
		formattime string
		title string
		catename string
		datas []interface{}
	)
	offset := c.DefaultQuery("page","1")
	size := c.DefaultQuery("size","10")
	intOffset,_ := strconv.Atoi(offset)
	intSize,_ := strconv.Atoi(size)
	userid,_ := c.Get("userid")
	row ,err:=Db.Query("SELECT a.id, a.resource_unqiue_code, a.repo_cate, a.createtime,b.ar_title,c.catename FROM sparrow.sp_recycle_bin a  " +
		"JOIN sparrow.sp_article b ON a.resource_unqiue_code = b.ar_unique_code   " +
		"JOIN  sparrow.sp_repo_cate c  ON   a.repo_cate = c.id  where a.userid = ? LIMIT ?,?",userid,(intOffset-1)*intSize,intSize)
	countRow := Db.QueryRow("SELECT count(1) as count from sparrow.sp_recycle_bin where a.userid = ?",userid)
	countRow.Scan(&count)
	if err != nil{
		code = 1
		msg = err.Error()
	}else {
		LOOP:for row.Next(){
			var data = make(map[string]interface{})
			err =row.Scan(&id,&resourceUniqueCode,&respCate,&createtime,&title,&catename)
			if err != nil{
				goto LOOP
			}
			fmt.Println(createtime)
			formattime = createtime.Format("2006-01-02 13:04")
			data["id"] = id
			data["resource_unique_code"] = resourceUniqueCode
			data["resp_cate"] = respCate
			data["deletetime"] = formattime
			data["title"] = title
			data["catename"] = catename
			datas = append(datas,data)
		}
		defer row.Close()
		code = 0
		msg = "success"

	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":datas,
		"msg":msg,
		"count":count,
	})
}


// 彻底删除数据
func DelRecycle(c *gin.Context){
	id := c.Param("id")
	_,err :=Db.Exec("delete  from sparrow.sp_recycle_bin   WHERE id=?",id)
	if err != nil{
		code = 1
		msg = err.Error()
	}else{
		code = 0
		msg = "success"
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":"",
		"msg":msg,
	})
}
