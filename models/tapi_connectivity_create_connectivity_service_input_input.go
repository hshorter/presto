// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiConnectivityCreateConnectivityServiceInputInput tapi connectivity create connectivity service input input
// swagger:model tapi.connectivity.CreateConnectivityServiceInputInput
type TapiConnectivityCreateConnectivityServiceInputInput struct {
	NrpInterfaceCreateConnectivityServiceInputAugmentation1

	// none
	ConnConstraint *TapiConnectivityConnectivityConstraint `json:"conn-constraint,omitempty"`

	// none
	EndPoint []*TapiConnectivityCreateconnectivityserviceInputEndPoint `json:"end-point"`

	// none
	ResilienceConstraint []*TapiConnectivityResilienceConstraint `json:"resilience-constraint"`

	// none
	State string `json:"state,omitempty"`

	// none
	TopoConstraint *TapiConnectivityTopologyConstraint `json:"topo-constraint,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiConnectivityCreateConnectivityServiceInputInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 NrpInterfaceCreateConnectivityServiceInputAugmentation1
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.NrpInterfaceCreateConnectivityServiceInputAugmentation1 = aO0

	// AO1
	var dataAO1 struct {
		ConnConstraint *TapiConnectivityConnectivityConstraint `json:"conn-constraint,omitempty"`

		EndPoint []*TapiConnectivityCreateconnectivityserviceInputEndPoint `json:"end-point"`

		ResilienceConstraint []*TapiConnectivityResilienceConstraint `json:"resilience-constraint"`

		State string `json:"state,omitempty"`

		TopoConstraint *TapiConnectivityTopologyConstraint `json:"topo-constraint,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ConnConstraint = dataAO1.ConnConstraint

	m.EndPoint = dataAO1.EndPoint

	m.ResilienceConstraint = dataAO1.ResilienceConstraint

	m.State = dataAO1.State

	m.TopoConstraint = dataAO1.TopoConstraint

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiConnectivityCreateConnectivityServiceInputInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.NrpInterfaceCreateConnectivityServiceInputAugmentation1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ConnConstraint *TapiConnectivityConnectivityConstraint `json:"conn-constraint,omitempty"`

		EndPoint []*TapiConnectivityCreateconnectivityserviceInputEndPoint `json:"end-point"`

		ResilienceConstraint []*TapiConnectivityResilienceConstraint `json:"resilience-constraint"`

		State string `json:"state,omitempty"`

		TopoConstraint *TapiConnectivityTopologyConstraint `json:"topo-constraint,omitempty"`
	}

	dataAO1.ConnConstraint = m.ConnConstraint

	dataAO1.EndPoint = m.EndPoint

	dataAO1.ResilienceConstraint = m.ResilienceConstraint

	dataAO1.State = m.State

	dataAO1.TopoConstraint = m.TopoConstraint

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi connectivity create connectivity service input input
func (m *TapiConnectivityCreateConnectivityServiceInputInput) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NrpInterfaceCreateConnectivityServiceInputAugmentation1
	if err := m.NrpInterfaceCreateConnectivityServiceInputAugmentation1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnConstraint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResilienceConstraint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopoConstraint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiConnectivityCreateConnectivityServiceInputInput) validateConnConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnConstraint) { // not required
		return nil
	}

	if m.ConnConstraint != nil {
		if err := m.ConnConstraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conn-constraint")
			}
			return err
		}
	}

	return nil
}

func (m *TapiConnectivityCreateConnectivityServiceInputInput) validateEndPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.EndPoint) { // not required
		return nil
	}

	for i := 0; i < len(m.EndPoint); i++ {
		if swag.IsZero(m.EndPoint[i]) { // not required
			continue
		}

		if m.EndPoint[i] != nil {
			if err := m.EndPoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("end-point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiConnectivityCreateConnectivityServiceInputInput) validateResilienceConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.ResilienceConstraint) { // not required
		return nil
	}

	for i := 0; i < len(m.ResilienceConstraint); i++ {
		if swag.IsZero(m.ResilienceConstraint[i]) { // not required
			continue
		}

		if m.ResilienceConstraint[i] != nil {
			if err := m.ResilienceConstraint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resilience-constraint" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiConnectivityCreateConnectivityServiceInputInput) validateTopoConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.TopoConstraint) { // not required
		return nil
	}

	if m.TopoConstraint != nil {
		if err := m.TopoConstraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topo-constraint")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityCreateConnectivityServiceInputInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityCreateConnectivityServiceInputInput) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityCreateConnectivityServiceInputInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
