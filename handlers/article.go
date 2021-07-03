package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"sparrow/utils"
	"strconv"
	"time"
)

var article struct{
	Title string `json:"title"`
	Repoid string	`json:"repoid"`
	Content string `json:"content"`
	Puretext string `json:"puretext"`
	Aid string `json:"aid"`
}
// 添加文章
func AddArticle(c *gin.Context){
	var (
		id int64
		respUniqueCode string
		rJson map[string]interface{}
		private int
	)
	form := article
	if err = c.ShouldBind(&form);err !=nil{
		utils.Logger.Error("表单提交错误")
		rJson = ReturnData(1,"",err.Error())
	}else {
		userid,_ := c.Get("userid")
		// 存在即更新
		if form.Aid != ""{
			_,err =Db.Exec("update sparrow.sp_article set  ar_title= ?,ar_content = ?,ar_pure_content = ?  where ar_unique_code = ?",
				form.Title,form.Content,form.Puretext,form.Aid)
			if err != nil{
				rJson = ReturnData(1,"",err.Error())
			}else{
				rJson = ReturnData(0,"","success")
			}
		}else {
			fmt.Println("insert")
		GenCode:
			for{
				respUniqueCode = utils.GetRandomString(10)
				row:=Db.QueryRow("select id from sp_article where ar_unique_code = ?",respUniqueCode)
				err = row.Scan(&id)
				if err == nil{
					utils.Logger.Error("文章的唯一标识冲突，重新生成")
					goto GenCode
				}else {
					break
				}
			}
			// 查询知识库是否公开
			privateRow:= Db.QueryRow("select repo_status from sp_repository where repo_unique_code = ?",form.Repoid)
			err:= privateRow.Scan(&private)
			if err == nil{
				privateStr := strconv.Itoa(private)
				_,err =Db.Exec("insert into sparrow.sp_article (ar_title,ar_cate,ar_content,repo_unique_code,userid,ar_pure_content,ar_unique_code,private) values (?,1,?,?,?,?,?,?)",
					form.Title,form.Content,form.Repoid,userid,form.Puretext,respUniqueCode,privateStr)
				if err != nil{
					utils.Logger.Error(err.Error())
					rJson = ReturnData(1,"",err.Error())
				}else {
					rJson = ReturnData(0,"","success")
				}
			}else {
				rJson = ReturnData(1,"","查询知识库状态错误")
			}
		}
	}

	c.JSON(http.StatusOK,rJson)
}

