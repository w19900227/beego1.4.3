package baseController

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/session"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) ServeError(err error) {
	this.Data["json"] = struct {
		Error string `json:"Error"`
	}{err.Error()}
	this.Ctx.Output.SetStatus(500)
	this.ServeJson()
}

func (this *BaseController) AjaxResponse(status string, erron string, errmsg string, data interface{}) {
	response := struct {
		Status string
		Erron  string
		Errmsg string
		Data   interface{}
	}{
		Status : status,
		Erron  : erron,
		Errmsg : errmsg,
		Data   : data,
	}

	this.Data["json"] = response
	this.ServeJson()
}
