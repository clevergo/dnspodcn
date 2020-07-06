// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package dnspodcn

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	name       = "Go DNSPod Client"
	version    = "dev"
	apiBaseURL = "https://dnsapi.cn"
)

// Client is an API client.
type Client struct {
	// The ID of application.
	AppID string
	// The API token of application.
	AppToken string
	// The base URL of API.
	BaseURL string
	// The language of API response, en or zh.
	Language string
}

// NewClient returns a client with the given app ID, token and default settings.
func NewClient(appID, appToken string) *Client {
	return &Client{
		AppID:    appID,
		AppToken: appToken,
		BaseURL:  apiBaseURL,
		Language: "en",
	}
}

func (c *Client) newRequest(ctx context.Context, req Request) (*http.Request, error) {
	vs, err := toValues(req)
	if err != nil {
		return nil, err
	}
	// Common parameters.
	vs.Set("format", "json")
	vs.Set("login_token", fmt.Sprintf("%s,%s", c.AppID, c.AppToken))
	vs.Set("lang", c.Language)
	fmt.Println(vs.Encode())
	body := strings.NewReader(vs.Encode())
	url := fmt.Sprintf("%s/%s", c.BaseURL, req.GetEndpoint())
	httpReq, err := http.NewRequestWithContext(ctx, req.GetMethod(), url, body)
	if err != nil {
		return nil, err
	}
	// API doesn't support JSON request.
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Set("User-Agent", fmt.Sprintf("%s/%s", name, version))
	return httpReq, nil
}

// Do makes a request and returns the response.
func (c *Client) Do(ctx context.Context, req Request, resp Response) error {
	httpReq, err := c.newRequest(ctx, req)
	if err != nil {
		return err
	}
	result, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer result.Body.Close()

	if err = json.NewDecoder(result.Body).Decode(resp); err != nil {
		return err
	}
	if err = resp.Error(); err != nil {
		return err
	}
	return nil
}

// CreateRecord creates a DNS record.
func (c *Client) CreateRecord(ctx context.Context, req *CreateRecordRequest) (*CreateRecordResponse, error) {
	resp := &CreateRecordResponse{}
	err := c.Do(ctx, req, resp)
	return resp, err
}

// QueryRecords queries DNS records.
func (c *Client) QueryRecords(ctx context.Context, req *QueryRecordsRequest) (*QueryRecordsResponse, error) {
	resp := &QueryRecordsResponse{}
	err := c.Do(ctx, req, resp)
	return resp, err
}

// DeleteRecord deletes a DNS record.
func (c *Client) DeleteRecord(ctx context.Context, req *DeleteRecordRequest) (*DeleteRecordResponse, error) {
	resp := &DeleteRecordResponse{}
	err := c.Do(ctx, req, resp)
	return resp, err
}

// UpdateRecord updates a DNS record.
func (c *Client) UpdateRecord(ctx context.Context, req *UpdateRecordRequest) (*UpdateRecordResponse, error) {
	resp := &UpdateRecordResponse{}
	err := c.Do(ctx, req, resp)
	return resp, err
}
