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

// TapiConnectivityUpdateconnectivityserviceOutputService tapi connectivity updateconnectivityservice output service
// swagger:model tapi.connectivity.updateconnectivityservice.output.Service
type TapiConnectivityUpdateconnectivityserviceOutputService struct {
	NrpInterfaceServiceAugmentation1

	TapiCommonAdminStatePac

	TapiCommonGlobalClass

	TapiConnectivityConnectivityConstraint

	TapiConnectivityResilienceConstraint

	TapiConnectivityTopologyConstraint

	// none
	Connection []string `json:"connection"`

	// none
	Direction TapiCommonForwardingDirection `json:"direction,omitempty"`

	// none
	EndPoint []*TapiConnectivityUpdateconnectivityserviceOutputServiceEndPoint `json:"end-point"`

	// none
	LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiConnectivityUpdateconnectivityserviceOutputService) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 NrpInterfaceServiceAugmentation1
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.NrpInterfaceServiceAugmentation1 = aO0

	// AO1
	var aO1 TapiCommonAdminStatePac
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiCommonAdminStatePac = aO1

	// AO2
	var aO2 TapiCommonGlobalClass
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiCommonGlobalClass = aO2

	// AO3
	var aO3 TapiConnectivityConnectivityConstraint
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.TapiConnectivityConnectivityConstraint = aO3

	// AO4
	var aO4 TapiConnectivityResilienceConstraint
	if err := swag.ReadJSON(raw, &aO4); err != nil {
		return err
	}
	m.TapiConnectivityResilienceConstraint = aO4

	// AO5
	var aO5 TapiConnectivityTopologyConstraint
	if err := swag.ReadJSON(raw, &aO5); err != nil {
		return err
	}
	m.TapiConnectivityTopologyConstraint = aO5

	// AO6
	var dataAO6 struct {
		Connection []string `json:"connection"`

		Direction TapiCommonForwardingDirection `json:"direction,omitempty"`

		EndPoint []*TapiConnectivityUpdateconnectivityserviceOutputServiceEndPoint `json:"end-point"`

		LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO6); err != nil {
		return err
	}

	m.Connection = dataAO6.Connection

	m.Direction = dataAO6.Direction

	m.EndPoint = dataAO6.EndPoint

	m.LayerProtocolName = dataAO6.LayerProtocolName

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiConnectivityUpdateconnectivityserviceOutputService) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 7)

	aO0, err := swag.WriteJSON(m.NrpInterfaceServiceAugmentation1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiCommonAdminStatePac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.TapiCommonGlobalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.TapiConnectivityConnectivityConstraint)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	aO4, err := swag.WriteJSON(m.TapiConnectivityResilienceConstraint)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO4)

	aO5, err := swag.WriteJSON(m.TapiConnectivityTopologyConstraint)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO5)

	var dataAO6 struct {
		Connection []string `json:"connection"`

		Direction TapiCommonForwardingDirection `json:"direction,omitempty"`

		EndPoint []*TapiConnectivityUpdateconnectivityserviceOutputServiceEndPoint `json:"end-point"`

		LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`
	}

	dataAO6.Connection = m.Connection

	dataAO6.Direction = m.Direction

	dataAO6.EndPoint = m.EndPoint

	dataAO6.LayerProtocolName = m.LayerProtocolName

	jsonDataAO6, errAO6 := swag.WriteJSON(dataAO6)
	if errAO6 != nil {
		return nil, errAO6
	}
	_parts = append(_parts, jsonDataAO6)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi connectivity updateconnectivityservice output service
func (m *TapiConnectivityUpdateconnectivityserviceOutputService) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NrpInterfaceServiceAugmentation1
	if err := m.NrpInterfaceServiceAugmentation1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonAdminStatePac
	if err := m.TapiCommonAdminStatePac.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonGlobalClass
	if err := m.TapiCommonGlobalClass.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiConnectivityConnectivityConstraint
	if err := m.TapiConnectivityConnectivityConstraint.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiConnectivityResilienceConstraint
	if err := m.TapiConnectivityResilienceConstraint.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiConnectivityTopologyConstraint
	if err := m.TapiConnectivityTopologyConstraint.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLayerProtocolName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiConnectivityUpdateconnectivityserviceOutputService) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	if err := m.Direction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("direction")
		}
		return err
	}

	return nil
}

func (m *TapiConnectivityUpdateconnectivityserviceOutputService) validateEndPoint(formats strfmt.Registry) error {

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

func (m *TapiConnectivityUpdateconnectivityserviceOutputService) validateLayerProtocolName(formats strfmt.Registry) error {

	if swag.IsZero(m.LayerProtocolName) { // not required
		return nil
	}

	if err := m.LayerProtocolName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("layer-protocol-name")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityUpdateconnectivityserviceOutputService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityUpdateconnectivityserviceOutputService) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityUpdateconnectivityserviceOutputService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}