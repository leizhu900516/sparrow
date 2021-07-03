package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sparrow/utils"
)


//获取用户信息
// 1、用户基本资料
// 2、用户关注者
// 3、用户被关注数
// 4、用户文档数  知识库简介 走知识库api
// 5、用户团队数 团队列表走团队api
func GetUserinfo(c *gin.Context) {
	var (
		datas =make(map[string]interface{})
		followCount int
		beFollowCount int
		docCount int
		username string
		//avatarurl string
		avatarMd5 string
		desc string
	)
	userid ,_:= c.Get("userid")
	row := Db.QueryRow("select `username`,`avatar_md5`,`profile` from sparrow.sp_user where id = ?",userid)
	err := row.Scan(&username,&avatarMd5,&desc)
	if err != nil{
		fmt.Println(err.Error())
	}else{
		followRow := Db.QueryRow("select count(1) as count from sparrow.sp_follow where userid = ?",userid)
		err = followRow.Scan(&followCount)
		if err != nil{
			followCount = 0
		}
		//被关注
		befollowRow := Db.QueryRow("select count(1) as count from sparrow.sp_follow where follow_type = 0 and follow_id = ?",userid)
		err = befollowRow.Scan(&beFollowCount)
		if err != nil{
			beFollowCount = 0
		}
		docRow := Db.QueryRow("select count(1) as count from sparrow.sp_article where userid = ?",userid)
		err = docRow.Scan(&docCount)
		if err != nil{
			docCount = 0
		}
	}
	datas["username"] = username
	datas["followCount"] = followCount
	datas["beFollowCount"] = beFollowCount
	datas["docCount"] = docCount
	datas["avatarurl"] = "api/v1/img/"+avatarMd5
	datas["desc"] = desc
	datas["avatarMd5"] = avatarMd5
	datas["userid"] = userid
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":datas,
		"msg":msg,
	})
}



// 修改用户资料
type Userinfo struct {
	//Username string `json:"username"`
	Userdesc string `json:"userdesc"`
	Avatarurl string `json:"avatarurl"`
}
func UpdateUserinfo(c *gin.Context) {
	userinfo := Userinfo{}
	userid,_ := c.Get("userid")
	var RowsId int64
	err := c.ShouldBindJSON(&userinfo)
	fmt.Println(userinfo)
	if err != nil{
		code = 1
		msg = err.Error()
	}else {
		result,err := Db.Exec("update sparrow.sp_user set `profile` = ?,avatar_md5 = ? where id = ?",userinfo.Userdesc,userinfo.Avatarurl,userid)
		if err != nil{
			code = 1
			msg = err.Error()
		}else {
			code = 0
			msg = "success"
			RowsId,_  = result.RowsAffected()
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":RowsId,
		"msg":msg,
	})
}

//更新用户密码
type UserPasswd struct {
	OldPasswd string `json:"old_passwd"`
	NewPasswdOne string `json:"new_passwd_one"`
	NewPasswdTwo string `json:"new_passwd_two"`
}
func UpdateUserPasswd(c *gin.Context){
	userpasswd := UserPasswd{}
	err := c.ShouldBindJSON(&userpasswd)
	userid,_ :=c.Get("userid")
	var RowsId  int64
	var oldPassword string
	if err != nil{
		code = 1
		msg = err.Error()
	}else {
		if userpasswd.NewPasswdOne != userpasswd.NewPasswdTwo{
			code = 1
			msg = "两次密码不想等"
		}else{
			row := Db.QueryRow("select password from sparrow.sp_user where id = ?",userid)
			row.Scan(&oldPassword)
			if oldPassword != "" &&  !utils.Decryption(oldPassword,userpasswd.OldPasswd){
				code = 1
				msg = "原密码错误"
			}else{
				newPasswdStr := utils.Encryption(userpasswd.NewPasswdTwo)
				result,err := Db.Exec("update sparrow.sp_user set password = ? where id = ?",newPasswdStr,userid)
				if err != nil{
					code = 1
					msg = err.Error()
				}else {
					code = 0
					msg = "success"
					RowsId,_  = result.RowsAffected()
				}
			}

		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":RowsId,
		"msg":msg,
	})
}
