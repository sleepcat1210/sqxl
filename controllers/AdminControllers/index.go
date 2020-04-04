package AdminControllers

type IndexController struct {
	BaseController
}
func(c *IndexController)Get(){
	c.TplName = "back/admin/index.html"
}