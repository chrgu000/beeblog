package models

import (
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

func AddTopic(title, category, lable, content, attachment string) error {
	o := orm.NewOrm()
	lable = "$" + strings.Join(strings.Split(lable, " "), "#$") + "#"
	topic := &Topic{Title: title,
		Content:    content,
		Category:   category,
		Lable:      lable,
		Attachment: attachment,
		Created:    time.Now(),
		Updated:    time.Now(),
		ReplyTime:  time.Now()}
	_, err := o.Insert(topic)
	if err != nil {
		return err
	}

	// 添加或者更改分类列表
	cate := new(Category)
	rows := o.QueryTable("category")
	err = rows.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	} else {
		// 这个分类不存在
		cate.Created = time.Now()
		cate.Title = category
		cate.TopicCount = 1
		cate.TopicTime = time.Now()
		_, err = o.Insert(cate)
	}
	return err
}

func GetAllTopic() ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)

	rows := o.QueryTable("topic")
	_, err := rows.All(&topics)
	return topics, err
}

func GetOneTopic(id string) (*Topic, error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topic := new(Topic)

	row := o.QueryTable("topic")
	err = row.Filter("id", tid).One(topic)
	if err != nil {
		return nil, err
	}

	// 更新查看次数

	topic.Views++
	_, err = o.Update(topic)

	// 获取标签
	topic.Lable = strings.Replace(strings.Replace(topic.Lable, "#", " ", -1), "$", "", -1)
	return topic, err
}

func ModifyTopic(id, title, category, lable, content, attachment string) error {
	var oldCate, oldattachment string
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tid}
	err = o.Read(topic)
	if err != nil {
		return err
	}
	oldattachment = topic.Attachment
	oldCate = topic.Category
	topic.Title = title
	topic.Attachment = attachment
	topic.Lable = "$" + strings.Join(strings.Split(lable, " "), "$#") + "#"
	topic.Content = content
	topic.Updated = time.Now()
	topic.Category = category
	_, err = o.Update(topic)

	//-----------------------文章附近更改-------------------------------------
	if oldattachment != attachment && len(attachment) > 0 {
		os.Remove(path.Join("attachment", attachment))
	}
	//-----------------------文章分类更新-------------------------------------
	if oldCate == category {
		return err
	}
	// 将改文章对应的前后两个分类进行更新
	// 减少旧文章的数量
	cate := new(Category)
	rows := o.QueryTable("category")
	err = rows.Filter("title", oldCate).One(cate)
	if err == nil {
		cate.TopicCount--
		_, err = o.Update(cate)
		if err != nil {
			return err
		}
	}
	// 添加或者更改分类列表
	cate = new(Category)
	err = rows.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	} else {
		// 这个分类不存在
		cate.Created = time.Now()
		cate.Title = category
		cate.TopicCount = 1
		cate.TopicTime = time.Now()
		_, err = o.Insert(cate)
	}
	return err
}

func DeleteOneTopic(id string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tid}
	rows := o.QueryTable("topic")
	rows.Filter("id", tid).One(topic)
	_, err = o.Delete(topic)
	if err != nil {
		return err
	}
	// 更新分类
	category := topic.Category
	if len(category) > 0 {
		rows = o.QueryTable("category")
		cate := new(Category)
		err = rows.Filter("title", category).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
			if err != nil {
				return err
			}
		}
	}
	// 删除对应的回复
	rows = o.QueryTable("comment")
	_, err = rows.Filter("tid", id).Delete()
	return err
}
