// Code generated by go-swagger; DO NOT EDIT.

package tapi_common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDParams creates a new GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDParams object
// no default values defined in spec.
func NewGetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDParams() GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDParams {

	return GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDParams{}
}

// GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDParams contains all the bound params for the get data context topology UUID node node UUID node rule group node rule group UUID operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUID
type GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Id of node-rule-group
	  Required: true
	  In: path
	*/
	NodeRuleGroupUUID string
	/*Id of node
	  Required: true
	  In: path
	*/
	NodeUUID string
	/*Id of topology
	  Required: true
	  In: path
	*/
	UUID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDParams() beforehand.
func (o *GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rNodeRuleGroupUUID, rhkNodeRuleGroupUUID, _ := route.Params.GetOK("node-rule-group-uuid")
	if err := o.bindNodeRuleGroupUUID(rNodeRuleGroupUUID, rhkNodeRuleGroupUUID, route.Formats); err != nil {
		res = append(res, err)
	}

	rNodeUUID, rhkNodeUUID, _ := route.Params.GetOK("node-uuid")
	if err := o.bindNodeUUID(rNodeUUID, rhkNodeUUID, route.Formats); err != nil {
		res = append(res, err)
	}

	rUUID, rhkUUID, _ := route.Params.GetOK("uuid")
	if err := o.bindUUID(rUUID, rhkUUID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindNodeRuleGroupUUID binds and validates parameter NodeRuleGroupUUID from path.
func (o *GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDParams) bindNodeRuleGroupUUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.NodeRuleGroupUUID = raw

	return nil
}

// bindNodeUUID binds and validates parameter NodeUUID from path.
func (o *GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDParams) bindNodeUUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.NodeUUID = raw

	return nil
}

// bindUUID binds and validates parameter UUID from path.
func (o *GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDParams) bindUUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.UUID = raw

	return nil
}
