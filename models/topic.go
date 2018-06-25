package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type TopicModel struct {
	Id               int64
	Uid              int64
	Title            string
	Content          string `orm:size(6000)`
	Attachment       string
	Created          time.Time `orm:"index"`
	Updated          time.Time `orm:"index"`
	Views            int64     `orm:"index"`
	Author           string
	ReplyTime        time.Time
	ReplyCount       int64
	RepleyLastUserId int64
}

//insert one
/**
* param 添加时候接收和参数
**/
func (this *TopicModel) Insert(param map[string]interface{}) bool {
	fmt.Println(param)
	o := orm.NewOrm()
	sql := fmt.Sprintf("insert into topic ( title, content) value( %q, %q)", param["topic"], param["content"])
	_, err := o.Raw(sql).Exec()
	if err != nil {
		return false
	} else {
		return true
	}
}

//select
func (this *TopicModel) Select(where string, offset int, limit int) []orm.Params {
	limits := ""
	if limit > 0 {
		limits = fmt.Sprintf("limit %d, %d", offset, limit)
	}
	sql := fmt.Sprintf("select * from topic ")
	if where != "" {
		sql = sql + where
	}
	sql = sql + limits + " ORDER BY id DESC"
	fmt.Println(sql)
	var maps []orm.Params
	o := orm.NewOrm()
	o.Raw(sql).Values(&maps)
	return maps

}

func (this *TopicModel) DoDel(id int) bool {
	o := orm.NewOrm()
	sql := fmt.Sprintf("delete from topic where id=%d", id)
	_, err := o.Raw(sql).Exec()
	if err != nil {
		return false
	} else {
		return true
	}
}
