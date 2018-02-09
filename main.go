package main

import (
	_ "beeblog/models"
	_ "beeblog/routers"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {

	// 创建日志文件和附件目录
	err := os.Mkdir("logs", os.ModePerm)
	err = os.Mkdir("attachment", os.ModePerm)
	if err != nil {
		beego.Error(err)
	}

	logs.EnableFuncCallDepth(true)
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogger("file", `{"filename":"logs/beeblog.log"}`)

	beego.Run()

}
