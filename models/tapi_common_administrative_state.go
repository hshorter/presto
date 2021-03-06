// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TapiCommonAdministrativeState tapi common administrative state
// swagger:model tapi.common.AdministrativeState
type TapiCommonAdministrativeState string

const (

	// TapiCommonAdministrativeStateLOCKED captures enum value "LOCKED"
	TapiCommonAdministrativeStateLOCKED TapiCommonAdministrativeState = "LOCKED"

	// TapiCommonAdministrativeStateUNLOCKED captures enum value "UNLOCKED"
	TapiCommonAdministrativeStateUNLOCKED TapiCommonAdministrativeState = "UNLOCKED"
)

// for schema
var tapiCommonAdministrativeStateEnum []interface{}

func init() {
	var res []TapiCommonAdministrativeState
	if err := json.Unmarshal([]byte(`["LOCKED","UNLOCKED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiCommonAdministrativeStateEnum = append(tapiCommonAdministrativeStateEnum, v)
	}
}

func (m TapiCommonAdministrativeState) validateTapiCommonAdministrativeStateEnum(path, location string, value TapiCommonAdministrativeState) error {
	if err := validate.Enum(path, location, value, tapiCommonAdministrativeStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi common administrative state
func (m TapiCommonAdministrativeState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiCommonAdministrativeStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
