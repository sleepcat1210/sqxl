package main

import (
	_ "sqlx/routers"
	"github.com/astaxie/beego"
	_ "sqlx/models"
	_ "sqlx/sysinit"
	"github.com/astaxie/beego/toolbox"
	"sqlx/utils/filecache"
)

func main() {
	task:=toolbox.NewTask("delcache","*/2 * * * * *",func()error{
		filecache.ClearExpiredFiles()
		return nil
	})
	toolbox.AddTask("tsk",task)
	toolbox.StartTask()
	defer toolbox.StopTask()
	beego.Run()
}

