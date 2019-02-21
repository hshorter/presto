// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiConnectivityAuditGetconnectivityserviceauditdetailsOutputServiceauditEndPoint tapi connectivity audit getconnectivityserviceauditdetails output serviceaudit end point
// swagger:model tapi.connectivity.audit.getconnectivityserviceauditdetails.output.serviceaudit.EndPoint
type TapiConnectivityAuditGetconnectivityserviceauditdetailsOutputServiceauditEndPoint struct {
	NrpAuditInterfaceEndPointAugmentation1

	TapiConnectivityConnectivityServiceEndPoint
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiConnectivityAuditGetconnectivityserviceauditdetailsOutputServiceauditEndPoint) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 NrpAuditInterfaceEndPointAugmentation1
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.NrpAuditInterfaceEndPointAugmentation1 = aO0

	// AO1
	var aO1 TapiConnectivityConnectivityServiceEndPoint
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiConnectivityConnectivityServiceEndPoint = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiConnectivityAuditGetconnectivityserviceauditdetailsOutputServiceauditEndPoint) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.NrpAuditInterfaceEndPointAugmentation1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiConnectivityConnectivityServiceEndPoint)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi connectivity audit getconnectivityserviceauditdetails output serviceaudit end point
func (m *TapiConnectivityAuditGetconnectivityserviceauditdetailsOutputServiceauditEndPoint) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NrpAuditInterfaceEndPointAugmentation1
	if err := m.NrpAuditInterfaceEndPointAugmentation1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiConnectivityConnectivityServiceEndPoint
	if err := m.TapiConnectivityConnectivityServiceEndPoint.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityAuditGetconnectivityserviceauditdetailsOutputServiceauditEndPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityAuditGetconnectivityserviceauditdetailsOutputServiceauditEndPoint) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityAuditGetconnectivityserviceauditdetailsOutputServiceauditEndPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
