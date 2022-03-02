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

func GetDefaultResponse(w http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		writer:      w,
		contentType: "application/json",
	}

}

func (r *Response) NotFound(message string) {
	r.Status = http.StatusNotFound
	r.Data = nil
	r.Message = message
}

func (r *Response) Send() {
	r.writer.Header().Set("Content-Type", r.contentType)
	r.writer.WriteHeader(r.Status)

	output, _ := json.Marshal(&r)
	fmt.Fprintf(r.writer, string(output))
}