/*
获取文章通用函数
@hot 是否是热门文章
@filter 是否过滤用户文章：true 获取用户全部文章 false 获取所有的公开文章
@
*/
func getArticles(c *gin.Context,db *sql.DB,offset,size int,hot string,filter bool) []interface{}{
	var (
		followId int
		id int
		goodid int
		title string
		createtime *time.Time
		content string
		pureContent string
		desc string // 截取content的前50个字符为简介信息
		userid int
		loginUserInt64 int64
		username string
		liked int
		catename string
		ar_unique_code string
		datas []interface{}
		rows *sql.Rows
		err error
		userCache = make(map[int]map[string]interface{})
	)
	type userInfo struct {
		Username string
		Profile string
		Avatarurl string
	}
	var user = userInfo{}
	userIdStr ,_ := c.Get("userid")

	switch userIdStr.(type) {
	case int64:
		loginUserInt64,_  = userIdStr.(int64)
	default:
		utils.Logger.Error("解析用户类型出错")
		fmt.Println(reflect.TypeOf(userIdStr))
		panic("解析用户类型出错")
	}
	fmt.Println("loginUserid=",loginUserInt64)
	if hot =="hot"{
		// 获取热门文章
		rows,err = db.Query("SELECT id, ar_title,createtime,liked,ar_unique_code " +
			"FROM sp_article where state=1  and  private = 1  order by liked desc limit ?,?",(offset-1)*size,size)
	}else if hot == "" {
		if filter{
			rows,err = db.Query("SELECT a.id, a.ar_title,a.ar_content,a.ar_pure_content,a.createtime, b.catename,a.userid,a.liked,a.ar_unique_code,c.username " +
				"FROM sp_article a JOIN sp_repo_cate b ON a.ar_cate =b.id  left join sp_user c on a.userid = c.id  " +
				"where a.state=1  and a.userid = ? " +
				"order by a.createtime desc limit ?,?",loginUserInt64,(offset-1)*size,size)
		} else {
			rows,err = db.Query("SELECT a.id, a.ar_title,a.ar_content,a.ar_pure_content,a.createtime, b.catename,a.userid,a.liked,a.ar_unique_code,c.username " +
				"FROM sp_article a JOIN sp_repo_cate b ON a.ar_cate =b.id  left join sp_user c on a.userid = c.id  where a.state=1 and  a.private = 1 " +
				"order by a.createtime desc limit ?,?",(offset-1)*size,size)
		}

	}
	if err != nil{
		utils.Logger.Error("获取文章信息失败:"+err.Error())
		return datas
	}
	for rows.Next(){
		var row = make(map[string]interface{})
		if hot =="hot" {
			err = rows.Scan(&id,&title,&createtime,&liked,&ar_unique_code)
		}else {
			err = rows.Scan(&id,&title,&content,&pureContent,&createtime,&catename,&userid,&liked,&ar_unique_code,&username)
		}
		if err != nil{
			utils.Logger.Error("获取文章信息失败:"+err.Error())
		}else {
			// 查询用户信息
			// 本地做缓存
			hasUser,ok:=userCache[userid]
			if ok{
				row["user"] = hasUser
			}else {
				var (
					followCount int
					beFollowCount int
				)
				userinfo := Db.QueryRow("select `username`,`profile`,`avatar_md5` from sp_user where id = ?",userid)
				err = userinfo.Scan(&user.Username,&user.Profile,&user.Avatarurl)
				if err == nil{
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
					var _user = make(map[string]interface{})
					userCache[userid] = _user
					_user["username"] = user.Username
					_user["followCount"] = followCount
					_user["beFollowCount"] = beFollowCount
					_user["profile"] = string([]rune(user.Profile)[:20])
					_user["avatarurl"] = "api/v1/img/"+user.Avatarurl
					row["user"] = _user
				}else {
					utils.Logger.Error("获取用户信息失败:"+err.Error())
					row["user"] = nil
				}
			}
			// 查询是否关注该用户
			//fmt.Println("userid=",userid,"loginUserInt64=",loginUserInt64)
			followInfo := Db.QueryRow("select id from sp_follow where follow_id = ? and userid = ? and follow_type = 0",userid,loginUserInt64)
			err = followInfo.Scan(&followId)
			if err != nil{
				row["isFollow"] = false
				utils.Logger.Error("获取是否关注该用户信息出错",zap.String("error",err.Error()))
			}else{
				row["isFollow"] = true
			}
			// 查询是否点赞
			goodrow := Db.QueryRow("select id from sparrow.sp_good_history where userid = ? and aritcle_unique_code = ?",userid,ar_unique_code)
			err:= goodrow.Scan(&goodid)
			if err != nil{
				row["isGood"] = false
			}else {
				row["isGood"] = true
			}

			row["id"] = id
			row["title"] = title
			row["content"] = content
			//row["pure_content"] = pureContent
			if len([]rune(pureContent)) <50{
				if len([]rune(pureContent)) == 0 {
					desc = "文章暂无介绍信息"
				}else{
					desc = pureContent
				}
			}else{
				desc  = string([]rune(pureContent)[:50])
			}
			row["desc"] = desc
			row["liked"] = liked
			row["createtime"] = createtime.Format("2006-01-02 15:04")
			row["userid"] = userid
			row["catename"] = catename
			row["ar_unique_code"] = ar_unique_code
			row["username"] = username
			datas = append(datas,row)
		}
	}
	defer rows.Close()
	return datas
}

// 获取文章列表
func GetArticle(c *gin.Context){
	var (
		datas []interface{}
		count int
	)
	offset := c.DefaultQuery("page","1")
	size := c.DefaultQuery("size","10")
	intOffset,_ := strconv.Atoi(offset)
	intSize,_ := strconv.Atoi(size)
	hot := c.Query("flag")
	datas = getArticles(c,Db,intOffset,intSize,hot,false)

	countRow := Db.QueryRow("select count(1) as count from sparrow.sp_article where state=1")
	_ = countRow.Scan(&count)
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":datas,
		"msg":msg,
		"count":count,
	})
}
// 获取自己所有的文章列表
func GetMySelfArticle(c *gin.Context){
	var (
		datas []interface{}
		count int
	)
	offset := c.DefaultQuery("page","1")
	size := c.DefaultQuery("size","10")
	intOffset,_ := strconv.Atoi(offset)
	intSize,_ := strconv.Atoi(size)
	hot := c.Query("flag")
	userid,_ :=c.Get("userid")
	datas = getArticles(c,Db,intOffset,intSize,hot,true)

	countRow := Db.QueryRow("select count(1) as count from sparrow.sp_article where state=1 and userid = ?",userid)
	_ = countRow.Scan(&count)
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":datas,
		"msg":msg,
		"count":count,
	})
}


