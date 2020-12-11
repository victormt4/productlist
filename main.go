package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"productlist/routes"
	"productlist/utils/errorutils"
)

func main() {

	err := godotenv.Load(".env")
	errorutils.ExitOnError(err)

	routes.RegisterRoutes()

	addr := "localhost:9000"

	fmt.Println("Listening on", addr)

	err = http.ListenAndServe(addr, nil)
	errorutils.ExitOnError(err)
}
