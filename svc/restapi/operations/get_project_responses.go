// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetProjectOKCode is the HTTP code returned for type GetProjectOK
const GetProjectOKCode int = 200

/*GetProjectOK successful request

swagger:response getProjectOK
*/
type GetProjectOK struct {

	/*
	  In: Body
	*/
	Payload *GetProjectOKBody `json:"body,omitempty"`
}

// NewGetProjectOK creates GetProjectOK with default headers values
func NewGetProjectOK() *GetProjectOK {

	return &GetProjectOK{}
}

// WithPayload adds the payload to the get project o k response
func (o *GetProjectOK) WithPayload(payload *GetProjectOKBody) *GetProjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project o k response
func (o *GetProjectOK) SetPayload(payload *GetProjectOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
