package routes

import (
	"net/http"
	"productlist/controllers/productcontroller"
)

func RegisterRoutes() {
	http.HandleFunc("/", productcontroller.Index)
}
