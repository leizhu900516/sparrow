package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sparrow/utils"
	"strconv"
	"time"
)

// 新增文件夹==文件分类
func AddFileCate(c *gin.Context){
	type Filecate struct {
		Filedir string `json:"filedir"`
		Repositoryid string `json:"repositoryid"`
		Parentid int `json:"parentid"`
	}
	var (
		filecate = Filecate{}
		data int64
		rJson map[string]interface{}
	)

	err := c.ShouldBindJSON(&filecate)

	if err != nil{
		rJson = ReturnData(1,"",err.Error())
	}else {
		result ,err :=Db.Exec("insert into sparrow.sp_filecate (`name`,`parent_id`,`repo_unique_code`) values (?,?,?)",
			filecate.Filedir,filecate.Parentid,filecate.Repositoryid)
		if err != nil{
			rJson = ReturnData(1,"",err.Error())
		}else {
			data,_ = result.LastInsertId()
			rJson = ReturnData(0,data,"success")
		}
	}
	c.JSON(http.StatusOK,rJson)
}


// 获取资源列表
func GetFiles(c *gin.Context)  {
	var (
		md5str string
		data []interface{}
		fileCount int
	)
	userid,_ := c.Get("userid")
	offset := c.DefaultQuery("page","1")
	size := c.DefaultQuery("size","10")
	intOffset,_ := strconv.Atoi(offset)
	intSize,_ := strconv.Atoi(size)
	rows,err := Db.Query("select a.id,a.doc_name,a.repo_unique_code,a.createtime,a.doc_md5,b.repo_name,f.filetype from sp_doc a left join " +
		" sparrow.sp_repository b  on  a.repo_unique_code = b.repo_unique_code  left join sp_files as f on a.doc_md5=f.md5 " +
		"where a.userid = ? limit ?,?",userid,(intOffset-1)*intSize,intSize)
	if err != nil{
		code = 1
		msg = err.Error()
	}else {
		for rows.Next(){
			var (
				id int
				filename,filetype,respName sql.NullString
				repoUniqueCode sql.NullString
				createtime *time.Time
			)
			_data := make(map[string]interface{})
			rows.Scan(&id,&filename,&repoUniqueCode,&createtime,&md5str,&respName,&filetype)
			_data["id"] = id
			_data["filename"] = filename.String
			_data["filetype"] = filetype.String
			_data["repositoryid"] = repoUniqueCode.String
			_data["createtime"] = createtime.Format("2006-01-02 15:04")
			_data["repositoryname"] = respName.String
			_data["md5str"] = md5str
			data = append(data,_data)

		}

		defer rows.Close()
		row := Db.QueryRow("select count(1) as count from sparrow.sp_files where userid = ?",userid)
		if row != nil{
			err := row.Scan(&fileCount)
			if err != nil{
				fileCount = 0
			}
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"count":fileCount,
		"msg":msg,
	})
}


func DelFiles(c *gin.Context){
	var (
		fileurl string
	)
	id := c.Param("id")
	fileInfo := Db.QueryRow("select fileurl from sparrow.sp_files where  id=?",id)
	if fileInfo != nil{
		fileInfo.Scan(&fileurl)
	}
	_,err =Db.Exec("delete  from sparrow.sp_files   WHERE id = ?",id)
	if err != nil{
		code = 1
		msg = err.Error()
	}else{
		if utils.CheckFileState(fileurl) {
			err := os.Remove(fileurl)
			if err != nil{
				code = 1
				msg = err.Error()
			}else{
				code = 0
				msg = "success"
			}
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":"",
		"msg":msg,
	})
}
