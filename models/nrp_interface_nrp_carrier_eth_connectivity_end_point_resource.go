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

// NrpInterfaceNrpCarrierEthConnectivityEndPointResource nrp interface nrp carrier eth connectivity end point resource
// swagger:model nrp.interface.NrpCarrierEthConnectivityEndPointResource
type NrpInterfaceNrpCarrierEthConnectivityEndPointResource struct {

	// List of one or more C-VLAN ID values. A C-Tagged Frame, whose C-VLAN ID value matches an entry in this attribute, maps to the Connectivity Service End Point.
	//                         It is possible to specify whether untagged and priority tagged frames are included in the mapping.
	//                         Type=LIST: all listed VLAN IDs. Type=EXCEPT: all VLAN IDs except the listed ones. Type=ALL, all VLAN IDs, hence vlanId list is not applicable.
	CeVlanIDListAndUntag *NrmConnectivityVlanIDListAndUntag `json:"ce-vlan-id-list-and-untag,omitempty"`

	// This attribute represents the relationship between the Connectivity Service End Point and a Color Identifier.
	ColorIdentifier *MefCommonColorIdentifier `json:"color-identifier,omitempty"`

	// This attribute represents the relationship between the Connectivity Service End Point and the Class of Service Identifier(s).
	CosIdentifierList []*MefCommonCosIdentifier `json:"cos-identifier-list"`

	// The Class of Service (CoS) is used to specify ingress Bandwidth Profiles. The CoS Mapping Type is one of SEP (Service End Point) based, PCP based or DSCP based.
	CosMappingType MefCommonTypesCosOrEecMappingType `json:"cos-mapping-type,omitempty"`

	// This attribute represents the relationship between the Connectivity Service End Point and the Egress Equivalence Class Identifier(s).
	EecIdentifierList []*MefCommonEecIdentifier `json:"eec-identifier-list"`

	// The Egress Equivalence Class (EEC) is used to specify Egress Bandwidth Profiles. The EEC Mapping Type is one of SEP (Service End Point) based, PCP based or DSCP based.
	//                         When the list of EEC Identifier is empty,  this attribute shall be unset. Otherwise it shall be set.
	EecMappingType MefCommonTypesCosOrEecMappingType `json:"eec-mapping-type,omitempty"`

	// This attribute denotes the relationship between a Connectivity Service End Point and the bandwidth profile flow. It describes egress policing on all egress EI Frames mapped to a given End Point.
	EgressBwpFlow *MefCommonBwpFlow `json:"egress-bwp-flow,omitempty"`

	// This attribute represents the relationship between the End Point and the Egress Map(s).  This attribute is a set of mappings that determine the content of the S-Tag or C-Tag of an egress EI Frame.
	//
	EgressMapList []*MefCommonEgressMap `json:"egress-map-list"`

	//  This attribute denotes the relationship between a Connectivity Service End Point and the bandwidth profile flow. It describes ingress policing on all ingress EI Frames mapped to a given End Point.
	IngressBwpFlow *MefCommonBwpFlow `json:"ingress-bwp-flow,omitempty"`

	// This attribute specifies the subset of the Bridge Reserved Addresses that are filtered (i.e. L2CP Frames with this destination address are Peered or Discarded but not Passed) at a L2CP Decision Point.
	L2cpAddressSet MefCommonTypesL2cpAddressSet `json:"l2cp-address-set,omitempty"`

	// This attribute applies only to End Points with Trunk End Point Role. It identifies the S-VLAN ID of frames mapped to either a Leaf End Point or a Trunk End Point (via the Leaf S-VLAN ID value) of the Connectivity Service.
	LeafSvlanID *MefCommonTypesVlanID `json:"leaf-svlan-id,omitempty"`

	// This attribute applies only to End Points with Trunk End Point Role. It identifies the S-VLAN ID of frames mapped to either a Root End Point or a Trunk End Point (via the Root S-VLAN ID value) of the Connectivity Service.
	RootSvlanID *MefCommonTypesVlanID `json:"root-svlan-id,omitempty"`

	// List of one or more S-VLAN ID values. An S-Tagged Frame, whose S-VLAN ID value matches an entry in this attribute, maps to the Connectivity Service End Point.
	//                         Type=LIST: all listed VLAN IDs. Type=EXCEPT: all VLAN IDs except the listed ones. Type=ALL, all VLAN IDs, hence vlanId list is not applicable.
	//
	SVlanIDList *MefCommonTypesVlanIDListing `json:"s-vlan-id-list,omitempty"`

	// This attribute limits the number of source MAC Addresses that can be used in ingress EI Frames mapped to the Connectivity Service End Point of all types over a time interval. When not present, the number of source MAC addresses is unlimited.
	//                         Two independent parameters control the behavior of this attribute: N : A positive integer and t : A time interval.
	//                         This attribute operates by maintaining a list of maximum length N of source MAC addresses which are aged-out of the list if not seen in a time interval t. If an ingress Service Frame arrives with a new source MAC address when the list is full, the Service Frame is discarded.
	SourceMacAddressLimit *MefCommonTypesSourceMacAddressLimit `json:"source-mac-address-limit,omitempty"`
}

