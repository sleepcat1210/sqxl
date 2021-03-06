package routers

import (
	"github.com/astaxie/beego"
	"sqlx/controllers/AdminControllers"
)

func init() {
    admin()
}
func admin(){
	beego.Router("/", &AdminControllers.MainController{})
	//后台
	beego.Router("/admin",&AdminControllers.IndexController{})
	//品牌
	beego.Router("/admin/brand", &AdminControllers.BrandController{},"get:GetList")
	beego.Router("/admin/brand/add", &AdminControllers.BrandController{},"get:AddShow;post:AddBrand")
	beego.Router("/admin/brand/delete",&AdminControllers.BrandController{},"get:DelBrand")
	beego.Router("/admin/brand/delall",&AdminControllers.BrandController{},"get:DelAll")
	beego.Router("/admin/brand/edit",&AdminControllers.BrandController{},"get:EditShow;post:UpdateBrand")
	//图片上传接口
	beego.Router("/upload/one", &AdminControllers.UploadController{},"post:UploadOne")
	//商品分类
	beego.Router("/admin/category",&AdminControllers.CategoryController{},"get:GetList")
	beego.Router("/admin/category/add",&AdminControllers.CategoryController{},"get:AddShow;post:AddCategory")
	beego.Router("/admin/category/delete",&AdminControllers.CategoryController{},"get:Del")
	beego.Router("/admin/category/delall",&AdminControllers.CategoryController{},"get:DelAll")
	beego.Router("/admin/category/edit",&AdminControllers.CategoryController{},"get:Edit;post:UpdateCategory")
	beego.Router("/admin/goodstype",&AdminControllers.GoodsTypeController{},"get:GetList")
	beego.Router("/admin/goodstype/add",&AdminControllers.GoodsTypeController{},"get:GetAdd;post:AddType")
	beego.Router("/admin/attr",&AdminControllers.AttributeController{},"get:GetList")
	beego.Router("/admin/attr/add",&AdminControllers.AttributeController{},"get:Add;post:AddAttr")
	//商品
	beego.Router("/admin/goods",&AdminControllers.GoodsController{},"get:GetList")
	beego.Router("/admin/goods/add",&AdminControllers.GoodsController{},"get:AddShow;post:AddGoods")
	beego.Router("/admin/goods/attr",&AdminControllers.GoodsController{},"get:GetAttr")
	beego.Router("/admin/goods/setattr",&AdminControllers.GoodsController{},"get:SetAddr;post:SetSpu")
}