package models

import (
	"github.com/astaxie/beego/orm"
)

type Video struct {
	Id             int
	Title          string
	SubTitle       string
	AddTime        int64
	Img            string
	Img1           string
	ChannelId      int
	Status         int
	TypeId         int
	RegionId       int
	UserId         int
	EpisodesCount  int
	EpisodesUpdate int
	IsEnd          int
	IsHot          int
	IsRecommend    int
	Comment        int
}

type Episodes struct {
	Id      int
	Title   string
	AddTime int64
	Num     int
	VideoId int
	PlayUrl string
	Status  int
	Comment int
}

func GetChannelHotList(channelId int) (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video

	num, err := o.QueryTable("video").Filter("status", 1).Filter("is_hot", 1).Filter("channel_id", channelId).OrderBy("-episodes_update").Limit(9).All(&videos)
	return num, videos, err
}

func GetChannelRecommendRegionList(channelId, regionId int) (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video

	num, err := o.QueryTable("video").Filter("status", 1).Filter("is_hot", 1).Filter("channel_id", channelId).Filter("region_id", regionId).Filter("is_recommend", 1).OrderBy("-episodes_update").Limit(9).All(&videos)
	return num, videos, err
}

func GetChannelRecommendTypeList(channelId, typeId int) (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video

	num, err := o.QueryTable("video").Filter("status", 1).Filter("is_hot", 1).Filter("channel_id", channelId).Filter("type_id", typeId).Filter("is_recommend", 1).OrderBy("-episodes_update").Limit(9).All(&videos)
	return num, videos, err
}

func GetChannelVideoList(channelId int, typeId int, regionId int, end string, sort string, offset int, limit int) (int64, []orm.Params, error) {
	o := orm.NewOrm()
	var videos []orm.Params

	qs := o.QueryTable("video")
	qs = qs.Filter("channel_id", channelId)
	qs = qs.Filter("status", 1)

	if regionId > 0 {
		qs = qs.Filter("region_id", regionId)
	}

	if end == "n" {
		qs = qs.Filter("is_end", 0)
	} else if end == "y" {
		qs = qs.Filter("is_end", 1)
	}

	if sort == "episodes_update" {
		qs = qs.OrderBy("-episodes_update")
	} else if sort == "comment" {
		qs = qs.OrderBy("-comment")
	} else if sort == "addTime" {
		qs = qs.OrderBy("-add_time")
	} else {
		qs = qs.OrderBy("-add_time")
	}

	num, _ := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	qs = qs.Limit(limit, offset)
	_, err := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")

	return num, videos, err
}

func GetVideoInfo(videoId int) (Video, error) {
	o := orm.NewOrm()
	video := Video{Id: videoId}
	err := o.Read(&video)
	if err != nil {
		return Video{}, err
	}
	return video, nil
}

func GetVideoEpisodesList(videoId int) (int64, []Episodes, error) {
	o := orm.NewOrm()
	var episodes []Episodes
	num, err := o.QueryTable("video_episodes").Filter("status", 1).Filter("video_id", videoId).OrderBy("-num").All(&episodes)
	return num, episodes, err
}
