package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type RoleData struct {
	RdId      int `orm:"pk;auto"`
	RdName    string
	RdLogo    string
	Order     int
	CreatedAt int
	UpdatedAt int
	Status    int
	RdMessage string
}

func (r RoleData) FindRoleById(id int) []orm.Params {

	sql := fmt.Sprintf("select * from role_data where rd_id= %d", id)
	o := orm.NewOrm()
	data := []orm.Params{}
	o.Raw(sql).Values(&data)
	return data
	// 运行一下
	// var maps []orm.Params
	// num, err :=o.QueryTable("role_data").Values(&maps)
	// if err == nil {
	// 	fmt.Printf("Result Nums: %d\n", num)
	// 	for _, m := range maps {
	// 		fmt.Println(m["Id"], m["Name"])
	// 	}
	// }
	// return maps
}

//删除数据
func (r RoleData) Delete(id int) bool {
	sql := fmt.Sprintf("delete from role_data where rd_id= %d", id)
	o := orm.NewOrm()
	_, err := o.Raw(sql).Exec()

	if err != nil {
		return false
	}
	return true
}
