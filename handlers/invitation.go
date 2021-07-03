package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sparrow/utils"
	"strconv"
	"strings"
	"time"
)

// 邀请入团队
// 验证是否过期
// 验证团队id是否正确
func Invitation(c *gin.Context)  {
	var (
		id int
		data interface{}
		groupMember string
		groupName string
		groupDesc string
		groupUniqueCode string
		respName string
		respDesc string
		groupUserid int64
		respList []interface{}
		rJson map[string]interface{}
	)
	userid ,_:= c.Get("userid")
	enCode := c.Param("encode")
	isJoin := c.DefaultQuery("join","false")
	state,deCode :=utils.NewCFBDecrypter(enCode)
	//fmt.Println(deCode)
	if !state{
		rJson = ReturnData(1,"","邀请码错误")
	}else {
		sep := strings.Split(deCode,".")
		fmt.Println(">>>>",sep)
		beforeTime,_ := strconv.ParseInt(sep[0],10,64)
		localTime := time.Now().Unix()
		if( localTime - beforeTime) > 0 {
			rJson = ReturnData(1,"","邀请已经过期")
		}else{
			// 是否为加入组功能 isJoin为ture 是 将此人加入到组
			if isJoin == "true"{
				var mid int64
				row:= Db.QueryRow("select id from sp_user_group_member where userid=?  and sp_user_group_member = ?",userid,sep[1])
				err = row.Scan(&mid)
				if err == nil{
					rJson = ReturnData(2,"","已经在组内")
				}else {
					result,err := Db.Exec("insert into sp_user_group_member (userid,group_unique_code) values (?,?)",userid,sep[1])
					if  err != nil{
						utils.Logger.Error("添加团队成员失败",zap.String("error",err.Error()))
						rJson = ReturnData(1,"",err.Error())
					}else{
						data ,_= result.LastInsertId()
						rJson = ReturnData(0,data,"success")
					}
				}
			}else{
				// 邀请页面逻辑
				row:=Db.QueryRow("select id,group_member,group_name,group_desc,group_unique_code,userid " +
					"from sparrow.sp_user_group where group_unique_code = ?",sep[1])
				err := row.Scan(&id,&groupMember,&groupName,&groupDesc,&groupUniqueCode,&groupUserid)
				if err != nil{
					rJson = ReturnData(1,"","组不存在")
				}else{
					var mid int64
					row:= Db.QueryRow("select id from sp_user_group_member where userid=?  and group_unique_code = ?",userid,sep[1])
					err = row.Scan(&mid)
					if err == nil{
						rJson = ReturnData(2,"","已经在组内")
					}else {
						// 获取组的知识库信息
						repoRows,err := Db.Query("select repo_name,repo_desc from sp_repository where repo_user_group = ?",id)
						if err != nil{
							rJson = ReturnData(1,"",err.Error())
						}else{
							for repoRows.Next(){
								_data := make(map[string]string)
								repoRows.Scan(&respName,&respDesc)
								_data["respname"] = respName
								_data["respdesc"] = respDesc
								respList = append(respList,_data)
							}
							defer repoRows.Close()

							_data := make(map[string]interface{})
							_data["gid"] = id
							_data["groupname"] = groupName
							_data["groupdesc"] = groupDesc
							_data["resplist"] = respList
							_data["groupUniqueCode"] = groupUniqueCode
							code = 0
							msg = "success"
							data = _data
							rJson = ReturnData(0,data,"success")
						}
					}
				}
			}
		}
	}
	c.JSON(http.StatusOK,rJson)
}


// 生成邀请链接，规则如下
// 过期时间 10位+组唯一字符串8+随机字符串10，用"."进行拼接
func GenerateInvitationEncode(c *gin.Context) {
	expiredDay := c.Query("expired")
	teamfalg := c.Query("flag")
	expiredDayInt,err := strconv.Atoi(expiredDay)
	if err != nil{
		expiredDayInt = 3
	}
	localTimestemp := time.Now().AddDate(0,0,expiredDayInt).Unix()
	localTimeTempstr := strconv.FormatInt(localTimestemp,10)
	enStr := localTimeTempstr + "." + teamfalg + "." + utils.GetRandomString(10)
	fmt.Println("enStr=",enStr)
	enCode := utils.NewCFBEncrypter(enStr)
	fmt.Println("邀请码长度=",len(enCode))
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"data":enCode,
		"msg":"success",
	})

}
