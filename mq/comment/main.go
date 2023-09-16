package main

import (
	"encoding/json"
	"ffyouku/services/mq"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	beego.LoadAppConfig("ini", "../../conf/app.conf")
	defaultdb := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)

	mq.ConsumerDlx("fyouku.comment.count", "fyouku_comment_count", "fyouku.comment.count.dlx", "fyouku_comment_count_dlx", 10000, callback)
}

func callback(s string) {
	type Data struct {
		VideoId    int
		EpisodesId int
	}

	var data Data
	err := json.Unmarshal([]byte(s), &data)
	if err == nil {
		o := orm.NewOrm()
		//修改视频的总论数
		o.Raw("update video set comment=comment+1 where id=?", data.VideoId).Exec()
		//修改剧集的总论数
		o.Raw("update video_episodes set comment=comment+1 where id=?", data.EpisodesId).Exec()

		//更新redis排行榜
		videoObj := map[string]int{
			"VideoId": data.VideoId,
		}
		videoJson, _ := json.Marshal(videoObj)
		mq.Publish("", "fyouku_top", string(videoJson))
	}

	fmt.Println("msg is: %s\n", s)
}
