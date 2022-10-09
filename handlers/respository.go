package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"sparrow/utils"
	"strconv"
	"time"
)

//获取知识库列表
// flag值为 hot、all
func Getrespositorys(c *gin.Context) {
	var (
		id             int
		name           string
		catename       string
		cateid         string
		groupid        int
		replist        []interface{}
		followSum      int
		docSum         int
		respUniqueCode string
		rJson          map[string]interface{}
	)
	flag := c.Query("flag")     // flag类型
	repoType := c.Query("type") // flag类型
	userid, _ := c.Get("userid")
	if flag != "" {
		//获取热门知识库
		if flag == "hot" {
			var repocodes = []string{}
			var repocode string
			var liked int
			rows, err := Db.Query("SELECT DISTINCT(repo_unique_code),count(liked) AS l FROM sparrow.sp_article GROUP BY repo_unique_code ORDER BY l DESC limit 10")
			if err != nil {
				utils.Logger.Error("获取热门知识库失败", zap.String("error", err.Error()))
			} else {
				for rows.Next() {
					_ = rows.Scan(&repocode, &liked)
					if repocode != "" {
						repocodes = append(repocodes, repocode)
					}
				}
			}
			if len(repocodes) == 0 {
				rJson = ReturnData(1, "", "data is null")
			} else {
				//repoidsStr := strings.Join(repocodes,",")
				repoidsStr := ""
				for index, code := range repocodes {
					if index+1 == len(repocodes) {
						repoidsStr += "'" + code + "'"
					} else {
						repoidsStr += "'" + code + "'" + ","
					}

				}
				sql := fmt.Sprintf("SELECT id, repo_name,repo_unique_code FROM sparrow.sp_repository WHERE repo_unique_code  in (%s) and repo_status = 1", repoidsStr)
				rows, err := Db.Query(sql)
				if err != nil {
					utils.Logger.Error("获取热门知识库错误", zap.String("error", err.Error()))
					rJson = ReturnData(1, "", err.Error())
				} else {
					for rows.Next() {
						var temdata = make(map[string]interface{})
						rows.Scan(&id, &name, &respUniqueCode)
						temdata["id"] = id
						temdata["name"] = name
						temdata["repo_unique_code"] = respUniqueCode
						replist = append(replist, temdata)
					}
					defer rows.Close()
					rJson = ReturnData(0, replist, "success")
				}
			}
		}
	} else {
		/*
			1、先获取用户所有组
			2、获取团队的知识库
			3、获取自己的知识库
		*/
		var (
			groupUniqueCode string
			gId             int64
		)
		rows, err := Db.Query("select a.group_unique_code,b.id from sp_user_group_member  a left join sp_user_group b on "+
			" a.group_unique_code = b.group_unique_code where a.userid = ?", userid)
		if err != nil {
			utils.Logger.Error("获取用户组错误", zap.String("error", err.Error()))
			rJson = ReturnData(1, "", err.Error())
		} else {
			for rows.Next() {
				err = rows.Scan(&groupUniqueCode, &gId)
				if err == nil {
					var sql string
					sql = "SELECT a.id, a.repo_name,b.catename,b.id,a.repo_user_group,a.repo_unique_code FROM sp_repository a " +
						" left join sparrow.sp_repo_cate b on a.repo_cate = b.id WHERE a.repo_user_group =? and a.userid != ? "
					if repoType == "doc" {
						sql = sql + " and a.repo_cate = 1"
					}
					repoRows, err := Db.Query(sql, gId, userid)
					if err != nil {
						utils.Logger.Error("读取组知识库错误", zap.String("error", err.Error()))
					} else {
						for repoRows.Next() {
							var temdata = make(map[string]interface{})
							err = repoRows.Scan(&id, &name, &catename, &cateid, &groupid, &respUniqueCode)
							fmt.Println(id, name, catename, cateid)
							if err == nil {
								followRow := Db.QueryRow("select count(1) as count from sparrow.sp_follow where follow_type = 1 and follow_id = ?", id)
								docRow := Db.QueryRow("select count(1) as count from sparrow.sp_article where repository_id = ?", id)

								_ = followRow.Scan(&followSum)
								_ = docRow.Scan(&docSum)
								temdata["id"] = id
								temdata["name"] = name
								temdata["followsum"] = followSum
								temdata["docsum"] = docSum
								temdata["catename"] = catename
								temdata["cateid"] = cateid
								temdata["groupid"] = groupid
								temdata["repo_unique_code"] = respUniqueCode
								replist = append(replist, temdata)
							}
						}
						defer repoRows.Close()
					}
				}
			}
			defer rows.Close()
			// 获取自己的知识库
			var selfsql string
			selfsql = "SELECT a.id, a.repo_name,b.catename,b.id,a.repo_user_group,a.repo_unique_code FROM sparrow.sp_repository as a " +
				" join sparrow.sp_repo_cate as b on a.repo_cate = b.id WHERE a.userid  =? and a.state = 1"
			if repoType == "doc" {
				selfsql = selfsql + " and a.repo_cate = 1"
			}
			selfRows, err := Db.Query(selfsql, userid)
			if err != nil {
				utils.Logger.Error("获取知识库错误", zap.String("error", err.Error()))
				rJson = ReturnData(1, "", err.Error())
			} else {
				for selfRows.Next() {
					var temdata = make(map[string]interface{})
					err = selfRows.Scan(&id, &name, &catename, &cateid, &groupid, &respUniqueCode)
					if err == nil {
						followRow := Db.QueryRow("select count(1) as count from sparrow.sp_follow where follow_type = 1 and follow_id = ?", id)
						docRow := Db.QueryRow("select count(1) as count from sparrow.sp_article where repository_id = ?", id)

						_ = followRow.Scan(&followSum)
						_ = docRow.Scan(&docSum)
						temdata["id"] = id
						temdata["name"] = name
						temdata["followsum"] = followSum
						temdata["docsum"] = docSum
						temdata["catename"] = catename
						temdata["cateid"] = cateid
						temdata["groupid"] = groupid
						temdata["repo_unique_code"] = respUniqueCode
						replist = append(replist, temdata)
					}

				}
				defer rows.Close()
			}
			rJson = ReturnData(0, replist, "success")
		}
	}

	c.JSON(http.StatusOK, rJson)
}

