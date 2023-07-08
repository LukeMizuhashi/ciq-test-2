package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Query struct {
	Label string `json:"label"`
	Query string `json:"query"`
}

func main() {
	file, err := ioutil.ReadFile("queries.json")
	if err != nil {
		log.Fatal(err)
	}

	var queries []Query
	err = json.Unmarshal(file, &queries)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	for _, q := range queries {
		var result int
		err = db.QueryRow(q.Query).Scan(&result)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf(q.Label, result)
	}
}

