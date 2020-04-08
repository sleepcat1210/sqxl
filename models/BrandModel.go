package models

import "time"

type Brand struct {
	BrandId int64 `orm:"auto;pk";json:"brand_id"`
	BrandName string `orm:"size(50);description:(品牌名称)";json:"brand_name"`
	Telephone string `orm:"size(11);description(联系电话)";json:"telephone"`
	BrandWeb string `orm:"size(100);description(品牌网络)";json:"brand_web"`
	BrandLogo string `orm:"size(100);description(品牌logo)";json:"brand_logo"`
	BrandDesc  string `orm:"size(150);description(品牌描述)";json:"brand_desc"`
	BrandStatus  int8 `orm:"default(0);description(品牌状态0禁用1启用)";json:"brand_status"`
	BrandSzsx	string `orm:"size(10);description(品牌首字母缩写)";json:"brand_szsx"`
	BrandOrder int	`orm:"description(品牌排序)";json:"brand_order"`
	Goods []*Goods `orm:"reverse(many)"`
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)";json:"created_time"`
	UpdatedTime time.Time `orm:"auto_now;type(datetime)";json:"updated_time"`
}
func (b *Brand) TableEngine() string {
	return "INNODB"
}