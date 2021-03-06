package AdminControllers

import (
	"github.com/astaxie/beego"
	"io"
	"sqlx/utils/filecache"
)

type BaseController struct {
	beego.Controller
}
func(b *BaseController)Finish(){
	controllerName,actionName:=b.GetControllerAndAction()
	if filecache.NeedWrite(controllerName,actionName,b.Ctx.Input.Params()){
		render,err:=b.RenderString()
		if nil==err && len(render)>0{
			filecache.Write(controllerName,actionName,b.Ctx.Input.Params(),&render)
		}
	}
}
func(b *BaseController)Prepare(){
controllerName,actionName:=b.GetControllerAndAction()
if filecache.InCacheList(controllerName,actionName,b.Ctx.Input.Params()){//如果在队列里
	contentPtr,err:=filecache.Read(controllerName,actionName,b.Ctx.Input.Params())
	if nil==err && len(*contentPtr)>0{
		io.WriteString(b.Ctx.ResponseWriter,*contentPtr)
		b.StopRun()
	}
}
}
//返回结果
func(b *BaseController)JsonResult(errcode int,errmsg string,data ...interface{}){
		jsonData :=make(map[string]interface{},3)
		jsonData["code"]=errcode
		jsonData["message"]=errmsg
		if len(data)>0 && data[0]!=nil{
			jsonData["data"]=data[0]
		}
	b.Data["json"] = jsonData
	b.ServeJSON()
}
