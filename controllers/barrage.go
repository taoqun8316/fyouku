package controllers

import (
	"encoding/json"
	"ffyouku/models"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
)

type BarrageController struct {
	beego.Controller
}

type WsData struct {
	CurrentTime int
	EpisodesId  int
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// BarrageWs 获取弹幕
func (this *BarrageController) BarrageWs() {
	var (
		conn     *websocket.Conn
		err      error
		data     []byte
		barrages []models.BarrageData
	)

	if conn, err = upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil); err != nil {
		goto ERR
	}
	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		var wsData WsData
		json.Unmarshal(data, &wsData)
		endTime := wsData.CurrentTime + 60
		//获取弹幕数据
		_, barrages, err = models.BarrageList(wsData.EpisodesId, wsData.CurrentTime, endTime)
		if err == nil {
			if err = conn.WriteJSON(barrages); err != nil {
				goto ERR
			}
		}
	}

ERR:
	conn.Close()
}

// Save 保存弹幕
func (this *BarrageController) Save() {
	content := this.GetString("content")
	uid, _ := this.GetInt("uid")
	episodesId, _ := this.GetInt("episodesId")
	videoId, _ := this.GetInt("videoId")
	currentTime, _ := this.GetInt("currentTime")

	if content == "" {
		this.Data["json"] = ReturnError(4001, "内容不能为空")
		this.ServeJSON()
	}
	if uid == 0 {
		this.Data["json"] = ReturnError(4002, "请先登录")
		this.ServeJSON()
	}
	if episodesId == 0 {
		this.Data["json"] = ReturnError(4003, "必须指定评论剧集ID")
		this.ServeJSON()
	}
	if videoId == 0 {
		this.Data["json"] = ReturnError(4004, "必须指定视频ID")
		this.ServeJSON()
	}
	if currentTime == 0 {
		this.Data["json"] = ReturnError(4004, "必须指定视频播放时间")
		this.ServeJSON()
	}

	err := models.SaveBarrage(content, uid, episodesId, videoId, currentTime)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "请求数据成功", "", 1)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(5000, err)
		this.ServeJSON()
	}
}
