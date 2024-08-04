package main

import (
	"micgofiber/db"
	"micgofiber/lib"
	"micgofiber/router"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		dbConfig := lib.NewDB()

		if os.Args[1] == "migrate" {
			db.Migrate(dbConfig)
		}
		if os.Args[1] == "reset-migrate" {
			db.ResetMigrate(dbConfig)
		}
	} else {
		app := lib.NewApp()
		router.InitApp(app)
		app.App.Listen(":" + app.Port)
	}
}
