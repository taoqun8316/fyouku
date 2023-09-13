package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	orm.RegisterDataBase("default", "mysql", "root:root@/fyouku?charset=utf8", 30)

	orm.RegisterModel(new(User), new(Advert), new(Video))

	orm.RunSyncdb("default", false, true)
}
