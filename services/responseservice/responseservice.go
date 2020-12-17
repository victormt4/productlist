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

func JsonResponseWithConfig(w http.ResponseWriter, data interface{}, config ResponseConfig) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(config.HttpStatusCode)
	json.NewEncoder(w).Encode(data)
}

func JsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetResponseServiceWriter(w http.ResponseWriter) ResponseWriter {
	return ResponseWriter{
		writer: w,
	}
}