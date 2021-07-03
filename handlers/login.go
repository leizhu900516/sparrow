package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"sparrow/utils"
	_"sparrow/utils"
	"sparrow/models"
)

//type loginParams struct {
//	Username string 	`form:"username" json:"username"`
//	Password string 	`form:"password" json:"password"`
//}

func Login(c *gin.Context){
	var userinfo models.UserInfo
	if err := c.ShouldBindJSON(&userinfo);err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(userinfo)
	var msg string
	var token string = "testtokenstr"
	//var err error
	var code uint8
	var username string
	var password string
	var userid uint
	row := Db.QueryRow("select id,username,password from sp_user where username = ?",userinfo.Username)
	err = row.Scan(&userid,&username,&password)
	if err != nil{
		fmt.Printf("scan failed, err:%v",err)
		msg = "认证失败,用户名和密码错误"
		code  = 1
	}else {
		authPwd := utils.Decryption(password,userinfo.Password)
		if authPwd{
			msg="认证成功"
			code  = 0
			token,err =utils.NewJwt(10,userid)
		}else{
			msg = "认证失败,用户名和密码错误"
			code  = 1
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":token,
		"msg":msg,
	})
}

// 用户注册
func Register(c *gin.Context) {
	var (
		id int
		userinfo models.UserInfo
		LastInsertId int64
		err error
	)
	if err := c.ShouldBindJSON(&userinfo);err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":code,
			"data":"",
			"msg":err.Error(),
		})
	}
	row := Db.QueryRow("select id from sp_user where username = ?",userinfo.Username)
	if err = row.Scan(&id);err != nil{
		utils.Logger.Error("注册用户错误,可能用户已存在,或者查询数据库错误"+err.Error())
		if id != 0{
			code = 1
			msg = "用户名已被注册"
		}else {
			passwdHash, _ := bcrypt.GenerateFromPassword([]byte(userinfo.Password), bcrypt.DefaultCost)
			result,err := Db.Exec("insert into  sp_user (username,password) values (?,?) ",
				userinfo.Username,passwdHash)

			if err != nil{
				code = 1
				msg = err.Error()
			}else {
				// 同时创建一个默认知识库
				respUniqueCode := utils.GetRandomString(8)
				fmt.Println("respUniqueCode=",respUniqueCode)
				LastInsertId,_ = result.LastInsertId()
				fmt.Println("LastInsertId=",LastInsertId)

				resultRepo,err := Db.Exec("insert into sp_repository (repo_name,repo_cate,repo_user_group,userid,repo_unique_code) " +
					"values ('默认知识库',1,-1,?,?)",LastInsertId,respUniqueCode)
				if err != nil{
					utils.Logger.Error("创建默认知识库失败"+err.Error())
				}else{
					LastInsertIdRepo,_ := resultRepo.LastInsertId()
					fmt.Println(LastInsertIdRepo)
					utils.Logger.Info("创建默认知识库成功" )
				}
				code = 0
				msg = "success"

			}
		}
	}else {
		code = 1
		msg = "用户名已被注册，去登录"
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":LastInsertId,
		"msg":msg,
	})
}
