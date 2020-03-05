package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver"
)

var (
	dbstring = "/pml"
	DB, _ = sql.Open("mysql", dbstring)
)

func init() {
	err := DB.Ping()
	if err != nil {
		log.Fatal("E: Can't access database!\n")
	}
}