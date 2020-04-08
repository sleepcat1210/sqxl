package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Goods struct {
	GoodsId int64 `orm:"auto;pk";json:"goods_id"`
	GoodsName string `orm:"size(200);description(分类名称)";json:"goods_name"`
	GoodsSn	string `orm:"size(100);description(商品编号)";json:"goods_sn"`
	ClickCount int `orm:"description(点击次数)";json:"click_count"`
	StoreCount	int `orm:"description(库存)";json:"store_count"`
	CommentCount int `orm:"description(评论数)";json:"comment_count"`
	ShopPrice  float64 `orm:"digits(10);decimals(2);description(本店价格)";json:"shop_price"`
	MarketPrice float64 `orm:"digits(10);decimals(2);description(市场价格)";json:"market_price"`
	CostPrice	float64 `orm:"digits(10);decimals(2);description(成本价)";json:"cost_price"`
	OriginalImg	string `orm:"size(100); description(商品原图)";json:"original_img"`
	Keywords	string `orm:"size(100);description(关键词)";json:"keywords"`
	GoodsRemark	string `orm:"size(200);description(描述)";json:"goods_remark"`
	IsOnSale	int `orm:"default(0);description(是否上架)";json:"is_on_sale"`
	IsRecommend int `orm:"default(0);description(是否推荐)";json:"is_recommend"`
	IsNew	int `orm:"default(0);description(是否新品)";json:"is_new"`
	IsHot	int `orm:"default(0);description(是否热卖)";json:"is_hot"`
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)";json:"created_time"`
	UpdatedTime time.Time `orm:"auto_now;type(datetime)";json:"updated_time"`
	Brand *Brand `orm:"rel(fk);description(品牌)";json:"brand_id"`
	Category *Category `orm:"rel(fk);description(商品分类)";json:"category_id"`
	GoodsType *GoodsType `orm:"rel(fk);description(商品类型);null";json:"type_id"`
}
func (b *Goods) TableEngine() string {
	return "INNODB"
}
//获取商品列表
func(this *Goods)GetList()[]Goods{
	o:=orm.NewOrm()
	var goods []Goods
	o.QueryTable("SqxlGoods").RelatedSel("Brand").RelatedSel("Category").All(&goods)
	return goods
}

