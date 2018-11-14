// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiCommonGetServiceInterfacePointDetailsOutput tapi common get service interface point details output
// swagger:model tapi.common.GetServiceInterfacePointDetailsOutput
type TapiCommonGetServiceInterfacePointDetailsOutput struct {

	// output
	Output *TapiCommonGetServiceInterfacePointDetailsOutputOutput `json:"output,omitempty"`
}

// Validate validates this tapi common get service interface point details output
func (m *TapiCommonGetServiceInterfacePointDetailsOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiCommonGetServiceInterfacePointDetailsOutput) validateOutput(formats strfmt.Registry) error {

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
func (m *TapiCommonGetServiceInterfacePointDetailsOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiCommonGetServiceInterfacePointDetailsOutput) UnmarshalBinary(b []byte) error {
	var res TapiCommonGetServiceInterfacePointDetailsOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiCommonGetServiceInterfacePointDetailsOutputOutput tapi common get service interface point details output output
// swagger:model TapiCommonGetServiceInterfacePointDetailsOutputOutput
type TapiCommonGetServiceInterfacePointDetailsOutputOutput struct {

	// none
	Sip *TapiCommonGetserviceinterfacepointdetailsOutputSip `json:"sip,omitempty"`
}

// Validate validates this tapi common get service interface point details output output
func (m *TapiCommonGetServiceInterfacePointDetailsOutputOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSip(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiCommonGetServiceInterfacePointDetailsOutputOutput) validateSip(formats strfmt.Registry) error {

	if swag.IsZero(m.Sip) { // not required
		return nil
	}

	if m.Sip != nil {
		if err := m.Sip.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output" + "." + "sip")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiCommonGetServiceInterfacePointDetailsOutputOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiCommonGetServiceInterfacePointDetailsOutputOutput) UnmarshalBinary(b []byte) error {
	var res TapiCommonGetServiceInterfacePointDetailsOutputOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
