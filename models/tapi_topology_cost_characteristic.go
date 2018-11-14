// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiTopologyCostCharacteristic tapi topology cost characteristic
// swagger:model tapi.topology.CostCharacteristic
type TapiTopologyCostCharacteristic struct {

	// The cost may vary based upon some properties of the TopologicalEntity. The rules for the variation are conveyed by the costAlgorithm.
	CostAlgorithm string `json:"cost-algorithm,omitempty"`

	// The cost characteristic will related to some aspect of the TopologicalEntity (e.g. $ cost, routing weight). This aspect will be conveyed by the costName.
	CostName string `json:"cost-name,omitempty"`

	// The specific cost.
	CostValue string `json:"cost-value,omitempty"`
}

// Validate validates this tapi topology cost characteristic
func (m *TapiTopologyCostCharacteristic) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiTopologyCostCharacteristic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiTopologyCostCharacteristic) UnmarshalBinary(b []byte) error {
	var res TapiTopologyCostCharacteristic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}