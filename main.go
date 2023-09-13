package main

import (
	"ffyouku/models"
	_ "ffyouku/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.Init()

	beego.Run()
}