func  GoodArticle(c *gin.Context)  {
	type aid struct {
		Uniquecode string
	}
	var (
		id int64

		isGood bool = false

		aidObj  = aid{}
	)
	userid,_ := c.Get("userid")
	err := c.ShouldBindJSON(&aidObj)
	if err != nil{
		fmt.Println(err.Error())
	}else {
		fmt.Println(aidObj)
		//查询是否点过赞
		row := Db.QueryRow("select id from sparrow.sp_good_history where userid = ? and aritcle_unique_code = ?",userid,aidObj.Uniquecode)
		err := row.Scan(&id)
		fmt.Println(err)
		if err != nil{
			_,err1 := Db.Exec("update sp_article set liked = liked+1 where ar_unique_code = ?",aidObj.Uniquecode)
			_,err2 := Db.Exec("insert into  sp_good_history (`userid`,`aritcle_unique_code`) values (?,?)",userid,aidObj.Uniquecode)
			if err1 != nil || err2 != nil{
				code = 1
				msg = err1.Error() +err2.Error()
			}else {
				code =0
				msg = "success"
			}
		}else {
			isGood = true
			code =1
			msg = "您已经点过赞"
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":isGood,
		"msg":msg,
	})
}

func GetArticleDesc(c *gin.Context){
	var (
		title string
		createtime *time.Time
		content string
		formattime string
		reponame string
		repocode string
		wikiuserid int64
		datas = make(map[string]interface{})
		collectId int64
		isCollect bool = false
	)
	userid ,_:= c.Get("userid")
	id := c.Param("id")
	// 查询是否收藏过
	collectRow := Db.QueryRow("select id from sparrow.sp_collect where userid = ? and resource_unique_code = ?",userid,id)
	err :=collectRow.Scan(&collectId)
	if err == nil{
		isCollect = true
	}
	sql := "SELECT  a.ar_title, a.ar_content, a.createtime,b.repo_name,b.repo_unique_code,a.userid  FROM sparrow.sp_article a " +
		"left join sparrow.sp_repository b on a.repo_unique_code = b.repo_unique_code WHERE a.ar_unique_code=?"
	row :=Db.QueryRow(sql,id)
	err =row.Scan(&title,&content,&createtime,&reponame,&repocode,&wikiuserid)
	if err != nil{
		utils.Logger.Error("获取文章详情出错"+err.Error())
	}
	formattime = createtime.Format("2006-01-02 13:04")
	datas["title"] = title
	datas["content"] = content
	datas["createtime"] = formattime
	datas["reponame"] = reponame
	datas["repocode"] = repocode
	datas["iscollect"] = isCollect
	if userid == wikiuserid {
		datas["isauthor"] = true
	}else {
		datas["isauthor"] = false
	}
	if err != nil{

	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":datas,
		"msg":"",
	})
}

