package controllers

import (
	"strconv"

	"github.com/astaxie/beego"

	"httptest/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["Home"] = true
	this.TplNames = "index.html"
}

func (this *IndexController) Post() {
	name := this.Input().Get("name")
	pwd := this.Input().Get("pwd")
	agestr := this.Input().Get("age")
	phone := this.Input().Get("phone")
	mobile := this.Input().Get("mobile")
	qq := this.Input().Get("qq")
	email := this.Input().Get("email")
	age, _ := strconv.Atoi(agestr)

	// 方法1
	// id, _ := models.AddUser(name, pwd, int16(age))
	// models.AddUserConn(int(id), phone, mobile, email, qq)

	// 方法2 interface
	user := &models.User{
		Name: name,
		Pwd:  pwd,
		Age:  int16(age),
	}
	id, _ := user.Add()
	user.Id = int(id)
	user_conn := &models.UserConn{
		Phone:  phone,
		Mobile: mobile,
		Email:  email,
		Qq:     qq,
		User:   user,
	}
	user_conn.Add()

	this.Data["Home"] = true
	this.TplNames = "index.html"
	return
}
