package handlers

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"net/http"
	"sparrow/utils"
	"strings"
	"time"
)


// 上传文件，得到md5值，根据日期进行规整 例如 20201020/jajgkagjgagja
// 返回文件地址
// 文件存储下来
// 上传图片专用函数
func BookUpload(c *gin.Context){
	var (
		code int = 0
		data string
		msg string = "success"
		files []*multipart.FileHeader
		fileurl []string
		md5List []string
		filename string
		size string
		ip string = "127.0.0.1"
		domain string = "127.0.0.1"
	)
	fileType := c.PostForm("filetype")
	userid, _ := c.Get("userid")
	if fileType == "img"{

	}
	if fileType == "zip" {

	}
	//filename:= c.PostForm("filename")
	filecate:= c.PostForm("filecate")
	//fmt.Println(">>>",filename,filecate,)
	file,err := c.FormFile("file")
	avator,err := c.FormFile("avator")
	files = append(files,file)
	files = append(files,avator)
	if err != nil{
		code = 1
		msg = "fail"
	}
	for index,f := range files{
		if index == 0{
			filename = f.Filename
		}
		_filename := f.Filename
		size =  utils.FormatFileSize(f.Size)
		ext :=strings.Split(_filename,".")[1]
		md5c := md5.New()
		out,_ := f.Open()
		io.Copy(md5c,out)
		Md5Str := hex.EncodeToString(md5c.Sum(nil))
		defer out.Close()
		nowDate :=  time.Now().Format("2006-01-02")
		parentPath := DataDir+nowDate
		utils.MakeDir(parentPath)
		var DataBuffer bytes.Buffer
		DataBuffer.WriteString(parentPath)
		DataBuffer.WriteString("/")
		DataBuffer.WriteString(Md5Str)
		DataBuffer.WriteString(".")
		DataBuffer.WriteString(ext)
		FullPath := DataBuffer.String()
		fmt.Println(FullPath)
		exits := utils.Isfile(FullPath)
		if !exits{
			if err := c.SaveUploadedFile(f,FullPath); err != nil {
				utils.Logger.Error("保存失败 Error:%s"+err.Error())
				//c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
				code = 1
				msg = "fail"
			}
		}
		fileurl = append(fileurl,FullPath)
		md5List = append(md5List,Md5Str)
		_,err := Db.Exec("insert into sparrow.sp_files (`filename`,`filetype`,`md5`,`fileurl`,`ip`,`domain`,`userid`,`size`,`repo_unique_code`,`file_dir_level`) " +
			"values (?,?,?,?,?,?,?,?,?,?)",_filename,fileType,Md5Str,FullPath,ip,domain,userid,size,"avator",0)
		if err != nil{
			utils.Logger.Error("删除文件出错",zap.String("filename",filename))
		}
	}
	if len(fileurl) == 2 && len(md5List) == 2{
		_,err =Db.Exec("insert into sp_book(`b_name`,`b_url`,`b_avator_url`,`b_cate_id`,`b_userid`,`book_md5`,`avator_md5`) values(?,?,?,?,?,?,?)",
			filename,fileurl[0],fileurl[1],filecate,userid,md5List[0],md5List[1])
		if err != nil{
			utils.Logger.Error("插入图书信息错误"+err.Error())
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"msg":msg,
	})
}

//  知识库下载文件接口
func Download(c *gin.Context){
	var (
		filename string
		fileurl string
	)
	fileMd5 := c.Param("md5")
	if fileMd5 == ""{
		c.JSON(http.StatusOK,gin.H{
			"code":0,
			"data":"",
			"msg":"md5不存在",
		})
	}
	fileInfo := Db.QueryRow("select filename,fileurl from sparrow.sp_files where md5 = ?",fileMd5)
	if fileInfo == nil{
		c.JSON(http.StatusOK,gin.H{
			"code":0,
			"data":"",
			"msg":"文件不存在",
		})
	}
	err :=fileInfo.Scan(&filename,&fileurl)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":0,
			"data":"",
			"msg":err.Error(),
		})
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(fileurl)
}


// 图书的专用下载接口
func BookDownload(c *gin.Context) {

	var (
		filename string
		fileurl string
	)
	fileId := c.Param("id")
	if fileId == ""{
		c.JSON(http.StatusOK,gin.H{
			"code":0,
			"data":"",
			"msg":"请求错误",
		})
	}else {
		fileInfo := Db.QueryRow("select b_name,b_url from sparrow.sp_book where id = ?",fileId)
		if fileInfo == nil{
			c.JSON(http.StatusOK,gin.H{
				"code":0,
				"data":"",
				"msg":"文件不存在",
			})
		}else {
			err :=fileInfo.Scan(&filename,&fileurl)
			if err != nil{
					c.JSON(http.StatusOK,gin.H{
						"code":1,
						"data":"",
						"msg":err.Error(),
					})
			}else {
				_,err :=Db.Exec("update  sparrow.sp_book set download= download+1 where id = ?",fileId)
				if err != nil{
					utils.Logger.Error("下载文件，新增下载次数失败 Error:%s"+err.Error())
				}

				fileurlList := strings.Split(fileurl,".")
				fullFileName :=  filename+"."+fileurlList[1]
				fmt.Println(fullFileName)
				c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s",fullFileName))
				c.Writer.Header().Add("Content-Type", "application/octet-stream")
				c.File(fileurl)
			}
		}
	}
}

