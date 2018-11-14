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

// TapiConnectivityConnectionEndPoint tapi connectivity connection end point
// swagger:model tapi.connectivity.ConnectionEndPoint
type TapiConnectivityConnectionEndPoint struct {
	TapiCommonGlobalClass

	TapiCommonOperationalStatePac

	TapiCommonTerminationPac

	// none
	ClientNodeEdgePoint []*TapiTopologyOwnedNodeEdgePointRef `json:"client-node-edge-point"`

	// The orientation of defined flow at the EndPoint.
	ConnectionPortDirection TapiCommonPortDirection `json:"connection-port-direction,omitempty"`

	// Each EP of the FC has a role (e.g., working, protection, protected, symmetric, hub, spoke, leaf, root)  in the context of the FC with respect to the FC function.
	ConnectionPortRole TapiCommonPortRole `json:"connection-port-role,omitempty"`

	// none
	ConnectivityServiceEndPoint string `json:"connectivity-service-end-point,omitempty"`

	// none
	LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

	// none
	ParentNodeEdgePoint []*TapiTopologyOwnedNodeEdgePointRef `json:"parent-node-edge-point"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiConnectivityConnectionEndPoint) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonGlobalClass
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonGlobalClass = aO0

	// AO1
	var aO1 TapiCommonOperationalStatePac
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiCommonOperationalStatePac = aO1

	// AO2
	var aO2 TapiCommonTerminationPac
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiCommonTerminationPac = aO2

	// AO3
	var dataAO3 struct {
		ClientNodeEdgePoint []*TapiTopologyOwnedNodeEdgePointRef `json:"client-node-edge-point"`

		ConnectionPortDirection TapiCommonPortDirection `json:"connection-port-direction,omitempty"`

		ConnectionPortRole TapiCommonPortRole `json:"connection-port-role,omitempty"`

		ConnectivityServiceEndPoint string `json:"connectivity-service-end-point,omitempty"`

		LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

		ParentNodeEdgePoint []*TapiTopologyOwnedNodeEdgePointRef `json:"parent-node-edge-point"`
	}
	if err := swag.ReadJSON(raw, &dataAO3); err != nil {
		return err
	}

	m.ClientNodeEdgePoint = dataAO3.ClientNodeEdgePoint

	m.ConnectionPortDirection = dataAO3.ConnectionPortDirection

	m.ConnectionPortRole = dataAO3.ConnectionPortRole

	m.ConnectivityServiceEndPoint = dataAO3.ConnectivityServiceEndPoint

	m.LayerProtocolName = dataAO3.LayerProtocolName

	m.ParentNodeEdgePoint = dataAO3.ParentNodeEdgePoint

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiConnectivityConnectionEndPoint) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 4)

	aO0, err := swag.WriteJSON(m.TapiCommonGlobalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiCommonOperationalStatePac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.TapiCommonTerminationPac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	var dataAO3 struct {
		ClientNodeEdgePoint []*TapiTopologyOwnedNodeEdgePointRef `json:"client-node-edge-point"`

		ConnectionPortDirection TapiCommonPortDirection `json:"connection-port-direction,omitempty"`

		ConnectionPortRole TapiCommonPortRole `json:"connection-port-role,omitempty"`

		ConnectivityServiceEndPoint string `json:"connectivity-service-end-point,omitempty"`

		LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

		ParentNodeEdgePoint []*TapiTopologyOwnedNodeEdgePointRef `json:"parent-node-edge-point"`
	}

	dataAO3.ClientNodeEdgePoint = m.ClientNodeEdgePoint

	dataAO3.ConnectionPortDirection = m.ConnectionPortDirection

	dataAO3.ConnectionPortRole = m.ConnectionPortRole

	dataAO3.ConnectivityServiceEndPoint = m.ConnectivityServiceEndPoint

	dataAO3.LayerProtocolName = m.LayerProtocolName

	dataAO3.ParentNodeEdgePoint = m.ParentNodeEdgePoint

	jsonDataAO3, errAO3 := swag.WriteJSON(dataAO3)
	if errAO3 != nil {
		return nil, errAO3
	}
	_parts = append(_parts, jsonDataAO3)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi connectivity connection end point
func (m *TapiConnectivityConnectionEndPoint) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonGlobalClass
	if err := m.TapiCommonGlobalClass.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonOperationalStatePac
	if err := m.TapiCommonOperationalStatePac.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonTerminationPac
	if err := m.TapiCommonTerminationPac.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientNodeEdgePoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionPortDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionPortRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLayerProtocolName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentNodeEdgePoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiConnectivityConnectionEndPoint) validateClientNodeEdgePoint(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientNodeEdgePoint) { // not required
		return nil
	}

	for i := 0; i < len(m.ClientNodeEdgePoint); i++ {
		if swag.IsZero(m.ClientNodeEdgePoint[i]) { // not required
			continue
		}

		if m.ClientNodeEdgePoint[i] != nil {
			if err := m.ClientNodeEdgePoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("client-node-edge-point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiConnectivityConnectionEndPoint) validateConnectionPortDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionPortDirection) { // not required
		return nil
	}

	if err := m.ConnectionPortDirection.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("connection-port-direction")
		}
		return err
	}

	return nil
}

func (m *TapiConnectivityConnectionEndPoint) validateConnectionPortRole(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionPortRole) { // not required
		return nil
	}

	if err := m.ConnectionPortRole.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("connection-port-role")
		}
		return err
	}

	return nil
}

func (m *TapiConnectivityConnectionEndPoint) validateLayerProtocolName(formats strfmt.Registry) error {

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

func (m *TapiConnectivityConnectionEndPoint) validateParentNodeEdgePoint(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentNodeEdgePoint) { // not required
		return nil
	}

	for i := 0; i < len(m.ParentNodeEdgePoint); i++ {
		if swag.IsZero(m.ParentNodeEdgePoint[i]) { // not required
			continue
		}

		if m.ParentNodeEdgePoint[i] != nil {
			if err := m.ParentNodeEdgePoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parent-node-edge-point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityConnectionEndPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityConnectionEndPoint) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityConnectionEndPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
