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

// MefCommonTypesEgressMapType mef common types egress map type
// swagger:model mef.common.types.EgressMapType
type MefCommonTypesEgressMapType string

const (

	// MefCommonTypesEgressMapTypeCNCTAGPCP captures enum value "CN_C_TAG_PCP"
	MefCommonTypesEgressMapTypeCNCTAGPCP MefCommonTypesEgressMapType = "CN_C_TAG_PCP"

	// MefCommonTypesEgressMapTypeCCCTAGPCP captures enum value "CC_C_TAG_PCP"
	MefCommonTypesEgressMapTypeCCCTAGPCP MefCommonTypesEgressMapType = "CC_C_TAG_PCP"

	// MefCommonTypesEgressMapTypeCCCTAGDEI captures enum value "CC_C_TAG_DEI"
	MefCommonTypesEgressMapTypeCCCTAGDEI MefCommonTypesEgressMapType = "CC_C_TAG_DEI"

	// MefCommonTypesEgressMapTypeCNSTAGPCP captures enum value "CN_S_TAG_PCP"
	MefCommonTypesEgressMapTypeCNSTAGPCP MefCommonTypesEgressMapType = "CN_S_TAG_PCP"

	// MefCommonTypesEgressMapTypeCCSTAGPCP captures enum value "CC_S_TAG_PCP"
	MefCommonTypesEgressMapTypeCCSTAGPCP MefCommonTypesEgressMapType = "CC_S_TAG_PCP"

	// MefCommonTypesEgressMapTypeCCSTAGDEI captures enum value "CC_S_TAG_DEI"
	MefCommonTypesEgressMapTypeCCSTAGDEI MefCommonTypesEgressMapType = "CC_S_TAG_DEI"
)

// for schema
var mefCommonTypesEgressMapTypeEnum []interface{}

func init() {
	var res []MefCommonTypesEgressMapType
	if err := json.Unmarshal([]byte(`["CN_C_TAG_PCP","CC_C_TAG_PCP","CC_C_TAG_DEI","CN_S_TAG_PCP","CC_S_TAG_PCP","CC_S_TAG_DEI"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mefCommonTypesEgressMapTypeEnum = append(mefCommonTypesEgressMapTypeEnum, v)
	}
}

func (m MefCommonTypesEgressMapType) validateMefCommonTypesEgressMapTypeEnum(path, location string, value MefCommonTypesEgressMapType) error {
	if err := validate.Enum(path, location, value, mefCommonTypesEgressMapTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this mef common types egress map type
func (m MefCommonTypesEgressMapType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMefCommonTypesEgressMapTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
