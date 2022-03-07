package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	writer      http.ResponseWriter
}

// crea una respuesta por defecto
func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		writer:      w,
		contentType: "application/json",
	}

}

// func receiver Devuelve un mensaje de error cuando no encuentra información
func (r *Response) NotFound(message string) {
	r.Status = http.StatusNotFound
	r.Message = message
}

// func receiver Envia mensaje Json solicitado
func (r *Response) Send() {
	r.writer.Header().Set("Content-Type", r.contentType)
	r.writer.WriteHeader(r.Status)
	output, _ := json.Marshal(&r)
	fmt.Fprintf(r.writer, string(output))
}

func (r *Response) UnprocessableEntity(message string) {
	r.Status = http.StatusUnprocessableEntity
	r.Message = message
}

// Envia un mensaje cuando no se encuentra información
func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound("Send Not Found")
	response.Send()
}

// Envia la información
func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}

// envia mensaje error cuando no encuentre user
func SendUnprocessableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity("UnprocessableEntity")
	response.Send()
}
