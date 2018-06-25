package admin

import "myapp/controllers"

type WelcomeController struct {
	controllers.BaseController
}

func (this *WelcomeController) Get() {
	this.TplName = "admin/welcome.html"
}
