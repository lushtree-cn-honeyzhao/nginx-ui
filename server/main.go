package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	_ "github.com/beego/beego/v2/server/web/session/redis"
	_ "server/config"
	"server/db"
	"server/models"
	_ "server/routers"
)

func init() {
	gob.Register(models.User{})
}

func main() {

	db.Init()
	beego.Run()

}
