package main

import (
	"database/sql"
	"log"
	"os"
	// go get github.com/lib/pq
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

	// method Ping: DB 연결 확인
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
