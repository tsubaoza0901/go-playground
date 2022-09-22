package driver

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB ...
func InitDB() *gorm.DB {
	dsn := "root:root@tcp(db:3306)/goplayground?parseTime=True&loc=Local"
	return initDB(dsn)
}

// InitTestDB ...
func InitTestDB() *gorm.DB {
	dsn := "root:root@tcp(test-db:3306)/goplaygroundtest?parseTime=True&loc=Local"
	return initDB(dsn)
}

func initDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// NowFunc: setTimeZone(),
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
