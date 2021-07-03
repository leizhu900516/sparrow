package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sparrow/utils"
	"time"
)

func Collect(c *gin.Context) {
	var (
		collecttime *time.Time
		id int
		resourcetype int
		resourceid string
		resourceName string
		datas []interface{}
		catename string
		//respositoryname string
		whereSql string
		rJson  map[string]interface{}
	)
	var flag = c.DefaultQuery("flag","all")
	switch flag {
	case "doc":
		whereSql = " where resource_type = 1"
	case "file":
		whereSql = " where resource_type = 2"
	case "team":
		whereSql = " where resource_type = 3"
	case "repository":
		whereSql = " where resource_type = 4"
	default:
		break
	}
	if flag != "all"{
		whereSql =  "SELECT id, resource_type,resource_unique_code, createtime FROM sparrow.sp_collect" + whereSql
	}else {
		whereSql = "SELECT id, resource_type,resource_unique_code, createtime FROM sparrow.sp_collect"
	}
	fmt.Println(whereSql)
	rows,err := Db.Query(whereSql)
	if err != nil{
		utils.Logger.Error("查询收藏失败",zap.String("error",err.Error()))
		rJson = ReturnData(1,"",err.Error())
	}else{
		for rows.Next(){
			data := make(map[string]interface{})
			err := rows.Scan(&id,&resourcetype,&resourceid,&collecttime)
			if err == nil{ // todo获取收藏显示错误问题
				switch resourcetype {
				case 1: //文档
					row:= Db.QueryRow("select a.ar_title,b.repo_name from sparrow.sp_article a join sparrow.sp_repository b on a.repo_unique_code = b.repo_unique_code where a.ar_unique_code=?",resourceid)
					_ = row.Scan(&resourceName,&catename)
				case 2: //资源
					row:= Db.QueryRow("select a.filename,b.repo_name from sparrow.sp_files a join sparrow.sp_repository b on a.repo_unique_code = b.repo_unique_code where a.id=?",resourceid)
					_ = row.Scan(&resourceName)
				case 3: // 团队
					row := Db.QueryRow("select group_name from sparrow.sp_user_group where group_unique_code=?",resourceid)
					_ = row.Scan(&resourceName)
				case 4: // 知识库
					row := Db.QueryRow("select a.repo_name,b.catename from sparrow.sp_repository a join sparrow.sp_repo_cate b on a.id = b.id where id=?",resourceid)
					_ = row.Scan(&resourceName)
				}
				fmt.Println(">>>",resourcetype,resourceName)
				data["id"] = id
				data["resourcename"] = resourceName
				data["resourcetype"] = resourcetype
				data["collecttime"] = collecttime.Format("2006-01-02 03:04")
				data["catename"] = catename
				datas = append(datas,data)
			}
		}
		defer rows.Close()
		rJson = ReturnData(0,datas,"success")
	}
	c.JSON(http.StatusOK,rJson)
}


// 新增收藏
type collect struct {
	DocCode string `josn:"docCode"`
	Resourcetype int `json:"Resourcetype"`
}
func AddCollect(c *gin.Context){
	userid,_ := c.Get("userid")
	params := collect{}
	err := c.ShouldBind(&params)
	if err != nil{
		code = 1
		msg = err.Error()
	}else{
		_,err :=Db.Exec("insert into sparrow.sp_collect (resource_type,resource_unique_code,userid) values (?,?,?)",
			params.Resourcetype,params.DocCode,userid)
		if err !=nil{
			code = 1
			msg = err.Error()
		}else{
			code = 0
			msg = "收藏成功"
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":"",
		"msg":msg,
	})
}


// 取消收藏
func DelCollect(c *gin.Context){
	id:= c.Param("id")
	if id == ""{
		code = 1
		msg = "id不能为空"
	}else {
		_,err := Db.Exec("delete from sparrow.sp_collect  where id = ?",id)
		if err != nil{
			code = 1
			msg = err.Error()
		}else{
			code = 0
			msg = "取消成功"
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":"",
		"msg":msg,
	})
}
