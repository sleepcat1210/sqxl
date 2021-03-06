package filecache

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego"
	"errors"
	"time"
	"strings"
)

var(
	BasePath string =""//文件路径
	ExpirseSec int64=0//过期时间
	store *cache.FileCache=nil
	cacheMap map[string]bool=nil
	paramMap map[string][]string=nil
)
//初始化
func InitCache(){
	store =&cache.FileCache{CachePath:BasePath}
	filecacheList:=beego.AppConfig.Strings("filecache_list")
	//初始化静态化配置列表
	cacheMap=make(map[string]bool)
	for _,v:=range filecacheList{
		cacheMap[v]=true
	}
	paramMap=make(map[string][]string)
	filecacheMap,_:=beego.AppConfig.GetSection("filecache_param")
	for k,v:=range filecacheMap{
		sv:=strings.Split(v,";")
		paramMap[k]=sv
	}
	}
//是否存在缓存列表
func InCacheList(controllerName,actionName string,param map[string]string)bool{
	keyName:=cacheKey(controllerName,actionName,param)
	if f:=cacheMap[keyName];f{
		return f
	}
	return false
}
//是否写缓存
func NeedWrite(controllerName,actionName string,param map[string]string)bool{
	if InCacheList(controllerName,actionName,param){
		keyName:=cacheKey(controllerName,actionName,param)
		if len(store.Get(keyName).(string))>0{
			return false
		}else{
			return true
		}
	}
	return false
}
//写
func Write(controllerName,actionName string,param map[string]string,content *string)error{
	keyname :=cacheKey(controllerName,actionName,param)
	if len(keyname)==0{
		return errors.New("未找到缓存key")
	}
	err:=store.Put(keyname,*content,time.Duration(ExpirseSec)*time.Second)
	return err
}
//读
func Read(controllerName,actionName string,param map[string]string)(*string,error){
	keyname:=cacheKey(controllerName,actionName,param)
	if len(keyname)==0{
		return nil,errors.New("未找到缓存key")
	}
	content:=store.Get(keyname).(string)
	return &content,nil
}
//设置缓存key
func cacheKey(controllerName,actionName string,paramArray ...map[string]string)string{
	if len(controllerName)>0 && len(actionName)>0{
		rtnstr:=strings.ToLower(controllerName)+"_"+strings.ToLower(actionName)
		if len(paramArray)>0{
			for _,v:=range paramMap[rtnstr]{
				rtnstr =rtnstr +"_"+strings.ReplaceAll(v,":","")+"_"+paramArray[0][v]
			}
		}
		return rtnstr
	}
	return ""
}
//删除
func ClearExpiredFiles(){
	for k,_:=range cacheMap{
		if store.IsExist(k){
			content:=store.Get(k).(string)
			if len(content)==0{
				store.Delete(k)
			}
		}
	}
}