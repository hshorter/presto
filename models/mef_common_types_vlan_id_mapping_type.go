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

// MefCommonTypesVlanIDMappingType mef common types vlan Id mapping type
// swagger:model mef.common.types.VlanIdMappingType
type MefCommonTypesVlanIDMappingType string

const (

	// MefCommonTypesVlanIDMappingTypeALL captures enum value "ALL"
	MefCommonTypesVlanIDMappingTypeALL MefCommonTypesVlanIDMappingType = "ALL"

	// MefCommonTypesVlanIDMappingTypeEXCEPT captures enum value "EXCEPT"
	MefCommonTypesVlanIDMappingTypeEXCEPT MefCommonTypesVlanIDMappingType = "EXCEPT"

	// MefCommonTypesVlanIDMappingTypeLIST captures enum value "LIST"
	MefCommonTypesVlanIDMappingTypeLIST MefCommonTypesVlanIDMappingType = "LIST"
)

// for schema
var mefCommonTypesVlanIdMappingTypeEnum []interface{}

func init() {
	var res []MefCommonTypesVlanIDMappingType
	if err := json.Unmarshal([]byte(`["ALL","EXCEPT","LIST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mefCommonTypesVlanIdMappingTypeEnum = append(mefCommonTypesVlanIdMappingTypeEnum, v)
	}
}

func (m MefCommonTypesVlanIDMappingType) validateMefCommonTypesVlanIDMappingTypeEnum(path, location string, value MefCommonTypesVlanIDMappingType) error {
	if err := validate.Enum(path, location, value, mefCommonTypesVlanIdMappingTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this mef common types vlan Id mapping type
func (m MefCommonTypesVlanIDMappingType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMefCommonTypesVlanIDMappingTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
