package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	pgURL := os.Getenv("PGURL")
	if pgURL == "" {
		log.Fatal("PGURL empty")
	}

	db, err := sql.Open("postres", pgURL)
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
