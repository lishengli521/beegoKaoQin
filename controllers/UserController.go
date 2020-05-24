package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"mysqlgo/models"
	"mysqlgo/utils"
)
type JSONS1 struct {
	//必须的大写开头
	Code string
	Msg  string
	Data []models.User `json:"user"`//key重命名,最外面是反引号
}
type UserController struct {
	beego.Controller
}
func (c *UserController) Get() {
	var user =models.GetAll()
	test := utils.RDSTest
	test.SetEx("users",user,10)
	var users =test.Get("users")
	//转成字符串
	fmt.Println(string(users))
	var aType []models.User
	 //json转实体对象
	json.Unmarshal(users,&aType)
	data := &JSONS1{"200", "获取成功",  aType}
	c.Data["json"] = data
	c.ServeJSON()
}
