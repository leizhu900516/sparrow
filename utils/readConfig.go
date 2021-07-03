package utils

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"go.uber.org/zap"
)
var (
	err error
	Config *goconfig.ConfigFile
)

const ConfigName  = "sparrow"
func ParseConfig(config string) {
	fmt.Println("获取配置文件",config)
	Logger.Info("开始解析配置文件",zap.String("file=",config))
	if config == ""{
		Config,err = goconfig.LoadConfigFile("config/config.ini")
		if err !=nil{
			Logger.Error("配置文件解析错误",zap.String("error",err.Error()))
			panic("配置文件解析错误"+err.Error())
		}
	}else {
		Config,err = goconfig.LoadConfigFile(config)
		if err !=nil{
			Logger.Error("配置文件解析错误",zap.String("error",err.Error()))
			panic("配置文件解析错误"+err.Error())
		}
	}
	Logger.Info("解析配置文件成功")
}
//func init(){
//	Config,err = goconfig.LoadConfigFile("utils/config.ini")
//	if err !=nil{
//		panic("配置文件解析错误")
//	}
//}