type repository struct {
	Id              int
	Resp_name       string
	Resp_desc       string
	Resp_cate       string
	Resp_status     int
	Resp_user_group int
	Createtime      *time.Time
	Userid          int64
	Username        string
}

/*
获取知识库详情
1、知识库是否公开
2、不公开等话，用户是否属于该知识库
3、不属于返回error
4、用户是否收藏过该知识库
*/
func GetrespositoryInfo(c *gin.Context) {
	repoflag := c.Param("repoflag")
	userid, _ := c.Get("userid")
	var (
		repositoryInfo = repository{}
		groupUserid    int64
		followId       int64
		data           = make(map[string]interface{})
		rJson          = make(map[string]interface{})
	)
	row := Db.QueryRow("select a.id,a.repo_name,a.repo_desc,a.repo_cate,a.repo_status,"+
		"a.repo_user_group,a.createtime,a.userid,b.username from sp_repository a left join sp_user b on a.userid = b.id  where a.repo_unique_code = ?", repoflag)
	err := row.Scan(&repositoryInfo.Id, &repositoryInfo.Resp_name, &repositoryInfo.Resp_desc,
		&repositoryInfo.Resp_cate, &repositoryInfo.Resp_status,
		&repositoryInfo.Resp_user_group, &repositoryInfo.Createtime, &repositoryInfo.Userid, &repositoryInfo.Username)
	if err != nil {
		rJson = ReturnData(1, "", err.Error())
	} else {
		followStateFunc := func() bool {
			followRow := Db.QueryRow("select id from sp_follow where userid = ? and follow_id = ?", userid, repoflag)
			if err = followRow.Scan(&followId); err != nil {
				return false
			} else {
				return true
			}
		}
		data = utils.Struct2Map(repositoryInfo)
		// 该用户自己的知识库获取知识库为开放状态 1open 0 private 判断用户是否为所属组拥有者 是
		if repositoryInfo.Resp_status == 1 {
			if repositoryInfo.Userid == userid {
				log.Println("知识库开放并属于该用户")
				data["isAuthorOrMember"] = true
				data["follow"] = true
				rJson = ReturnData(0, data, "success")
			} else {
				exitsRow := Db.QueryRow("select id from sp_user_group_member where group_unique_code = ? and userid = ?", repoflag, userid)
				err := exitsRow.Scan(&groupUserid)
				if err == nil && groupUserid != 0 {
					log.Println("知识库包含该用户")
					data["isAuthorOrMember"] = true
					data["follow"] = followStateFunc()

				} else {
					data["isAuthorOrMember"] = false
				}
				rJson = ReturnData(0, data, "success")
			}
		} else {
			if repositoryInfo.Userid == userid {
				log.Println("知识库开放并属于该用户")
				data["isAuthorOrMember"] = true
				data["follow"] = true
				rJson = ReturnData(0, data, "success")
			} else {
				// 否 判断是否属于该知识库组
				exitsRow := Db.QueryRow("select id from sp_user_group_member where group_unique_code = ? and userid = ?", repoflag, userid)
				err := exitsRow.Scan(&groupUserid)
				if err == nil && groupUserid != 0 {
					log.Println("知识库包含该用户")
					data = utils.Struct2Map(repositoryInfo)
					data["isAuthorOrMember"] = true
					data["follow"] = followStateFunc()
					rJson = ReturnData(0, data, "success")
				} else {
					rJson = ReturnData(2, "", "没有权限查看该知识库")
				}
			}
		}
	}
	c.JSON(http.StatusOK, rJson)
}

