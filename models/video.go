package models

import (
	"encoding/json"
	redisClient "ffyouku/services/redis"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
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
	EpisodesUpdate int64
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

// 增加redis缓存-获取视频详情
func RedisGetVideoInfo(videoId int) (Video, error) {
	var video Video
	conn := redisClient.PoolConnect()
	defer conn.Close()
	//定义redis key
	redisKey := "video:id:" + strconv.Itoa(videoId)
	//判断redis中是否存在
	exists, err := redis.Bool(conn.Do("exists", redisKey))

	if exists {
		values, _ := redis.Values(conn.Do("hgetall", redisKey))
		err = redis.ScanStruct(values, &video)
	} else {
		o := orm.NewOrm()
		video.Id = videoId
		err = o.Read(&video)
		if err == nil {
			//保存redis
			_, err = conn.Do("hmset", redis.Args{redisKey}.AddFlat(video))
			if err == nil {
				conn.Do("expire", redisKey, 86400)
			}
		}
	}
	return video, err
}

func RedisGetVideoEpisodesList(videoId int) (int64, []Episodes, error) {
	var (
		episodes []Episodes
		num      int64
		err      error
	)
	conn := redisClient.PoolConnect()
	redisKey := "video:episodes:videoId:" + strconv.Itoa(videoId)
	defer conn.Close()
	exists, err := redis.Bool(conn.Do("exists", redisKey))
	if exists {
		num, err = redis.Int64(conn.Do("llen", redisKey))
		if err == nil {
			values, _ := redis.Values(conn.Do("lrange", redisKey, "0", "-1"))
			var episodesInfo Episodes
			for _, v := range values {
				json.Unmarshal(v.([]byte), &episodesInfo)
				episodes = append(episodes, episodesInfo)
			}
		}
	} else {
		o := orm.NewOrm()
		num, err = o.QueryTable("video_episodes").Filter("status", 1).Filter("video_id", videoId).OrderBy("-num").All(&episodes)
		if err == nil {
			for _, v := range episodes {
				jsonValue, _ := json.Marshal(v)
				conn.Do("rpush", redisKey, jsonValue)
			}
			conn.Do("expire", redisKey, 86400)
			return num, episodes, nil
		}
	}
	return num, episodes, err
}

func GetChannelTop(channelId int) (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video
	num, err := o.QueryTable("video").Filter("status", 1).Filter("channel_id", channelId).OrderBy("-comment").Limit(10).All(&videos)
	return num, videos, err
}

func GetTypeTop(typeId int) (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video
	num, err := o.QueryTable("video").Filter("status", 1).Filter("type_id", typeId).OrderBy("-comment").Limit(10).All(&videos)
	return num, videos, err
}

func GetUserVideo(uid int) (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video
	num, err := o.QueryTable("video").Filter("user_id", uid).OrderBy("-add_time").Limit(10).All(&videos)
	return num, videos, err
}

func SaveVideo(title string, subTitle string, channelId int, regionId int, typeId int, playUrl string, uid int) error {
	o := orm.NewOrm()
	var video Video
	current := time.Now().Unix()

	video.Title = title
	video.SubTitle = subTitle
	video.ChannelId = channelId
	video.RegionId = regionId
	video.TypeId = typeId
	video.Img = ""
	video.Img1 = ""
	video.EpisodesCount = 1
	video.IsEnd = 1
	video.Status = 1
	video.UserId = uid
	video.Comment = 0
	video.EpisodesUpdate = current
	video.AddTime = current
	videoId, err := o.Insert(video)

	if err == nil {
		//修改视频的总论数
		o.Raw("insert into video_episodes (title,add_time,num,video_id,play_url,status,comment) values(?,?,?,?,?,?,?,?)",
			subTitle, current, 1, videoId, playUrl, 1, 0).Exec()
	}
	return err
}
