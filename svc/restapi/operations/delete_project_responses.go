// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteProjectOKCode is the HTTP code returned for type DeleteProjectOK
const DeleteProjectOKCode int = 200

/*DeleteProjectOK successfully deleted

swagger:response deleteProjectOK
*/
type DeleteProjectOK struct {

	/*success
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewDeleteProjectOK creates DeleteProjectOK with default headers values
func NewDeleteProjectOK() *DeleteProjectOK {

	return &DeleteProjectOK{}
}

// WithPayload adds the payload to the delete project o k response
func (o *DeleteProjectOK) WithPayload(payload string) *DeleteProjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project o k response
func (o *DeleteProjectOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}