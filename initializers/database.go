package initializers

import (
	"log"
	"os"

	"github.com/moosashah/go-crud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB", err.Error())
		os.Exit(1)
	}
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	DB.Migrator().DropTable(&models.Tournament{})
	DB.AutoMigrate(&models.Tournament{})
	log.Println("Connected to Database Successfully")
}
