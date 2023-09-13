package controllers

import (
	"crypto/md5"
	"encoding/hex"
	beego "github.com/beego/beego/v2/server/web"
)

type CommonController struct {
	beego.Controller
}

type JsonStruct struct {
	Code  int         `json:code`
	Msg   interface{} `json:msg`
	Items interface{} `json:items`
	Count int64       `json:count`
}

func (*CommonController) ReturnSuccess(code int, msg interface{}, item interface{}, count int64) (json *JsonStruct) {
	json = &JsonStruct{
		Code:  code,
		Msg:   msg,
		Items: item,
		Count: count,
	}
	return
}

func (*CommonController) ReturnError(code int, msg interface{}) (json *JsonStruct) {
	json = &JsonStruct{
		Code: code,
		Msg:  msg,
	}
	return
}

func (*CommonController) MD5V(password string) string {
	h := md5.New()
	md5code, _ := beego.AppConfig.String("md5code")
	h.Write([]byte(password + md5code))
	return hex.EncodeToString(h.Sum(nil))
}
