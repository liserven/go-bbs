package home

import (
	"myapp/controllers"
)

type TopicController struct {
	controllers.BaseController
}

//文章列表展示
func (this *TopicController) Get() {
	this.TplName = "home/topic/index.html"
}
