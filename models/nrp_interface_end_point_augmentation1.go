// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NrpInterfaceEndPointAugmentation1 nrp interface end point augmentation1
// swagger:model nrp.interface.EndPointAugmentation1
type NrpInterfaceEndPointAugmentation1 struct {

	// none
	NrpCarrierEthConnectivityEndPointResource *NrpInterfaceNrpCarrierEthConnectivityEndPointResource `json:"nrp-carrier-eth-connectivity-end-point-resource,omitempty"`
}

// Validate validates this nrp interface end point augmentation1
func (m *NrpInterfaceEndPointAugmentation1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNrpCarrierEthConnectivityEndPointResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NrpInterfaceEndPointAugmentation1) validateNrpCarrierEthConnectivityEndPointResource(formats strfmt.Registry) error {

	if swag.IsZero(m.NrpCarrierEthConnectivityEndPointResource) { // not required
		return nil
	}

	if m.NrpCarrierEthConnectivityEndPointResource != nil {
		if err := m.NrpCarrierEthConnectivityEndPointResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nrp-carrier-eth-connectivity-end-point-resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NrpInterfaceEndPointAugmentation1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NrpInterfaceEndPointAugmentation1) UnmarshalBinary(b []byte) error {
	var res NrpInterfaceEndPointAugmentation1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}