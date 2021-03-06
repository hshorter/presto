// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MefCommonDscpEecIDPac mef common dscp eec Id pac
// swagger:model mef.common.DscpEecIdPac
type MefCommonDscpEecIDPac struct {

	// This attribute is a list of DSCP values that maps to the EEC Name.
	DscpValueList []int64 `json:"dscp-value-list"`

	// This attribute specifies the IP version for the DSCP. It can be IPV4, IPV6 or IPV4_AND_IPV6.
	IPVersion MefCommonTypesIPVersion `json:"ip-version,omitempty"`
}

// Validate validates this mef common dscp eec Id pac
func (m *MefCommonDscpEecIDPac) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MefCommonDscpEecIDPac) validateIPVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.IPVersion) { // not required
		return nil
	}

	if err := m.IPVersion.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ip-version")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MefCommonDscpEecIDPac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MefCommonDscpEecIDPac) UnmarshalBinary(b []byte) error {
	var res MefCommonDscpEecIDPac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
