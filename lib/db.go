package lib

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Db *gorm.DB
}

func NewDB() *DBConfig {
	if err := godotenv.Load("../.env"); err != nil {
		if err := godotenv.Load(); err != nil {
			log.Fatalln("Failed to get .env file")
		}
	}

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASS") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE") +
		" TimeZone=" + os.Getenv("DB_TIMEZONE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to Connect Database")
	}

	return &DBConfig{
		Db: db,
	}
}
