package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"productlist/repo/productrepo"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)

	addr := "localhost:9000"

	fmt.Println("Listening on", addr)

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		fmt.Println("Error on starting server")
		fmt.Println(err)
		os.Exit(1)
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	products := productrepo.GetAll()

	temp.ExecuteTemplate(w, "Index", products)
}
