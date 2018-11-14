// Code generated by go-swagger; DO NOT EDIT.

package tapi_common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDataContextNwTopologyServiceHandlerFunc turns a function with the right signature into a get data context nw topology service handler
type GetDataContextNwTopologyServiceHandlerFunc func(GetDataContextNwTopologyServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDataContextNwTopologyServiceHandlerFunc) Handle(params GetDataContextNwTopologyServiceParams) middleware.Responder {
	return fn(params)
}

// GetDataContextNwTopologyServiceHandler interface for that can handle valid get data context nw topology service params
type GetDataContextNwTopologyServiceHandler interface {
	Handle(GetDataContextNwTopologyServiceParams) middleware.Responder
}

// NewGetDataContextNwTopologyService creates a new http.Handler for the get data context nw topology service operation
func NewGetDataContextNwTopologyService(ctx *middleware.Context, handler GetDataContextNwTopologyServiceHandler) *GetDataContextNwTopologyService {
	return &GetDataContextNwTopologyService{Context: ctx, Handler: handler}
}

/*GetDataContextNwTopologyService swagger:route GET /data/context/nw-topology-service/ tapi-common getDataContextNwTopologyService

returns tapi.topology.NetworkTopologyService

*/
type GetDataContextNwTopologyService struct {
	Context *middleware.Context
	Handler GetDataContextNwTopologyServiceHandler
}

func (o *GetDataContextNwTopologyService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDataContextNwTopologyServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
