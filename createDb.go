package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
  fmt.Print("a");
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS log (
    timestamp timestamp,
    username text,
    operation text CHECK (operation IN ('upload', 'download')),
    size integer
  )`)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

    fmt.Println(row)
		_, err = db.Exec("INSERT INTO log (timestamp, username, operation, size) VALUES (?, ?, ?, ?)", row[0], row[1], row[2], row[3])
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("CSV data successfully loaded into the SQLite database.")
}

