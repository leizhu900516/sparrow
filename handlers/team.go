package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"sparrow/utils"
	"strconv"
	"time"
)

// 获取团队列表
func GetTeam(c *gin.Context) {
	var (
		id int
		flag string
		name string
		groupdesc string
		avatorurl string
		datas []interface{}
	)
	userid,_ := c.Get("userid")
	rows,err :=Db.Query("SELECT b.id,a.group_unique_code,b.group_name,b.group_desc,c.fileurl " +
		"FROM sp_user_group_member a left join sp_user_group b on a.group_unique_code = b.group_unique_code" +
		" left join sp_files c on b.group_avator_url_md5 = c.md5  WHERE a.userid = ? ",userid)
	if err != nil{
		code = 1
		msg = err.Error()
	}else{
		for rows.Next(){
			var row =  make(map[string]interface{})
			rows.Scan(&id,&flag,&name,&groupdesc,&avatorurl)
			row["id"] = id
			row["flag"] = flag
			row["groupdesc"] = groupdesc
			row["avatorurl"] = avatorurl
			row["name"] = name
			datas = append(datas,row)
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


//获取团队详细信息
type teamInfo struct {
	Id int
	Group_unique_code string
	Group_name string
	Group_desc string
	Group_status int
	Group_avator_url_id string
	Createtime *time.Time
}

func GetTeamInfo(c *gin.Context){
	var (
		teaminfo = teamInfo{}
		data  =make(map[string]interface{})
	)
	groupflag := c.Param("flag")
	rows := Db.QueryRow("select id,group_unique_code,group_name,group_desc,group_status,group_avator_url_md5,createtime from sparrow.sp_user_group where group_unique_code = ?",groupflag)
	if err != nil{
		code =1
		msg =  err.Error()
	}else {

		rows.Scan(&teaminfo.Id,&teaminfo.Group_unique_code,
			&teaminfo.Group_name,&teaminfo.Group_desc,
			&teaminfo.Group_status,&teaminfo.Group_avator_url_id,&teaminfo.Createtime)
		data = utils.Struct2Map(teaminfo)

	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"msg":msg,
	})
}
//获取组的组员信息
func GetTeamMember(c *gin.Context) {
	var (
		groupMemberInfoSlice []interface{}
		uid int64
		isleader string
		//group_leader_id int
		usertype string
		username string
		rJson map[string]interface{}
	)
	groupflag := c.Param("teamflag")
	if groupflag != "" {
		rows,err := Db.Query("select userid,isleader from sp_user_group_member where group_unique_code = ?",groupflag)
		if err != nil{
			utils.Logger.Error("读取组成员失败",zap.String("error",err.Error()))
			rJson = ReturnData(1,"",err.Error())
		}else {
			for rows.Next(){
				rows.Scan(&uid,&isleader)
				if isleader == "1"{
					usertype = "组长"
				}else {
					usertype = "组员"
				}
				_data := make(map[string]interface{})
				userRow:=Db.QueryRow("select id,username from sparrow.sp_user where id = ?",uid)
				err := userRow.Scan(&uid,&username)
				if err == nil{
					_data["uid"] = uid
					_data["username"] = username
					_data["usertype"] = usertype
					groupMemberInfoSlice = append(groupMemberInfoSlice,_data)
				}

			}
			defer rows.Close()
		}
		rJson = ReturnData(0,groupMemberInfoSlice,"success")
	}else {
		rJson = ReturnData(1,"","团队标识不能为空")
	}
	c.JSON(http.StatusOK,rJson)
}
// 删除团队某一用户
func DelTeamMember(c *gin.Context) {
	var (
		groupMemberStr string
		userid int64
		groupMemberSlice []int
	)
	groupflag := c.Param("teamflag")
	uid := c.Param("uid")
	uidInt64 ,err := strconv.ParseInt(uid,10,64)
	uidInt ,err := strconv.Atoi(uid)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"msg":"参数错误",
			"data":"",
		})
	}
	if groupflag != "" && uid != "" {
		row := Db.QueryRow("select group_member,userid from sparrow.sp_user_group where group_unique_code = ?",groupflag)
		row.Scan(&groupMemberStr,&userid)
		if (userid == uidInt64) {
			code = 1
			msg = "不能删除组长"
		}else {
			err :=json.Unmarshal([]byte(groupMemberStr),&groupMemberSlice)
			if err != nil{
				code = 1
				msg = err.Error()
			}else{
				for i :=0;i<len(groupMemberSlice);i++{
					if groupMemberSlice[i] == uidInt{
						groupMemberSlice = append(groupMemberSlice[:i],groupMemberSlice[i+1:]...)
					}
				}
				groupMember,_ :=json.Marshal(groupMemberSlice)
				_,err := Db.Exec("update sparrow.sp_user_group set group_member = ?  where group_unique_code = ?",groupMember,groupflag)
				if err != nil{
					code = 1
					msg = err.Error()
				}else{
					code = 0
					msg = "success"
				}
			}
		}
	}else {
		code = 1
		msg = "参数错误"
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
		"data":"",
	})
}
// 获取组的知识库列表
func GetTeamRepository(c *gin.Context) {
	var (
		id int64
		respname string
		resp_desc string
		resp_cate int
		createtime *time.Time
		respuniquecode string
		replist []interface{}
	)
	groupid := c.Query("groupid")
	rows,err := Db.Query("select id,repo_name,repo_desc,repo_unique_code,createtime,repo_cate from sparrow.sp_repository where repo_user_group = ?",groupid)
	if err != nil{
		code = 1
		msg = err.Error()
	}
	for rows.Next(){
			_data := make(map[string]interface{})
			_ =rows.Scan(&id,&respname,&resp_desc,&respuniquecode,&createtime,&resp_cate)
			_data["id"] = id
			_data["respname"] = respname
			_data["cateid"] = resp_cate
			_data["resp_desc"] = resp_desc
			_data["createtime"] = createtime.Format("2006-01-02 15:04")
			_data["respuniquecode"] = respuniquecode
			replist = append(replist,_data)
	}
	defer rows.Close()
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":replist,
		"msg":msg,
	})
}

