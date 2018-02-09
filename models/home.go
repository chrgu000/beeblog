package models

import (
	"github.com/astaxie/beego/orm"
)

func GetAllTopicOrderByTime(cate, lable string) ([]*Topic, error) {

	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	rows := o.QueryTable("topic")
	var err error
	// 获取某个分类下的所有文章
	if len(cate) > 0 {
		rows = rows.Filter("category", cate)

	}
	if len(lable) > 0 {
		rows = rows.Filter("lable__contains", "$"+lable+"#")
	}
	_, err = rows.OrderBy("-updated").All(&topics)

	return topics, err
}
