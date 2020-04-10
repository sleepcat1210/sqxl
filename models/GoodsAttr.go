package models
type GoodsAttr struct {
	GoodsAttrId int64 `orm:"auto;pk";json:"goods_attr_id"`
	AttrName string `orm:"size(100);description(属性值)"`
	Goods  *Goods `orm:"rel(fk);description:(商品)";json:"goods_id"`
	GoodsType  *GoodsType `orm:"rel(fk);description:(商品类型)";json:"type_id"`
	Attribute  *Attribute `orm:"rel(fk);description:(属性)";json:"attr_id"`
}
func (this *GoodsAttr) TableEngine() string {
	return "INNODB"
}