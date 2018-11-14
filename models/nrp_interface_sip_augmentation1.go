// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NrpInterfaceSipAugmentation1 nrp interface sip augmentation1
// swagger:model nrp.interface.SipAugmentation1
type NrpInterfaceSipAugmentation1 struct {

	// none
	NrpCarrierEthEnniNResource *NrpInterfaceNrpCarrierEthEnniNResource `json:"nrp-carrier-eth-enni-n-resource,omitempty"`

	// none
	NrpCarrierEthInniNResource *NrpInterfaceNrpCarrierEthInniNResource `json:"nrp-carrier-eth-inni-n-resource,omitempty"`

	// none
	NrpCarrierEthUniNResource *NrpInterfaceNrpCarrierEthUniNResource `json:"nrp-carrier-eth-uni-n-resource,omitempty"`
}

// Validate validates this nrp interface sip augmentation1
func (m *NrpInterfaceSipAugmentation1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNrpCarrierEthEnniNResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNrpCarrierEthInniNResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNrpCarrierEthUniNResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NrpInterfaceSipAugmentation1) validateNrpCarrierEthEnniNResource(formats strfmt.Registry) error {

	if swag.IsZero(m.NrpCarrierEthEnniNResource) { // not required
		return nil
	}

	if m.NrpCarrierEthEnniNResource != nil {
		if err := m.NrpCarrierEthEnniNResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nrp-carrier-eth-enni-n-resource")
			}
			return err
		}
	}

	return nil
}

func (m *NrpInterfaceSipAugmentation1) validateNrpCarrierEthInniNResource(formats strfmt.Registry) error {

	if swag.IsZero(m.NrpCarrierEthInniNResource) { // not required
		return nil
	}

	if m.NrpCarrierEthInniNResource != nil {
		if err := m.NrpCarrierEthInniNResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nrp-carrier-eth-inni-n-resource")
			}
			return err
		}
	}

	return nil
}

func (m *NrpInterfaceSipAugmentation1) validateNrpCarrierEthUniNResource(formats strfmt.Registry) error {

	if swag.IsZero(m.NrpCarrierEthUniNResource) { // not required
		return nil
	}

	if m.NrpCarrierEthUniNResource != nil {
		if err := m.NrpCarrierEthUniNResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nrp-carrier-eth-uni-n-resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NrpInterfaceSipAugmentation1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NrpInterfaceSipAugmentation1) UnmarshalBinary(b []byte) error {
	var res NrpInterfaceSipAugmentation1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}