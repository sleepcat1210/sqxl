package AdminControllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}
func(c *IndexController)Get(){
	c.TplName = "back/admin/index.html"
}