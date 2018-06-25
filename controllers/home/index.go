package home

import (
	"myapp/controllers"
	"myapp/models"
)

type IndexController struct {
	controllers.BaseController
}

var (
	TopicM models.TopicModel
)

func (i *IndexController) Get() {
	list := TopicM.Select("", 0 ,0)
	i.Data["list"] = list
	i.TplName = "home/index.html"
}
