// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package dnspodcn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusResponseError(t *testing.T) {
	resp := StatusResponse{
		Code: CodeSuccess,
	}
	assert.Nil(t, resp.Error())

	resp.Code = "-1"
	assert.NotNil(t, resp.Error())
}

func TestBasicResponseError(t *testing.T) {
	resp := BasicResponse{
		Status: StatusResponse{
			Code: CodeSuccess,
		},
	}
	assert.Nil(t, resp.Error())

	resp.Status.Code = "-1"
	assert.NotNil(t, resp.Error())
}
