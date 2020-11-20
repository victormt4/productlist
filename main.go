package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
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
	temp.ExecuteTemplate(w, "Index", nil)
}
