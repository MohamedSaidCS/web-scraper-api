package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=admin dbname=go sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = DB.Ping()

	if err != nil {
		panic(err)
	}

	createArticlesTable()

}

func createArticlesTable() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS articles (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			link TEXT NOT NULL UNIQUE,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)

	if err != nil {
		panic(err)
	}
}
