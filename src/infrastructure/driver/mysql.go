package driver

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB ...
func InitDB() *gorm.DB {
	dsn := "root:root@tcp(db:3306)/goplayground?parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// NowFunc: setTimeZone(),
	})
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
