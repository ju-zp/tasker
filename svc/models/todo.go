package models

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"github.com/jinzhu/gorm"
)

// Todo todo
//
// swagger:model todo
type Todo struct {
	gorm.Model

	// done
	// Required: true
	Done *bool `json:"done"`

	// task it is associated with
	// Required: true
	TaskID *int64 `json:"taskId"`

	// the content of the todo
	// Required: true
	Todo *string `json:"todo"`
}

// Validate validates this todo
func (m *Todo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTodo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Todo) validateDone(formats strfmt.Registry) error {

	if err := validate.Required("done", "body", m.Done); err != nil {
		return err
	}

	return nil
}

func (m *Todo) validateTaskID(formats strfmt.Registry) error {

	if err := validate.Required("taskId", "body", m.TaskID); err != nil {
		return err
	}

	return nil
}

func (m *Todo) validateTodo(formats strfmt.Registry) error {

	if err := validate.Required("todo", "body", m.Todo); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Todo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Todo) UnmarshalBinary(b []byte) error {
	var res Todo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
