package models

type Attribute struct {
	AttrId int64 `orm:"auto;pk";json:"attr_id"`
	AttrName string `orm:"size(100) description:(属性名称)";json:"attr_name"`
	AttrType int `orm:"default(0);description(属性类型0固有属性1销售属性)";json:"attr_type"`
	AttrInputType int	`orm:"default(0);description(0手工录入1列表)";json:"attr_input_type"`
	AttrValues string `orm:"size(100);description(可选值列表)";json:"attr_values"`
	GoodsType *GoodsType `orm:"rel(fk);description:(类型)";json:"type_id"`
}