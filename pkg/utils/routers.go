package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ResponseParametizer struct {
	Headers HeaderSetters
	Body    []byte
	Status  int
}

type HeaderSetters map[string]string

func CreateParametizer() ResponseParametizer {
	parametizer := ResponseParametizer{
		Headers: HeaderSetters{},
		Status:  http.StatusOK,
	}

	return parametizer
}

func (parametizer *ResponseParametizer) SetBody(body []byte) {
	parametizer.Body = body
}

func (parametizer *ResponseParametizer) SetStatus(status int) {
	parametizer.Status = status
}

func (headers HeaderSetters) SetHeader(key string, value string) {
	headers[key] = value
}

func (parametizer ResponseParametizer) Respond(writer http.ResponseWriter) {
	for key, value := range parametizer.Headers {
		writer.Header().Add(key, value)
	}

	writer.WriteHeader(parametizer.Status)

	writer.Write(parametizer.Body)
}

func ParseBody(request *http.Request, expectedModel interface{}) {
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(body), expectedModel)

	if err != nil {
		panic(err)
	}
}
