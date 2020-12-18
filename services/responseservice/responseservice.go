package responseservice

type DefaultResponseStructure struct {
	Message string
	Data    interface{}
}

type TemplateResponse struct {
	TemplateName string
	TemplateData interface{}
}

type ResponseConfig struct {
	HttpStatusCode int
}

type Response interface {
	GetData() interface{}
	GetConfig() ResponseConfig
}

type response struct {
	data   interface{}
	config ResponseConfig
}

func (r response) GetData() interface{} {
	return r.data
}

func (r response) GetConfig() ResponseConfig {
	return r.config
}

func Success(data interface{}) Response {

	data = formatData(data, "Success")

	return response{
		data:   data,
		config: ResponseConfig{HttpStatusCode: 200},
	}
}

func NotFound(data interface{}) Response {

	data = formatData(data, "Not found")

	return response{
		data:   data,
		config: ResponseConfig{HttpStatusCode: 404},
	}
}

func NotAllowed(data interface{}) Response {

	data = formatData(data, "Method not allowed")

	return response{
		data:   data,
		config: ResponseConfig{HttpStatusCode: 405},
	}
}

func ServerError(data interface{}) Response {

	data = formatData(data, "Server error")

	return response{
		data:   data,
		config: ResponseConfig{HttpStatusCode: 500},
	}
}

func UnprocessableEntity(data interface{}) Response {

	data = formatData(data, "Unprocessable Entity")

	return response{
		data:   data,
		config: ResponseConfig{HttpStatusCode: 422},
	}
}

func formatData(t interface{}, defaultMessage string) interface{} {
	switch i := t.(type) {
	case nil:
		data := map[string]string{
			"Message": defaultMessage,
		}
		return data
	case string:
		data := map[string]string{
			"Message": i,
		}
		return data
	default:
		return t
	}
}
