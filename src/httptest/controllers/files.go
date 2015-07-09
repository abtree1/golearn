package controllers

import (
	"github.com/astaxie/beego"
)

type FilesController struct {
	beego.Controller
}

func (this *FilesController) Get() {
	this.Data["Files"] = true
	this.TplNames = "index.html"
}
