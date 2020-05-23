package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"mysqlgo/models"
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
	fmt.Println("--------------------")
	fmt.Println(user)
	fmt.Println("--------------------")
	 data := &JSONS1{"200", "获取成功",  user}
	c.Data["json"] = data
	c.ServeJSON()
	//c.Ctx.WriteString("mysqldo=====test")
}
