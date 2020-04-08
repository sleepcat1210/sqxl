package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

type Category struct {
	CategoryId int64 `orm:"auto;pk";json:"category_id"`
	CategoryName string `orm:"size(50);description(分类名称)";json:"category_name"`
	Pid int64 `orm:"default(0);description(父级)";json:"pid"`
	CategoryLogo string `orm:"size(100);description(分类logo)";json:"category_logo"`
	CategoryOrder int	`orm:"description(分类排序)";json:"category_order"`
	CategoryStatus int8 `orm:"default(0);description(分类状态0禁用1启用)";json:"category_status"`
	GoodsType []*GoodsType `orm:"reverse(many)"`
	Goods []*Goods `orm:"reverse(many)"`

}
type CategoryList struct{
	CategoryId int64 `json:"category_id"`
	CategoryName string `json:"category_name"`
	Pid int64 `json:"pid"`
	CategoryLogo string `json:"category_logo"`
	CategoryOrder int	`json:"category_order"`
	CategoryStatus int8 `json:"category_status"`
	Children []*CategoryList	`json:"children"`
}
type CategoryTree struct{
	Category
	Level int
	Htmlstring string
}



func(c *Category)GetCategory()[]*CategoryList{
	return c.GetList(0)
}
func(c *Category)GetList(pid int64)[]*CategoryList{
    o:=orm.NewOrm()
    var category []Category
    o.QueryTable("SqxlCategory").Filter("Pid",pid).OrderBy("CategoryOrder").All(&category)
    var category_menu []*CategoryList
    for _,v:=range category{
    	child :=v.GetList(v.CategoryId)
    	nodes :=&CategoryList{
			CategoryId:v.CategoryId,
			CategoryName:v.CategoryName,
			Pid:v.Pid,
			CategoryLogo:v.CategoryLogo,
			CategoryOrder:v.CategoryOrder,
			CategoryStatus:v.CategoryStatus,
		}
		nodes.Children=child
		category_menu=append(category_menu,nodes)
	}
	return category_menu

}
func(c *Category)GetTree()(tree []CategoryTree){
	o:=orm.NewOrm()
	var category []Category
	o.QueryTable("SqxlCategory").All(&category)
	return c.GetTrees(category,0,0)
}
func(c *Category)GetTrees(data []Category,pid int64,level int)(tree []CategoryTree){
	for _,val:=range data{
		if val.Pid==pid{
			tree =append(tree,CategoryTree{
				Category:val,
				Level:level,
				Htmlstring:strings.Repeat("&nbsp",level*3),
			})
			tree =append(tree,c.GetTrees(data,val.CategoryId,level+1)...)
		}
	}
	return tree
}
