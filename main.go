package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./calendar.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM plants")
	checkErr(err)

	w.Header().Set("Content-Type", "text/html")
	json, err := json.Marshal(rows)
	checkErr(err)

	w.Write(json)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
