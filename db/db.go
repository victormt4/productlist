package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	connectionStr := "user= dbname= password=host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		panic(err.Error())
	}

	return db
}
