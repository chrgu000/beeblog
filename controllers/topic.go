package controllers

import (
	"beeblog/models"
	"path"
	"strings"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	if checkCookies(this.Ctx) {
		this.Data["IsLogin"] = true
	}
	this.Data["IsTopic"] = true
	topics, err := models.GetAllTopic()

	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
	this.TplName = "topic.html"
}

// 添加和更新的实际操作
func (this *TopicController) Post() {
	if !checkCookies(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")
	lable := this.Input().Get("lable")

	var attachment string
	_, h, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}
	if h != nil {
		attachment = h.Filename
		beego.Info("上传附近名称", attachment)

		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}
	tid := this.Input().Get("id")
	if len(tid) > 0 {
		// 更新
		err = models.ModifyTopic(tid, title, category, lable, content, attachment)
	} else {
		// 添加
		err = models.AddTopic(title, category, lable, content, attachment)
	}
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

// add topic 界面的显示
func (this *TopicController) Add() {
	if !checkCookies(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	this.Data["IsLogin"] = true
	this.TplName = "topic_add.html"
}

// 查看
func (this *TopicController) View() {

	tid := this.Ctx.Input.Param("0")
	beego.Info("显示文章", tid)
	topic, err := models.GetOneTopic(tid)
	if err != nil {
		beego.Error("failed---------------------", err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	comments, err := models.GetAllComment(tid)
	if err != nil {
		beego.Error(err)
		// 失败不显示回复内容
	}
	this.Data["Lables"] = strings.Split(topic.Lable, " ")
	this.Data["Comments"] = comments
	this.Data["Tid"] = tid
	this.Data["IsLogin"] = checkCookies(this.Ctx)
	this.TplName = "topic_view.html"
}

// 显示修改界面
func (this *TopicController) Modify() {
	if !checkCookies(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	tid := this.Input().Get("tid")
	topic, err := models.GetOneTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Lables"] = strings.Split(topic.Lable, " ")
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
	this.TplName = "topic_modify.html"
}

// 删除
func (this *TopicController) Delete() {
	if !checkCookies(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	tid := this.Input().Get("tid")
	err := models.DeleteOneTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Redirect("/topic", 302)
}
