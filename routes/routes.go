package routes

import (
	"net/http"
	"productlist/controllers/productcontroller"
	"productlist/middleware"
)

func RegisterRoutes(mux *http.ServeMux) {

	mux.Handle("/", middleware.RecoverMiddleware(http.HandlerFunc(productcontroller.Index)))
	mux.Handle("/add", middleware.RecoverMiddleware(http.HandlerFunc(productcontroller.Add)))
	mux.Handle("/remove", middleware.RecoverMiddleware(http.HandlerFunc(productcontroller.Remove)))
}
