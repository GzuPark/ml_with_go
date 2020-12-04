package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	 _ "github.com/lib/pq"
)

func main() {
	pgHost := os.Getenv("PGHOST")
	pgDB := os.Getenv("PGDATABASE")
	pgUser := os.Getenv("PGUSER")
	pgPW := os.Getenv("PGPASSWORD")

	var connection string = fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable", 
		pgHost, pgDB, pgUser, pgPW)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// method Ping: DB 연결 확인
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success to connect 'iris' database.")
}
