package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	// 判断当前是退出还是登录操作
	isexit := this.Input().Get("exit") == "true"
	if isexit {
		maxAge := -1
		this.Ctx.SetCookie("username", "", maxAge, "/")
		this.Ctx.SetCookie("password", "", maxAge, "/")
		this.Ctx.Redirect(301, "/")
		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	uname := this.Input().Get("username")
	pwd := this.Input().Get("password")
	autologin := this.Input().Get("autologin") == "on"
	// 验证通过 设置cookie
	if beego.AppConfig.String("username") == uname &&
		beego.AppConfig.String("password") == pwd {
		// 浏览器关闭 cookie 失效
		maxAge := 0
		if autologin {
			// 长时间的有效
			maxAge = 1<<31 - 1
		}
		// 设置cookie
		this.Ctx.SetCookie("username", uname, maxAge, "/")
		this.Ctx.SetCookie("password", pwd, maxAge, "/")
	}
	// 重镜像到首页
	this.Ctx.Redirect(301, "/")
	return
}
