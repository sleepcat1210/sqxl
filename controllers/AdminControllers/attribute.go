package AdminControllers
type AttributeController struct{
	BaseController
}
func(this *AttributeController)GetList(){
	this.TplName="back/attr/index.html"
}
