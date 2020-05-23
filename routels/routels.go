package routels

import (
	"github.com/astaxie/beego"
	"mysqlgo/controllers"
)

func init()  {
	beego.Router("/user",&controllers.UserController{})
	beego.Router("/meun",&controllers.MeunController{})

}