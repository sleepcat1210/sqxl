package AdminControllers

import "sqlx/models"

type AttributeController struct{
	BaseController
}
func(this *AttributeController)GetList(){
	typeId,_:=this.GetInt64("typeid")
	this.Data["typeid"]=typeId
	attribute:=models.Attribute{}
	attributes,_:=attribute.GetList(typeId)
	this.Data["attributes"]=attributes
	this.TplName="back/attr/index.html"
}

//添加
func(this *AttributeController)Add(){
	typeId,_:=this.GetInt64("typeid")
	this.Data["typeid"]=typeId
	this.TplName="back/attr/add.html"
}
func(this *AttributeController)AddAttr(){
	typeid,_:=this.GetInt64("typeid")
	GoodsType:=&models.GoodsType{
		TypeId:typeid,
	}
	attr_name:=this.GetString("attr_name")
	attr_type,_:=this.GetInt("attr_type")
	attr_input_type,_:=this.GetInt("attr_input_type")
	attr_values:=this.GetString("attr_values")
	 attr :=models.Attribute{}
	 _,err:=attr.Add(attr_name,attr_type,attr_input_type,attr_values,GoodsType)
	 if err !=nil{
	 	this.JsonResult(0,"添加失败")
	 }else{
	 	this.JsonResult(200,"添加成功")
	 }
}
//获取属性列表
