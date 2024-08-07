package db

import (
	"micgofiber/lib"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Check       bool   `json:"check"`
}

func TodoMigration(config *lib.DBConfig) {
	config.Db.AutoMigrate(&Todo{})
}

func TodoDropMigration(config *lib.DBConfig) {
	if (config.Db.Migrator().HasTable(&Todo{})) {
		config.Db.Migrator().DropTable(&Todo{})
	}
}
