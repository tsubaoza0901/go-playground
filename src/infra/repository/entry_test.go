package repository

import (
	"database/sql"
	"go-playground/m/v1/src/infra/driver"
	"log"
	"os"
	"testing"
	"time"

	"gorm.io/gorm"
)

var testTimeDate = time.Date(2021, 1, 20, 10, 10, 00, 0, time.Local)

func TestMain(m *testing.M) {
	beforeAll(m)

	exitCode := m.Run()

	afterAll(m)

	os.Exit(exitCode)
}

var db *gorm.DB

func beforeAll(m *testing.M) {
	db = driver.InitTestDB("test")
}

func afterAll(m *testing.M) {
	conn, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	defer truncateTable(conn)
}

func truncateTable(conn *sql.DB) {
	// TODO: 現状だとusersテーブルしかデータ削除できないため、他のテーブルも作成する場合はやり方の検討が必要
	if _, err := conn.Exec("TRUNCATE TABLE users"); err != nil {
		log.Fatal(err)
	}
}
