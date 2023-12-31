package models

import (
	"encoding/json"
	"ffyouku/services/mq"
	"github.com/astaxie/beego/orm"
	"time"
)

type Comment struct {
	Id          int
	Content     string
	AddTime     int64
	UserId      int
	Status      int
	Stamp       int
	PraiseCount int
	EpisodesId  int
	VideoId     int
}

func GetCommentList(episodesId, offset, limit int) (int64, []Comment, error) {
	o := orm.NewOrm()
	var comments []Comment
	num, err := o.QueryTable("comment").Filter("episodes_id", episodesId).Limit(limit, offset).All(&comments)
	return num, comments, err
}

func SaveContent(content string, uid, episodesId, videoId int) error {
	o := orm.NewOrm()
	comment := Comment{
		Content:    content,
		AddTime:    time.Now().Unix(),
		UserId:     uid,
		Stamp:      0,
		Status:     1,
		EpisodesId: episodesId,
		VideoId:    videoId,
	}
	_, err := o.Insert(comment)
	if err != nil {
		return err
	}
	//修改视频的总论数
	o.Raw("update video set comment=comment+1 where id=?", videoId).Exec()
	//修改剧集的总论数
	o.Raw("update video_episodes set comment=comment+1 where id=?", episodesId).Exec()

	//更新redis排行榜
	videoObj := map[string]int{
		"VideoId": videoId,
	}
	videoJson, _ := json.Marshal(videoObj)
	mq.Publish("", "fyouku_top", string(videoJson))

	//延迟增加评论数
	videoCountObj := map[string]int{
		"VideoId":    videoId,
		"EpisodesId": episodesId,
	}
	videCountJson, _ := json.Marshal(videoCountObj)
	mq.PublishDlx("fyouku.comment.count", string(videCountJson))

	return nil
}
