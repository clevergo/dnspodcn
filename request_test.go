// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package dnspodcn

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicRequest(t *testing.T) {
	cases := []struct {
		method, endpoint string
	}{
		{http.MethodGet, "/get"},
		{http.MethodPost, "/post"},
	}
	for _, test := range cases {
		req := NewBasicRequest(test.method, test.endpoint)
		assert.Equal(t, test.method, req.Method)
		assert.Equal(t, test.method, req.GetMethod())
		assert.Equal(t, test.endpoint, req.Endpoint)
		assert.Equal(t, test.endpoint, req.GetEndpoint())
	}
}
