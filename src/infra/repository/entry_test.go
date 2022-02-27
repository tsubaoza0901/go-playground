package repository

import (
	"database/sql"
	"fmt"
	"go-playground/m/v1/src/infra/driver"
	"go-playground/m/v1/src/infra/rdb"
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
	defer truncateTable(conn) // テスト実行後、test用DBの全テーブルレコードを削除
}

func truncateTable(conn *sql.DB) {
	names := rdb.AllTableNames()
	for _, name := range names {
		if _, err := conn.Exec(fmt.Sprintf("TRUNCATE TABLE %s", name)); err != nil {
			log.Fatal(err)
		}
	}
}
