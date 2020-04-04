package AdminControllers

import (
	"sqlx/utils"
	"github.com/astaxie/beego/orm"
	"sqlx/models"
	"strings"
)

type BrandController struct {
	BaseController
}
func(b *BrandController)GetList(){
	var brands []models.Brand
	o:=orm.NewOrm()
	o.QueryTable("SqxlBrand").OrderBy("-BrandOrder").All(&brands)
	b.Data["brands"]=brands
	b.TplName="back/brand/index.html"
}
//添加品牌
func(b *BrandController)AddShow(){
	b.TplName="back/brand/add.html"
}
//提交数据
func (b *BrandController)AddBrand(){
	brand_name :=b.GetString("brand_name")
	telephone :=b.GetString("telephone")
	brand_web :=b.GetString("brand_web")
	brand_desc :=b.GetString("brand_desc")
	brand_logo :=b.GetString("brand_logo")
	brand_status,_:=b.GetInt8("brand_status")
	brand_order,_:=b.GetInt("brand_order")
	brand_szsx,_:=utils.SzmPinyin(brand_name)
	o :=orm.NewOrm()
	brand :=models.Brand{
		BrandName:brand_name,
		BrandDesc:brand_desc,
		BrandWeb:brand_web,
		BrandStatus:brand_status,
		Telephone:telephone,
		BrandLogo:brand_logo,
		BrandSzsx:brand_szsx,
		BrandOrder:brand_order,
	}
	result,err:=o.Insert(&brand)
	resp := make(map[string]interface{})
	if err !=nil{
		resp["code"]="0"
		resp["msg"]="添加错误"
		resp["data"]=""
		b.Data["json"] = resp
		b.ServeJSON()
	}
	if result>0{
		resp["code"]="200"
		resp["msg"]="添加成功"
		resp["data"]=""
		b.Data["json"] = resp
		b.ServeJSON()
	}

}
//删除
func(b *BrandController)DelBrand(){
	brandId,_:=b.GetInt64("id")
	o:=orm.NewOrm()
	var brand models.Brand
	brand.BrandId=brandId
	_,err:=o.Delete(&brand)
	resp := make(map[string]interface{})
	if err !=nil{
		resp["code"]="0"
		resp["msg"]="删除错误"
		resp["data"]=""
		b.Data["json"] = resp
		b.ServeJSON()
	}else{
		resp["code"]="200"
		resp["msg"]="删除成功"
		resp["data"]=""
		b.Data["json"] = resp
		b.ServeJSON()
	}
}
//删除选择
func(b *BrandController)DelAll(){
 ids :=b.GetString("ids")
 id:=strings.Split(ids,",")
 o:=orm.NewOrm()
 _,err:=o.QueryTable("SqxlBrand").Filter("BrandId__in",id).Delete()
	resp := make(map[string]interface{})
	if err !=nil{
		resp["code"]="0"
		resp["msg"]="删除错误"
		resp["data"]=""
		b.Data["json"] = resp
		b.ServeJSON()
	}else{
		resp["code"]="200"
		resp["msg"]="删除成功"
		resp["data"]=""
		b.Data["json"] = resp
		b.ServeJSON()
	}
}
//编辑
func(b *BrandController)EditShow(){
    id,_ :=b.GetInt64("id")
    o:=orm.NewOrm()
    brand:= models.Brand{}
    brand.BrandId=id
    o.Read(&brand,"BrandId")
    b.Data["brand"]=brand
    b.TplName="back/brand/edit.html"
}
//编辑更新
func(b *BrandController)UpdateBrand(){
	brand_name :=b.GetString("brand_name")
	telephone :=b.GetString("telephone")
	brand_web :=b.GetString("brand_web")
	brand_desc :=b.GetString("brand_desc")
	brand_logo :=b.GetString("brand_logo")
	brand_status,_:=b.GetInt8("brand_status")
	brand_order,_:=b.GetInt("brand_order")
	brand_id,_:=b.GetInt64("brand_id")
	brand_szsx,_:=utils.SzmPinyin(brand_name)
	o :=orm.NewOrm()
	brand :=models.Brand{
		BrandName:brand_name,
		BrandDesc:brand_desc,
		BrandWeb:brand_web,
		BrandStatus:brand_status,
		Telephone:telephone,
		BrandLogo:brand_logo,
		BrandSzsx:brand_szsx,
		BrandOrder:brand_order,
		BrandId:brand_id,
	}
	result,err:=o.Update(&brand)
	resp := make(map[string]interface{})
	if err !=nil{
		resp["code"]="0"
		resp["msg"]="编辑错误"
		resp["data"]=""
		b.Data["json"] = resp
		b.ServeJSON()
	}
	if result>0{
		resp["code"]="200"
		resp["msg"]="编辑成功"
		resp["data"]=""
		b.Data["json"] = resp
		b.ServeJSON()
	}
}