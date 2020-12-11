package productcontroller

import (
	"html/template"
	"net/http"
	"productlist/repo/productrepo"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	products := productrepo.GetAll()

	temp.ExecuteTemplate(w, "Index", products)
}
