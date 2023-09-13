package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	Id      int
	Content string
	AddTime int64
}

type MessageUser struct {
	Id        int
	UserId    int64
	MessageId int64
	AddTime   int64
	Status    int
}

func SendMessageDo(content string) (int64, error) {
	o := orm.NewOrm()
	message := Message{
		Content: content,
		AddTime: time.Now().Unix(),
	}
	insert, err := o.Insert(&message)
	if err != nil {
		return 0, err
	}
	return insert, nil
}

func SendMessageUser(userId int64, messageId int64) error {
	o := orm.NewOrm()
	messageUser := MessageUser{
		UserId:    userId,
		MessageId: messageId,
		Status:    1,
		AddTime:   time.Now().Unix(),
	}
	_, err := o.Insert(&messageUser)
	return err
}