// 删除文章
func DelArticle(c *gin.Context) {
	var ar_cate int
	ar_unique_code := c.Param("ar_unique_code")
	if ar_unique_code == ""{
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"data":"",
			"msg":"id不能为空",
		})
	}
	userid,_:=c.Get("userid")
	row := Db.QueryRow("select ar_cate from sparrow.sp_article where ar_unique_code = ?",ar_unique_code)
	err = row.Scan(&ar_cate)
	if err != nil{
		code = 1
		msg = err.Error()
	}else{
		_,err :=Db.Exec("update  sparrow.sp_article  set state=0 WHERE ar_unique_code=?",ar_unique_code)

		if err != nil{
			code = 1
			msg = err.Error()
		}else{
			_,err =Db.Exec("insert into  sparrow.sp_recycle_bin (`resource_unqiue_code`,`repo_cate`,`userid`) values (?,?,?)",ar_unique_code,ar_cate,userid)
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

//更新回收站的文章状态到正常
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	recycleId := c.Query("rid")
	var err error
	tx,_ := Db.Begin()
	_,err1 :=tx.Exec("update  sparrow.sp_article  set state=1 WHERE ar_unique_code=?",id)
	_,err2 :=Db.Exec("delete  from sparrow.sp_recycle_bin   WHERE id=?",recycleId)
	if err1 != nil || err2 != nil{
		err =tx.Rollback()
		utils.Logger.Error("恢复回收站的文章状态出错",zap.String("error:",  err.Error()))
		code = 1
		msg = "出错了"
	}else {
		err = tx.Commit()
		if err != nil{
			code = 1
			msg = err.Error()
		}else {
			code = 0
			msg = "success"
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":"",
		"msg":msg,
	})
}

// 获取用户关注的文章和知识库列表信息
func GetArticleFollow(c *gin.Context) {
	var (
		id int
		title string
		createtime *time.Time
		content string
		desc string // 截取content的前50个字符为简介信息
		liked int
		catename string
		articleList []interface{}
		teamList []interface{}
		datas = make(map[string]interface{})
		userid interface{}

		followtype int
		followid int
		articleIds []int
		teamIds []int

		groupflag int
		groupname string
	)
	userid,_ = c.Get("userid")
	offset := c.DefaultQuery("page","1")
	size := c.DefaultQuery("size","10")
	intOffset,_ := strconv.Atoi(offset)
	intSize,_ := strconv.Atoi(size)
	fmt.Println(offset,size)
	rows,err := Db.Query("SELECT id, follow_type, follow_id FROM sparrow.sp_follow where userid = ?",userid)
	if err != nil{
		code = 1
		msg = "获取文章列表失败"
	}
	for rows.Next(){
		_ = rows.Scan(&followtype,&followid)
		switch followid {
		case 0:
			articleIds = append(articleIds,followid)
		case 1:
			teamIds = append(teamIds,followid)
		}
	}
	defer rows.Close()
	articleRows,err:= Db.Query("SELECT a.id, a.ar_title,a.ar_content,a.createtime, b.catename,a.userid,a.liked " +
	"FROM sparrow.sp_article a JOIN sparrow.sp_repo_cate b ON a.ar_cate =b.id " +
		"where a.id in ? order by a.createtime desc limit ?,?",articleIds,(intOffset-1)*intSize,intSize)
	if err != nil{
		utils.Logger.Debug("获取文章列表数据失败", zap.String("error", err.Error()))
	}
	for articleRows.Next(){
		var row = make(map[string]interface{})
		_ = articleRows.Scan(&id,&title,&content,&createtime,&catename,&userid,&liked)
		row["id"] = id
		row["title"] = title
		row["content"] = content
		if len([]rune(content)) <50{
			desc = content
		}else{
			desc  = string([]rune(content)[:50])
		}
		row["desc"] = desc
		row["liked"] = liked
		row["createtime"] = createtime.Format("2006-01-02 15:04")
		row["userid"] = userid
		row["catename"] = catename
		articleList = append(articleList,row)
	}
	defer articleRows.Close()
	teamRows,err:= Db.Query("SELECT id, group_unique_code, group_name, createtime FROM sparrow.sp_user_group where id in ?",teamIds)
	if err != nil{
		utils.Logger.Debug("获取团队列表数据失败", zap.String("error", err.Error()))
	}
	for teamRows.Next(){
		var teamrow = make(map[string]interface{})
		_ = teamRows.Scan(&id,&groupflag,&groupname,&createtime)
		teamrow["id"] = id
		teamrow["groupflag"] = groupflag
		teamrow["groupname"] = groupname
		teamrow["createtime"] = createtime.Format("2006-01-02 15:04")
		teamList = append(teamList,teamrow)
	}
	datas["article"] = articleList
	datas["team"] = teamList
	defer teamRows.Close()
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":datas,
		"msg":msg,
	})
}

