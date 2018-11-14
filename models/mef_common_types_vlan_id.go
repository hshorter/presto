// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MefCommonTypesVlanID mef common types vlan Id
// swagger:model mef.common.types.VlanId
type MefCommonTypesVlanID struct {

	// This is the Vlan ID value.
	VlanID int64 `json:"vlan-id,omitempty"`
}

// Validate validates this mef common types vlan Id
func (m *MefCommonTypesVlanID) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MefCommonTypesVlanID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MefCommonTypesVlanID) UnmarshalBinary(b []byte) error {
	var res MefCommonTypesVlanID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
