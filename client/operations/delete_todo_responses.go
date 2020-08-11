// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteTodoReader is a Reader for the DeleteTodo structure.
type DeleteTodoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTodoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTodoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTodoOK creates a DeleteTodoOK with default headers values
func NewDeleteTodoOK() *DeleteTodoOK {
	return &DeleteTodoOK{}
}

/*DeleteTodoOK handles this case with default header values.

Delete a todo
*/
type DeleteTodoOK struct {
	Payload string
}

func (o *DeleteTodoOK) Error() string {
	return fmt.Sprintf("[DELETE /todo/{todoId}][%d] deleteTodoOK  %+v", 200, o.Payload)
}

func (o *DeleteTodoOK) GetPayload() string {
	return o.Payload
}

func (o *DeleteTodoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
