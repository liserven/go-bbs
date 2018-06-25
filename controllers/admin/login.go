package admin

import "myapp/controllers"

type LoginController struct {
	controllers.BaseController
}

func (this *LoginController) Get() {
	this.TplName = "admin/login.html"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	if username == "" {
		this.JsonResult(0, "账号不能为空")
		return
	}
	pwd := this.GetString("pwd")
	if pwd == "" {
		this.JsonResult(0, "密码不能为空")
		return
	}

	this.JsonResult(1, "登录成功")
	return

}
