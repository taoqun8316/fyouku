package controllers

import (
	"ffyouku/models"
	"github.com/astaxie/beego"
)

type TopController struct {
	beego.Controller
}

func (this *TopController) ChannelTop() {
	channelId, _ := this.GetInt("channelId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	num, videos, err := models.RedisGetChannelTop(channelId)
	if err != nil {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
	this.Data["json"] = ReturnSuccess(0, "请求数据成功", videos, num)
	this.ServeJSON()
}

func (this *TopController) TypeTop() {
	typeId, _ := this.GetInt("typeId")

	if typeId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定内型")
		this.ServeJSON()
	}
	num, videos, err := models.RedisGetTypeTop(typeId)
	if err != nil {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
	this.Data["json"] = ReturnSuccess(0, "请求数据成功", videos, num)
	this.ServeJSON()
}
