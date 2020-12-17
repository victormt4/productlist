package responseservice

import (
	"encoding/json"
	"net/http"
)

type DefaultResponseStructure struct {
	Message string
	Data    interface{}
}

type ResponseConfig struct {
	HttpStatusCode int
}

type ResponseWriter struct {
	writer http.ResponseWriter
}

type Service interface {
	SendSuccess(data interface{})
	SendNotAllowed(data interface{})
	SendServerError(data interface{})
	SendUnprocessableEntity(data interface{})
	SendNotFound(data interface{})
}

func (r ResponseWriter) SendSuccess(data interface{}) {

	data = formatData(data, "Success")

	JsonResponse(r.writer, data)
}

func (r ResponseWriter) SendNotFound(data interface{}) {

	data = formatData(data, "Not found")

	JsonResponseWithConfig(r.writer, data, ResponseConfig{
		HttpStatusCode: 404,
	})
}

func (r ResponseWriter) SendNotAllowed(data interface{}) {

	data = formatData(data, "Method not allowed")

	JsonResponseWithConfig(r.writer, data, ResponseConfig{
		HttpStatusCode: 403,
	})
}

func (r ResponseWriter) SendServerError(data interface{}) {

	data = formatData(data, "Server error")

	JsonResponseWithConfig(r.writer, data, ResponseConfig{
		HttpStatusCode: 500,
	})
}

func (r ResponseWriter) SendUnprocessableEntity(data interface{}) {

	data = formatData(data, "Unprocessable Entity")

	JsonResponseWithConfig(r.writer, data, ResponseConfig{
		HttpStatusCode: 422,
	})
}

func formatData(t interface{}, defaultMessage string) interface{} {
	switch i := t.(type) {
	case nil:
		data := make(map[string]string)
		data["Message"] = defaultMessage
		return data
	case string:
		data := make(map[string]string)
		data["Message"] = i
		return data
	default:
		return t
	}
}

func JsonResponseWithConfig(w http.ResponseWriter, data interface{}, config ResponseConfig) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(config.HttpStatusCode)
	json.NewEncoder(w).Encode(data)
}

func JsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//GetResponseServiceWriter wrap a http.ResponseWriter in a service with methods to print json responses
func GetResponseServiceWriter(w http.ResponseWriter) Service {
	return ResponseWriter{
		writer: w,
	}
}
