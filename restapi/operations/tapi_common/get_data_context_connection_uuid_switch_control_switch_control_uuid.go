// Code generated by go-swagger; DO NOT EDIT.

package tapi_common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDHandlerFunc turns a function with the right signature into a get data context connection UUID switch control switch control UUID handler
type GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDHandlerFunc func(GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDHandlerFunc) Handle(params GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDParams) middleware.Responder {
	return fn(params)
}

// GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDHandler interface for that can handle valid get data context connection UUID switch control switch control UUID params
type GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDHandler interface {
	Handle(GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDParams) middleware.Responder
}

// NewGetDataContextConnectionUUIDSwitchControlSwitchControlUUID creates a new http.Handler for the get data context connection UUID switch control switch control UUID operation
func NewGetDataContextConnectionUUIDSwitchControlSwitchControlUUID(ctx *middleware.Context, handler GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDHandler) *GetDataContextConnectionUUIDSwitchControlSwitchControlUUID {
	return &GetDataContextConnectionUUIDSwitchControlSwitchControlUUID{Context: ctx, Handler: handler}
}

/*GetDataContextConnectionUUIDSwitchControlSwitchControlUUID swagger:route GET /data/context/connection={uuid}/switch-control={switch-control-uuid}/ tapi-common getDataContextConnectionUuidSwitchControlSwitchControlUuid

returns tapi.connectivity.SwitchControl

*/
type GetDataContextConnectionUUIDSwitchControlSwitchControlUUID struct {
	Context *middleware.Context
	Handler GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDHandler
}

func (o *GetDataContextConnectionUUIDSwitchControlSwitchControlUUID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDataContextConnectionUUIDSwitchControlSwitchControlUUIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
