package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Get() {
	this.Data["IsLogin"] = this.GetString("id")
	this.TplName = "error.html"
}
