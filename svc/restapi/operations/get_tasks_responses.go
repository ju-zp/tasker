// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ju-zp/tasker/svc/models"
)

// GetTasksOKCode is the HTTP code returned for type GetTasksOK
const GetTasksOKCode int = 200

/*GetTasksOK Get all the tasks

swagger:response getTasksOK
*/
type GetTasksOK struct {

	/*array of all the tasks
	  In: Body
	*/
	Payload []*models.Task `json:"body,omitempty"`
}

// NewGetTasksOK creates GetTasksOK with default headers values
func NewGetTasksOK() *GetTasksOK {

	return &GetTasksOK{}
}

// WithPayload adds the payload to the get tasks o k response
func (o *GetTasksOK) WithPayload(payload []*models.Task) *GetTasksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tasks o k response
func (o *GetTasksOK) SetPayload(payload []*models.Task) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTasksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Task, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTasksBadRequestCode is the HTTP code returned for type GetTasksBadRequest
const GetTasksBadRequestCode int = 400

/*GetTasksBadRequest Incomplete data

swagger:response getTasksBadRequest
*/
type GetTasksBadRequest struct {

	/*Body of the request was incomplete
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetTasksBadRequest creates GetTasksBadRequest with default headers values
func NewGetTasksBadRequest() *GetTasksBadRequest {

	return &GetTasksBadRequest{}
}

// WithPayload adds the payload to the get tasks bad request response
func (o *GetTasksBadRequest) WithPayload(payload string) *GetTasksBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tasks bad request response
func (o *GetTasksBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTasksBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
