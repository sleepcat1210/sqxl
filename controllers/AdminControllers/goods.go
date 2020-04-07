package AdminControllers

import (
	"sqlx/models"
	"encoding/json"
	"strings"
	"fmt"
)

type GoodsController struct {
	BaseController
}
func(this *GoodsController)GetList(){
	this.TplName="back/goods/index.html"
}
//添加产品
func(this *GoodsController)AddShow(){
	//获取商品固有属性
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

	fmt.Println(arrMap)
	this.Data["attrMap"]=arrMap
    this.Data["attrs"]=attrs
	this.TplName="back/goods/add.html"
}