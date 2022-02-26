package driver

import (
	"log"

	"github.com/DATA-DOG/go-txdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB ...
func InitDB() *gorm.DB {
	dsn := "root:root@tcp(db:3306)/goplayground?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// NowFunc: setTimeZone(),
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitTestDB(name string) *gorm.DB {
	dsn := "root:root@tcp(db:3306)/goplayground_test?parseTime=True&loc=Local"
	txdb.Register(name, "mysql", dsn)
	dialector := mysql.New(mysql.Config{
		DriverName: name,
		DSN:        dsn,
	})
	db, err := gorm.Open(dialector, &gorm.Config{
		// NowFunc: setTimeZone(),
	})
	if err != nil {
		log.Println(err)
		panic("faild to connect database")
	}
	return db
}

// このやり方だとカラムにwith timezoneの設定を入れないとJSTにならないようなので、一旦timezoneはmain.goでグローバルに設定
// const location = "Asia/Tokyo"

// // setTimeZone ...
// func setTimeZone() func() time.Time {
// 	loc, err := time.LoadLocation(location)
// 	if err != nil {
// 		loc = time.FixedZone(location, 9*60*60)
// 	}
// 	return func() time.Time {
// 		return time.Now().In(loc)
// 	}
// }
