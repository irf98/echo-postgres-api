package db

import (
	"awesomeProject/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func Init() *gorm.DB {
	var err error
	dsn := "host=localhost user=postgres password=password dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
