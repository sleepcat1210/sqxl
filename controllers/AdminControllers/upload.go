package AdminControllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"strconv"
	"os"
)

type UploadController struct {
	beego.Controller
}
//上传一张图片
func(u *UploadController)UploadOne(){
	file,info,_ :=u.GetFile("file")
	ext:=path.Ext(info.Filename)
	timenow := time.Now().Unix()
	topath :="static/upload/"+time.Now().Format("2006/01/02/")
	//判断文件
	if _,err :=os.Stat(topath); os.IsNotExist(err){
		os.MkdirAll(topath,0666)
	}
	topath =topath+strconv.FormatInt(timenow,10)+ext
	defer file.Close()
	u.SaveToFile("file",topath)
	resp := make(map[string]interface{})
	resp["code"]="200"
	resp["msg"]="上传成功"
	resp["data"]=topath
	u.Data["json"] = resp
	u.ServeJSON()
}