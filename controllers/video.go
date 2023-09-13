package controllers

import (
	"ffyouku/models"
	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

// ChannelAdvert 顶部广告功能
func (this *VideoController) ChannelAdvert() {
	channelId, _ := this.GetInt("channelId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}

	num, videos, err := models.GetChannelAdvert(channelId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "请求数据成功", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "请求数据失败")
		this.ServeJSON()
	}
}

// 频道页-获取正在热播
func (this *VideoController) ChannelHotList() {
	channelId, _ := this.GetInt("channelId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}

	num, videos, err := models.GetChannelHotList(channelId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

// ChannelRecommendRegionList 频道页-根据频道地区获取推荐视频
func (this *VideoController) ChannelRecommendRegionList() {
	channelId, _ := this.GetInt("channelId")
	regionId, _ := this.GetInt("regionId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}

	if regionId == 0 {
		this.Data["json"] = ReturnError(4002, "必须指定地区")
		this.ServeJSON()
	}

	num, videos, err := models.GetChannelRecommendRegionList(channelId, regionId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

// ChannelRecommendTypeList 频道页-根据频道类型获取视频推荐
func (this *VideoController) ChannelRecommendTypeList() {
	channelId, _ := this.GetInt("channelId")
	typeId, _ := this.GetInt("typeId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}

	if typeId == 0 {
		this.Data["json"] = ReturnError(4002, "必须指定类型")
		this.ServeJSON()
	}

	num, videos, err := models.GetChannelRecommendTypeList(channelId, typeId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}
