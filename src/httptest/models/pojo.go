package models

type User struct {
	Id       int
	Name     string
	Pwd      string
	Age      int16
	UserConn []*UserConn `orm:"reverse(many)"`
}

type UserConn struct {
	Id     int
	Phone  string
	Mobile string
	Email  string
	Qq     string
	User   *User `orm:"rel(fk)"`
}

type ICrud interface {
	Add() (int, error)
}
