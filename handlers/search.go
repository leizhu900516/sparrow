package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sparrow/utils"
	"strconv"
	"time"
)

/*
搜索 todo走全文索引？ 全文索引需要数据库支持
resourceFlat 几种类型
1、文档(文件默认不在全局搜索) 默认
2、知识库
3、团队
4、用户
*/
func Search(c *gin.Context){
	var (
		rJson map[string]interface{}
		sql string
		countSql string

		count int = 0
		id int64
		uniqueCode string
		title string
		desc string
		createtime *time.Time
		repoName string

		username string
		profile string
		avatarmd5 string

		result []map[string]interface{}
		data  = make(map[string][]map[string]interface{})

	)
	kw := c.DefaultQuery("kw","")
	resourceFlag := c.DefaultQuery("flag","1")
	offset := c.DefaultQuery("page","1")
	size := c.DefaultQuery("size","10")
	intOffset,_ := strconv.Atoi(offset)
	intSize,_ := strconv.Atoi(size)
	switch resourceFlag {
	case "1":
		countSql = "select count(1) from sp_article  where ar_title like '%"+kw+"%' or ar_pure_content like '%"+kw+"%'"
		sql = "select a.id,a.ar_unique_code,a.ar_title,a.ar_pure_content,a.createtime,b.repo_name from sp_article a left join " +
			" sp_repository b on a.repo_unique_code = b.repo_unique_code WHERE a.ar_title like '%"+kw+"%' or a.ar_pure_content like '%"+kw+"%' limit ?,?"
	case "2":
		countSql = "select count(1) from sp_repository where repo_status = 0 and repo_name like '%"+kw+"%' or repo_desc like '%"+kw+"%'"
		sql = "select id,repo_unique_code,repo_name,repo_desc,createtime from sp_repository where repo_status = 0 and repo_name like '%"+kw+"%' or repo_desc like '%"+kw+"%' limit ?,?"
	case "3":
		countSql = "select count(1) from sp_user_group where group_name like '%"+kw+"%' or group_desc like '%"+kw+"%'"
		sql = "select id,group_unique_code,group_name,group_desc,createtime from sp_user_group where group_name like '%"+kw+"%' or group_desc like '%"+kw+"%' limit ?,?"
	case "4":
		countSql = "select count(1) from sp_user where username like '%"+kw+"%'"
		sql = "select username,profile,avatar_md5 from sp_user where username like '%"+kw+"%' limit ?,?"
	}
	fmt.Println(sql)

	switch resourceFlag {
	case "1":
		countRow := Db.QueryRow(countSql)
		err = countRow.Scan(&count)
		if err != nil{
			utils.Logger.Error("统计搜索总数错误",zap.String("error",err.Error()))
		}
		rows,err := Db.Query(sql,(intOffset-1)*intSize,intSize)
		if err != nil{
			rJson = ReturnData(1,"",err.Error())
		}else {
			for rows.Next(){
				_data := make(map[string]interface{})
				err = rows.Scan(&id,&uniqueCode,&title,&desc,&createtime,&repoName)
				if err != nil{
					utils.Logger.Error("解析查询数据出错",zap.String("error",err.Error()))
				}else {
					_data["id"] = id
					_data["uniqueCode"] = uniqueCode
					_data["title"] = title
					if len([]rune(desc)) >50 {
						_data["desc"] = string([]rune(desc)[:50])
					}else {
						_data["desc"] = desc
					}
					_data["repoName"] = repoName
					_data["createtime"] = createtime.Format("2006-01-02 03:04")
				}
				result = append(result,_data)
			}
			defer rows.Close()
			data["doc"] = result
			rJson = ReturnDataTwo(0,data,"success",count)
		}
	case "2","3":
		countRow := Db.QueryRow(countSql)
		err = countRow.Scan(&count)
		if err != nil{
			utils.Logger.Error("统计搜索总数错误",zap.String("error",err.Error()))
		}
		rows,err := Db.Query(sql,(intOffset-1)*intSize,intSize)
		if err != nil{
			rJson = ReturnData(1,"",err.Error())
		}else {
			for rows.Next(){
				_data := make(map[string]interface{})
				err = rows.Scan(&id,&uniqueCode,&title,&desc,&createtime)
				if err != nil{
					utils.Logger.Error("解析查询数据出错",zap.String("error",err.Error()))
				}else {
					_data["id"] = id
					_data["uniqueCode"] = uniqueCode
					_data["title"] = title
					_data["desc"] = desc
					_data["createtime"] = createtime.Format("2006-01-02 03:04")
				}
				result = append(result,_data)
			}
			defer rows.Close()
			data["doc"] = result
			rJson = ReturnDataTwo(0,data,"success",count)
		}
	case "4":
		countRow := Db.QueryRow(countSql)
		err = countRow.Scan(&count)
		if err != nil{
			utils.Logger.Error("统计搜索总数错误",zap.String("error",err.Error()))
		}
		rows,err := Db.Query(sql,(intOffset-1)*intSize,intSize)
		if err != nil{
			rJson = ReturnData(1,"",err.Error())
		}else {
			for rows.Next(){
				_data := make(map[string]interface{})
				err = rows.Scan(&username,&profile,&avatarmd5)
				if err != nil{
					utils.Logger.Error("解析查询数据出错",zap.String("error",err.Error()))
				}else {
					_data["username"] = username
					_data["profile"] = profile
					_data["avatarmd5"] = avatarmd5
				}
				result = append(result,_data)

			}
			defer rows.Close()
			data["user"] = result
			rJson = ReturnDataTwo(0,data,"success",count)
		}
	default:
		utils.Logger.Error("查询参数错误",zap.String("参数kw",kw))
		rJson = ReturnData(1,"","查询参数错误")
	}

	c.JSON(http.StatusOK,rJson)
}