func IsFollow() {

}

// 获取指定知识库的知识列表
// 判断知识库分类 1、文档 2、资料管理
func GetRidrespositorys(c *gin.Context) {
	var (
		resp_cate      int
		id             int
		title          string
		username       string
		repositoryName string
		ar_unique_code string
		md5Str         string

		filename string

		size           string
		fileDirLevel   string
		fileExt        string
		fileDirLevelId int
		filecateName   string
		count          int
		createtime     *time.Time
		datas          []interface{}
		rJson          map[string]interface{}
	)
	repoflag := c.Param("repoflag")
	fileDirLevel = c.DefaultQuery("file_dir_level", "0")
	kw := c.DefaultQuery("kw", "")
	offset := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("size", "10")
	intOffset, _ := strconv.Atoi(offset)
	intSize, _ := strconv.Atoi(pageSize)
	fileDirLevelId, _ = strconv.Atoi(fileDirLevel)
	row := Db.QueryRow("select repo_cate from sparrow.sp_repository where repo_unique_code = ?", repoflag)
	err := row.Scan(&resp_cate)
	if err == nil {
		switch resp_cate {
		case 1: // 文档
			countRow := Db.QueryRow("select count(1) from sp_article where repo_unique_code = ? and state = 1 ", repoflag)
			if err := countRow.Scan(&count); err != nil {
				utils.Logger.Error("统计"+repoflag+"知识库知识总数出错", zap.String("error", err.Error()))
			}
			sql := "select a.id,a.ar_title,a.createtime,b.username,r.repo_name,a.ar_unique_code from sparrow.sp_article a " +
				"left join sparrow.sp_user b on a.userid=b.id left join sparrow.sp_repository r on " +
				"a.repo_unique_code = r.repo_unique_code where a.repo_unique_code = ? and a.state = 1"
			if kw != "" {
				sql += " and a.ar_title like '" + kw + "%'"
			}
			sql += fmt.Sprintf(" limit %d,%d", intSize*(intOffset-1), intSize)
			rows, err := Db.Query(sql, repoflag)
			if err != nil {
				utils.Logger.Error("统计文档信息列表出错", zap.String("error", err.Error()))
				rJson = ReturnData(1, "", err.Error())
			} else {
				for rows.Next() {
					data := make(map[string]interface{})
					rows.Scan(&id, &title, &createtime, &username, &repositoryName, &ar_unique_code)
					data["id"] = id
					data["title"] = title
					data["username"] = username
					data["repository_name"] = repositoryName
					data["ar_unique_code"] = ar_unique_code
					data["createtime"] = createtime.Format("2006-01-02 15:04")
					datas = append(datas, data)
				}
				rows.Close()
				rJson = ReturnDataTwo(0, datas, "success", count)
			}
		case 2: // 文件
			countRow := Db.QueryRow("select count(1) from sp_doc where repo_unique_code = ? and state = 1 ", repoflag)
			if err := countRow.Scan(&count); err != nil {
				utils.Logger.Error("统计"+repoflag+"知识库知识总数出错", zap.String("error", err.Error()))
			}
			//sql := "select s1.id,s1.filename,s1.filetype,s1.createtime,s1.size,s2.name,s2.id from sparrow.sp_files s1 " +
			//	"left join sparrow.sp_filecate s2 on s1.file_dir_level = s2.id where s1.repo_unique_code = ? and s1.file_dir_level = ?"
			sqlStr := "select d.id,d.doc_md5,d.doc_name,f.filetype,d.createtime,f.size,s2.name,f.id from sp_doc as d left join sp_files f on d.doc_md5 = f.md5 " +
				" left join sp_filecate s2 on d.file_dir_level = s2.id where d.repo_unique_code = ? and d.file_dir_level = ?"
			if kw != "" {
				sqlStr += " and d.doc_name like '" + kw + "%'"
			}
			sqlStr += fmt.Sprintf(" limit %d,%d", intSize*(intOffset-1), intSize)

			rows, err := Db.Query(sqlStr, repoflag, fileDirLevelId)
			fmt.Println(sqlStr)
			if err != nil {
				log.Println(err.Error())
				rJson = ReturnData(1, "", err.Error())
			} else {
				for rows.Next() {
					tmpfileExt := sql.NullString{String: "", Valid: false}
					data := make(map[string]interface{})
					rows.Scan(&id, &md5Str, &filename, &tmpfileExt, &createtime, &size, &filecateName, &fileDirLevelId)
					data["id"] = id
					data["filename"] = filename
					data["createtime"] = createtime.Format("2006-01-02 15:04")
					data["size"] = size
					//data["filecatename"] = ""
					data["filecatename"] = filecateName
					data["file_dir_level"] = fileDirLevelId
					data["fileext"] = fileExt
					data["md5Str"] = md5Str
					data["filetype"] = "file"
					datas = append(datas, data)
				}
				rows.Close()
				rJson = ReturnDataTwo(0, datas, "success", count)
			}
			// 获取同级的目录
			filepathRows, err := Db.Query("select `id`,`name`,`createtime` from sp_filecate where repo_unique_code = ? and parent_id = ?", repoflag, fileDirLevelId)
			if err != nil {
				log.Println(err.Error())
				rJson = ReturnData(1, "", err.Error())
			} else {
				fmt.Println(">>>>>>")
				for filepathRows.Next() {
					data := make(map[string]interface{})
					filepathRows.Scan(&id, &username, &createtime)
					data["id"] = id
					data["filename"] = username
					data["createtime"] = createtime.Format("2006-01-02 15:04")
					data["size"] = ""
					data["filecatename"] = "文件夹"
					data["filecateid"] = ""
					data["filetype"] = "dir"
					datas = append(datas, data)
				}
				defer filepathRows.Close()
				rJson = ReturnDataTwo(0, datas, "success", count)
			}
		}
	} else {
		utils.Logger.Error("获取知识库分类信息出错", zap.String("error", err.Error()))
		rJson = ReturnData(1, "", err.Error())
	}
	c.JSON(http.StatusOK, rJson)
}

