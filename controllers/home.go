package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	beego.Info("获取首页")
	this.Data["IsLogin"] = checkCookies(this.Ctx)
	this.Data["IsHome"] = true
	// 获取某个分类的所有文章
	topics, err := models.GetAllTopicOrderByTime(this.Input().Get("cate"), this.Input().Get("lable"))
	if err != nil {
		beego.Error(err.Error())
	}

	categories, err := models.GetAllCatagroy()
	if err != nil {
		beego.Error(err.Error())
	}

	this.Data["Categories"] = categories
	this.Data["Topics"] = topics
	this.TplName = "home.html"
}
