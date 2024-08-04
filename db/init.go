package db

import "micgofiber/lib"

func Migrate(config *lib.DBConfig) {
	TodoMigration(config)
}

func ResetMigrate(config *lib.DBConfig) {
	TodoDropMigration(config)
	Migrate(config)
}
