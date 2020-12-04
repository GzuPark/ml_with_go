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

	res, err := db.Exec("UPDATE iris SET species = 'setosa' WHERE species = 'Iris-setosa'")
	if err != nil {
		log.Fatal(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("affected = %d\n", rowCount)
}
