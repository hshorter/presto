// Code generated by go-swagger; DO NOT EDIT.

package tapi_path_computation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsOptimizeP2PPathHandlerFunc turns a function with the right signature into a post operations optimize p2 p path handler
type PostOperationsOptimizeP2PPathHandlerFunc func(PostOperationsOptimizeP2PPathParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsOptimizeP2PPathHandlerFunc) Handle(params PostOperationsOptimizeP2PPathParams) middleware.Responder {
	return fn(params)
}

// PostOperationsOptimizeP2PPathHandler interface for that can handle valid post operations optimize p2 p path params
type PostOperationsOptimizeP2PPathHandler interface {
	Handle(PostOperationsOptimizeP2PPathParams) middleware.Responder
}

// NewPostOperationsOptimizeP2PPath creates a new http.Handler for the post operations optimize p2 p path operation
func NewPostOperationsOptimizeP2PPath(ctx *middleware.Context, handler PostOperationsOptimizeP2PPathHandler) *PostOperationsOptimizeP2PPath {
	return &PostOperationsOptimizeP2PPath{Context: ctx, Handler: handler}
}

/*PostOperationsOptimizeP2PPath swagger:route POST /operations/optimize-p-2-p-path/ tapi-path-computation postOperationsOptimizeP2PPath

PostOperationsOptimizeP2PPath post operations optimize p2 p path API

*/
type PostOperationsOptimizeP2PPath struct {
	Context *middleware.Context
	Handler PostOperationsOptimizeP2PPathHandler
}

func (o *PostOperationsOptimizeP2PPath) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsOptimizeP2PPathParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
