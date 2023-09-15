package routers

import (
	"ffyouku/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register/save", &controllers.UserController{}, "post:SaveRegister") //注册
	beego.Router("/login/do", &controllers.UserController{}, "*:LoginDo")              //登陆

	beego.Router("/channel/advert", &controllers.VideoController{}, "post:ChannelAdvert")                       //顶部广告功能
	beego.Router("/channel/hot", &controllers.VideoController{}, "get:ChannelHotList")                          //正在热播功能
	beego.Router("/channel/recommend/region", &controllers.VideoController{}, "get:ChannelRecommendRegionList") //根据频道地区获取推荐视频
	beego.Router("/channel/recommend/type", &controllers.VideoController{}, "get:ChannelRecommendTypeList")     //根据频道类型获取视频推荐
	beego.Router("/channel/video", &controllers.VideoController{}, "get:ChannelVideos")                         //根据传入参数获取视频列表
	beego.Router("/video/info", &controllers.VideoController{}, "get:VideoInfo")                                //获取视频详情
	beego.Router("/video/episodes/list", &controllers.VideoController{}, "get:VideoEpisodesList")               //获取视频剧集列表
	beego.Router("/user/video", &controllers.VideoController{}, "get:UserVideo")                                //我的视频管理
	beego.Router("/video/saves", &controllers.VideoController{}, "post:VideoSave")                              //保存用户上传视频信息

	beego.Router("/channel/region", &controllers.BaseController{}, "get:ChannelRegion") //获取频道下地区
	beego.Router("/channel/type", &controllers.BaseController{}, "get:ChannelType")     //获取频道下类型

	beego.Router("/comment/list", &controllers.CommentController{}, "get:List") //获取评论列表
	beego.Router("/comment/save", &controllers.CommentController{}, "get:Save") //发表评论功能

	beego.Router("/channel/top", &controllers.TopController{}, "get:ChannelTop") //频道排行榜接口
	beego.Router("/type/top", &controllers.TopController{}, "get:TypeTop")       //类型排行榜接口

	beego.Router("/barrage/ws", &controllers.BarrageController{}, "get:BarrageWs") //获取弹幕信息接口
	beego.Router("/send/message", &controllers.BarrageController{}, "post:Save")   //批量发送消息接口

}
