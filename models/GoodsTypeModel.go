package models

import "github.com/astaxie/beego/orm"

type GoodsType struct {
	TypeId int64 `orm:"auto;pk";json:"type_id"`
	TypeName string `orm:"size(100);description:(类型名称)";json:"type_name"`
	Category	*Category `orm:"rel(fk);description:(所属分类)";json:"category_id"`
	Attribute []*Attribute `orm:"reverse(many)"`
	Goods []*Goods `orm:"reverse(many)"`
}

//添加
func(this *GoodsType)Add(type_name string,category *Category)(int64,error){
	o:=orm.NewOrm()
	this.TypeName=type_name
	this.Category=category
	return o.Insert(this)
}
//获取列表
func(this *GoodsType)GetList()(goodsType []GoodsType,err error){
	o:=orm.NewOrm()
	_,err=o.QueryTable("SqxlGoodsType").RelatedSel("Category").All(&goodsType)
	return
}