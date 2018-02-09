package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

func AddComment(tid, nickname, content string) error {
	o := orm.NewOrm()

	comment := Comment{Tid: tid,
		NickName: nickname,
		Content:  content,
		Created:  time.Now()}
	_, err := o.Insert(&comment)
	if err != nil {
		return err
	}
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	// 更新这个文章最后一次回复的时间等
	topic := &Topic{Id: tidNum}
	//rows := o.QueryTable("topic")
	if err = o.Read(topic); err == nil {
		topic.ReplyCount++
		topic.ReplyTime = time.Now()
		_, err = o.Update(topic)
	}
	return err
}

func GetAllComment(tid string) ([]*Comment, error) {
	o := orm.NewOrm()
	comments := make([]*Comment, 0)

	rows := o.QueryTable("comment")
	_, err := rows.Filter("tid", tid).All(&comments)
	return comments, err
}

func DeleteComment(cid string) error {
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	var tid string
	comment := &Comment{Id: cidNum}
	if err = o.Read(comment); err == nil {
		tid = comment.Tid
		_, err = o.Delete(comment)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	//  更新文章的相关的回复数和回复时间
	comments := make([]*Comment, 0)

	rows := o.QueryTable("comment")
	_, err = rows.Filter("tid", tid).OrderBy("-created").All(&comments)
	if err != nil {
		return err
	}
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	topic := &Topic{Id: tidNum}

	if err = o.Read(topic); err == nil {
		commentlen := int64(len(comments))
		topic.ReplyCount = commentlen
		if commentlen > 0 {
			topic.ReplyTime = comments[0].Created
		} else {
			topic.ReplyTime = time.Now()
		}
		_, err = o.Update(topic)
	}
	return err

}
