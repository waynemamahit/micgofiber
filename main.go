package main

import (
	"micgofiber/db"
	"micgofiber/lib"
	"micgofiber/router"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "migrate" {
			db.Migrate()
		}
	} else {
		app := lib.NewApp()
		router.InitApp(app)
		app.App.Listen(":" + app.Port)
	}
}
