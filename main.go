package main

import (
	"go-playground/m/v1/src/infra/driver"
	"go-playground/m/v1/src/infra/repository"
	"log"
)

func main() {
	db := driver.InitDB()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	user := repository.User{
		Name: "山田 太郎",
		Age:  30,
	}

	if err := user.CreateUser(db, user); err != nil {
		log.Fatal(err)
	}
}
