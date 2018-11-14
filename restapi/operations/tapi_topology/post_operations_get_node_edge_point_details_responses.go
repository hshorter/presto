// Code generated by go-swagger; DO NOT EDIT.

package tapi_topology

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/damianoneill/presto/models"
)

// PostOperationsGetNodeEdgePointDetailsOKCode is the HTTP code returned for type PostOperationsGetNodeEdgePointDetailsOK
const PostOperationsGetNodeEdgePointDetailsOKCode int = 200

/*PostOperationsGetNodeEdgePointDetailsOK Correct response

swagger:response postOperationsGetNodeEdgePointDetailsOK
*/
type PostOperationsGetNodeEdgePointDetailsOK struct {

	/*
	  In: Body
	*/
	Payload *models.TapiTopologyGetNodeEdgePointDetailsOutput `json:"body,omitempty"`
}

// NewPostOperationsGetNodeEdgePointDetailsOK creates PostOperationsGetNodeEdgePointDetailsOK with default headers values
func NewPostOperationsGetNodeEdgePointDetailsOK() *PostOperationsGetNodeEdgePointDetailsOK {

	return &PostOperationsGetNodeEdgePointDetailsOK{}
}

// WithPayload adds the payload to the post operations get node edge point details o k response
func (o *PostOperationsGetNodeEdgePointDetailsOK) WithPayload(payload *models.TapiTopologyGetNodeEdgePointDetailsOutput) *PostOperationsGetNodeEdgePointDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post operations get node edge point details o k response
func (o *PostOperationsGetNodeEdgePointDetailsOK) SetPayload(payload *models.TapiTopologyGetNodeEdgePointDetailsOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOperationsGetNodeEdgePointDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOperationsGetNodeEdgePointDetailsCreatedCode is the HTTP code returned for type PostOperationsGetNodeEdgePointDetailsCreated
const PostOperationsGetNodeEdgePointDetailsCreatedCode int = 201

/*PostOperationsGetNodeEdgePointDetailsCreated No response

swagger:response postOperationsGetNodeEdgePointDetailsCreated
*/
type PostOperationsGetNodeEdgePointDetailsCreated struct {
}

// NewPostOperationsGetNodeEdgePointDetailsCreated creates PostOperationsGetNodeEdgePointDetailsCreated with default headers values
func NewPostOperationsGetNodeEdgePointDetailsCreated() *PostOperationsGetNodeEdgePointDetailsCreated {

	return &PostOperationsGetNodeEdgePointDetailsCreated{}
}

// WriteResponse to the client
func (o *PostOperationsGetNodeEdgePointDetailsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostOperationsGetNodeEdgePointDetailsBadRequestCode is the HTTP code returned for type PostOperationsGetNodeEdgePointDetailsBadRequest
const PostOperationsGetNodeEdgePointDetailsBadRequestCode int = 400

/*PostOperationsGetNodeEdgePointDetailsBadRequest Internal error

swagger:response postOperationsGetNodeEdgePointDetailsBadRequest
*/
type PostOperationsGetNodeEdgePointDetailsBadRequest struct {
}

// NewPostOperationsGetNodeEdgePointDetailsBadRequest creates PostOperationsGetNodeEdgePointDetailsBadRequest with default headers values
func NewPostOperationsGetNodeEdgePointDetailsBadRequest() *PostOperationsGetNodeEdgePointDetailsBadRequest {

	return &PostOperationsGetNodeEdgePointDetailsBadRequest{}
}

// WriteResponse to the client
func (o *PostOperationsGetNodeEdgePointDetailsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