// Validate validates this nrp interface nrp carrier eth connectivity end point resource
func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCeVlanIDListAndUntag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColorIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCosIdentifierList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCosMappingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEecIdentifierList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEecMappingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEgressBwpFlow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEgressMapList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngressBwpFlow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateL2cpAddressSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeafSvlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootSvlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSVlanIDList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceMacAddressLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateCeVlanIDListAndUntag(formats strfmt.Registry) error {

	if swag.IsZero(m.CeVlanIDListAndUntag) { // not required
		return nil
	}

	if m.CeVlanIDListAndUntag != nil {
		if err := m.CeVlanIDListAndUntag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ce-vlan-id-list-and-untag")
			}
			return err
		}
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateColorIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.ColorIdentifier) { // not required
		return nil
	}

	if m.ColorIdentifier != nil {
		if err := m.ColorIdentifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("color-identifier")
			}
			return err
		}
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateCosIdentifierList(formats strfmt.Registry) error {

	if swag.IsZero(m.CosIdentifierList) { // not required
		return nil
	}

	for i := 0; i < len(m.CosIdentifierList); i++ {
		if swag.IsZero(m.CosIdentifierList[i]) { // not required
			continue
		}

		if m.CosIdentifierList[i] != nil {
			if err := m.CosIdentifierList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cos-identifier-list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateCosMappingType(formats strfmt.Registry) error {

	if swag.IsZero(m.CosMappingType) { // not required
		return nil
	}

	if err := m.CosMappingType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cos-mapping-type")
		}
		return err
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateEecIdentifierList(formats strfmt.Registry) error {

	if swag.IsZero(m.EecIdentifierList) { // not required
		return nil
	}

	for i := 0; i < len(m.EecIdentifierList); i++ {
		if swag.IsZero(m.EecIdentifierList[i]) { // not required
			continue
		}

		if m.EecIdentifierList[i] != nil {
			if err := m.EecIdentifierList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eec-identifier-list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateEecMappingType(formats strfmt.Registry) error {

	if swag.IsZero(m.EecMappingType) { // not required
		return nil
	}

	if err := m.EecMappingType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("eec-mapping-type")
		}
		return err
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateEgressBwpFlow(formats strfmt.Registry) error {

	if swag.IsZero(m.EgressBwpFlow) { // not required
		return nil
	}

	if m.EgressBwpFlow != nil {
		if err := m.EgressBwpFlow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("egress-bwp-flow")
			}
			return err
		}
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateEgressMapList(formats strfmt.Registry) error {

	if swag.IsZero(m.EgressMapList) { // not required
		return nil
	}

	for i := 0; i < len(m.EgressMapList); i++ {
		if swag.IsZero(m.EgressMapList[i]) { // not required
			continue
		}

		if m.EgressMapList[i] != nil {
			if err := m.EgressMapList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("egress-map-list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateIngressBwpFlow(formats strfmt.Registry) error {

	if swag.IsZero(m.IngressBwpFlow) { // not required
		return nil
	}

	if m.IngressBwpFlow != nil {
		if err := m.IngressBwpFlow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ingress-bwp-flow")
			}
			return err
		}
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateL2cpAddressSet(formats strfmt.Registry) error {

	if swag.IsZero(m.L2cpAddressSet) { // not required
		return nil
	}

	if err := m.L2cpAddressSet.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("l2cp-address-set")
		}
		return err
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateLeafSvlanID(formats strfmt.Registry) error {

	if swag.IsZero(m.LeafSvlanID) { // not required
		return nil
	}

	if m.LeafSvlanID != nil {
		if err := m.LeafSvlanID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("leaf-svlan-id")
			}
			return err
		}
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateRootSvlanID(formats strfmt.Registry) error {

	if swag.IsZero(m.RootSvlanID) { // not required
		return nil
	}

	if m.RootSvlanID != nil {
		if err := m.RootSvlanID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("root-svlan-id")
			}
			return err
		}
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateSVlanIDList(formats strfmt.Registry) error {

	if swag.IsZero(m.SVlanIDList) { // not required
		return nil
	}

	if m.SVlanIDList != nil {
		if err := m.SVlanIDList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s-vlan-id-list")
			}
			return err
		}
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) validateSourceMacAddressLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceMacAddressLimit) { // not required
		return nil
	}

	if m.SourceMacAddressLimit != nil {
		if err := m.SourceMacAddressLimit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source-mac-address-limit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NrpInterfaceNrpCarrierEthConnectivityEndPointResource) UnmarshalBinary(b []byte) error {
	var res NrpInterfaceNrpCarrierEthConnectivityEndPointResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}