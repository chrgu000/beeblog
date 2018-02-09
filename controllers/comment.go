package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) Add() {
	tid := this.Input().Get("tid")
	content := this.Input().Get("content")
	nickname := this.Input().Get("nickname")
	err := models.AddComment(tid, nickname, content)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)
}

func (this *CommentController) Delete() {
	if !checkCookies(this.Ctx) {
		return
	}
	cid := this.Input().Get("cid")
	tid := this.Input().Get("tid")
	err := models.DeleteComment(cid)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)
}
