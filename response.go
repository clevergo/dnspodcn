// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package dnspodcn

import "fmt"

// Response is an interface.
type Response interface {
	// Error returns the error of the status of request.
	Error() error
}

// StatusResponse contains the infomation of status field.
type StatusResponse struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

// Error returns the error of the response.
func (resp StatusResponse) Error() error {
	if resp.Code == CodeSuccess {
		return nil
	}
	return fmt.Errorf("%s, code %s", resp.Message, resp.Code)
}

// BasicResponse contains the basic fields of response.
type BasicResponse struct {
	Status StatusResponse `json:"status"`
}

// Error returns the error of the response.
func (resp BasicResponse) Error() error {
	return resp.Status.Error()
}
