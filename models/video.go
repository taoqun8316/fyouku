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
