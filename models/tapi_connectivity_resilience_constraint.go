// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiConnectivityResilienceConstraint tapi connectivity resilience constraint
// swagger:model tapi.connectivity.ResilienceConstraint
type TapiConnectivityResilienceConstraint struct {

	// This attribute indicates the time, in milliseconds, between declaration of signal degrade or signal fail, and the initialization of the protection switching algorithm.
	HoldOffTime int32 `json:"hold-off-time,omitempty"`

	// Is operating such that switching at both ends of each flow acorss the FC is coordinated at both ingress and egress ends.
	IsCoordinatedSwitchingBothEnds *bool `json:"is-coordinated-switching-both-ends,omitempty"`

	// Temporarily prevents any switch action to be taken and, as such, freezes the current state.
	//                     Until the freeze is cleared, additional near-end external commands are rejected and fault condition changes and received APS messages are ignored.
	//                     All administrative controls of any aspect of protection are rejected.
	IsFrozen *bool `json:"is-frozen,omitempty"`

	// The resource is configured to temporarily not be available for use in the protection scheme(s) it is part of.
	//                     This overrides all other protection control states including forced.
	//                     If the item is locked out then it cannot be used under any circumstances.
	//                     Note: Only relevant when part of a protection scheme.
	IsLockOut *bool `json:"is-lock-out,omitempty"`

	// Indicate which layer this resilience parameters package configured for.
	LayerProtocol TapiCommonLayerProtocolName `json:"layer-protocol,omitempty"`

	// Used to limit the maximum swtich times. When work fault disappears , and traffic return to the original work path, switch counter reset.
	MaxSwitchTimes int32 `json:"max-switch-times,omitempty"`

	// none
	ResilienceType *TapiTopologyResilienceType `json:"resilience-type,omitempty"`

	//  The coordination mechanism between multi-layers.
	RestorationCoordinateType TapiConnectivityCoordinateType `json:"restoration-coordinate-type,omitempty"`

	// none
	RestorePriority int32 `json:"restore-priority,omitempty"`

	// Indcates whether the protection scheme is revertive or non-revertive.
	ReversionMode TapiConnectivityReversionMode `json:"reversion-mode,omitempty"`

	// If the protection system is revertive, this attribute specifies the time, in minutes, to wait after a fault clears on a higher priority (preferred) resource before reverting to the preferred resource.
	WaitToRevertTime *int32 `json:"wait-to-revert-time,omitempty"`
}

// Validate validates this tapi connectivity resilience constraint
func (m *TapiConnectivityResilienceConstraint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLayerProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResilienceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestorationCoordinateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReversionMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiConnectivityResilienceConstraint) validateLayerProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.LayerProtocol) { // not required
		return nil
	}

	if err := m.LayerProtocol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("layer-protocol")
		}
		return err
	}

	return nil
}

func (m *TapiConnectivityResilienceConstraint) validateResilienceType(formats strfmt.Registry) error {

	if swag.IsZero(m.ResilienceType) { // not required
		return nil
	}

	if m.ResilienceType != nil {
		if err := m.ResilienceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resilience-type")
			}
			return err
		}
	}

	return nil
}

func (m *TapiConnectivityResilienceConstraint) validateRestorationCoordinateType(formats strfmt.Registry) error {

	if swag.IsZero(m.RestorationCoordinateType) { // not required
		return nil
	}

	if err := m.RestorationCoordinateType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("restoration-coordinate-type")
		}
		return err
	}

	return nil
}

func (m *TapiConnectivityResilienceConstraint) validateReversionMode(formats strfmt.Registry) error {

	if swag.IsZero(m.ReversionMode) { // not required
		return nil
	}

	if err := m.ReversionMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reversion-mode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityResilienceConstraint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityResilienceConstraint) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityResilienceConstraint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
