// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixFrontgateInfo openpitrix frontgate info
// swagger:model openpitrixFrontgateInfo
type OpenpitrixFrontgateInfo struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this openpitrix frontgate info
func (m *OpenpitrixFrontgateInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixFrontgateInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixFrontgateInfo) UnmarshalBinary(b []byte) error {
	var res OpenpitrixFrontgateInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
