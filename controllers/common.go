package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"time"
)

type JsonStruct struct {
	Code  int         `json:code`
	Msg   interface{} `json:msg`
	Items interface{} `json:items`
	Count int64       `json:count`
}

func ReturnSuccess(code int, msg interface{}, item interface{}, count int64) (json *JsonStruct) {
	json = &JsonStruct{
		Code:  code,
		Msg:   msg,
		Items: item,
		Count: count,
	}
	return
}

func ReturnError(code int, msg interface{}) (json *JsonStruct) {
	json = &JsonStruct{
		Code: code,
		Msg:  msg,
	}
	return
}

func MD5V(password string) string {
	h := md5.New()
	h.Write([]byte(password + beego.AppConfig.String("md5code")))
	return hex.EncodeToString(h.Sum(nil))
}

func DateFormat(t int64) string {
	videoTime := time.Unix(t, 0)
	return videoTime.Format("1992-12-10")
}
