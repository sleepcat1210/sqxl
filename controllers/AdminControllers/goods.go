package AdminControllers

import (
	"sqlx/models"
	"encoding/json"
	"strings"
	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
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
	typeid,_:=this.GetInt64("typeid")
	attr:=new(models.Attribute)
	attrm,_:=attr.GetAttr(typeid)
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
//设置商品属性
func(this *GoodsController)SetAddr(){
	goods_id,_:=this.GetInt64("goods_id")
	o:=orm.NewOrm()
	goods:=models.Goods{
		GoodsId:goods_id,
	}
	o.Read(&goods,"GoodsId")
	cateType:=[]models.GoodsType{}
	o.QueryTable("SqxlGoodsType").Filter("Category",goods.Category).All(&cateType)
	this.Data["cateType"]=cateType
	this.Data["goods_id"]=goods_id
	this.TplName="back/goods/setattr.html"
}
func(this *GoodsController)SetSpu(){
	goods_id,_:=this.GetInt64("goods_id")
	type_id,_:=this.GetInt64("type_id")
	attrs:=make(map[string]string,0)
	this.Ctx.Input.Bind(&attrs,"attr")
	var attrIds string
	o:=orm.NewOrm()
	o.Begin()
	is_delete :=true
	goods:=&models.Goods{GoodsId:goods_id}
	goodsType:=&models.GoodsType{TypeId:type_id}
	attr:=models.GoodsAttr{Goods:goods}
	o.Delete(&attr,"Goods")
	for k,v:=range attrs{
		ks:=strings.Split(k,"-")//0-属性-
		if len(ks[1])>0 &&v!=""{
			if ks[0]=="1"{
				if !strings.Contains(attrIds,ks[1]){
					attrIds =attrIds+ks[1]+","
				}
			}
			attr_id,_:=strconv.ParseInt(ks[1],10,64)
			goodsAttr:=models.GoodsAttr{}
			attrbute:=&models.Attribute{AttrId:attr_id}
			goodsAttr.Attribute=attrbute
			goodsAttr.Goods=goods
			goodsAttr.AttrName=v
			goodsAttr.GoodsType=goodsType
			_,err:=o.Insert(&goodsAttr)
			if err !=nil{
				is_delete=false
				o.Rollback()
				break
			}
		}
	}
	if is_delete{
		o.Commit()
	}
	attrIds=strings.TrimRight(attrIds,",")
	attrArr:=strings.Split(attrIds,",")
	attrbutes:=[]models.Attribute{}
	o.QueryTable("SqxlAttribute").Filter("AttrId__in",attrArr).OrderBy("AttrId").All(&attrbutes)
	goods_attrs:=make(map[int][]models.GoodsAttr)
	var arrindex int=0
	for _,val:=range attrbutes{
		GoodsAttr:=[]models.GoodsAttr{}
		qs:=o.QueryTable("SqxlGoodsAttr")
		qs.RelatedSel("Attribute").Filter("Attribute__AttrId",val.AttrId).Filter("Goods",goods).OrderBy("GoodsAttrId").All(&GoodsAttr)
		goods_attrs[arrindex]=GoodsAttr
		arrindex++

	}
	goodsattr:=this.GetSku(goods_attrs,0,2)
	fmt.Println(goodsattr)

	//fmt.Println(goodsattr)
	this.Data["goods_attrs"]=goodsattr
	this.Data["attrbutes"]=attrbutes
	this.TplName="back/goods/spu.html"
}
func(this *GoodsController)GetSku(goodsAttr map[int][]models.GoodsAttr,offset int,index int)(goodsattr map[int][]models.GoodsAttr){
		//lenattrType:=len(goods_attrs)
		arrIndex:=len(goodsAttr)//属性分类长度
		goodsattr=make(map[int][]models.GoodsAttr)
		var kk int =0
		for  i:=0;i<index;i++{
			for _,val:=range goodsAttr[offset]{

				if offset<arrIndex-1{
					offset++
				}
				goodsattr[index][kk]=append(goodsattr[index][kk],val)
				fmt.Println(goodsattr)
				//goodsattr=append(goodsattr,this.GetSku(goodsAttr,offset ,index)...)
			}
			kk++
		}

		return goodsattr
}