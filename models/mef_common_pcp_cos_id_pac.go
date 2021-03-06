// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MefCommonPcpCosIDPac mef common pcp cos Id pac
// swagger:model mef.common.PcpCosIdPac
type MefCommonPcpCosIDPac struct {

	// This attribute is a list of PCP values that map to the CoS Name.
	PcpValueList []int64 `json:"pcp-value-list"`
}

// Validate validates this mef common pcp cos Id pac
func (m *MefCommonPcpCosIDPac) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MefCommonPcpCosIDPac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MefCommonPcpCosIDPac) UnmarshalBinary(b []byte) error {
	var res MefCommonPcpCosIDPac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
