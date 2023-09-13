package models

import (
	"github.com/astaxie/beego/orm"
)

type Region struct {
	Id        int
	Name      string
	AddTime   int64
	Status    int
	ChannelId int
	Sort      int
}

type Type struct {
	Id        int
	Name      string
	AddTime   int64
	Status    int
	ChannelId int
	Sort      int
}

func GetChannelRegion(channelId int) (int64, []Region, error) {
	o := orm.NewOrm()
	var regions []Region

	num, err := o.QueryTable("channel_region").Filter("status", 1).Filter("channel_id", channelId).OrderBy("sort").Limit(9).All(&regions)
	return num, regions, err
}

func GetChannelType(channelId int) (int64, []Type, error) {
	o := orm.NewOrm()
	var types []Type

	num, err := o.QueryTable("channel_type").Filter("status", 1).Filter("channel_id", channelId).OrderBy("sort").Limit(9).All(&types)
	return num, types, err
}
