package main

import (
	_ "beeblog/models"
	_ "beeblog/routers"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.EnableFuncCallDepth(true)
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogger("file", `{"filename":"logs/beeblog.log"}`)

	// 启动 beego
	err := os.Mkdir("logs", os.ModePerm)
	err = os.Mkdir("attachment", os.ModePerm)
	if err != nil {
		beego.Error(err)
	}
	beego.Run()

}
