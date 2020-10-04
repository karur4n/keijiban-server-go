package models

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func newDB() (database *sql.DB) {
	database, err := sql.Open("mysql", "root:@/go_keijiban?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return
}
