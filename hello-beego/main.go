package main

import (
	"./controllers"
	"github.com/titan-group/beego"
	"github.com/titan-group/beego/context"
)

/*
func init() {
	beego.Router("/", &controllers.MainController{})
}
*/


func TT(ctx *context.Context){
	ctx.Output.Body ( []byte("it is a good day to die"))
}

func main() {
	beego.HttpPort=8888
	beego.Router("/a", &controllers.MainController{})
	beego.Get("/a", TT)
	beego.Run()
}
