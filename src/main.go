package main

import (
	"embed"
	"log"
	"time"

	backend "go-playground/m/v1/framework"
	"go-playground/m/v1/infrastructure/driver"

	"github.com/pressly/goose/v3"
)

const location = "Asia/Tokyo"

// SetTimeZone ...
func SetTimeZone() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

//go:embed infrastructure/migrations/*.sql
var embedMigrations embed.FS

func main() {
	SetTimeZone()

	db := driver.InitDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	goose.SetBaseFS(embedMigrations)

	backend.NewApp(db).Start()
}
