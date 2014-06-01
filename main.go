package main

import (
	"github.com/astaxie/beego"
	"github.com/ulricqin/gen-winning-numbers/g"
	"github.com/ulricqin/gen-winning-numbers/models"
	_ "github.com/ulricqin/gen-winning-numbers/routers"
)

func main() {
	g.InitEnv()
	models.InitBusinessCfg()
	beego.Run()
}
