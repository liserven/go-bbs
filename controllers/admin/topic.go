package admin

import (
	"myapp/controllers"
	"myapp/models"
	"strconv"
)

type TopicController struct {
	controllers.BaseController
}

func (this *TopicController) TopicList() {
	var topicModel models.TopicModel
	topicList := topicModel.Select("", 0, 0)
	this.Data["list"] = topicList
	this.TplName = "admin/topic_list.html"
}

func (this *TopicController) TopicAdd() {
	this.TplName = "admin/topic_add.html"
}

func (this *TopicController) TopicDoAdd() {
	topic := this.GetString("topic")
	content := this.GetString("content")
	maps := make(map[string]interface{}, 10)
	maps["topic"] = topic
	maps["content"] = content
	var topicModel models.TopicModel
	err := topicModel.Insert(maps)
	if !err {
		this.JsonResult(0, "add  error!")

	} else {
		this.JsonResult(1, "add ok")
	}
	return
}

func (this *TopicController) TopicDel() {
	id := this.GetString("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		this.JsonResult(0, "id有误")
	}
	var topicModel models.TopicModel
	isSuccess := topicModel.DoDel(intId)
	if !isSuccess {
		this.JsonResult(0, "删除失败")
	} else {
		this.JsonResult(1, "删除成功")
	}
}