// 获取关注的人或者知识库的最近三天的最新动态
func FollowMoments(c *gin.Context){
	var (
		followType int
		followId string

		articleUniqueCode string
		articleName string
		createtime *time.Time
		username string
		avatarMd5 string
		rJson map[string]interface{}
		//moments = make(map[string]map[string]interface{})
		moments []map[string]interface{}
	)
	userid ,ok:= c.Get("userid")
	if ok{
		rows,err := Db.Query("select follow_type,follow_id from sp_follow where userid = ?",userid)
		if err != nil{
			rJson = ReturnData(1,"",err.Error())
		}else {
			for rows.Next(){
				if err := rows.Scan(&followType,&followId);err == nil{
					fmt.Println(followId,followType)
					threeBeforeTime := time.Now().AddDate(0,0,-3).Format("2006-01-02 03:04:05")
					fmt.Println("threeday=",threeBeforeTime)
					var threeRows *sql.Rows
					var err2 error
					switch followType {
					case 0:// 如果是用户，获取用户最新发布的文章和用户资料
						//userinfo := Db.QueryRow("select * from sp_user where id = ?",followId)
						sql:= fmt.Sprintf("select a.ar_unique_code,a.ar_title,a.createtime,b.username,b.avatar_md5 from sp_article a " +
							"left join sp_user b on a.userid = b.id  where a.userid = %s and a.createtime>= '%s' limit 1",followId,threeBeforeTime)
						fmt.Println(sql)
						threeRows ,err2 = Db.Query(sql)
					case 1:// 如果是知识库，获取该知识库最新的三天的知识
						sql := fmt.Sprintf("select a.ar_unique_code,a.ar_title,a.createtime,b.repo_name,b.repo_desc from sp_article a left join sp_repository b on " +
							"a.repo_unique_code = b.repo_unique_code where a.repo_unique_code = %s and createtime>= '%s' limit 1",followId,threeBeforeTime)
						threeRows ,err2 = Db.Query(sql)
					default:
						utils.Logger.Error("关注类型出错")
						errors.New("关注类型出错")
					}
					if err2 != nil{
						utils.Logger.Error("获取关注信息错误",zap.String("error",err.Error()))
					}else {
						for threeRows.Next() {
							fmt.Println(threeRows.Columns())
							if err3 := threeRows.Scan(&articleUniqueCode, &articleName, &createtime,&username,&avatarMd5); err3 != nil {
								utils.Logger.Error("解析关注信息错误",zap.String("error",err3.Error()))
							}else {
								fmt.Println(articleName, articleUniqueCode, createtime)
								_data := make(map[string]interface{})
								switch followType {
								case 0:
									var avatorurl string
									if avatarMd5 != "" {
										avatorurl = "api/v1/img/"+avatarMd5
									} else {
										avatorurl = ""
									}
									_data["username"] = username
									_data["avator"] = avatorurl
									_data["unique_code"] = articleUniqueCode
									_data["title"] = articleName
									_data["createtime"] = createtime.Format("2006-01-02 03:04")
									_data["flag"] = "user"
									moments= append(moments,_data)
								case 1:
									_data["username"] = username
									_data["avator"] = avatarMd5
									_data["unique_code"] = articleUniqueCode
									_data["title"] = articleName
									_data["createtime"] = createtime.Format("2006-01-02 03:04")
									_data["flag"] = "repo"
									moments= append(moments,_data)
								}

							}
						}
						defer threeRows.Close()
					}
				}
			}
			rJson = ReturnData(0,moments,"success")
			defer rows.Close()
		}

	}

	c.JSON(http.StatusOK,rJson)
}

//添加关注
func AddFollow(c *gin.Context){
	type Follow struct {
		Flag string `json:"flag"`
		Id int `json:"id"`
	}
	var (
		follow = Follow{}
		rJson  map[string]interface{}
		id int
	)

	err := c.ShouldBind(&follow)
	if err != nil{
		rJson = ReturnData(1,"",err.Error())
	}else {
		userid,_:=c.Get("userid")
		hasFollow := Db.QueryRow("select id from sp_follow where userid = ? and follow_id = ? and follow_type = 0",userid,follow.Id)
		err = hasFollow.Scan(&id)
		if err != nil && id == 0{
			switch follow.Flag {
			case "user":
				_,err =Db.Exec("insert into sp_follow (userid,follow_type,follow_id) values (?,?,?)",userid,0,follow.Id)
			case "repo":
				_,err =Db.Exec("insert into sp_follow (userid,follow_type,follow_id) values (?,?,?)",userid,1,follow.Id)
			default:
				rJson = ReturnData(1,"","收藏类型错误")
			}
			if err != nil{
				rJson = ReturnData(1,"",err.Error())
			}else{
				rJson = ReturnData(0,"","success")
			}
		} else {
			utils.Logger.Error("已经关注")
			rJson = ReturnData(1,"","已经关注")
		}
	}
	c.JSON(http.StatusOK,rJson)
}
