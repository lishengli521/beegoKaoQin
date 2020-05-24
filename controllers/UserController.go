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
	values,_:=json.Marshal(user)
	test.Set("users",string(values))
	var users,_=test.Get("users")
	var aType []models.User
	json.Unmarshal(users.([]byte),&aType)
	fmt.Println("---------redis-存值----------")
	fmt.Println(aType)
	fmt.Println("---------redis取值-----------")
	data := &JSONS1{"200", "获取成功",  aType}
	c.Data["json"] = data
	c.ServeJSON()
}
