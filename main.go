package main

import (
	"micgofiber/cmd"
	"micgofiber/lib"
	"micgofiber/router"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		cmd.Init()
	} else {
		app := lib.NewApp()
		router.InitApp(app)
		app.App.Listen(":" + app.Port)
	}
}
