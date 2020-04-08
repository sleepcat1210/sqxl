package AdminControllers

import (
	"sqlx/models"
	"encoding/json"
	"strings"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type GoodsController struct {
	BaseController
}
func(this *GoodsController)GetList(){
	good :=models.Goods{}
	goods:=good.GetList()
	fmt.Println(goods)
	this.Data["goods"]=goods
	this.TplName="back/goods/index.html"
}
//添加产品
func(this *GoodsController)AddShow(){
	//获取商品固有属性
	category:=models.Category{}
	categorys:=category.GetTree()
	brand:=models.Brand{}
	brands,_:=brand.GetList()

	this.Data["brands"]=brands
	this.Data["category_list"]=categorys
	this.TplName="back/goods/add.html"
}
func(this *GoodsController)GetAttr(){
	attr:=new(models.Attribute)
	attrs,_:=attr.GetAttr(0)
	attrm,_:=attr.GetAttr(1)
	arrMap:=make(map[int]map[string]interface{})
	for k,v:=range attrm{
		j, _ := json.Marshal(v)
		attrss:=make(map[string]interface{})
		json.Unmarshal(j, &attrss)
		if v.AttrInputType==1&&v.AttrValues!=""{
			attrss["AttrValues"]=strings.Split(v.AttrValues,"|")
		}
		arrMap[k]=attrss
	}
	this.Data["attrMap"]=arrMap
	this.Data["attrs"]=attrs
	this.TplName="back/goods/attr.html"
}
//添加商品
func(this *GoodsController)AddGoods(){
	goods_name:=this.GetString("goods_name")
	goods_sn:=this.GetString("goods_sn")
	category_id,_:=this.GetInt64("category_id")
	category:=&models.Category{CategoryId:category_id}
	store_count,_:=this.GetInt("store_count")
	original_img:=this.GetString("original_img")
	is_on_sale,_:=this.GetInt("is_on_sale")
	is_new,_:=this.GetInt("is_new")
	is_recommend,_:=this.GetInt("is_recommend")
	is_hot,_:=this.GetInt("is_hot")
	keywords:=this.GetString("keywords")
	goods_remark:=this.GetString("goods_remark")
	shop_price,_:=this.GetFloat("shop_price")
	market_price,_:=this.GetFloat("market_price")
	cost_price,_:=this.GetFloat("cost_price")
	brand_id,_:=this.GetInt64("brand_id")
	brand:=&models.Brand{BrandId:brand_id}
	o:=orm.NewOrm()
	goods :=models.Goods{}
	goods.Category=category
	goods.CostPrice=cost_price
	goods.GoodsName=goods_name
	goods.GoodsSn=goods_sn
	goods.ShopPrice=shop_price
	goods.StoreCount=store_count
	goods.OriginalImg=original_img
	goods.IsOnSale=is_on_sale
	goods.IsNew=is_new
	goods.IsRecommend=is_recommend
	goods.IsHot=is_hot
	goods.Keywords=keywords
	goods.GoodsRemark=goods_remark
	goods.MarketPrice=market_price
	goods.Brand=brand
	_,err:= o.Insert(&goods)
	if err !=nil{
		fmt.Println(err)
		this.JsonResult(0,"添加失败")
	}else{
		this.JsonResult(200,"添加成功")
	}
}