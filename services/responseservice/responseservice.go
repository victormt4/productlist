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
}

func (r ResponseWriter) SendSuccess(data interface{}) {
	JsonResponse(r.writer, data)
}

func (r ResponseWriter) SendNotAllowed(data interface{}) {

	if data == nil {
		data = DefaultResponseStructure{
			Message: "Method not allowed",
			Data:    nil,
		}
	}

	JsonResponseWithConfig(r.writer, data, ResponseConfig{
		HttpStatusCode: 403,
	})
}

func (r ResponseWriter) SendServerError(data interface{}) {
	JsonResponseWithConfig(r.writer, data, ResponseConfig{
		HttpStatusCode: 500,
	})
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
