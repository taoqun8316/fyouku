package controllers

import (
	"ffyouku/models"
	"ffyouku/utils"
	"github.com/astaxie/beego"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type UserController struct {
	beego.Controller
}

// SaveRegister 用户注册功能
func (this *UserController) SaveRegister() {
	var (
		mobile   string
		password string
		err      error
	)
	mobile = this.GetString("mobile")
	password = this.GetString("password")

	if mobile == "" {
		this.Data["json"] = ReturnError(4001, "手机号不能为空")
		this.ServeJSON()
	}

	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		this.Data["json"] = ReturnError(4002, "手机号格式不正确")
		this.ServeJSON()
	}

	if password == "" {
		this.Data["json"] = ReturnError(4001, "密码不能为空")
		this.ServeJSON()
	}

	status := models.IsUserMobile(mobile)
	if status {
		this.Data["json"] = ReturnError(4002, "手机号已经被注册")
		this.ServeJSON()
	} else {
		err = models.UserSave(mobile, MD5V(password))
		if err != nil {
			this.Data["json"] = ReturnError(5000, err)
			this.ServeJSON()
		} else {
			this.Data["json"] = ReturnSuccess(0, "注册成功", nil, 0)
			this.ServeJSON()
		}
	}
}

// LoginDo 用户登录
func (this *UserController) LoginDo() {
	mobile := this.GetString("mobile")
	password := this.GetString("password")

	if mobile == "" {
		this.Data["json"] = ReturnError(4001, "手机号不能为空")
		this.ServeJSON()
	}

	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		this.Data["json"] = ReturnError(4002, "手机号格式不正确")
		this.ServeJSON()
	}

	if password == "" {
		this.Data["json"] = ReturnError(4001, "密码不能为空")
		this.ServeJSON()
	}

	uid, name := models.IsMobileLogin(mobile, MD5V(password))
	if uid != 0 {
		this.Data["json"] = ReturnSuccess(0, "登录成功", map[string]interface{}{"uid": uid, "username": name}, 1)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "手机号或密码不正确")
		this.ServeJSON()
	}
}

// SendMessageDo 批量发送消息接口
func (this *UserController) SendMessageDo() {
	uids := this.GetString("uids")
	content := this.GetString("content")

	if uids == "" {
		this.Data["json"] = ReturnError(4001, "请填写接收人")
		this.ServeJSON()
	}
	if content == "" {
		this.Data["json"] = ReturnError(4002, "请填写发送内容")
		this.ServeJSON()
	}

	mid, err := models.SendMessageDo(content)
	if err != nil {
		uidConfig := strings.Split(uids, ",")
		for _, v := range uidConfig {
			userId, _ := strconv.Atoi(v)
			//models.SendMessageUser(userId, mid)
			models.SendMessageUserMq(userId, mid)
		}
		this.Data["json"] = ReturnSuccess(0, "发送消息成功", "", 0)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "发送消息失败")
		this.ServeJSON()
	}
}

// 上传视频文件
func (this *UserController) UploadVideo() {
	var (
		err   error
		title string
	)
	r := this.Ctx.Request
	uid := r.FormValue("uid")
	file, header, _ := r.FormFile("file") //获取文件流
	b, _ := ioutil.ReadAll(file)          //转换文件流为二进制

	//生成文件名
	filename := strings.Split(header.Filename, ".")
	filename[0] = utils.GetVideoName(uid)
	//文件保存的位置
	var fileDir = "/Users/taoqun/code_project/go_project/fyouku/static/uploads" + filename[0] + "." + filename[1]
	//播放地址
	var playUrl = "/static/uploads/" + filename[0] + "." + filename[1]
	err = ioutil.WriteFile(fileDir, b, 0777)
	if err != nil {
		title = utils.ReturnError(500, "上传失败")
	} else {
		title = utils.ReturnSuccess(0, playUrl, nil, 1)
	}
	this.Ctx.WriteString(title)
}