func Find(val string,slice []string) ( bool) {
	for _, item := range slice {
		if item == val {
			return  true
		}
	}
	return  false
}

func Upload(c *gin.Context){
	var (
		//fileurl []string
		filepath string
		fileType string
		ip string = "127.0.0.1"
		domain string = "127.0.0.1"
		size string
		userid interface{}
		imgType =[]string{"jpg","png","gif","webp","svg","apng","jpeg"}
		//othersFileType =[]string{"zip","tar","gz","pdf"}
		returnData = make(map[string]interface{})
	)
	userid ,_ = c.Get("userid")
	file_dir_level:= c.PostForm("file_dir_level")  //获取当前目录层级
	repoUniqueCode:= c.PostForm("repo_unique_code")  //获取当前目录层级
	fmt.Println(">>>",file_dir_level,filepath)
	file,err := c.FormFile("file")
	if err != nil{
		returnData = ReturnData(1,"","上传文件失败" + err.Error())
		c.JSON(http.StatusOK,returnData)
		return
	}

	_filename := file.Filename
	filesep := strings.Split(_filename,".")
	fileType = filesep[1]
	if Find(fileType,imgType) {
		filepath = DataDir + "img/"
	}else{
		filepath = DataDir + "file/"
	}
	//if Find(fileType,othersFileType) {
	//	filepath = "files/file/"
	//}
	size =  utils.FormatFileSize(file.Size)
	ext :=strings.Split(_filename,".")[1]
	md5c := md5.New()
	out,_ := file.Open()
	io.Copy(md5c,out)
	nowDate :=  time.Now().Format("2006-01-02")
	parentPath := filepath+nowDate
	fmt.Println("parentPath=",parentPath)
	utils.MakeDir(parentPath)
	Md5Str := hex.EncodeToString(md5c.Sum(nil))
	var DataBuffer bytes.Buffer
	DataBuffer.WriteString(parentPath)
	DataBuffer.WriteString("/")
	DataBuffer.WriteString(Md5Str)
	DataBuffer.WriteString(".")
	DataBuffer.WriteString(ext)
	FullPath := DataBuffer.String()
	fmt.Println(FullPath)
	exits := utils.Isfile(FullPath)
	if !exits{
		if err := c.SaveUploadedFile(file,FullPath); err != nil {
			c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
			returnData = ReturnData(1,"",err.Error())
		} else {
			fmt.Println("写入文件库")
			//fileurl = append(fileurl,FullPath)
			//  写入文件库 先判断是否存在该文件
			var fileId int
			hasMd5Row := Db.QueryRow("select id from sp_files where md5 = ?",Md5Str)
			if err = hasMd5Row.Scan(&fileId);err != nil{
				utils.Logger.Info("文件已经存在，只需更新文档表即可")
				_data :=joinImgPath(Md5Str,fileType,FullPath,filepath,imgType)
				returnData = ReturnData(1,_data,"文件已存在")
			}else {
				var filesId int64
				_exits:= Db.QueryRow("select id from sp_files where md5 = ?",Md5Str)
				if err := _exits.Scan(&filesId);err != nil{
					_,err := Db.Exec("insert into sparrow.sp_files (`filename`,`filetype`,`md5`,`fileurl`,`ip`,`domain`,`userid`,`size`,`repo_unique_code`,`file_dir_level`) " +
						"values (?,?,?,?,?,?,?,?,?,?)",_filename,fileType,Md5Str,FullPath,ip,domain,userid,size,repoUniqueCode,file_dir_level)
					if err != nil{
						utils.Logger.Error("写入文件出错",zap.String("error",err.Error()))
						returnData = ReturnData(1,"",err.Error())
					}else{
						_data :=joinImgPath(Md5Str,fileType,FullPath,filepath,imgType)
						returnData = ReturnData(0,_data,"success")
						//data ,_= result.RowsAffected()
					}
				} else {
					_data :=joinImgPath(Md5Str,fileType,FullPath,filepath,imgType)
					returnData = ReturnData(0,_data,"success")
				}

			}

		}
	}else{
		_data :=joinImgPath(Md5Str,fileType,FullPath,filepath,imgType)
		returnData = ReturnData(1,_data,"文件已存在")
	}
	// 写入资料知识库，先判断是否存在
	var docId int
	_exits:= Db.QueryRow("select id from sp_doc where doc_md5 = ? and userid = ? and file_dir_level = ?",Md5Str,userid,file_dir_level)
	if err = _exits.Scan(&docId);err != nil{
		_,err = Db.Exec("insert into sp_doc (doc_md5,doc_name,repo_unique_code,file_dir_level,userid) value (?,?,?,?,?)",Md5Str,_filename,repoUniqueCode,file_dir_level,userid)
		if err != nil{
			utils.Logger.Error("写入资料知识库错误",zap.String("error",err.Error()))
		}
	}else {
		utils.Logger.Info("已经存在文件记录")
	}
	c.JSON(http.StatusOK,returnData)
}

// 拼接图片地址
func joinImgPath(Md5Str string,fileType string,FullPath string,filepath string,imgType []string) map[string]interface{}{
	_data := make(map[string]interface{})
	_data["md5"] = Md5Str
	if Find(fileType,imgType){
		_data["url"] =  "api/v1/img/" + Md5Str
	}else {
		_data["url"] =  strings.Replace(FullPath,filepath,"",-1)
	}
	return ReturnData(0,_data,"success")
}
