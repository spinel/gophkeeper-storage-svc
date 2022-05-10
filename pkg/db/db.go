package db

import (
	"log"

	"github.com/spinel/gophkeeper-storage-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Entity{})
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.CreditCard{})

	return Handler{db}
}
