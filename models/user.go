package models

import (
	redisClient "ffyouku/services/redis"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"strconv"
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

type UserInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
	AddTime int64  `json:"addTime"`
	Avatar  string `json:"avatar"`
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

func RedisGetUserInfo(uid int) (UserInfo, error) {
	var user UserInfo
	conn := redisClient.PoolConnect()
	redisKey := "user:id:" + strconv.Itoa(uid)
	defer conn.Close()
	exists, err := redis.Bool(conn.Do("exists", redisKey))
	if exists {
		res, _ := redis.Values(conn.Do("hgetall", redisKey))
		err = redis.ScanStruct(res, &user)
	} else {
		o := orm.NewOrm()
		user.Id = uid
		err = o.Read(&user)
		if err == nil {
			_, err = conn.Do("hmset", user)
			if err == nil {
				conn.Do("expire", redisKey, 86400)
			}
			return user, nil
		}
	}
	return user, err
}
