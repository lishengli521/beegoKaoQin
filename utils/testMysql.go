package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
var Db *sqlx.DB

func init() {
	fmt.Println("InitMysql....")
	//注册数据库驱动
	driverName := beego.AppConfig.String("driverName")
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/cmsproject?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	database, err := sqlx.Open(driverName, dbConn)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
}
func GetDb() *sqlx.DB {
	return Db
}
