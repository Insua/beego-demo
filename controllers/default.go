package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego-app.test"
	c.Data["Email"] = "insuaxe@gmail.com"
	c.TplName = "index.tpl"
}

type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {
	this.Ctx.WriteString("hello")
}
