package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func checkCookies(ctx *context.Context) bool {
	username := ctx.GetCookie("username")
	password := ctx.GetCookie("password")
	return beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password
}
