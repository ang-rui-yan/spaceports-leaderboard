package database

import (
	"fmt"
	"log"
	"os"

	"spaceports-leaderboard/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {

	var dsn string

    if os.Getenv("IS_PROD") == "true"{
		log.Println("Running in production mode")
        dsn = os.Getenv("DATABASE_URL")
    } else {
		log.Println("Running in development mode")
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Singapore",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		)
    }

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Leaderboard{})

	DB = Dbinstance{
		Db: db,
	}
}
