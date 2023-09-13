package controllers

import (
	"ffyouku/models"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// ChannelRegion 获取频道地区列表
func (this *BaseController) ChannelRegion() {
	channelId, _ := this.GetInt("channelId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}

	num, regions, err := models.GetChannelRegion(channelId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", regions, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

// ChannelType 获取频道类型列表
func (this *BaseController) ChannelType() {
	channelId, _ := this.GetInt("channelId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}

	num, regions, err := models.GetChannelType(channelId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", regions, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}
