package utils

import (
	"fmt"
	"log"
	"os"
	"runtime"
)



func MakeDir(path string){
	fmt.Println(path)
	_,err = os.Stat(path)
	if err != nil{
		err =os.MkdirAll(path,0755)
		if err != nil{
			log.Fatal("创建文件目录失败",err.Error())
		}
	}
}
// 创建文件存储路径
func CreateFilePath(){
	filepath,err := Config.GetValue(ConfigName,"filepath")
	if err != nil{
		log.Fatal("读取文件配置目录失败",err.Error())
	}
	systemName := runtime.GOOS
	switch systemName {
	case "darwin":
		MakeDir("./data")
	case "linux":
		MakeDir(filepath)
	default:
		MakeDir(filepath)
	}

}
