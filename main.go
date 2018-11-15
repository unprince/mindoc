package main

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/memcache"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lifei6671/mindoc/commands"
	"github.com/lifei6671/mindoc/conf"
	"github.com/lifei6671/mindoc/controllers"
	_ "github.com/lifei6671/mindoc/routers"
)

func main() {

	commands.RegisterCommand()
	commands.ResolveCommand(os.Args[1:])

	commands.RegisterFunction()

	beego.ErrorController(&controllers.ErrorController{})

	fmt.Printf("MinDoc version => %s\nbuild time => %s\nstart directory => %s\n%s\n", conf.VERSION, conf.BUILD_TIME, os.Args[0], conf.GO_VERSION)

	beego.Run()
}
