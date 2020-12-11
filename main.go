package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"html/template"
	"net/http"
	"productlist/repo/productrepo"
	"productlist/utils/errorutils"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	err := godotenv.Load(".env")
	errorutils.ExitOnError(err)

	http.HandleFunc("/", index)

	addr := "localhost:9000"

	fmt.Println("Listening on", addr)

	err = http.ListenAndServe(addr, nil)
	errorutils.ExitOnError(err)
}

func index(w http.ResponseWriter, r *http.Request) {

	products := productrepo.GetAll()

	temp.ExecuteTemplate(w, "Index", products)
}
