// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package dnspodcn

import (
	"net/http"
	"strconv"
	"time"

	"github.com/libdns/libdns"
)

const (
	defaultRecordLine = "默认"
)

// CreateRecordRequest is a request for creating a DNS record.
type CreateRecordRequest struct {
	BasicRequest
	Domain       string `json:"domain,omitempty"`
	DomainID     string `json:"domain_id,omitempty"`
	Subdomain    string `json:"sub_domain,omitempty"`
	RecordType   string `json:"record_type,omitempty"`
	RecordLine   string `json:"record_line,omitempty"`
	RecordLineID string `json:"record_line_id,omitempty"`
	Value        string `json:"value,omitempty"`
	MX           string `json:"mx,omitempty"`
	TTL          string `json:"ttl,omitempty"`
	Weight       string `json:"weight,omitempty"`
}

// NewCreateRecordRequest returns a request for creating a DNS record.
func NewCreateRecordRequest(domain, subdomain, recordType, value string) *CreateRecordRequest {
	return &CreateRecordRequest{
		BasicRequest: NewBasicRequest(http.MethodPost, "Record.Create"),
		Domain:       domain,
		Subdomain:    subdomain,
		RecordType:   recordType,
		Value:        value,
		RecordLine:   defaultRecordLine,
	}
}

// CreateRecordResponse is a response of creating a DNS record.
type CreateRecordResponse struct {
	BasicResponse
	Record RecordResponse `json:"record"`
}

// RecordResponse contains the infomation of DNS record field.
type RecordResponse struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Line    string `json:"line"`
	LineID  string `json:"line_id"`
	TTL     string `json:"ttl"`
	Value   string `json:"value"`
	Weight  string `json:"weight"`
	MX      string `json:"mx"`
	Enabled string `json:"enabled"`
	Remark  string `json:"remark"`
}

func (resp RecordResponse) libdnsRecord() libdns.Record {
	ttl, _ := strconv.Atoi(resp.TTL)
	return libdns.Record{
		ID:    resp.ID,
		Type:  resp.Type,
		Name:  resp.Name,
		Value: resp.Value,
		TTL:   time.Second * time.Duration(ttl),
	}
}

// QueryRecordsRequest is a request for querying DNS records.
type QueryRecordsRequest struct {
	BasicRequest
	Domain       string `json:"domain,omitempty"`
	DomainID     string `json:"domain_id,omitempty"`
	Subdomain    string `json:"sub_domain,omitempty"`
	RecordType   string `json:"record_type,omitempty"`
	RecordLine   string `json:"record_line,omitempty"`
	RecordLineID string `json:"record_line_id,omitempty"`
	Offset       string `json:"offset,omitempty"`
	Length       string `json:"length,omitempty"`
	Keyword      string `json:"keyword,omitempty"`
}

// NewQueryRecordsRequest returns a request for querying DNS records.
func NewQueryRecordsRequest(domain string) *QueryRecordsRequest {
	return &QueryRecordsRequest{
		BasicRequest: NewBasicRequest(http.MethodPost, "Record.List"),
		Domain:       domain,
	}
}

// QueryRecordsResponse is the response of querying DNS records.
type QueryRecordsResponse struct {
	BasicResponse
	Domain struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"domain"`
	Info struct {
		SubDomains  string `json:"sub_domains"`
		RecordTotal string `json:"record_total"`
		RecordsNum  string `json:"records_num"`
	} `json:"info"`
	Records []RecordResponse `json:"records"`
}

// Error returns the error of response.
func (resp QueryRecordsResponse) Error() error {
	if resp.Status.Code == CodeSuccess || resp.Status.Code == CodeNoRecords {
		return nil
	}

	return resp.Status.Error()
}

// DeleteRecordRequest is a request for deleting a DNS record.
type DeleteRecordRequest struct {
	BasicRequest
	Domain   string `json:"domain,omitempty"`
	DomainID string `json:"domain_id,omitempty"`
	RecordID string `json:"record_id,omitempty"`
}

// NewDeleteRecordRequest returns a request for deleting a DNS record.
func NewDeleteRecordRequest(domain, recordID string) *DeleteRecordRequest {
	return &DeleteRecordRequest{
		BasicRequest: NewBasicRequest(http.MethodPost, "Record.Remove"),
		Domain:       domain,
		RecordID:     recordID,
	}
}

// DeleteRecordResponse is the response of deleting a DNS record.
type DeleteRecordResponse struct {
	BasicResponse
}

// UpdateRecordRequest is a request for updating a DNS record.
type UpdateRecordRequest struct {
	BasicRequest
	Domain       string `json:"domain,omitempty"`
	DomainID     string `json:"domain_id,omitempty"`
	Subdomain    string `json:"sub_domain,omitempty"`
	RecordID     string `json:"record_id,omitempty"`
	RecordType   string `json:"record_type,omitempty"`
	RecordLine   string `json:"record_line,omitempty"`
	RecordLineID string `json:"record_line_id,omitempty"`
	Value        string `json:"value,omitempty"`
	MX           string `json:"mx,omitempty"`
	TTL          string `json:"ttl,omitempty"`
	Weight       string `json:"weight,omitempty"`
}

// NewUpdateRecordRequest returns a request for updating a DNS record.
func NewUpdateRecordRequest(domain, recordID, recordType, value string) *UpdateRecordRequest {
	return &UpdateRecordRequest{
		BasicRequest: NewBasicRequest(http.MethodPost, "Record.Modify"),
		Domain:       domain,
		RecordID:     recordID,
		RecordType:   recordType,
		Value:        value,
		RecordLine:   defaultRecordLine,
	}
}

// UpdateRecordResponse is the response of updating a DNS record.
type UpdateRecordResponse struct {
	BasicResponse
	Record struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Value  string `json:"value"`
		Status string `json:"status"`
	} `json:"record"`
}
