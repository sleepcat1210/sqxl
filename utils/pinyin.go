package utils

import (
	"github.com/Chain-Zhang/pinyin"
	"strings"
)

//汉字转化拼音
func HzToPinyin(word string)(py string,err error){
	res,err:=pinyin.New(word).Split("").Mode(pinyin.InitialsInCapitals).Convert()
	if err !=nil{
		return "", err
	}
	return res,nil
}
//获取首字母
func SzmPinyin(word string)(py string,err error){
	res,err:=HzToPinyin(word)
	if err !=nil{
		return "",err
	}
   strup:=strings.ToUpper(res)
   return strup[:1],nil


}
