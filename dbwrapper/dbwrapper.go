package dbwrapper

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"productlist/utils/errorutils"
)

func GetDB() *sql.DB {

	connectionStr := "user=%s dbname=%s password=%s host=%s sslmode=disable"
	connectionStr = fmt.Sprintf(connectionStr, os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"))

	db, err := sql.Open("postgres", connectionStr)
	errorutils.PanicOnError(err)

	return db
}
