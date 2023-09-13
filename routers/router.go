package routers

import (
	"ffyouku/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.UserController{})

}
