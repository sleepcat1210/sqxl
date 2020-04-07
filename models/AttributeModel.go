package models

import "github.com/astaxie/beego/orm"

type Attribute struct {
	AttrId int64 `orm:"auto;pk";json:"attr_id"`
	AttrName string `orm:"size(100) description:(属性名称)";json:"attr_name"`
	AttrType int `orm:"default(0);description(属性类型0固有属性1销售属性)";json:"attr_type"`
	AttrInputType int	`orm:"default(0);description(0手工录入1列表)";json:"attr_input_type"`
	AttrValues string `orm:"size(100);description(可选值列表)";json:"attr_values"`
	GoodsType *GoodsType `orm:"rel(fk);description:(类型)";json:"type_id"`
}

func(this *Attribute)Add(attr_name string,attr_type int,attr_input_type int,attr_values string,GoodsType *GoodsType)(int64,error){
     this.GoodsType=GoodsType
     this.AttrName=attr_name
     this.AttrType=attr_type
     this.AttrInputType=attr_input_type
     this.AttrValues=attr_values
     o:=orm.NewOrm()
     return o.Insert(this)
}
func(this *Attribute)GetList(typeId int64)(attribute []Attribute,err error){
	o:=orm.NewOrm()
	_,err=o.QueryTable("SqxlAttribute").RelatedSel("GoodsType").Filter("GoodsType__TypeId",typeId).All(&attribute)
	return
}
//获取属性
func(this *Attribute)GetAttr(attr_type int)(attribute []Attribute,err error){
	o:=orm.NewOrm()
	_,err=o.QueryTable("SqxlAttribute").RelatedSel("GoodsType").Filter("AttrType",attr_type).All(&attribute)
	return
}