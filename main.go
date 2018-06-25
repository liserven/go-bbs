package main

import (
	"myapp/models"
	_ "myapp/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("username")+":root@tcp(localhost:3306)/qz?charset=utf8")
	orm.RegisterModel(
		new(models.RoleData),
	)
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.InsertFilter("/delete_role", beego.BeforeRouter, UserFilter)

	beego.Run()
}

var UserFilter = func(ctx *context.Context) {
	ctx.Redirect(302, "/error?id=1")
}

//也可以了
