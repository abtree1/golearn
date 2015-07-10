package routers

import (
	"github.com/astaxie/beego"

	"httptest/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/files", &controllers.FilesController{})
}
