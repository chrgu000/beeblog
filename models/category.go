package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// 分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{
		Title:     name,
		Created:   time.Now(),
		TopicTime: time.Now()}

	row := o.QueryTable("category")
	err := row.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	_, err = o.Insert(cate)
	return err
}

func GetAllCatagroy() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	row := o.QueryTable("category")
	_, err := row.All(&cates)
	return cates, err
}

func DelCategory(idstr string) error {
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{
		Id: id,
	}
	_, err = o.Delete(cate)
	return err

}
