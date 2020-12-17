package productcontroller

import (
	"html/template"
	"net/http"
	"productlist/model"
	"productlist/repo/productrepo"
	"productlist/services/responseservice"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	products := productrepo.GetAll()

	temp.ExecuteTemplate(w, "Index", products)
}

func Add(w http.ResponseWriter, r *http.Request) {

	resService := responseservice.GetResponseServiceWriter(w)

	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")

		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			price = 0.0
		}

		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			quantity = 0
		}

		product := productrepo.Insert(model.Product{
			Name:        name,
			Description: description,
			Price:       price,
			Quantity:    quantity,
		})

		resService.SendSuccess(responseservice.DefaultResponseStructure{
			Message: "Product successfully registered", Data: product,
		})

	} else {
		resService.SendNotAllowed(nil)
	}
}
