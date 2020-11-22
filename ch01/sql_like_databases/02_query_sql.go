package main

import (
	"database/sql"
	"fmt"
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

	// sql.Rows 값에 대한 pointer 반환되므로 defer를 사용하여 Close
	rows, err := db.Query(`
		SELECT sepal_length as sLength,
			   sepal_width as sWidth,
			   petal_length as pLength,
			   petal_width as pWidth
		FROM iris
		WHERE species =$1`, "Iris-setosa")
	
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Scan을 사용하여 pointer로 반환된 값을 변수에 매핑
	for rows.Next() {
		var (
			sLength float64
			sWidth  float64
			pLength float64
			pWidth  float64
		)

		if err := rows.Scan(&sLength, &sWidth, &pLength, &pWidth); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%.2f, %.2f, %.2f, %.2f\n", sLength, sWidth, pLength, pWidth)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
