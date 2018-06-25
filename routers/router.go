package routers

import (
	"myapp/controllers"
	"myapp/controllers/admin"
	"myapp/controllers/home"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &admin.RoleController{})
	beego.Router("/delete_role", &admin.RoleController{}, "get:DeleteRole")
	beego.Router("/error", &controllers.ErrorController{})

	//登录界面
	beego.Router("/admin_login_index", &admin.LoginController{})

	//首页
	beego.Router("/admin_index", &admin.IndexController{})

	beego.Router("/welcome", &admin.WelcomeController{})

	//前台登录

	beego.Router("/login", &home.LoginController{})
	beego.Router("/", &home.IndexController{})

	beego.Router("/test", &home.TestController{})

	beego.Router("/topic_list", &admin.TopicController{}, "get:TopicList")
	beego.Router("/topic_add", &admin.TopicController{}, "get:TopicAdd")
	beego.Router("/topic_do_add", &admin.TopicController{}, "post:TopicDoAdd")
	beego.Router("/topic_del", &admin.TopicController{}, "post:TopicDel")

}
