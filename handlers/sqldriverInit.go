package handlers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"sparrow/utils"
	"time"
)

const mysqlname  = utils.ConfigName
var (
	Db *sql.DB
	err error
	DataDir string
	//Config *goconfig.ConfigFile = utils.Config
)

func DbConnectInit()  {
	username,_ := utils.Config.GetValue(mysqlname,"username")
	password,_ := utils.Config.GetValue(mysqlname,"password")
	host,_ := utils.Config.GetValue(mysqlname,"host")
	db,_ := utils.Config.GetValue(mysqlname,"db")
	port,_ := utils.Config.GetValue(mysqlname,"port")
	DataDir,_ = utils.Config.GetValue(mysqlname,"datadir")
	if string(DataDir[len(DataDir)-1]) != "/" {
		DataDir += "/"
	}
	utils.Logger.Info("start mysql connect", zap.String("mysql address ",  host))
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True",username,password,host,port,db)
	Db,err = sql.Open("mysql",dataSourceName)
	if err != nil{
		utils.Logger.Error("start mysql connect fail", zap.String("mysql address ",  host))
		panic(err.Error())

	}
	err := Db.Ping()
	if err != nil{
		panic(err.Error())
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	Db.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	Db.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	Db.SetConnMaxLifetime(time.Hour)
	//defer Db.Close()
	Db.Stats()
}
