package AdminControllers

import "sqlx/models"

type GoodsTypeController struct{
	BaseController
}
//首页
func(this *GoodsTypeController)GetList(){
	var goodsType=new(models.GoodsType)
	goodsTypes,_:=goodsType.GetList()
	this.Data["goodsTypes"]=goodsTypes
	this.TplName="back/goodstype/index.html"
}
func(this *GoodsTypeController)GetAdd(){
	category:=models.Category{}
	categorys:=category.GetTree()
	this.Data["category_list"]=categorys
	this.TplName="back/goodstype/add.html"
}
func(this *GoodsTypeController)AddType(){
	type_name:=this.GetString("type_name")
	pid,_:=this.GetInt64("pid")
	category:=&models.Category{
		CategoryId:pid,	}
	goodsType:=models.GoodsType{}
	_,err:=goodsType.Add(type_name,category)
	if err !=nil{
		this.JsonResult(0,"添加失败")
	}else{
		this.JsonResult(200,"添加成功")
	}
}