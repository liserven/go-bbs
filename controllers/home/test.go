package home

import "myapp/controllers"

type TestController struct {
	controllers.BaseController
}

func (this *TestController) Get() {
	this.TplName = "home/index.html"
}
