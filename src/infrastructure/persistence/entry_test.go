package persistence_test

import (
	"database/sql"
	"go-playground/m/v1/infrastructure/driver"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/pressly/goose/v3"
	"gorm.io/gorm"
)

var testDB *gorm.DB
var sqlDB *sql.DB
var targetMigrationFileCount int

func TestMain(m *testing.M) {
	beforeAll(m)

	exitCode := m.Run()

	afterAll(m)

	os.Exit(exitCode)
}

func beforeAll(m *testing.M) {
	var err error

	testDB = driver.InitTestDB()
	sqlDB, err = testDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	count, err := countMigrationfileList()
	if err != nil {
		log.Fatal(err)
	}

	targetMigrationFileCount = count

	log.Println("goose up for start test")
	if err := goose.Up(sqlDB, "../../../infrastructure/migrations"); err != nil {
		panic(err)
	}
}

func afterAll(m *testing.M) {
	defer sqlDB.Close()

	count, err := countMigrationfileList()
	if err != nil {
		log.Fatal(err)
	}

	// NOTE: goose up時からmigration file数が変更になっている場合は手動でdownが必要な可能性があるためログを出力
	if targetMigrationFileCount != count {
		log.Println("The number of migration file don't match.")
	}

	log.Println("goose down for close test")
	for count > 0 {
		if err := goose.Down(sqlDB, "../../../infrastructure/migrations"); err != nil {
			panic(err)
		}
		count -= 1
	}
}

func countMigrationfileList() (count int, err error) {
	var files []string
	if err := filepath.Walk("../../../infrastructure/migrations", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return len(files), nil
}
