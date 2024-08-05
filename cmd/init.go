package cmd

import (
	"micgofiber/db"
	"micgofiber/lib"
	"os"
)

func Init() {
	dbConfig := lib.NewDB()

	if os.Args[1] == "migrate" {
		db.Migrate(dbConfig)
	}

	if os.Args[1] == "reset-migrate" {
		db.ResetMigrate(dbConfig)
	}
}
