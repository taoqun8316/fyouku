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

// ChannelHotList 频道页-获取正在热播
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

// ChannelVideos 根据传入参数获取视频列表
func (this *VideoController) ChannelVideos() {
	channelId, _ := this.GetInt("channelId")
	regionId, _ := this.GetInt("regionId")
	typeId, _ := this.GetInt("typeId")
	end := this.GetString("end")
	sort := this.GetString("sort")
	limit, _ := this.GetInt("limit")
	offset, _ := this.GetInt("offset")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}

	if regionId == 0 {
		this.Data["json"] = ReturnError(4002, "必须指定地区")
		this.ServeJSON()
	}

	if typeId == 0 {
		this.Data["json"] = ReturnError(4003, "必须指定类型")
		this.ServeJSON()
	}

	if limit == 0 {
		limit = 12
	}

	num, videos, err := models.GetChannelVideoList(channelId, typeId, regionId, end, sort, offset, limit)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

// VideoInfo 获取视频详情
func (this *VideoController) VideoInfo() {
	videoId, _ := this.GetInt("videoId")
	if videoId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频ID")
		this.ServeJSON()
	}

	video, err := models.RedisGetVideoInfo(videoId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", video, 1)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

// VideoEpisodesList 获取视频剧集列表
func (this *VideoController) VideoEpisodesList() {
	videoId, _ := this.GetInt("videoId")
	if videoId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频ID")
		this.ServeJSON()
	}

	num, episodes, err := models.RedisGetVideoEpisodesList(videoId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", episodes, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

// UserVideo 我的视频管理
func (this *VideoController) UserVideo() {
	uid, _ := this.GetInt("uid")
	if uid == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定用户")
		this.ServeJSON()
	}

	num, videos, err := models.GetUserVideo(uid)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

// VideoSave 保存用户上传视频信息
func (this *VideoController) VideoSave() {
	playUrl := this.GetString("playUrl")
	title := this.GetString("title")
	subTitle := this.GetString("subTitle")
	channelId, _ := this.GetInt("channelId")
	typeId, _ := this.GetInt("typeId")
	regionId, _ := this.GetInt("regionId")
	uid, _ := this.GetInt("uid")

	if uid == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定用户")
		this.ServeJSON()
	}

	if playUrl == "" {
		this.Data["json"] = ReturnError(4002, "视频地址不能为空")
		this.ServeJSON()
	}
	err := models.SaveVideo(title, subTitle, channelId, regionId, typeId, playUrl, uid)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", "", 0)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "保存失败")
		this.ServeJSON()
	}
}