//新增团队 team or group
type Team struct {
	Name string `json:"name"`
	Desc string `josn:"desc"`
	Auth int `json:"auth"`
	Avator string `json:"avator"`
}
func AddTeam(c *gin.Context){

	var (
		//data string
		groupFlag string
		id int
		rJons map[string]interface{}
	)

	team := Team{}
	userid,_ := c.Get("userid")
	err := c.ShouldBind(&team)
	fmt.Println("进来了吗",team)
	if err != nil {
		rJons = ReturnData(1,"",err.Error())
	}else {
		GenRandomStr:for{
			groupFlag =utils.GetRandomString(10)
			fmt.Println("groupFlag=",groupFlag)
			row := Db.QueryRow("select id from sparrow.sp_user_group where group_unique_code = ?",groupFlag)
			err:=row.Scan(&id)
			fmt.Println("id=",id)
			if err == nil{
				utils.Logger.Error("组唯一字符串已存在，重新生成")
				time.Sleep(1*time.Second)
				goto GenRandomStr
			}else {
				break
			}
		}
		trx ,_:= Db.Begin()
		groupRow,err :=trx.Exec("insert into sparrow.sp_user_group (group_unique_code,group_name,group_desc,group_status,group_avator_url_md5,userid) values (?,?,?,?,?,?)",
			groupFlag,team.Name,team.Desc,team.Auth,team.Avator,userid)
		if  err != nil{
			log.Printf(err.Error())
		}
		memberRow,err := trx.Exec("insert into sp_user_group_member (userid,group_unique_code,isleader) values (?,?,'1')",userid,groupFlag)
		if  err != nil{
			log.Printf(err.Error())
		}
		i1, _ := groupRow.RowsAffected()
		i2, _ := memberRow.RowsAffected()
		if i1 > 0 && i2 > 0 {
			err  = trx.Commit()
			if err != nil{
				utils.Logger.Error("新增团队提交失败",zap.String("error",err.Error()))
				rJons = ReturnData(1,"",err.Error())
			}else {
				rJons = ReturnData(0,groupFlag,"success")
			}
		}else {
			err = trx.Rollback()
			if err != nil{
				utils.Logger.Error("新增团队回滚失败",zap.String("error",err.Error()))
			}
			rJons = ReturnData(1,"",err.Error())
		}
	}
	c.JSON(http.StatusOK,rJons)
}
