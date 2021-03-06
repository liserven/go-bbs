package controllers

import (
	"encoding/json"
	"io"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) JsonResult(errCode int, errMsg string, data ...interface{}) {
	jsonData := make(map[string]interface{}, 3)
	jsonData["error"] = errCode
	jsonData["msg"] = errMsg
	if len(data) > 0 && data[0] != nil {
		jsonData["data"] = data[0]
	}
	resultJson, err := json.Marshal(jsonData)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")
	io.WriteString(c.Ctx.ResponseWriter, string(resultJson))
	c.StopRun()

}
