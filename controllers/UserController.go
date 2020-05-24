package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"mysqlgo/models"
	"mysqlgo/utils"
)

type UserController struct {
	beego.Controller
}
func (c *UserController) Get() {
	var user =models.GetAll()
	test := utils.RDSTest
	test.SetEx("users",user,100)
	var users =test.Get("users")
	//转成字符串
	//fmt.Println(string(users))
	var aType []models.User
	 //json转实体对象
	json.Unmarshal(users,&aType)
	c.Data["json"] =utils.Succee(aType)
	c.ServeJSON()
}
