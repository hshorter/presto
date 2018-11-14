// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiCommonContextServiceInterfacePoint tapi common context service interface point
// swagger:model tapi.common.context.ServiceInterfacePoint
type TapiCommonContextServiceInterfacePoint struct {
	NrpInterfaceServiceInterfacePointAugmentation1

	TapiCommonServiceInterfacePoint
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiCommonContextServiceInterfacePoint) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 NrpInterfaceServiceInterfacePointAugmentation1
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.NrpInterfaceServiceInterfacePointAugmentation1 = aO0

	// AO1
	var aO1 TapiCommonServiceInterfacePoint
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiCommonServiceInterfacePoint = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiCommonContextServiceInterfacePoint) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.NrpInterfaceServiceInterfacePointAugmentation1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiCommonServiceInterfacePoint)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi common context service interface point
func (m *TapiCommonContextServiceInterfacePoint) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NrpInterfaceServiceInterfacePointAugmentation1
	if err := m.NrpInterfaceServiceInterfacePointAugmentation1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonServiceInterfacePoint
	if err := m.TapiCommonServiceInterfacePoint.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiCommonContextServiceInterfacePoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiCommonContextServiceInterfacePoint) UnmarshalBinary(b []byte) error {
	var res TapiCommonContextServiceInterfacePoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}