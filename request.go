// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package dnspodcn

// Request is an interface that should be implemented by all API requests.
type Request interface {
	// GetEndpoint returns the endpoint of API without the basic URL.
	GetEndpoint() string
	// GetMethod returns the HTTP request method.
	GetMethod() string
}

// BasicRequest contains the basic fields of API request.
type BasicRequest struct {
	// The HTTP request method.
	Method string `json:"-"`

	// The endpoint of API.
	Endpoint string `json:"-"`
}

// NewBasicRequest returns a basic request with the given method and endpoint.
func NewBasicRequest(method, endpoint string) BasicRequest {
	return BasicRequest{
		Method:   method,
		Endpoint: endpoint,
	}
}

// GetMethod returns the HTTP request method.
func (req BasicRequest) GetMethod() string {
	return req.Method
}

// GetEndpoint returns the endpoint of API.
func (req BasicRequest) GetEndpoint() string {
	return req.Endpoint
}
