package routes

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"productlist/controllers/productcontroller"
	"productlist/services/responseservice"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func RegisterRoutes(mux *http.ServeMux) {

	mux.Handle("/", globalHandler(productcontroller.Index))
	mux.Handle("/add", globalHandler(productcontroller.Add))
	mux.Handle("/remove", globalHandler(productcontroller.Remove))
}

func globalHandler(f func(r *http.Request) responseservice.Response) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			r := recover()
			if r != nil {

				var err error

				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("unknown error")
				}

				errorRes := responseservice.ServerError(err.Error())
				writer.Header().Set("Content-Type", "application/jsonresponseservice")
				writer.WriteHeader(errorRes.GetConfig().HttpStatusCode)
				json.NewEncoder(writer).Encode(errorRes.GetData())

			}
		}()

		res := f(request)

		switch resModel := res.GetData().(type) {
		case responseservice.TemplateResponse:
			temp.ExecuteTemplate(writer, resModel.TemplateName, resModel.TemplateData)
		default:
			writer.Header().Set("Content-Type", "application/jsonresponseservice")
			writer.WriteHeader(res.GetConfig().HttpStatusCode)
			json.NewEncoder(writer).Encode(resModel)
		}
	}
}
