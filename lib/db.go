package lib

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Db *gorm.DB
}

func NewDB() *DBConfig {
	dsn := "host=localhost user=postgres password=nprogrammer21 dbname=micgofiber port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &DBConfig{
		Db: db,
	}
}
