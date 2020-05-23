package controllers

import (
	"github.com/astaxie/beego"
	"mysqlgo/models"
)

type JSONS struct {
	//必须的大写开头
	Code string
	Msg  string
	Data []models.Meun `json:"meun"`//key重命名,最外面是反引号
}
type MeunController struct {
	beego.Controller
}
func (c *MeunController) Get() {
	var meun =models.GetMeunList(0)
	data := &JSONS{"200", "获取成功",  meun}
	c.Data["json"] = data
	c.ServeJSON()
	//c.Ctx.WriteString("mysqldo=====test")
}
