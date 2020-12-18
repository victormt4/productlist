package productcontroller

import (
	"net/http"
	"productlist/dbwrapper"
	"productlist/model"
	"productlist/repo/productrepo"
	"productlist/services/responseservice"
	"strconv"
)

func Index(r *http.Request) responseservice.Response {

	db := dbwrapper.GetDB()
	defer db.Close()

	repo := productrepo.GetRepo(db)

	products := repo.GetAll()

	return responseservice.Success(responseservice.TemplateResponse{
		TemplateName: "Index",
		TemplateData: products,
	})
}

func Add(r *http.Request) responseservice.Response {

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

		return responseservice.Success(responseservice.DefaultResponseStructure{
			Message: "Product successfully registered", Data: product,
		})

	} else {
		return responseservice.NotAllowed(nil)
	}
}

func Remove(r *http.Request) responseservice.Response {

	if r.Method != "DELETE" {
		return responseservice.NotAllowed(nil)
	} else {

		id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)

		if err != nil {
			return responseservice.UnprocessableEntity("id is invalid")
		} else {

			db := dbwrapper.GetDB()
			defer db.Close()

			repo := productrepo.GetRepo(db)

			product := repo.Find(id)

			if product == nil {
				return responseservice.NotFound("Product not found")
			} else {
				repo.Remove(id)
				return responseservice.Success("Product removed")
			}
		}
	}
}
