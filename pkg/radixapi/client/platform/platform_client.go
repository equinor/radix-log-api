// Code generated by go-swagger; DO NOT EDIT.

package platform

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new platform API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for platform API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	RegisterApplication(params *RegisterApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegisterApplicationOK, error)

	SearchApplications(params *SearchApplicationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchApplicationsOK, error)

	ShowApplications(params *ShowApplicationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShowApplicationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  RegisterApplication creates an application registration
*/
func (a *Client) RegisterApplication(params *RegisterApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegisterApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterApplicationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "registerApplication",
		Method:             "POST",
		PathPattern:        "/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegisterApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RegisterApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for registerApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchApplications gets applications by name n o t e doesn t get application summary latest job environments
*/
func (a *Client) SearchApplications(params *SearchApplicationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchApplicationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchApplicationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchApplications",
		Method:             "POST",
		PathPattern:        "/applications/_search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchApplicationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchApplicationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchApplications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ShowApplications lists the applications n o t e doesn t get application summary latest job environments
*/
func (a *Client) ShowApplications(params *ShowApplicationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShowApplicationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShowApplicationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "showApplications",
		Method:             "GET",
		PathPattern:        "/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ShowApplicationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ShowApplicationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for showApplications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}