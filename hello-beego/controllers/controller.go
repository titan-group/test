package controllers

import (
	"github.com/titan-group/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["message"] = "Fuck out"
	this.TplNames = "index.tpl"
	fmt.Println("============= OK =============")
}

