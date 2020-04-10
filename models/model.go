package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/sqlx?charset=utf8&loc=Asia%2FShanghai")


	orm.RegisterModelWithPrefix("sqxl_",new(Brand),new(Category),new(GoodsType),new(Attribute),new(Goods),new(GoodsAttr))
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
}