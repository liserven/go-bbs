package admin

import "myapp/controllers"

type IndexController struct {
	controllers.BaseController
}

func (this *IndexController) Get() {
	this.TplName = "admin/index.html"
}
