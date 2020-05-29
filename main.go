package main

import (
	"database/sql"
	"log"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/BYUSoT/Aizon/job"
)

var (
	dbstring = "dallin@/pml"
	DB, _ = sql.Open("mysql", dbstring)
)

func main() {
	err := DB.Ping()
	if err != nil {
		log.Fatal("E: Can't access database!\n", err)
	}
	job.Setup(DB)
	current := job.Current()
	for _, v := range current {
		fmt.Printf("Job: %+v\n", v)
	}

}


