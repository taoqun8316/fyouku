package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Barrage struct {
	Id          int
	Content     string
	AddTime     int64
	UserId      int
	Status      int
	CurrentTime int
	EpisodesId  int
	VideoId     int
}

type BarrageData struct {
	Id          int    `json:"id"`
	Content     string `json:"content"`
	CurrentTime int    `json:"currentTime"`
}

func BarrageList(episodesId, startTime, endTime int) (int64, []BarrageData, error) {
	o := orm.NewOrm()
	var barrages []BarrageData
	num, err := o.Raw("select id,content,`current_time` from barrage where status=1 and episodes_id=? and `current_time`>=? and `current_time`<?", episodesId, startTime, endTime).QueryRows(&barrages)
	return num, barrages, err
}

func SaveBarrage(content string, uid, episodesId, videoId, currentTime int) error {
	o := orm.NewOrm()
	barrage := Barrage{
		Content:     content,
		AddTime:     time.Now().Unix(),
		UserId:      uid,
		CurrentTime: currentTime,
		Status:      1,
		EpisodesId:  episodesId,
		VideoId:     videoId,
	}
	_, err := o.Insert(barrage)
	if err != nil {
		return err
	}
	return nil
}
