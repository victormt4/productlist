package productcontroller

import (
	"html/template"
	"net/http"
	"productlist/dbwrapper"
	"productlist/model"
	"productlist/repo/productrepo"
	"productlist/services/responseservice"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	db := dbwrapper.GetDB()
	defer db.Close()

	repo := productrepo.GetRepo(db)

	products := repo.GetAll()

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

		db := dbwrapper.GetDB()
		defer db.Close()

		repo := productrepo.GetRepo(db)

		product := repo.Insert(model.Product{
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

func Remove(w http.ResponseWriter, r *http.Request) {

	res := responseservice.GetResponseServiceWriter(w)

	if r.Method != "DELETE" {
		res.SendNotAllowed(nil)
	} else {

		id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)

		if err != nil {
			res.SendUnprocessableEntity("id is invalid")
		} else {

			db := dbwrapper.GetDB()
			defer db.Close()

			repo := productrepo.GetRepo(db)

			product := repo.Find(id)

			if product == nil {
				res.SendNotFound("Product not found")
			} else {
				repo.Remove(id)
				res.SendSuccess("Product removed")
			}
		}
	}
}
