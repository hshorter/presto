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

// TapiTopologyContextAugmentation3 tapi topology context augmentation3
// swagger:model tapi.topology.ContextAugmentation3
type TapiTopologyContextAugmentation3 struct {

	// none
	NwTopologyService *TapiTopologyNetworkTopologyService `json:"nw-topology-service,omitempty"`

	// none
	Topology []*TapiTopologyTopology `json:"topology"`
}

// Validate validates this tapi topology context augmentation3
func (m *TapiTopologyContextAugmentation3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNwTopologyService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopology(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiTopologyContextAugmentation3) validateNwTopologyService(formats strfmt.Registry) error {

	if swag.IsZero(m.NwTopologyService) { // not required
		return nil
	}

	if m.NwTopologyService != nil {
		if err := m.NwTopologyService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nw-topology-service")
			}
			return err
		}
	}

	return nil
}

func (m *TapiTopologyContextAugmentation3) validateTopology(formats strfmt.Registry) error {

	if swag.IsZero(m.Topology) { // not required
		return nil
	}

	for i := 0; i < len(m.Topology); i++ {
		if swag.IsZero(m.Topology[i]) { // not required
			continue
		}

		if m.Topology[i] != nil {
			if err := m.Topology[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topology" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiTopologyContextAugmentation3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiTopologyContextAugmentation3) UnmarshalBinary(b []byte) error {
	var res TapiTopologyContextAugmentation3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}