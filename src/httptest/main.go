package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	_ "httptest/models"
	_ "httptest/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/go_test?charset=utf8")
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	// o := orm.NewOrm()
	// //o.Using("default")

	// var user models.User
	// user.Id = 1
	// user.Name = "admin"
	// user.Age = 20

	// var user_conn models.UserConn
	// user_conn.Id = 1
	// user_conn.Phone = "023-1231123"
	// user_conn.Mobile = "12312312312"
	// user_conn.Email = "admin@test.com"
	// user_conn.User = &user

	// fmt.Println(o.Insert(&user))
	// fmt.Println(o.Insert(&user_conn))

	beego.Run()
}
