// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiConnectivityUpdateConnectivityServiceOutput tapi connectivity update connectivity service output
// swagger:model tapi.connectivity.UpdateConnectivityServiceOutput
type TapiConnectivityUpdateConnectivityServiceOutput struct {

	// output
	Output *TapiConnectivityUpdateConnectivityServiceOutputOutput `json:"output,omitempty"`
}

// Validate validates this tapi connectivity update connectivity service output
func (m *TapiConnectivityUpdateConnectivityServiceOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiConnectivityUpdateConnectivityServiceOutput) validateOutput(formats strfmt.Registry) error {

	if swag.IsZero(m.Output) { // not required
		return nil
	}

	if m.Output != nil {
		if err := m.Output.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityUpdateConnectivityServiceOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityUpdateConnectivityServiceOutput) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityUpdateConnectivityServiceOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiConnectivityUpdateConnectivityServiceOutputOutput tapi connectivity update connectivity service output output
// swagger:model TapiConnectivityUpdateConnectivityServiceOutputOutput
type TapiConnectivityUpdateConnectivityServiceOutputOutput struct {

	// none
	Service *TapiConnectivityUpdateconnectivityserviceOutputService `json:"service,omitempty"`
}

// Validate validates this tapi connectivity update connectivity service output output
func (m *TapiConnectivityUpdateConnectivityServiceOutputOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiConnectivityUpdateConnectivityServiceOutputOutput) validateService(formats strfmt.Registry) error {

	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output" + "." + "service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityUpdateConnectivityServiceOutputOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityUpdateConnectivityServiceOutputOutput) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityUpdateConnectivityServiceOutputOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
