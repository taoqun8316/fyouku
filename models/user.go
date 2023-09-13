package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@/fyouku?charset=utf8", 30)

	orm.RegisterModel(new(User))

	orm.RunSyncdb("default", false, true)
}

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
	err := o.Read(&user, "Mobile")

	if err != orm.ErrNoRows {
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
