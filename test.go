package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "test:test@tcp(mysql:3306)/"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var result string
	err = db.QueryRow("SELECT NOW()").Scan(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current time:", result)
}
