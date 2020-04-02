package main

import (
	_ "sqlx/routers"
	"github.com/astaxie/beego"
	_ "sqlx/models"
)

func main() {
	beego.Run()
}