/*
删除指定知识库
*/
func Delrespositorys(c *gin.Context) {
	var (
		rJson map[string]interface{}
	)
	repoflag := c.Param("repoflag")
	_, err := Db.Exec("update  sp_repository  set state = 0 where repo_unique_code = ? ", repoflag)
	if err != nil {
		rJson = ReturnData(1, "", err.Error())
	} else {
		/*
			删除该知识库的所有内容
			1、清除sp_doc、sp_article
			2、todo 用户关注或者收藏的怎么处理？
		*/
		cleanAllDoc := func(repoUniqueCode string) {
			var id uint64
			//清除sp_doc
			docRows, err := Db.Query("select id from sp_doc where repo_unique_code= ? and state = 1", repoflag)
			if err != nil {
				utils.Logger.Error("获取sp_doc知识库数据失败", zap.String("error", err.Error()))
			} else {
				for docRows.Next() {
					if err := docRows.Scan(&id); err != nil {
						_, err = Db.Exec("update sp_doc set state = 0 where id= ?", id)
						if err != nil {
							utils.Logger.Error("删除知识库sp_doc数据失败", zap.String("error", err.Error()))
						}
					}
				}
				defer docRows.Close()
			}
			//清除sp_article
			articleRows, err := Db.Query("select id from sp_article where repo_unique_code= ? and state = 1", repoflag)
			if err != nil {
				utils.Logger.Error("获取sp_article知识库数据失败", zap.String("error", err.Error()))
			} else {
				for articleRows.Next() {
					if err := docRows.Scan(&id); err != nil {
						_, err = Db.Exec("update sp_article set state = 0 where id= ?", id)
						if err != nil {
							utils.Logger.Error("删除知识库sp_article数据失败", zap.String("error", err.Error()))
						}
					}
				}
				defer articleRows.Close()
			}
		}
		go cleanAllDoc(repoflag)
		rJson = ReturnData(0, "", "success")
	}
	c.JSON(http.StatusOK, rJson)
}

//添加文档资源库
func Addrepository(c *gin.Context) {
	var addparams struct {
		Auth     int    `json:"auth"`
		Desc     string `json:"desc"`
		Reponame string `json:"reponame"`
		Repotype int    `json:"repotype"`
		Team     int    `json:"team"`
	}
	userid, _ := c.Get("userid")
	fmt.Println("usreid=", userid)
	var params = addparams
	err := c.ShouldBind(&params)
	if err != nil {
		log.Printf(err.Error())
	}
	respUniqueCode := utils.GetRandomString(8)
	//if params.Team
	_, err = Db.Exec("insert into sparrow.sp_repository ( repo_name, repo_desc, repo_cate, repo_status, repo_user_group,userid,repo_unique_code) values (?,?,?,?,?,?,?) ",
		params.Reponame, params.Desc, params.Repotype, params.Auth, params.Team, userid, respUniqueCode)
	if err != nil {
		code = 1
		msg = err.Error()
	} else {
		code = 0
		msg = "success"
	}
	fmt.Println(params)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": "",
		"msg":  msg,
	})
}
