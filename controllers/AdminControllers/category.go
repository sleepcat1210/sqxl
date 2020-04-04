package AdminControllers

import (
	"sqlx/models"
	"github.com/astaxie/beego/orm"
	"strings"
)

type CategoryController struct {
	BaseController
}
//商品分类
func(c *CategoryController)GetList(){
	o:=orm.NewOrm()
	var category []models.Category
	o.QueryTable("SqxlCategory").All(&category)
	c.Data["category"]=category
	c.TplName="back/category/index.html"
}
//添加分类
func(c *CategoryController)AddShow(){
	category :=models.Category{}
    category_list:=category.GetTree()
    c.Data["category_list"]=category_list
	c.TplName="back/category/add.html"
}
//添加
func(c *CategoryController)AddCategory(){
 category_name :=c.GetString("category_name")
 category_status,_:=c.GetInt8("category_status")
 category_order,_ :=c.GetInt("category_order")
 pid,_:=c.GetInt64("pid")
 category_logo:=c.GetString("category_logo")
 o:=orm.NewOrm()
 category:=models.Category{
 	CategoryName:category_name,
 	CategoryOrder:category_order,
 	CategoryStatus:category_status,
 	Pid:pid,
 	CategoryLogo:category_logo,
 }
 result,err :=o.Insert(&category)
	resp := make(map[string]interface{})
	if err !=nil{
		resp["code"]="0"
		resp["msg"]="添加错误"
		resp["data"]=""
		c.Data["json"] = resp
		c.ServeJSON()
	}
	if result>0{
		resp["code"]="200"
		resp["msg"]="添加成功"
		resp["data"]=""
		c.Data["json"] = resp
		c.ServeJSON()
	}
}
//删除
func(c *CategoryController)Del(){
	id,_:=c.GetInt64("id")
	o:=orm.NewOrm()
	var category models.Category
	category.CategoryId=id
	_,err:=o.Delete(&category,"CategoryId")
	resp := make(map[string]interface{})
	if err !=nil{
		resp["code"]="0"
		resp["msg"]="删除失败"
		resp["data"]=""
		c.Data["json"] = resp
		c.ServeJSON()
	}else{
		resp["code"]="200"
		resp["msg"]="删除成功"
		resp["data"]=""
		c.Data["json"] = resp
		c.ServeJSON()
	}

}
//删除选择
func(c *CategoryController)DelAll(){
	ids :=c.GetString("ids")
	id:=strings.Split(ids,",")
	o:=orm.NewOrm()
	_,err:=o.QueryTable("SqxlCategory").Filter("CategoryId__in",id).Delete()
	resp := make(map[string]interface{})
	if err !=nil{
		resp["code"]="0"
		resp["msg"]="删除错误"
		resp["data"]=""
		c.Data["json"] = resp
		c.ServeJSON()
	}else{
		resp["code"]="200"
		resp["msg"]="删除成功"
		resp["data"]=""
		c.Data["json"] = resp
		c.ServeJSON()
	}
}
//编辑
func(c *CategoryController)Edit(){
	category_id,_:=c.GetInt64("id")
	o:=orm.NewOrm()
	category:= models.Category{}
	category.CategoryId=category_id
	o.Read(&category,"CategoryId")
	category_list:=category.GetTree()
	c.Data["category_list"]=category_list
	c.Data["category"]=category
	c.TplName="back/category/edit.html"
}
func(c *CategoryController)UpdateCategory(){

	category_name :=c.GetString("category_name")
	category_status,_:=c.GetInt8("category_status")
	category_order,_ :=c.GetInt("category_order")
	pid,_:=c.GetInt64("pid")
	category_logo:=c.GetString("category_logo")
	category_id,_:=c.GetInt64("category_id")
	o:=orm.NewOrm()
	category:=models.Category{
		CategoryName:category_name,
		CategoryOrder:category_order,
		CategoryStatus:category_status,
		Pid:pid,
		CategoryLogo:category_logo,
		CategoryId:category_id,
	}
	_,err :=o.Update(&category)
	resp := make(map[string]interface{})
	if err !=nil{
		resp["code"]="0"
		resp["msg"]="修改错误"
		resp["data"]=""
		c.Data["json"] = resp
		c.ServeJSON()
	}else{
		resp["code"]="200"
		resp["msg"]="修改成功"
		resp["data"]=""
		c.Data["json"] = resp
		c.ServeJSON()
	}

}
