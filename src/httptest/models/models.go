package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User), new(UserConn))
}

func AddUser(name string, pwd string, age int16) (int64, error) {
	o := orm.NewOrm()
	user := &User{
		Name: name,
		Pwd:  pwd,
		Age:  age,
	}
	return o.Insert(user)
}

func AddUserConn(user_id int, phone string, mobile string, email string, qq string) (int64, error) {
	o := orm.NewOrm()
	user := &User{Id: user_id}
	//o.Read(user)
	user_conn := &UserConn{
		Phone:  phone,
		Mobile: mobile,
		Email:  email,
		Qq:     qq,
		User:   user,
	}
	return o.Insert(user_conn)
}

func (this *User) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(this)
}

func (this *UserConn) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(this)
}
