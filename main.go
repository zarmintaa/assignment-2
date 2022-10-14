package main

import (
	"assignment-2/models"
	"assignment-2/routers"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	dbPort   = "5432"
	user     = "postgres"
	password = "root"
	dbname   = "hacktiv8-assigment"
	db       *gorm.DB
	err      error
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err := db.Debug().AutoMigrate(models.Orders{}, models.Items{})
	if err != nil {
		panic(err)
	}
	err = routers.StartServer(db).Run()
	if err != nil {
		panic(err)
	}
}
