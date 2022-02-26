package main

import (
	"go-playground/m/v1/src/infra/driver"
	"go-playground/m/v1/src/infra/repository"
	"log"
	"time"

	"gorm.io/gorm"
)

func exec(db *gorm.DB) {
	// user := repository.User{
	// 	Name: "武田 太郎",
	// 	Age:  30,
	// }

	// if err := repository.CreateUser(db, user); err != nil {
	// 	log.Fatal(err)
	// }

	user := repository.User{
		Model: gorm.Model{
			ID: 3,
		},
		Name: "武田 太郎",
		Age:  20,
	}

	if err := repository.UpdateUser(db, user); err != nil {
		log.Fatal(err)
	}

	// if err := repository.TruncateUsersTable(db); err != nil {
	// 	log.Fatal(err)
	// }
}

const location = "Asia/Tokyo"

// SetTimeZone ...
func SetTimeZone() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	SetTimeZone()

	db := driver.InitDB()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	exec(db)
}
