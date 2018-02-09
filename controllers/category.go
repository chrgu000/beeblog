package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {

	op := this.Input().Get("op")
	// 添加分类
	if checkCookies(this.Ctx) {
		if op == "add" {
			name := this.Input().Get("name")
			if name != "" {
				err := models.AddCategory(name)
				if err != nil {
					beego.Error(err)
				}
			}
		}

		if op == "del" {
			id := this.Input().Get("id")
			if id != "" {
				err := models.DelCategory(id)
				if err != nil {
					beego.Error(err)
				}
			}
		}
		this.Data["IsLogin"] = true
	}

	Categories, err := models.GetAllCatagroy()
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Categories"] = Categories
	}
	// 不管是删除还是 添加最终都要get所有的分类并显示
	this.Data["IsCategory"] = true
	this.TplName = "category.html"
}
