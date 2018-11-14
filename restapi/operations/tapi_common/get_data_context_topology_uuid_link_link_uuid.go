// Code generated by go-swagger; DO NOT EDIT.

package tapi_common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDataContextTopologyUUIDLinkLinkUUIDHandlerFunc turns a function with the right signature into a get data context topology UUID link link UUID handler
type GetDataContextTopologyUUIDLinkLinkUUIDHandlerFunc func(GetDataContextTopologyUUIDLinkLinkUUIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDataContextTopologyUUIDLinkLinkUUIDHandlerFunc) Handle(params GetDataContextTopologyUUIDLinkLinkUUIDParams) middleware.Responder {
	return fn(params)
}

// GetDataContextTopologyUUIDLinkLinkUUIDHandler interface for that can handle valid get data context topology UUID link link UUID params
type GetDataContextTopologyUUIDLinkLinkUUIDHandler interface {
	Handle(GetDataContextTopologyUUIDLinkLinkUUIDParams) middleware.Responder
}

// NewGetDataContextTopologyUUIDLinkLinkUUID creates a new http.Handler for the get data context topology UUID link link UUID operation
func NewGetDataContextTopologyUUIDLinkLinkUUID(ctx *middleware.Context, handler GetDataContextTopologyUUIDLinkLinkUUIDHandler) *GetDataContextTopologyUUIDLinkLinkUUID {
	return &GetDataContextTopologyUUIDLinkLinkUUID{Context: ctx, Handler: handler}
}

/*GetDataContextTopologyUUIDLinkLinkUUID swagger:route GET /data/context/topology={uuid}/link={link-uuid}/ tapi-common getDataContextTopologyUuidLinkLinkUuid

returns tapi.topology.Link

*/
type GetDataContextTopologyUUIDLinkLinkUUID struct {
	Context *middleware.Context
	Handler GetDataContextTopologyUUIDLinkLinkUUIDHandler
}

func (o *GetDataContextTopologyUUIDLinkLinkUUID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDataContextTopologyUUIDLinkLinkUUIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
