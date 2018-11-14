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

// NrmConnectivityFrameDeliveryCondition nrm connectivity frame delivery condition
// swagger:model nrm.connectivity.FrameDeliveryCondition
type NrmConnectivityFrameDeliveryCondition struct {

	// Data Service Frame disposition.
	Action NrmConnectivityDeliveryActionType `json:"action,omitempty"`

	// When the value is conditionally, the conditions that determine whether a Data Service Frame is delivered or discarded MUST be specified. Conditions can be described in the name-value pair list, where name is used to represent the condition name or identifier, and value is used to represent the condition details associated to that name/identifier. Interoperability requires further standardization.
	DeliveryCondition []*TapiCommonNameAndValue `json:"delivery-condition"`
}

// Validate validates this nrm connectivity frame delivery condition
func (m *NrmConnectivityFrameDeliveryCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryCondition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NrmConnectivityFrameDeliveryCondition) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := m.Action.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action")
		}
		return err
	}

	return nil
}

func (m *NrmConnectivityFrameDeliveryCondition) validateDeliveryCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.DeliveryCondition) { // not required
		return nil
	}

	for i := 0; i < len(m.DeliveryCondition); i++ {
		if swag.IsZero(m.DeliveryCondition[i]) { // not required
			continue
		}

		if m.DeliveryCondition[i] != nil {
			if err := m.DeliveryCondition[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delivery-condition" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NrmConnectivityFrameDeliveryCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NrmConnectivityFrameDeliveryCondition) UnmarshalBinary(b []byte) error {
	var res NrmConnectivityFrameDeliveryCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
