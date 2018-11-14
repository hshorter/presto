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

// TapiConnectivityOwnedNodeEdgePointAugmentation1 tapi connectivity owned node edge point augmentation1
// swagger:model tapi.connectivity.OwnedNodeEdgePointAugmentation1
type TapiConnectivityOwnedNodeEdgePointAugmentation1 struct {

	// none
	ConnectionEndPoint []*TapiConnectivityConnectionEndPoint `json:"connection-end-point"`
}

// Validate validates this tapi connectivity owned node edge point augmentation1
func (m *TapiConnectivityOwnedNodeEdgePointAugmentation1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionEndPoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiConnectivityOwnedNodeEdgePointAugmentation1) validateConnectionEndPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionEndPoint) { // not required
		return nil
	}

	for i := 0; i < len(m.ConnectionEndPoint); i++ {
		if swag.IsZero(m.ConnectionEndPoint[i]) { // not required
			continue
		}

		if m.ConnectionEndPoint[i] != nil {
			if err := m.ConnectionEndPoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connection-end-point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityOwnedNodeEdgePointAugmentation1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityOwnedNodeEdgePointAugmentation1) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityOwnedNodeEdgePointAugmentation1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}