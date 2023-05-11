package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var DSN = "host=localhost user=postgres password=123456 dbname=cubicacion_web port=5432"

func DbConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error connecting with the database: " + err.Error())
		return
	}
	log.Println("Db connected!!!")
}
