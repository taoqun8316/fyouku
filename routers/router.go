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
	beego.Router("/channel/recommend/region", &controllers.VideoController{}, "get:ChannelRecommendRegionList") //日漫国漫推荐功能
	beego.Router("/channel/recommend/type", &controllers.VideoController{}, "get:ChannelRecommendTypeList")     //少女推荐功能

}
