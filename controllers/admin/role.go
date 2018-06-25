package admin

import (
	"myapp/controllers"
	"myapp/models"
	"strconv"
)

type RoleController struct {
	controllers.BaseController
}

func (this *RoleController) Get() {
	var roleModel models.RoleData

	id := this.GetString("id")
	intID, _ := strconv.Atoi(id)
	list := roleModel.FindRoleById(intID)
	if list == nil {
		this.JsonResult(1, "暂无数据", list)
	} else {
		this.JsonResult(0, "成功", list)

	}
	return
}

func (this *RoleController) DeleteRole() {
	id := this.GetString("id")
	intId, _ := strconv.Atoi(id)
	var roleModel models.RoleData
	err := roleModel.Delete(intId)
	if err {
		this.JsonResult(0, "ok")
	} else {
		this.JsonResult(1, "error")
	}
	return
}
