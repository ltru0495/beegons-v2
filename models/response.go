package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Content     interface{} `json:"content"`
	Message     string      `json:"message"`
	contentType string
	writer      http.ResponseWriter
}

type ApiData struct {
	Id        string           `json:"id"`
	Type      string           `json:"type`
	Parameter string           `json:"parameter"`
	Data      []CygnusDocument `json:"data"`
}

func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{Status: http.StatusOK, writer: w, contentType: "application/json"}
}

func (this *Response) NotFound() {
	this.Status = http.StatusNotFound
	this.Message = "Resource Not Found!"
}

func SendUnprocessableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}

func (this *Response) UnprocessableEntity() {
	this.Status = http.StatusUnprocessableEntity
	this.Message = "UnprocessableEntity"
}

func SendNoContent(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NoContent()
	response.Send()
}

func (this *Response) NoContent() {
	this.Status = http.StatusNoContent
	this.Message = "No Content!"

}

func SendData(w http.ResponseWriter, content interface{}) {
	response := CreateDefaultResponse(w)
	response.Content = content
	response.Send()
}

func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

func (this *Response) Send() {
	this.writer.Header().Set("Content-Type", this.contentType)
	this.writer.WriteHeader(this.Status)

	output, _ := json.Marshal(&this)
	fmt.Fprintf(this.writer, string(output))
}
