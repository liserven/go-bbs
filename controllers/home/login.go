package home

import "myapp/controllers"

type LoginController struct {
	controllers.BaseController
}

func (this *LoginController) Get() {
	this.TplName = "home/login/index.html"
}
