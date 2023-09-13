package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       int
	Name     string
	Password string
	Status   int
	AddTime  int64
	Mobile   string
	Avatar   string
}

func IsUserMobile(mobile string) bool {
	o := orm.NewOrm()
	user := User{Mobile: mobile}
	err := o.Read(&user)

	if err == orm.ErrNoRows {
		return false
	} else if err == orm.ErrMissPK {
		return false
	}
	return true
}

func UserSave(mobile, password string) error {
	o := orm.NewOrm()
	var user User

	user.Name = ""
	user.Password = password
	user.Mobile = mobile
	user.Status = 1
	user.AddTime = time.Now().Unix()

	_, err := o.Insert(&user)
	return err
}

func IsMobileLogin(mobile, password string) (int, string) {
	o := orm.NewOrm()
	var user User

	err := o.QueryTable("user").Filter("mobile", mobile).Filter("password", password).One(&user)
	if err == orm.ErrNoRows {
		return 0, ""
	} else if err == orm.ErrMissPK {
		return 0, ""
	}
	return user.Id, user.Name
}
