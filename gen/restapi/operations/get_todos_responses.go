// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ju-zp/tasker/svc/models"
)

// GetTodosOKCode is the HTTP code returned for type GetTodosOK
const GetTodosOKCode int = 200

/*GetTodosOK Get all todos

swagger:response getTodosOK
*/
type GetTodosOK struct {

	/*array of all the todos
	  In: Body
	*/
	Payload []*models.Todo `json:"body,omitempty"`
}

// NewGetTodosOK creates GetTodosOK with default headers values
func NewGetTodosOK() *GetTodosOK {

	return &GetTodosOK{}
}

// WithPayload adds the payload to the get todos o k response
func (o *GetTodosOK) WithPayload(payload []*models.Todo) *GetTodosOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todos o k response
func (o *GetTodosOK) SetPayload(payload []*models.Todo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodosOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Todo, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTodosBadRequestCode is the HTTP code returned for type GetTodosBadRequest
const GetTodosBadRequestCode int = 400

/*GetTodosBadRequest Incomplete data

swagger:response getTodosBadRequest
*/
type GetTodosBadRequest struct {

	/*Body of the request was Incomplete
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetTodosBadRequest creates GetTodosBadRequest with default headers values
func NewGetTodosBadRequest() *GetTodosBadRequest {

	return &GetTodosBadRequest{}
}

// WithPayload adds the payload to the get todos bad request response
func (o *GetTodosBadRequest) WithPayload(payload string) *GetTodosBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todos bad request response
func (o *GetTodosBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodosBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}