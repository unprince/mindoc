package main

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/lifei6671/mindoc/commands"
	"github.com/lifei6671/mindoc/controllers"
	"github.com/lifei6671/mindoc/conf"
)

func main() {

	commands.RegisterCommand()
	commands.ResolveCommand(os.Args[1:])

	commands.RegisterFunction()

	beego.ErrorController(&controllers.ErrorController{})

	fmt.Printf("MinDoc version => %s\nbuild time => %s\nstart directory => %s\n%s\n", conf.VERSION, conf.BUILD_TIME, os.Args[0], conf.GO_VERSION)

	beego.Run()
}
