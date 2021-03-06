// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DroneConfdBackendConfig See https://godoc.org/openpitrix.io/libconfd#BackendConfig
// swagger:model droneConfdBackendConfig
type DroneConfdBackendConfig struct {

	// client ca keys
	ClientCaKeys string `json:"client_ca_keys,omitempty"`

	// client cert
	ClientCert string `json:"client_cert,omitempty"`

	// client key
	ClientKey string `json:"client_key,omitempty"`

	// host
	Host []string `json:"host"`

	// password
	Password string `json:"password,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this drone confd backend config
func (m *DroneConfdBackendConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DroneConfdBackendConfig) validateHost(formats strfmt.Registry) error {

	if swag.IsZero(m.Host) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DroneConfdBackendConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DroneConfdBackendConfig) UnmarshalBinary(b []byte) error {
	var res DroneConfdBackendConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
