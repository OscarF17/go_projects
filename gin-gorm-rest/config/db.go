package config

import (
	"fmt"
	"gin-mod-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conncect() {
	dsn := "postgres://postgres:postgres@localhost:8081/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Esto falló")
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
