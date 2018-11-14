// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NrpInterfaceNrpCarrierEthConnectivityResource nrp interface nrp carrier eth connectivity resource
// swagger:model nrp.interface.NrpCarrierEthConnectivityResource
type NrpInterfaceNrpCarrierEthConnectivityResource struct {

	// When the value is conditionally, the conditions that determine whether a Data Service Frame is delivered or discarded MUST be specified. Conditions can be described in the name-value pair list.
	BroadcastFrameDelivery *NrmConnectivityFrameDeliveryCondition `json:"broadcast-frame-delivery,omitempty"`

	// This attribute can be used to preserve the value of the CE-VLAN DEI field in VLAN Tagged Service Frames across a Connectivity Service.
	CeVlanDeiPreservation *bool `json:"ce-vlan-dei-preservation,omitempty"`

	// This attribute describes a relationship between the format of the VLAN ID and related fields values of the frame at one Interface and the format and VLAN ID and related fields values of the corresponding frame at another Interface.
	//                         Used the MEF 7.3 OVC type (PRESERVE/STRIP/RETAIN) as it depends on EVC/OVC decomposition performed by SOFs.
	//
	CeVlanIDPreservation MefCommonTypesVlanIDPreservation `json:"ce-vlan-id-preservation,omitempty"`

	// This attribute can be used to preserve the value of the CE-VLAN PCP field in VLAN Tagged Service Frames across a Connectivity Service.
	CeVlanPcpPreservation *bool `json:"ce-vlan-pcp-preservation,omitempty"`

	// This attribute denotes the maximum frame size in bytes requested for the connectivity service.
	MaxFrameSize int64 `json:"max-frame-size,omitempty"`

	// When the value is conditionally, the conditions that determine whether a Data Service Frame is delivered or discarded MUST be specified. Conditions can be described in the name-value pair list.
	MulticastFrameDelivery *NrmConnectivityFrameDeliveryCondition `json:"multicast-frame-delivery,omitempty"`

	// A value of “true” indicates that the End Points can be added/removed during Connectivity Service lifecycle.
	MultipointCapable *bool `json:"multipoint-capable,omitempty"`

	// This attribute describes a relationship between the S-VLAN DEI value of a frame at one ENNI and the S-VLAN DEI of the corresponding frame at another ENNI supported by the Operator CEN where each ENNI has a Connectivity Service End Point that is associated by the Connectivity Service.
	SVlanDeiPreservation *bool `json:"s-vlan-dei-preservation,omitempty"`

	// This attribute describes a relationship between the S-VLAN PCP value of a frame at one ENNI and the S-VLAN PCP of the corresponding frame at another ENNI supported by the Operator CEN where each ENNI has a Connectivity Service End Point that is associated by the Connectivity Service.
	SVlanPcpPreservation *bool `json:"s-vlan-pcp-preservation,omitempty"`

	// When the value is conditionally, the conditions that determine whether a Data Service Frame is delivered or discarded MUST be specified. Conditions can be described in the name-value pair list.
	UnicastFrameDelivery *NrmConnectivityFrameDeliveryCondition `json:"unicast-frame-delivery,omitempty"`
}

// Validate validates this nrp interface nrp carrier eth connectivity resource
func (m *NrpInterfaceNrpCarrierEthConnectivityResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBroadcastFrameDelivery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCeVlanIDPreservation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMulticastFrameDelivery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnicastFrameDelivery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityResource) validateBroadcastFrameDelivery(formats strfmt.Registry) error {

	if swag.IsZero(m.BroadcastFrameDelivery) { // not required
		return nil
	}

	if m.BroadcastFrameDelivery != nil {
		if err := m.BroadcastFrameDelivery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("broadcast-frame-delivery")
			}
			return err
		}
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityResource) validateCeVlanIDPreservation(formats strfmt.Registry) error {

	if swag.IsZero(m.CeVlanIDPreservation) { // not required
		return nil
	}

	if err := m.CeVlanIDPreservation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ce-vlan-id-preservation")
		}
		return err
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityResource) validateMulticastFrameDelivery(formats strfmt.Registry) error {

	if swag.IsZero(m.MulticastFrameDelivery) { // not required
		return nil
	}

	if m.MulticastFrameDelivery != nil {
		if err := m.MulticastFrameDelivery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("multicast-frame-delivery")
			}
			return err
		}
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityResource) validateUnicastFrameDelivery(formats strfmt.Registry) error {

	if swag.IsZero(m.UnicastFrameDelivery) { // not required
		return nil
	}

	if m.UnicastFrameDelivery != nil {
		if err := m.UnicastFrameDelivery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unicast-frame-delivery")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NrpInterfaceNrpCarrierEthConnectivityResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NrpInterfaceNrpCarrierEthConnectivityResource) UnmarshalBinary(b []byte) error {
	var res NrpInterfaceNrpCarrierEthConnectivityResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
