package controllers

import (
	"ffyouku/models"
	beego "github.com/beego/beego/v2/server/web"
	"regexp"
)

type UserController struct {
	CommonController
	beego.Controller
}

func (this *UserController) SaveRegister() {
	var (
		mobile   string
		password string
		err      error
	)
	mobile = this.GetString("mobile")
	password = this.GetString("password")

	if mobile == "" {
		this.Data["json"] = this.ReturnError(4001, "手机号不能为空")
		this.ServeJSON()
	}

	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		this.Data["json"] = this.ReturnError(4002, "手机号格式不正确")
		this.ServeJSON()
	}

	if password == "" {
		this.Data["json"] = this.ReturnError(4001, "密码不能为空")
		this.ServeJSON()
	}

	status := models.IsUserMobile(mobile)
	if status {
		this.Data["json"] = this.ReturnError(4002, "手机号已经被注册")
		this.ServeJSON()
	} else {
		err = models.UserSave(mobile, this.MD5V(password))
		if err != nil {
			this.Data["json"] = this.ReturnError(5000, err)
			this.ServeJSON()
		} else {
			this.Data["json"] = this.ReturnSuccess(0, "手机号格式不正确", nil, 0)
			this.ServeJSON()
		}
	}
}
