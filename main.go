package main

import (
	_ "sqlx/routers"
	"github.com/astaxie/beego"
	_ "sqlx/models"
	_ "sqlx/sysinit"
)

func main() {
	beego.Run()
}

